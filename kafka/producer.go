// pkg/kafka/producer.go

package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// Producer is a Kafka producer
type Producer struct {
	producer sarama.SyncProducer
	topic    string
	disabled bool
}

// Event represents a Kafka event
type Event struct {
	EventType string      `json:"event_type"`
	Timestamp time.Time   `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}

// NewProducer creates a new Kafka producer
// NewProducer creates a new Kafka producer with retry logic
func NewProducer(config *KafkaConfig) (*Producer, error) {
    saramaConfig := config.NewSaramaConfig()
    
    var producer sarama.SyncProducer
    var err error
    
    // Retry for up to 30 seconds
    for i := 0; i < 6; i++ {
        producer, err = sarama.NewSyncProducer(config.Brokers, saramaConfig)
        if err == nil {
            break
        }
        
        log.Printf("Attempt %d: Failed to create Kafka producer: %v", i+1, err)
        time.Sleep(5 * time.Second)
    }
    
    if err != nil {
        log.Printf("Warning: Kafka unavailable, producer running in no-op mode: %v", err)
        return &Producer{disabled: true, topic: config.Topic}, nil
    }

    return &Producer{
        producer: producer,
        topic:    config.Topic,
    }, nil
}

// Close closes the producer
func (p *Producer) Close() error {
	if p.disabled {
		return nil
	}
	return p.producer.Close()
}

// SendMessage sends a message to Kafka
func (p *Producer) SendMessage(eventType string, payload interface{}) error {
	if p.disabled {
		return nil
	}
	event := Event{
		EventType: eventType,
		Timestamp: time.Now(),
		Payload:   payload,
	}
	
	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %v", err)
	}
	
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(jsonData),
		Key:   sarama.StringEncoder(eventType),
	}
	
	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	
	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	return nil
}

// SendPatientEvent sends a patient-related event to Kafka
func (p *Producer) SendPatientEvent(eventType string, patientID string, data interface{}) error {
	event := map[string]interface{}{
		"patient_id": patientID,
		"data":       data,
	}
	
	return p.SendMessage(eventType, event)
}


// SendAppointmentEvent sends an appointment-related event to Kafka
func (p *Producer) SendAppointmentEvent(eventType string, appointmentID string, data interface{}) error {
	event := map[string]interface{}{
		"appointment_id": appointmentID,
		"data":           data,
	}
	
	return p.SendMessage(eventType, event)
}

// SendQueueEvent sends a queue-related event to Kafka
func (p *Producer) SendQueueEvent(eventType string, queueID string, data interface{}) error {
	event := map[string]interface{}{
		"queue_id": queueID,
		"data":     data,
	}
	
	return p.SendMessage(eventType, event)
}


// SendStaffEvent sends a staff-related event to Kafka
func (p *Producer) SendStaffEvent(eventType string, staffID int32, data interface{}) error {
	event := map[string]interface{}{
		"staff_id": staffID,
		"data":     data,
	}
	
	return p.SendMessage(eventType, event)
}

// SendPayrollEvent sends a payroll-related event to Kafka
func (p *Producer) SendPayrollEvent(eventType string, payrollID int32, data interface{}) error {
	event := map[string]interface{}{
		"payroll_id": payrollID,
		"data":       data,
	}
	
	return p.SendMessage(eventType, event)
}

// SendLeaveEvent sends a leave-related event to Kafka
func (p *Producer) SendLeaveEvent(eventType string, leaveID int32, data interface{}) error {
	event := map[string]interface{}{
		"leave_id": leaveID,
		"data":     data,
	}
	
	return p.SendMessage(eventType, event)
}

// SendDepartmentEvent sends a department-related event to Kafka
func (p *Producer) SendDepartmentEvent(eventType string, departmentID int32, data interface{}) error {
	event := map[string]interface{}{
		"department_id": departmentID,
		"data":          data,
	}
	
	return p.SendMessage(eventType, event)
}

// SendAttendanceEvent sends an attendance-related event to Kafka
func (p *Producer) SendAttendanceEvent(eventType string, attendanceID int32, data interface{}) error {
	event := map[string]interface{}{
		"attendance_id": attendanceID,
		"data":          data,
	}
	
	return p.SendMessage(eventType, event)
}


// Medical Records specific event senders
func (p *Producer) SendDiagnosisEvent(eventType, diagnosisID string, data interface{}) error {
	event := map[string]interface{}{
		"diagnosis_id": diagnosisID,
		"data":         data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendNoteEvent(eventType, noteID string, data interface{}) error {
	event := map[string]interface{}{
		"note_id": noteID,
		"data":    data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendVitalSignEvent(eventType, vitalSignID string, data interface{}) error {
	event := map[string]interface{}{
		"vital_sign_id": vitalSignID,
		"data":          data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendBMIEvent(eventType, bmiID string, data interface{}) error {
	event := map[string]interface{}{
		"bmi_id": bmiID,
		"data":   data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendDocumentEvent(eventType, documentID string, data interface{}) error {
	event := map[string]interface{}{
		"document_id": documentID,
		"data":        data,
	}
	return p.SendMessage(eventType, event)
}


// SendAdmissionEvent sends an admission-related event to Kafka
func (p *Producer) SendAdmissionEvent(eventType string, admissionID string, data interface{}) error {
	event := map[string]interface{}{
		"admission_id": admissionID,
		"data":         data,
	}
	return p.SendMessage(eventType, event)
}

// SendBedEvent sends a bed-related event to Kafka
func (p *Producer) SendBedEvent(eventType string, bedID string, data interface{}) error {
	event := map[string]interface{}{
		"bed_id": bedID,
		"data":   data,
	}
	return p.SendMessage(eventType, event)
}

// SendWardEvent sends a ward-related event to Kafka
func (p *Producer) SendWardEvent(eventType string, wardID string, data interface{}) error {
	event := map[string]interface{}{
		"ward_id": wardID,
		"data":    data,
	}
	return p.SendMessage(eventType, event)
}

// SendDischargeEvent sends a discharge-related event to Kafka
func (p *Producer) SendDischargeEvent(eventType string, dischargeID string, data interface{}) error {
	event := map[string]interface{}{
		"discharge_id": dischargeID,
		"data":         data,
	}
	return p.SendMessage(eventType, event)
}

// SendTransferEvent sends a transfer-related event to Kafka
func (p *Producer) SendTransferEvent(eventType string, transferID string, data interface{}) error {
	event := map[string]interface{}{
		"transfer_id": transferID,
		"data":        data,
	}
	return p.SendMessage(eventType, event)
}

// Supply Chain specific event senders
func (p *Producer) SendStockItemEvent(eventType string, itemID string, data interface{}) error {
	event := map[string]interface{}{
		"item_id": itemID,
		"data":   data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendRequisitionEvent(eventType string, requisitionID string, data interface{}) error {
	event := map[string]interface{}{
		"requisition_id": requisitionID,
		"data":          data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendLPOEvent(eventType string, lpoID string, data interface{}) error {
	event := map[string]interface{}{
		"lpo_id": lpoID,
		"data":  data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendGRNEvent(eventType string, grnID string, data interface{}) error {
	event := map[string]interface{}{
		"grn_id": grnID,
		"data":  data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendStockAdjustmentEvent(eventType string, adjustmentID string, data interface{}) error {
	event := map[string]interface{}{
		"adjustment_id": adjustmentID,
		"data":         data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendInterStoreTransferEvent(eventType string, transferID string, data interface{}) error {
	event := map[string]interface{}{
		"transfer_id": transferID,
		"data":       data,
	}
	return p.SendMessage(eventType, event)
}

func (p *Producer) SendCreditNoteEvent(eventType string, creditNoteID string, data interface{}) error {
	event := map[string]interface{}{
		"credit_note_id": creditNoteID,
		"data":         data,
	}
	return p.SendMessage(eventType, event)
}

// SendMedicationEvent sends a medication-related event to Kafka
func (p *Producer) SendMedicationEvent(eventType string, medicationID string, data interface{}) error {
	event := map[string]interface{}{
		"medication_id": medicationID,
		"data":         data,
	}
	return p.SendMessage(eventType, event)
}

// SendPrescriptionEvent sends a prescription-related event to Kafka
func (p *Producer) SendPrescriptionEvent(eventType string, prescriptionID string, data interface{}) error {
	event := map[string]interface{}{
		"prescription_id": prescriptionID,
		"data":           data,
	}
	return p.SendMessage(eventType, event)
}

// SendStockRequestEvent sends a stock request-related event to Kafka
func (p *Producer) SendStockRequestEvent(eventType string, requestID string, data interface{}) error {
	event := map[string]interface{}{
		"request_id": requestID,
		"data":      data,
	}
	return p.SendMessage(eventType, event)
}

// SendStockMovementEvent sends a stock movement-related event to Kafka
func (p *Producer) SendStockMovementEvent(eventType string, movementID string, data interface{}) error {
	event := map[string]interface{}{
		"movement_id": movementID,
		"data":       data,
	}
	return p.SendMessage(eventType, event)
}

// SendDispensedMedicationEvent sends a dispensed medication-related event to Kafka
func (p *Producer) SendDispensedMedicationEvent(eventType string, dispenseID string, data interface{}, items interface{}) error {
	event := map[string]interface{}{
		"dispense_id": dispenseID,
		"data":       data,
		"items":      items,
	}
	return p.SendMessage(eventType, event)
}

// SendMedicationReturnEvent sends a medication return-related event to Kafka
func (p *Producer) SendMedicationReturnEvent(eventType string, returnID string, data interface{}) error {
	event := map[string]interface{}{
		"return_id": returnID,
		"data":     data,
	}
	return p.SendMessage(eventType, event)
}

// SendMedicationScheduleEvent sends a medication schedule-related event to Kafka
func (p *Producer) SendMedicationScheduleEvent(eventType string, scheduleID string, data interface{}) error {
	event := map[string]interface{}{
		"schedule_id": scheduleID,
		"data":       data,
	}
	return p.SendMessage(eventType, event)
}

// SendMedicationIntakeEvent sends a medication intake-related event to Kafka
func (p *Producer) SendMedicationIntakeEvent(eventType string, intakeID string, data interface{}) error {
	event := map[string]interface{}{
		"intake_id": intakeID,
		"data":     data,
	}
	return p.SendMessage(eventType, event)
}


// SendLabEvent sends a laboratory-related event to Kafka
func (p *Producer) SendLabEvent(eventType string, patientID string, data interface{}) error {
	event := map[string]interface{}{
		"patient_id": patientID,
		"data":       data,
	}
	return p.SendMessage(eventType, event)
}

// SendIssueEvent sends an issue-related event to Kafka
func (p *Producer) SendIssueEvent(eventType string, reportID int32, data interface{}) error {
	event := map[string]interface{}{
		"report_id": reportID,
		"data":      data,
	}
	return p.SendMessage(eventType, event)
}


// SendUserEvent sends a user-related event to Kafka
func (p *Producer) SendUserEvent(eventType string, userID string, data interface{}) error {
	event := map[string]interface{}{
		"user_id": userID,
		"data":    data,
	}
	return p.SendMessage(eventType, event)
}

// SendHospitalEvent sends a hospital-related event to Kafka
func (p *Producer) SendHospitalEvent(eventType string, hospitalID string, data interface{}) error {
	event := map[string]interface{}{
		"hospital_id": hospitalID,
		"data":        data,
	}
	return p.SendMessage(eventType, event)
}


// SendBillingEvent sends a billing-related event to Kafka
func (p *Producer) SendBillingEvent(eventType string, billingID string, data interface{}) error {
	event := map[string]interface{}{
		"billing_id": billingID,
		"data":       data,
	}
	return p.SendMessage(eventType, event)
}

// SendInvoiceEvent sends an invoice-related event to Kafka
func (p *Producer) SendInvoiceEvent(eventType string, invoiceID string, data interface{}) error {
	event := map[string]interface{}{
		"invoice_id": invoiceID,
		"data":       data,
	}
	return p.SendMessage(eventType, event)
}

// SendPaymentEvent sends a payment-related event to Kafka
func (p *Producer) SendPaymentEvent(eventType string, paymentID string, data interface{}) error {
	event := map[string]interface{}{
		"payment_id": paymentID,
		"data":       data,
	}
	return p.SendMessage(eventType, event)
}


func (p *Producer) SendServiceRecordEvent(eventType string, recordID, patientID, serviceType string, data interface{}) error {
    event := map[string]interface{}{
        "record_id":    recordID,
        "patient_id":   patientID,
        "service_type": serviceType,
        "data":         data,
    }
    return p.SendMessage(eventType, event)
}


func (p *Producer) SendPermissionEvent(eventType string, route string, data interface{}) error {
    event := map[string]interface{}{
        "route": route,
        "data": data,
    }
    return p.SendMessage(eventType, event)
}

func (p *Producer) SendActionPermissionEvent(eventType string, resource string, action string, data interface{}) error {
    event := map[string]interface{}{
        "resource": resource,
        "action":   action,
        "data":     data,
    }
    return p.SendMessage(eventType, event)
}