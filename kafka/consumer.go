// pkg/kafka/consumer.go

package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// MessageHandler is a function that handles a Kafka message
type MessageHandler func(message []byte) error

// Consumer is a Kafka consumer
type Consumer struct {
	consumer   sarama.ConsumerGroup
	topic      string
	handlers   map[string]MessageHandler
	ready      chan bool
	readyOnce  sync.Once
	mu         sync.Mutex
	disabled   bool
}

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct {
	consumer *Consumer
}

// NewConsumer creates a new Kafka consumer
func NewConsumer(config *KafkaConfig) (*Consumer, error) {
	saramaConfig := config.NewSaramaConfig()

	consumer, err := sarama.NewConsumerGroup(config.Brokers, config.GroupID, saramaConfig)
	if err != nil {
		log.Printf("Warning: Kafka unavailable, consumer running in no-op mode: %v", err)
		return &Consumer{
			disabled: true,
			topic:    config.Topic,
			handlers: make(map[string]MessageHandler),
		}, nil
	}

	return &Consumer{
		consumer: consumer,
		topic:    config.Topic,
		handlers: make(map[string]MessageHandler),
		ready:    make(chan bool),
	}, nil
}

// Close closes the consumer
func (c *Consumer) Close() error {
	if c.disabled {
		return nil
	}
	return c.consumer.Close()
}

// RegisterHandler registers a handler for a specific event type
func (c *Consumer) RegisterHandler(eventType string, handler MessageHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[eventType] = handler
}

// Start starts consuming messages
func (c *Consumer) Start(ctx context.Context) error {
	if c.disabled {
		return nil
	}
	// Create a new context with cancellation
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	// Handle signals for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	
	// Start consuming in a goroutine
	go func() {
		for {
			// Consume messages
			handler := &ConsumerGroupHandler{
				consumer: c,
			}
			
			err := c.consumer.Consume(ctx, []string{c.topic}, handler)
			if err != nil {
				log.Printf("Error from consumer: %v", err)
			}
			
			// Check if context was cancelled
			if ctx.Err() != nil {
				return
			}
			
			// Mark the consumer as ready
			c.ready = make(chan bool)
		}
	}()
	
	// Wait for consumer to be ready
	<-c.ready
	log.Println("Kafka consumer started")
	
	// Wait for signals
	select {
	case <-signals:
		log.Println("Received signal, shutting down")
		cancel()
	case <-ctx.Done():
		log.Println("Context cancelled, shutting down")
	}
	
	return nil
}

// Setup is run at the beginning of a new session
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	// Only close the ready channel once to avoid panic on channel close
	h.consumer.readyOnce.Do(func() {
		close(h.consumer.ready)
	})
	return nil
}

// Cleanup is run at the end of a session
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim is called for each consumer session to consume messages
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// Parse the event
		var event Event
		if err := json.Unmarshal(message.Value, &event); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			session.MarkMessage(message, "")
			continue
		}
		
		// Get the handler for this event type
		h.consumer.mu.Lock()
		handler, exists := h.consumer.handlers[event.EventType]
		h.consumer.mu.Unlock()
		
		if exists {
			// Handle the message
			if err := handler(message.Value); err != nil {
				log.Printf("Failed to handle message: %v", err)
			}
		} else {
			log.Printf("No handler registered for event type: %s", event.EventType)
		}
		
		// Mark the message as processed
		session.MarkMessage(message, "")
	}
	
	return nil
}