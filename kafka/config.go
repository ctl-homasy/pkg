// pkg/kafka/config.go

package kafka

import (
	"crypto/tls"
	"time"

	"github.com/IBM/sarama"
)

// KafkaConfig holds configuration for Kafka
type KafkaConfig struct {
	Brokers    []string
	Topic      string
	GroupID    string
	Username   string
	Password   string
	UseTLS     bool
	ClientType string // Added to distinguish between different services
}

// NewKafkaConfig creates a new Kafka configuration.
func NewKafkaConfig(brokers []string, topic, groupID, username, password string, useTLS bool, clientType string) *KafkaConfig {
	return &KafkaConfig{
		Brokers:    brokers,
		Topic:      topic,
		GroupID:    groupID,
		Username:   username,
		Password:   password,
		UseTLS:     useTLS,
		ClientType: clientType,
	}
}

// NewSaramaConfig creates a new Sarama configuration
func (c *KafkaConfig) NewSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	
	// Set producer config
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	
	// Set consumer config
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	
	config.ClientID = c.ClientType
	
	// Set version
	config.Version = sarama.V2_8_0_0
	
	// Set timeout
	config.Net.DialTimeout = 10 * time.Second
	config.Net.ReadTimeout = 10 * time.Second
	config.Net.WriteTimeout = 10 * time.Second
	
	// Set authentication if provided
	if c.Username != "" && c.Password != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = c.Username
		config.Net.SASL.Password = c.Password
		config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}
	
	// Set TLS if enabled
	if c.UseTLS {
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	
	return config
}