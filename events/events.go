// pkg/events/events.go

package events

import (
	"encoding/json"
	"fmt"
	"log"

	appointment_pb "github.com/ctl-homasy/pkg/shared/appointment-service/proto"
	hr_pb "github.com/ctl-homasy/pkg/shared/human-resource-service/proto"
	patient_pb "github.com/ctl-homasy/pkg/shared/patient-service/proto"
	medical_records_pb "github.com/ctl-homasy/pkg/shared/medical-records-service/proto"
	ward_pb "github.com/ctl-homasy/pkg/shared/ward-service/proto"
	supply_chain_pb "github.com/ctl-homasy/pkg/shared/supply-chain-service/proto"
	pharmacy_pb "github.com/ctl-homasy/pkg/shared/pharmacy-service/proto"
	laboratory_pb "github.com/ctl-homasy/pkg/shared/laboratory-service/proto"
	issue_pb "github.com/ctl-homasy/pkg/shared/issue-report-service/proto"
	user_pb "github.com/ctl-homasy/pkg/shared/user-service/proto"
	billing_pb "github.com/ctl-homasy/pkg/shared/billing-service/proto"
) 

// Event types
const (

	// User service events
	UserCreated    = "user.created"
	UserUpdated    = "user.updated"
	UserDeleted    = "user.deleted"
	HospitalCreated = "hospital.created"
	HospitalUpdated = "hospital.updated"
	HospitalDeleted = "hospital.deleted"

	// Appointment service events
	AppointmentCreated = "appointment.created"
	AppointmentUpdated = "appointment.updated"
	AppointmentDeleted = "appointment.deleted"
	AppointmentCancelled = "appointment.cancelled"
	AppointmentNoShow = "appointment.noshow"
	QueueCreated       = "queue.created"
	QueueStatusUpdated = "queue.status.updated"
	QueueCancelled     = "queue.cancelled"

	// Medical Records events
	DiagnosisCreated     = "diagnosis.created"
	DiagnosisUpdated     = "diagnosis.updated"
	DiagnosisDeleted     = "diagnosis.deleted"
	NoteCreated          = "note.created"
	NoteUpdated          = "note.updated"
	NoteDeleted          = "note.deleted"
	VitalSignRecorded    = "vitalsign.recorded"
	BMIRecorded          = "bmi.recorded"
	DocumentUploaded     = "document.uploaded"
	DocumentUpdated      = "document.updated"
	DocumentDeleted      = "document.deleted"
	DocumentVersionAdded = "document.version.added"
	DocumentShared       = "document.shared"

	// Patient service events
	PatientCreated   = "patient.created"
	PatientUpdated   = "patient.updated"
	PatientDeleted   = "patient.deleted"
	PatientRequested = "patient.requested"

	// HR service events
	StaffRegistered    = "staff.registered"
	StaffUpdated       = "staff.updated"
	StaffDeleted       = "staff.deleted"
	PayrollCreated     = "payroll.created"
	PayrollUpdated     = "payroll.updated"
	PayrollDeleted     = "payroll.deleted"
	LeaveRequested     = "leave.requested"
	LeaveApproved      = "leave.approved"
	LeaveRejected      = "leave.rejected"
	DepartmentCreated  = "department.created"
	DepartmentUpdated  = "department.updated"
	DepartmentDeleted  = "department.deleted"
	AttendanceRecorded = "attendance.recorded"
	BenefitAdded       = "benefit.added"
	BenefitRemoved     = "benefit.removed"

	// Ward service events
	AdmissionCreated     = "admission.created"
	AdmissionUpdated     = "admission.updated"
	AdmissionDeleted     = "admission.deleted"
	BedCreated           = "bed.created"
	BedStatusUpdated     = "bed.status.updated"
	BedDeleted           = "bed.deleted"
	WardCreated          = "ward.created"
	WardUpdated          = "ward.updated"
	WardDeleted          = "ward.deleted"
	DischargeCreated     = "discharge.created"
	DischargeUpdated     = "discharge.updated"
	PatientTransferred   = "patient.transferred"


	// Stock Item events
	StockItemCreated = "stock_item.created"
	StockItemUpdated = "stock_item.updated"
	StockItemDeleted = "stock_item.deleted"

	// Requisition events
	RequisitionCreated  = "requisition.created"
	RequisitionApproved = "requisition.approved"
	RequisitionRejected = "requisition.rejected"

	// LPO events
	LPOCreated  = "lpo.created"
	LPOApproved = "lpo.approved"
	LPORejected = "lpo.rejected"

	// GRN events
	GRNCreated        = "grn.created"
	GRNQualityChecked = "grn.quality_checked"

	// Stock Adjustment events
	StockAdjustmentCreated  = "stock_adjustment.created"
	StockAdjustmentApproved = "stock_adjustment.approved"
	StockAdjustmentRejected = "stock_adjustment.rejected"

	// Transfer events
	TransferCreated   = "transfer.created"
	TransferApproved  = "transfer.approved"
	TransferRejected  = "transfer.rejected"
	TransferReceived  = "transfer.received"

	// Credit Note events
	CreditNoteCreated  = "credit_note.created"
	CreditNoteApproved = "credit_note.approved"


	// Pharmacy Events
	MedicationCreated       = "medication.created"
	MedicationUpdated       = "medication.updated"
	MedicationDeleted       = "medication.deleted"
	PrescriptionCreated     = "prescription.created"
	PrescriptionUpdated     = "prescription.updated"
	PrescriptionDeleted     = "prescription.deleted"
	StockRequestCreated     = "stock_request.created"
	StockRequestApproved    = "stock_request.approved"
	StockMovementCreated    = "stock_movement.created"
	MedicationDispensed     = "medication.dispensed"
	MedicationReturned      = "medication.returned"
	MedicationScheduleCreated = "medication_schedule.created"
	MedicationIntakeRecorded = "medication_intake.recorded"

	// Laboratory service events
	LabRequestCreated = "lab.request.created"
	TestResultAdded   = "test.result.added"

	// Issue Report service events
	IssueReportCreated        = "issue.report.created"
	IssueReportStatusUpdated  = "issue.report.status.updated"


	InvoiceCreated    = "invoice.created"
	InvoiceUpdated    = "invoice.updated"
	InvoiceDeleted    = "invoice.deleted"
	InvoiceSent       = "invoice.sent"
	PaymentProcessed  = "payment.processed"
	PaymentReceived  = "payment.received"

    ServiceRecordCreated        = "service_record.created"
    ServiceRecordStatusUpdated  = "service_record.status_updated"
    ServiceRecordBilled         = "service_record.billed"
    ServiceRecordDeleted        = "service_record.deleted"
    ServiceRecordCancelled      = "service_record.cancelled"
    ServiceRecordPriceUpdated   = "service_record.price_updated"


	PermissionAdded = "permission.added"
    PermissionUpdated = "permission.updated"

	// Action Permission service events
	ActionPermissionAdded   = "action.permission.added"
	ActionPermissionUpdated = "action.permission.updated"
)

// Event structs

type PermissionEvent struct {
    Route string      `json:"route"`
    Data  interface{} `json:"data"`
}

// Billing event structs
type InvoiceEvent struct {
	InvoiceID string      `json:"invoice_id"`
	Data      interface{} `json:"data"`
}

type PaymentEvent struct {
	PaymentID string      `json:"payment_id"`
	Data      interface{} `json:"data"`
}

// Add these event structs
type ServiceRecordEvent struct {
    RecordID    string      `json:"record_id"`
    PatientID   string      `json:"patient_id"`
    ServiceType string      `json:"service_type"`
    Data        interface{} `json:"data"`
}


type UserEvent struct {
	UserID string      `json:"user_id"`
	Data   interface{} `json:"data"`
}

type HospitalEvent struct {
	HospitalID string      `json:"hospital_id"`
	Data       interface{} `json:"data"`
}

type AppointmentEvent struct {
	AppointmentID string      `json:"appointment_id"`
	Data          interface{} `json:"data"`
}

type QueueEvent struct {
	QueueID string      `json:"queue_id"`
	Data    interface{} `json:"data"`
}

type PatientEvent struct {
	PatientID string      `json:"patient_id"`
	Data      interface{} `json:"data"`
}

type StaffEvent struct {
	StaffID int32       `json:"staff_id"`
	Data    interface{} `json:"data"`
}

type PayrollEvent struct {
	PayrollID int32       `json:"payroll_id"`
	Data      interface{} `json:"data"`
}

type LeaveEvent struct {
	LeaveID int32       `json:"leave_id"`
	Data    interface{} `json:"data"`
}

type DepartmentEvent struct {
	DepartmentID int32       `json:"department_id"`
	Data         interface{} `json:"data"`
}

type AttendanceEvent struct {
	AttendanceID int32       `json:"attendance_id"`
	Data         interface{} `json:"data"`
}

type DiagnosisEvent struct {
	DiagnosisID string      `json:"diagnosis_id"`
	Data        interface{} `json:"data"`
}

type NoteEvent struct {
	NoteID string      `json:"note_id"`
	Data   interface{} `json:"data"`
}

type VitalSignEvent struct {
	VitalSignID string      `json:"vital_sign_id"`
	Data        interface{} `json:"data"`
}

type DocumentEvent struct {
	DocumentID string      `json:"document_id"`
	Data       interface{} `json:"data"`
}

type AdmissionEvent struct {
	AdmissionID string      `json:"admission_id"`
	Data        interface{} `json:"data"`
}

type BedEvent struct {
	BedID string      `json:"bed_id"`
	Data  interface{} `json:"data"`
}

type WardEvent struct {
	WardID string      `json:"ward_id"`
	Data   interface{} `json:"data"`
}

type DischargeEvent struct {
	DischargeID string      `json:"discharge_id"`
	Data        interface{} `json:"data"`
}

type TransferEvent struct {
	TransferID string      `json:"transfer_id"`
	Data       interface{} `json:"data"`
}

type StockItemEvent struct {
	ItemID string       `json:"item_id"`
	Data   interface{} `json:"data"`
}

type RequisitionEvent struct {
	RequisitionID string         `json:"requisition_id"`
	Data          interface{} `json:"data"`
}

type LPOEvent struct {
	LPOID string                  `json:"lpo_id"`
	Data  interface{} `json:"data"`
}

type GRNEvent struct {
	GRNID string                 `json:"grn_id"`
	Data  interface{} `json:"data"`
}

type StockAdjustmentEvent struct {
	AdjustmentID string                 `json:"adjustment_id"`
	Data         interface{} `json:"data"`
}

type InterStoreTransferEvent struct {
	TransferID string                  `json:"transfer_id"`
	Data       interface{} `json:"data"`
}

type CreditNoteEvent struct {
	CreditNoteID string              `json:"credit_note_id"`
	Data         interface{} `json:"data"`
}


// Pharmacy event structs
type MedicationEvent struct {
	MedicationID string      `json:"medication_id"`
	Data         interface{} `json:"data"`
}

type PrescriptionEvent struct {
	PrescriptionID string      `json:"prescription_id"`
	Data           interface{} `json:"data"`
}

type StockRequestEvent struct {
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

type StockMovementEvent struct {
	MovementID string      `json:"movement_id"`
	Data       interface{} `json:"data"`
}

type DispensedMedicationEvent struct {
	DispenseID string                               `json:"dispense_id"`
	Data       *pharmacy_pb.GetDispensedMedicationResponse `json:"data"`
	Items      []*DispensedMedicationItem            `json:"items"` // Added to track what was dispensed
}

type DispensedMedicationItem struct {
	MedicationID int32  `json:"medication_id"`
	Quantity     int32  `json:"quantity"`
	ItemID       string `json:"item_id,omitempty"` // For supply chain mapping
}

type MedicationReturnEvent struct {
	ReturnID string      `json:"return_id"`
	Data     interface{} `json:"data"`
}

type MedicationScheduleEvent struct {
	ScheduleID string      `json:"schedule_id"`
	Data       interface{} `json:"data"`
}

type MedicationIntakeEvent struct {
	IntakeID string      `json:"intake_id"`
	Data     interface{} `json:"data"`
}

// Laboratory service event structs

type LabEvent struct {
	PatientID string      `json:"patient_id"`
	Data      interface{} `json:"data"`
}


// Issue Report service event structs

type IssueEvent struct {
	ReportID int32       `json:"report_id"`
	Data     interface{} `json:"data"`
}


func PermissionAddedEvent(route string, roles, users []string) ([]byte, error) {
    event := PermissionEvent{
        Route: route,
        Data: map[string]interface{}{
            "roles": roles,
            "users": users,
        },
    }
    return json.Marshal(event)
}

func PermissionUpdatedEvent(route string, roles, users []string) ([]byte, error) {
    event := PermissionEvent{
        Route: route,
        Data: map[string]interface{}{
            "roles": roles,
            "users": users,
        },
    }
    return json.Marshal(event)
}


// Billing event creators
func InvoiceCreatedEvent(invoice *billing_pb.Invoice) ([]byte, error) {
	event := InvoiceEvent{
		InvoiceID: invoice.InvoiceId,
		Data:      invoice,
	}
	return json.Marshal(event)
}

func InvoiceUpdatedEvent(invoice *billing_pb.Invoice) ([]byte, error) {
	event := InvoiceEvent{
		InvoiceID: invoice.InvoiceId,
		Data:      invoice,
	}
	return json.Marshal(event)
}

func InvoiceDeletedEvent(invoiceID string) ([]byte, error) {
	event := InvoiceEvent{
		InvoiceID: invoiceID,
		Data:      nil,
	}
	return json.Marshal(event)
}

func PaymentProcessedEvent(payment *billing_pb.ProcessPaymentResponse) ([]byte, error) {
	event := PaymentEvent{
		PaymentID: payment.PaymentId,
		Data:      payment,
	}
	return json.Marshal(event)
}

// Add these event creators
func ServiceRecordCreatedEvent(record *billing_pb.ServiceRecord) ([]byte, error) {
    event := ServiceRecordEvent{
        RecordID:    record.RecordId,
        PatientID:   record.PatientId,
        ServiceType: record.ServiceType,
        Data:        record,
    }
    return json.Marshal(event)
}

func ServiceRecordStatusUpdatedEvent(recordID, status string) ([]byte, error) {
    event := ServiceRecordEvent{
        RecordID: recordID,
        Data: map[string]string{
            "status": status,
        },
    }
    return json.Marshal(event)
}

func ServiceRecordDeletedEvent(recordID, patientID, reason string) ([]byte, error) {
    event := ServiceRecordEvent{
        RecordID:  recordID,
        PatientID: patientID,
        Data: map[string]string{
            "reason": reason,
        },
    }
    return json.Marshal(event)
}

func ServiceRecordCancelledEvent(recordID, patientID, reason string, refundAmount float64) ([]byte, error) {
    event := ServiceRecordEvent{
        RecordID:  recordID,
        PatientID: patientID,
        Data: map[string]interface{}{
            "reason":         reason,
            "refund_amount":  refundAmount,
        },
    }
    return json.Marshal(event)
}

func ServiceRecordPriceUpdatedEvent(recordID, patientID string, oldAmount, newAmount float64, reason string) ([]byte, error) {
    event := ServiceRecordEvent{
        RecordID:  recordID,
        PatientID: patientID,
        Data: map[string]interface{}{
            "old_amount": oldAmount,
            "new_amount": newAmount,
            "reason":     reason,
        },
    }
    return json.Marshal(event)
}




// User service event creators
func UserCreatedEvent(user *user_pb.GetUserResponse) ([]byte, error) {
	event := UserEvent{
		UserID: user.Id,
		Data:   user,
	}
	return json.Marshal(event)
}

func UserUpdatedEvent(user *user_pb.GetUserResponse) ([]byte, error) {
	event := UserEvent{
		UserID: user.Id,
		Data:   user,
	}
	return json.Marshal(event)
}

func UserDeletedEvent(userID string) ([]byte, error) {
	event := UserEvent{
		UserID: userID,
		Data:   nil,
	}
	return json.Marshal(event)
}

func HospitalCreatedEvent(hospital *user_pb.GetHospitalResponse) ([]byte, error) {
	event := HospitalEvent{
		HospitalID: hospital.Id,
		Data:       hospital,
	}
	return json.Marshal(event)
}

func HospitalUpdatedEvent(hospital *user_pb.GetHospitalResponse) ([]byte, error) {
	event := HospitalEvent{
		HospitalID: hospital.Id,
		Data:       hospital,
	}
	return json.Marshal(event)
}

func HospitalDeletedEvent(hospitalID string) ([]byte, error) {
	event := HospitalEvent{
		HospitalID: hospitalID,
		Data:       nil,
	}
	return json.Marshal(event)
}

// Appointment service event creators
func AppointmentCreatedEvent(appointment *appointment_pb.GetAppointmentResponse) ([]byte, error) {
	event := AppointmentEvent{
		AppointmentID: appointment.AppointmentId,
		Data:          appointment,
	}
	return json.Marshal(event)
}

func AppointmentUpdatedEvent(appointment *appointment_pb.GetAppointmentResponse) ([]byte, error) {
	event := AppointmentEvent{
		AppointmentID: appointment.AppointmentId,
		Data:          appointment,
	}
	return json.Marshal(event)
}

func AppointmentDeletedEvent(appointmentID string) ([]byte, error) {
	event := AppointmentEvent{
		AppointmentID: appointmentID,
		Data:          nil,
	}
	return json.Marshal(event)
}

func QueueCreatedEvent(queue *appointment_pb.Queue) ([]byte, error) {
	event := QueueEvent{
		QueueID: queue.QueueId,
		Data:    queue,
	}
	return json.Marshal(event)
}

func QueueStatusUpdatedEvent(queueID string, newStatus string) ([]byte, error) {
	event := QueueEvent{
		QueueID: queueID,
		Data: map[string]string{
			"status": newStatus,
		},
	}
	return json.Marshal(event)
}

// Patient service event creators
func PatientCreatedEvent(patient *patient_pb.GetPatientResponse) ([]byte, error) {
	event := PatientEvent{
		PatientID: patient.Id,
		Data:      patient,
	}
	return json.Marshal(event)
}

func PatientUpdatedEvent(patient *patient_pb.GetPatientResponse) ([]byte, error) {
	event := PatientEvent{
		PatientID: patient.Id,
		Data:      patient,
	}
	return json.Marshal(event)
}

func PatientDeletedEvent(patientID string) ([]byte, error) {
	event := PatientEvent{
		PatientID: patientID,
		Data:      nil,
	}
	return json.Marshal(event)
}

func PatientRequestedEvent(patientID string) ([]byte, error) {
	event := PatientEvent{
		PatientID: patientID,
		Data:      nil,
	}
	return json.Marshal(event)
}

// HR service event creators
func StaffRegisteredEvent(staff *hr_pb.RegisterStaffResponse) ([]byte, error) {
	event := StaffEvent{
		StaffID: staff.StaffId,
		Data:    staff,
	}
	return json.Marshal(event)
}

func PayrollCreatedEvent(payroll *hr_pb.CreatePayrollResponse) ([]byte, error) {
	event := PayrollEvent{
		PayrollID: payroll.PayrollId,
		Data:      payroll,
	}
	return json.Marshal(event)
}

func LeaveRequestedEvent(leave *hr_pb.RequestLeaveResponse) ([]byte, error) {
	event := LeaveEvent{
		LeaveID: leave.LeaveId,
		Data:    leave,
	}
	return json.Marshal(event)
}


// Medical Records event creators
func DiagnosisCreatedEvent(diagnosis *medical_records_pb.Diagnosis) ([]byte, error) {
	event := DiagnosisEvent{
		DiagnosisID: diagnosis.Id,
		Data:        diagnosis,
	}
	return json.Marshal(event)
}

func NoteCreatedEvent(note *medical_records_pb.Note) ([]byte, error) {
	event := NoteEvent{
		NoteID: note.Id,
		Data:   note,
	}
	return json.Marshal(event)
}

func VitalSignRecordedEvent(vitalSign *medical_records_pb.VitalSign) ([]byte, error) {
	event := VitalSignEvent{
		VitalSignID: vitalSign.Id,
		Data:        vitalSign,
	}
	return json.Marshal(event)
}

func DocumentUploadedEvent(document *medical_records_pb.Document) ([]byte, error) {
	event := DocumentEvent{
		DocumentID: document.Id,
		Data:       document,
	}
	return json.Marshal(event)
}


// Ward service event creators
func AdmissionCreatedEvent(admission *ward_pb.Admission) ([]byte, error) {
	event := AdmissionEvent{
		AdmissionID: fmt.Sprintf("%d", admission.AdmissionId),
		Data:        admission,
	}
	return json.Marshal(event)
}

func AdmissionUpdatedEvent(admission *ward_pb.Admission) ([]byte, error) {
	event := AdmissionEvent{
		AdmissionID: fmt.Sprintf("%d", admission.AdmissionId),
		Data:        admission,
	}
	return json.Marshal(event)
}

func AdmissionDeletedEvent(admissionID string) ([]byte, error) {
	event := AdmissionEvent{
		AdmissionID: admissionID,
		Data:        nil,
	}
	return json.Marshal(event)
}

func BedCreatedEvent(bed *ward_pb.GetBedResponse) ([]byte, error) {
	event := BedEvent{
		BedID: fmt.Sprintf("%d", bed.BedId),
		Data:  bed,
	}
	return json.Marshal(event)
}

func BedStatusUpdatedEvent(bedID string, newStatus string) ([]byte, error) {
	event := BedEvent{
		BedID: bedID,
		Data: map[string]string{
			"status": newStatus,
		},
	}
	return json.Marshal(event)
}

func BedDeletedEvent(bedID string) ([]byte, error) {
	event := BedEvent{
		BedID: bedID,
		Data:  nil,
	}
	return json.Marshal(event)
}

func WardCreatedEvent(ward *ward_pb.GetWardResponse) ([]byte, error) {
	event := WardEvent{
		WardID: fmt.Sprintf("%d", ward.WardId),
		Data:   ward,
	}
	return json.Marshal(event)
}

func WardUpdatedEvent(ward *ward_pb.GetWardResponse) ([]byte, error) {
	event := WardEvent{
		WardID: fmt.Sprintf("%d", ward.WardId),
		Data:   ward,
	}
	return json.Marshal(event)
}

func WardDeletedEvent(wardID string) ([]byte, error) {
	event := WardEvent{
		WardID: wardID,
		Data:   nil,
	}
	return json.Marshal(event)
}

func DischargeCreatedEvent(discharge *ward_pb.Discharge) ([]byte, error) {
	event := DischargeEvent{
		DischargeID: fmt.Sprintf("%d", discharge.DischargeId),
		Data:        discharge,
	}
	return json.Marshal(event)
}

func DischargeUpdatedEvent(dischargeID string, reason string) ([]byte, error) {
	event := DischargeEvent{
		DischargeID: dischargeID,
		Data: map[string]string{
			"reason": reason,
		},
	}
	return json.Marshal(event)
}

func PatientTransferredEvent(transfer *ward_pb.Transfer) ([]byte, error) {
	event := TransferEvent{
		TransferID: fmt.Sprintf("%d", transfer.TransferId),
		Data:       transfer,
	}
	return json.Marshal(event)
}

// Supply Chain event creators
func StockItemCreatedEvent(item *supply_chain_pb.StockItem) ([]byte, error) {
    event := StockItemEvent{
        ItemID: item.Id,
        Data:   item,
    }
    return json.Marshal(event)
}

func StockItemUpdatedEvent(item *supply_chain_pb.StockItem) ([]byte, error) {
    event := StockItemEvent{
        ItemID: item.Id,
        Data:   item,
    }
    return json.Marshal(event)
}

func StockItemDeletedEvent(itemID string) ([]byte, error) {
    event := StockItemEvent{
        ItemID: itemID,
        Data:   nil,
    }
    return json.Marshal(event)
}

func RequisitionCreatedEvent(requisition *supply_chain_pb.Requisition) ([]byte, error) {
    event := RequisitionEvent{
        RequisitionID: requisition.Id,
        Data:          requisition,
    }
    return json.Marshal(event)
}

func RequisitionApprovedEvent(requisition *supply_chain_pb.Requisition) ([]byte, error) {
    event := RequisitionEvent{
        RequisitionID: requisition.Id,
        Data:          requisition,
    }
    return json.Marshal(event)
}

func LPOCreatedEvent(lpo *supply_chain_pb.LocalPurchaseOrder) ([]byte, error) {
    event := LPOEvent{
        LPOID: lpo.Id,
        Data:  lpo,
    }
    return json.Marshal(event)
}

func LPOApprovedEvent(lpo *supply_chain_pb.LocalPurchaseOrder) ([]byte, error) {
    event := LPOEvent{
        LPOID: lpo.Id,
        Data:  lpo,
    }
    return json.Marshal(event)
}

func GRNCreatedEvent(grn *supply_chain_pb.GoodsReceivedNote) ([]byte, error) {
    event := GRNEvent{
        GRNID: grn.Id,
        Data:  grn,
    }
    return json.Marshal(event)
}

func GRNQualityCheckedEvent(grn *supply_chain_pb.GoodsReceivedNote) ([]byte, error) {
    event := GRNEvent{
        GRNID: grn.Id,
        Data:  grn,
    }
    return json.Marshal(event)
}

func StockAdjustmentCreatedEvent(adjustment *supply_chain_pb.StockAdjustment) ([]byte, error) {
    event := StockAdjustmentEvent{
        AdjustmentID: adjustment.Id,
        Data:         adjustment,
    }
    return json.Marshal(event)
}

func StockAdjustmentApprovedEvent(adjustment *supply_chain_pb.StockAdjustment) ([]byte, error) {
    event := StockAdjustmentEvent{
        AdjustmentID: adjustment.Id,
        Data:         adjustment,
    }
    return json.Marshal(event)
}

func TransferCreatedEvent(transfer *supply_chain_pb.InterStoreTransfer) ([]byte, error) {
    event := InterStoreTransferEvent{
        TransferID: transfer.Id,
        Data:       transfer,
    }
    return json.Marshal(event)
}

func TransferApprovedEvent(transfer *supply_chain_pb.InterStoreTransfer) ([]byte, error) {
    event := InterStoreTransferEvent{
        TransferID: transfer.Id,
        Data:       transfer,
    }
    return json.Marshal(event)
}

func TransferReceivedEvent(transfer *supply_chain_pb.InterStoreTransfer) ([]byte, error) {
    event := InterStoreTransferEvent{
        TransferID: transfer.Id,
        Data:       transfer,
    }
    return json.Marshal(event)
}

func CreditNoteCreatedEvent(creditNote *supply_chain_pb.CreditNote) ([]byte, error) {
    event := CreditNoteEvent{
        CreditNoteID: creditNote.Id,
        Data:         creditNote,
    }
    return json.Marshal(event)
}

func CreditNoteApprovedEvent(creditNote *supply_chain_pb.CreditNote) ([]byte, error) {
    event := CreditNoteEvent{
        CreditNoteID: creditNote.Id,
        Data:         creditNote,
    }
    return json.Marshal(event)
}


// Medication event creators
func MedicationCreatedEvent(medication *pharmacy_pb.GetMedicationResponse) ([]byte, error) {
	event := MedicationEvent{
		MedicationID: fmt.Sprintf("%d", medication.MedicationId),
		Data:         medication,
	}
	return json.Marshal(event)
}

func MedicationUpdatedEvent(medication *pharmacy_pb.GetMedicationResponse) ([]byte, error) {
	event := MedicationEvent{
		MedicationID: fmt.Sprintf("%d", medication.MedicationId),
		Data:         medication,
	}
	return json.Marshal(event)
}

func MedicationDeletedEvent(medicationID string) ([]byte, error) {
	event := MedicationEvent{
		MedicationID: medicationID,
		Data:         nil,
	}
	return json.Marshal(event)
}

// Prescription event creators
func PrescriptionCreatedEvent(prescription *pharmacy_pb.GetSinglePrescriptionRequestResponse) ([]byte, error) {
	event := PrescriptionEvent{
		PrescriptionID: fmt.Sprintf("%d", prescription.PrescriptionRequest.PrescriptionId),
		Data:           prescription,
	}
	return json.Marshal(event)
}

func PrescriptionUpdatedEvent(prescription *pharmacy_pb.GetSinglePrescriptionRequestResponse) ([]byte, error) {
	event := PrescriptionEvent{
		PrescriptionID: fmt.Sprintf("%d", prescription.PrescriptionRequest.PrescriptionId),
		Data:           prescription,
	}
	return json.Marshal(event)
}

func PrescriptionDeletedEvent(prescriptionID string) ([]byte, error) {
	event := PrescriptionEvent{
		PrescriptionID: prescriptionID,
		Data:           nil,
	}
	return json.Marshal(event)
}

// Stock Request event creators
func StockRequestCreatedEvent(request *pharmacy_pb.GetStockRequestResponse) ([]byte, error) {
	event := StockRequestEvent{
		RequestID: fmt.Sprintf("%d", request.RequestId),
		Data:      request,
	}
	return json.Marshal(event)
}

func StockRequestApprovedEvent(request *pharmacy_pb.GetStockRequestResponse) ([]byte, error) {
	event := StockRequestEvent{
		RequestID: fmt.Sprintf("%d", request.RequestId),
		Data:      request,
	}
	return json.Marshal(event)
}

// Stock Movement event creators
func StockMovementCreatedEvent(movement *pharmacy_pb.GetStockMovementResponse) ([]byte, error) {
	event := StockMovementEvent{
		MovementID: fmt.Sprintf("%d", movement.StockMovement.MovementId),
		Data:       movement,
	}
	return json.Marshal(event)
}

// Dispensed Medication event creators
func MedicationDispensedEvent(dispensed *pharmacy_pb.GetDispensedMedicationResponse, items []*DispensedMedicationItem) ([]byte, error) {
	event := DispensedMedicationEvent{
		DispenseID: fmt.Sprintf("%d", dispensed.DispenseId),
		Data:       dispensed,
		Items:      items,
	}
	return json.Marshal(event)
}

// Medication Return event creators
func MedicationReturnedEvent(returnRecord *pharmacy_pb.GetMedicationReturnResponse) ([]byte, error) {
	event := MedicationReturnEvent{
		ReturnID: fmt.Sprintf("%d", returnRecord.MedicationReturn.ReturnId),
		Data:     returnRecord,
	}
	return json.Marshal(event)
}

// Medication Schedule event creators
func MedicationScheduleCreatedEvent(schedule *pharmacy_pb.GetMedicationScheduleResponse) ([]byte, error) {
	event := MedicationScheduleEvent{
		ScheduleID: fmt.Sprintf("%d", schedule.Schedule.ScheduleId),
		Data:       schedule,
	}
	return json.Marshal(event)
}

// Medication Intake event creators
func MedicationIntakeRecordedEvent(intake *pharmacy_pb.GetMedicationIntakeResponse) ([]byte, error) {
	event := MedicationIntakeEvent{
		IntakeID: fmt.Sprintf("%d", intake.Intake.IntakeId),
		Data:     intake,
	}
	return json.Marshal(event)
}


// Laboratory service event creators

func LabRequestCreatedEvent(request *laboratory_pb.RequestLabRequest) ([]byte, error) {
	event := LabEvent{
		PatientID: request.GetPatientInfo().GetPatientId(),
		Data:      request,
	}
	return json.Marshal(event)
}

func TestResultAddedEvent(result *laboratory_pb.AddTestResultRequest) ([]byte, error) {
	event := LabEvent{
		PatientID: result.GetPatientInfo().GetPatientId(),
		Data:      result,
	}
	return json.Marshal(event)
}


func IssueReportCreatedEvent(report *issue_pb.AddReportRequest, reportID int32) ([]byte, error) {
	event := IssueEvent{
		ReportID: reportID,
		Data:     report,
	}
	return json.Marshal(event)
}

func IssueReportStatusUpdatedEvent(statusUpdate *issue_pb.UpdateReportStatusRequest) ([]byte, error) {
	event := IssueEvent{
		ReportID: statusUpdate.GetId(),
		Data:     statusUpdate,
	}
	return json.Marshal(event)
}




// Event handlers
func HandleAppointmentCreated(data []byte) error {
	var event AppointmentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal appointment created event: %v", err)
	}
	log.Printf("Appointment created: %s", event.AppointmentID)
	// Process the event
	return nil
}

func HandleAppointmentUpdated(data []byte) error {
	var event AppointmentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal appointment updated event: %v", err)
	}
	log.Printf("Appointment updated: %s", event.AppointmentID)
	// Process the event
	return nil
}

func HandleAppointmentDeleted(data []byte) error {
	var event AppointmentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal appointment deleted event: %v", err)
	}
	log.Printf("Appointment deleted: %s", event.AppointmentID)
	// Process the event
	return nil
}

func HandleAppointmentCancelled(data []byte) error {
	var event AppointmentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal appointment cancelled event: %v", err)
	}
	log.Printf("Appointment cancelled: %s", event.AppointmentID)
	// Here, you would add logic to find the corresponding service_record
	// and update its status to 'VOID' or 'CANCELLED'.
	// e.g., serviceRecordRepo.VoidByReference(ctx, event.AppointmentID, "consultation")
	return nil
}

func HandleAppointmentNoShow(data []byte) error {
	var event AppointmentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal appointment no-show event: %v", err)
	}
	log.Printf("Appointment no-show: %s", event.AppointmentID)
	// Logic similar to cancellation, marking the service_record as void.
	// e.g., serviceRecordRepo.VoidByReference(ctx, event.AppointmentID, "consultation")
	return nil
}

func HandleQueueCreated(data []byte) error {
	var event QueueEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal queue created event: %v", err)
	}
	log.Printf("Queue created: %s", event.QueueID)
	// Process the event
	return nil
}

func HandleQueueStatusUpdated(data []byte) error {
	var event QueueEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal queue status updated event: %v", err)
	}
	log.Printf("Queue status updated: %s", event.QueueID)
	// Process the event
	return nil
}

func HandleQueueCancelled(data []byte) error {
	var event QueueEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal queue cancelled event: %v", err)
	}
	log.Printf("Queue cancelled: %s", event.QueueID)
	return nil
}

func HandlePatientCreated(data []byte) error {
	var event PatientEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal patient created event: %v", err)
	}
	log.Printf("Patient created: %s", event.PatientID)
	// Process the event
	return nil
}

func HandlePatientUpdated(data []byte) error {
	var event PatientEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal patient updated event: %v", err)
	}
	log.Printf("Patient updated: %s", event.PatientID)
	// Process the event
	return nil
}

func HandlePatientDeleted(data []byte) error {
	var event PatientEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal patient deleted event: %v", err)
	}
	log.Printf("Patient deleted: %s", event.PatientID)
	// Process the event
	return nil
}

func HandlePatientRequested(data []byte) error {
	var event PatientEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal patient requested event: %v", err)
	}
	log.Printf("Patient requested: %s", event.PatientID)
	// Process the event
	return nil
}

func HandleStaffRegistered(data []byte) error {
	var event StaffEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal staff registered event: %v", err)
	}
	log.Printf("Staff registered: %d", event.StaffID)
	// Process the event
	return nil
}

func HandleStaffUpdated(data []byte) error {
	var event StaffEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal staff updated event: %v", err)
	}
	log.Printf("Staff updated: %d", event.StaffID)
	// Process the event
	return nil
}

func HandleStaffDeleted(data []byte) error {
	var event StaffEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal staff deleted event: %v", err)
	}
	log.Printf("Staff deleted: %d", event.StaffID)
	// Process the event
	return nil
}

func HandlePayrollCreated(data []byte) error {
	var event PayrollEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal payroll created event: %v", err)
	}
	log.Printf("Payroll created: %d", event.PayrollID)
	// Process the event
	return nil
}

func HandleLeaveRequested(data []byte) error {
	var event LeaveEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal leave requested event: %v", err)
	}
	log.Printf("Leave requested: %d", event.LeaveID)
	// Process the event
	return nil
}

func HandleLeaveApproved(data []byte) error {
	var event LeaveEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal leave approved event: %v", err)
	}
	log.Printf("Leave approved: %d", event.LeaveID)
	// Process the event
	return nil
}

func HandleDepartmentCreated(data []byte) error {
	var event DepartmentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal department created event: %v", err)
	}
	log.Printf("Department created: %d", event.DepartmentID)
	// Process the event
	return nil
}

func HandleAttendanceRecorded(data []byte) error {
	var event AttendanceEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal attendance recorded event: %v", err)
	}
	log.Printf("Attendance recorded: %d", event.AttendanceID)
	// Process the event
	return nil
}




// Medical Records event handlers
func HandleDiagnosisCreated(data []byte) error {
	var event DiagnosisEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal diagnosis created event: %v", err)
	}
	log.Printf("Diagnosis created: %s", event.DiagnosisID)
	return nil
}

func HandleDiagnosisUpdated(data []byte) error {
	var event DiagnosisEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal diagnosis updated event: %v", err)
	}
	log.Printf("Diagnosis updated: %s", event.DiagnosisID)
	return nil
}

func HandleDiagnosisDeleted(data []byte) error {
	var event DiagnosisEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal diagnosis deleted event: %v", err)
	}
	log.Printf("Diagnosis deleted: %s", event.DiagnosisID)
	return nil
}

func HandleNoteCreated(data []byte) error {
	var event NoteEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal note created event: %v", err)
	}
	log.Printf("Note created: %s", event.NoteID)
	return nil
}

func HandleNoteUpdated(data []byte) error {
	var event NoteEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal note updated event: %v", err)
	}
	log.Printf("Note updated: %s", event.NoteID)
	return nil
}

func HandleNoteDeleted(data []byte) error {
	var event NoteEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal note deleted event: %v", err)
	}
	log.Printf("Note deleted: %s", event.NoteID)
	return nil
}

func HandleVitalSignRecorded(data []byte) error {
	var event VitalSignEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal vital sign recorded event: %v", err)
	}
	log.Printf("Vital sign recorded: %s", event.VitalSignID)
	return nil
}

func HandleDocumentUploaded(data []byte) error {
	var event DocumentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal document uploaded event: %v", err)
	}
	log.Printf("Document uploaded: %s", event.DocumentID)
	return nil
}

func HandleDocumentUpdated(data []byte) error {
	var event DocumentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal document updated event: %v", err)
	}
	log.Printf("Document updated: %s", event.DocumentID)
	return nil
}

func HandleDocumentDeleted(data []byte) error {
	var event DocumentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal document deleted event: %v", err)
	}
	log.Printf("Document deleted: %s", event.DocumentID)
	return nil
}


// Ward service event handlers
func HandleAdmissionCreated(data []byte) error {
	var event AdmissionEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal admission created event: %v", err)
	}
	log.Printf("Admission created: %s", event.AdmissionID)
	// Process the event
	return nil
}

func HandleAdmissionUpdated(data []byte) error {
	var event AdmissionEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal admission updated event: %v", err)
	}
	log.Printf("Admission updated: %s", event.AdmissionID)
	// Process the event
	return nil
}

func HandleAdmissionDeleted(data []byte) error {
	var event AdmissionEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal admission deleted event: %v", err)
	}
	log.Printf("Admission deleted: %s", event.AdmissionID)
	// Process the event
	return nil
}

func HandleBedCreated(data []byte) error {
	var event BedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal bed created event: %v", err)
	}
	log.Printf("Bed created: %s", event.BedID)
	// Process the event
	return nil
}

func HandleBedStatusUpdated(data []byte) error {
	var event BedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal bed status updated event: %v", err)
	}
	log.Printf("Bed status updated: %s", event.BedID)
	// Process the event
	return nil
}

func HandleBedDeleted(data []byte) error {
	var event BedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal bed deleted event: %v", err)
	}
	log.Printf("Bed deleted: %s", event.BedID)
	// Process the event
	return nil
}

func HandleWardCreated(data []byte) error {
	var event WardEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal ward created event: %v", err)
	}
	log.Printf("Ward created: %s", event.WardID)
	// Process the event
	return nil
}

func HandleWardUpdated(data []byte) error {
	var event WardEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal ward updated event: %v", err)
	}
	log.Printf("Ward updated: %s", event.WardID)
	// Process the event
	return nil
}

func HandleWardDeleted(data []byte) error {
	var event WardEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal ward deleted event: %v", err)
	}
	log.Printf("Ward deleted: %s", event.WardID)
	// Process the event
	return nil
}

func HandleDischargeCreated(data []byte) error {
	var event DischargeEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal discharge created event: %v", err)
	}
	log.Printf("Discharge created: %s", event.DischargeID)
	// Process the event
	return nil
}

func HandleDischargeUpdated(data []byte) error {
	var event DischargeEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal discharge updated event: %v", err)
	}
	log.Printf("Discharge updated: %s", event.DischargeID)
	// Process the event
	return nil
}

func HandlePatientTransferred(data []byte) error {
	var event TransferEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal patient transferred event: %v", err)
	}
	log.Printf("Patient transferred: %s", event.TransferID)
	// Process the event
	return nil
}


// Supply Chain event handlers
func HandleStockItemCreated(data []byte) error {
    var event StockItemEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal stock item created event: %v", err)
    }
    log.Printf("Stock item created: %s", event.ItemID)
    // Process the event (e.g., update search index, send notifications)
    return nil
}

func HandleStockItemUpdated(data []byte) error {
    var event StockItemEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal stock item updated event: %v", err)
    }
    log.Printf("Stock item updated: %s", event.ItemID)
    // Process the event
    return nil
}

func HandleStockItemDeleted(data []byte) error {
    var event StockItemEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal stock item deleted event: %v", err)
    }
    log.Printf("Stock item deleted: %s", event.ItemID)
    // Process the event
    return nil
}

func HandleRequisitionCreated(data []byte) error {
    var event RequisitionEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal requisition created event: %v", err)
    }
    log.Printf("Requisition created: %s", event.RequisitionID)
    // Process the event (e.g., notify approvers)
    return nil
}

func HandleRequisitionApproved(data []byte) error {
    var event RequisitionEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal requisition approved event: %v", err)
    }
    log.Printf("Requisition approved: %s", event.RequisitionID)
    // Process the event (e.g., trigger LPO creation)
    return nil
}

func HandleLPOCreated(data []byte) error {
    var event LPOEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal LPO created event: %v", err)
    }
    log.Printf("LPO created: %s", event.LPOID)
    // Process the event (e.g., notify suppliers)
    return nil
}

func HandleLPOApproved(data []byte) error {
    var event LPOEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal LPO approved event: %v", err)
    }
    log.Printf("LPO approved: %s", event.LPOID)
    // Process the event (e.g., trigger procurement process)
    return nil
}

func HandleGRNCreated(data []byte) error {
    var event GRNEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal GRN created event: %v", err)
    }
    log.Printf("GRN created: %s", event.GRNID)
    // Process the event (e.g., schedule quality check)
    return nil
}

func HandleGRNQualityChecked(data []byte) error {
    var event GRNEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal GRN quality checked event: %v", err)
    }
    log.Printf("GRN quality checked: %s", event.GRNID)
    // Process the event (e.g., update stock levels)
    return nil
}

func HandleStockAdjustmentCreated(data []byte) error {
    var event StockAdjustmentEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal stock adjustment created event: %v", err)
    }
    log.Printf("Stock adjustment created: %s", event.AdjustmentID)
    // Process the event (e.g., notify approvers)
    return nil
}

func HandleStockAdjustmentApproved(data []byte) error {
    var event StockAdjustmentEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal stock adjustment approved event: %v", err)
    }
    log.Printf("Stock adjustment approved: %s", event.AdjustmentID)
    // Process the event (e.g., update stock levels)
    return nil
}

func HandleTransferCreated(data []byte) error {
    var event TransferEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal transfer created event: %v", err)
    }
    log.Printf("Transfer created: %s", event.TransferID)
    // Process the event (e.g., notify destination store)
    return nil
}

func HandleTransferApproved(data []byte) error {
    var event TransferEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal transfer approved event: %v", err)
    }
    log.Printf("Transfer approved: %s", event.TransferID)
    // Process the event (e.g., prepare items for transfer)
    return nil
}

func HandleTransferReceived(data []byte) error {
    var event TransferEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal transfer received event: %v", err)
    }
    log.Printf("Transfer received: %s", event.TransferID)
    // Process the event (e.g., update stock levels)
    return nil
}

func HandleCreditNoteCreated(data []byte) error {
    var event CreditNoteEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal credit note created event: %v", err)
    }
    log.Printf("Credit note created: %s", event.CreditNoteID)
    // Process the event (e.g., notify accounts)
    return nil
}

func HandleCreditNoteApproved(data []byte) error {
    var event CreditNoteEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal credit note approved event: %v", err)
    }
    log.Printf("Credit note approved: %s", event.CreditNoteID)
    // Process the event (e.g., update supplier account)
    return nil
}


// Pharmacy event handlers
func HandleMedicationCreated(data []byte) error {
	var event MedicationEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal medication created event: %v", err)
	}
	log.Printf("Medication created: %s", event.MedicationID)
	// Add business logic here (e.g., update search index, send notifications)
	return nil
}

func HandleMedicationUpdated(data []byte) error {
	var event MedicationEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal medication updated event: %v", err)
	}
	log.Printf("Medication updated: %s", event.MedicationID)
	// Add business logic here
	return nil
}

func HandleMedicationDeleted(data []byte) error {
	var event MedicationEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal medication deleted event: %v", err)
	}
	log.Printf("Medication deleted: %s", event.MedicationID)
	// Add business logic here
	return nil
}

func HandlePrescriptionCreated(data []byte) error {
	var event PrescriptionEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal prescription created event: %v", err)
	}
	log.Printf("Prescription created: %s", event.PrescriptionID)
	// Add business logic here (e.g., notify pharmacy staff)
	return nil
}

func HandlePrescriptionUpdated(data []byte) error {
	var event PrescriptionEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal prescription updated event: %v", err)
	}
	log.Printf("Prescription updated: %s", event.PrescriptionID)
	// Add business logic here
	return nil
}

func HandleStockRequestCreated(data []byte) error {
	var event StockRequestEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal stock request created event: %v", err)
	}
	log.Printf("Stock request created: %s", event.RequestID)
	// Add business logic here (e.g., notify approvers)
	return nil
}

func HandleStockRequestApproved(data []byte) error {
	var event StockRequestEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal stock request approved event: %v", err)
	}
	log.Printf("Stock request approved: %s", event.RequestID)
	// Add business logic here (e.g., trigger procurement process)
	return nil
}

func HandleMedicationDispensed(data []byte) error {
	var event DispensedMedicationEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal medication dispensed event: %v", err)
	}
	log.Printf("Medication dispensed: %s", event.DispenseID)
	// Add business logic here (e.g., update patient records)
	return nil
}

func HandleMedicationReturned(data []byte) error {
	var event MedicationReturnEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal medication returned event: %v", err)
	}
	log.Printf("Medication returned: %s", event.ReturnID)
	// Add business logic here (e.g., update stock levels)
	return nil
}

func HandleMedicationScheduleCreated(data []byte) error {
	var event MedicationScheduleEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal medication schedule created event: %v", err)
	}
	log.Printf("Medication schedule created: %s", event.ScheduleID)
	// Add business logic here (e.g., schedule reminders)
	return nil
}

func HandleMedicationIntakeRecorded(data []byte) error {
	var event MedicationIntakeEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal medication intake recorded event: %v", err)
	}
	log.Printf("Medication intake recorded: %s", event.IntakeID)
	// Add business logic here (e.g., update adherence metrics)
	return nil
}



// Laboratory service event handlers
func HandleLabRequestCreated(data []byte) error {
	var event LabEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal lab request created event: %v", err)
	}
	log.Printf("Lab request created for patient: %s", event.PatientID)
	// Process the event
	return nil
}

func HandleTestResultAdded(data []byte) error {
	var event LabEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal test result added event: %v", err)
	}
	log.Printf("Test result added for patient: %s", event.PatientID)
	// Process the event
	return nil
}

// Issue Reporting event handlers
func HandleIssueReportCreated(data []byte) error {
	var event IssueEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal issue report created event: %v", err)
	}
	log.Printf("Issue report created: %d", event.ReportID)
	// Process the event
	return nil
}

func HandleIssueReportStatusUpdated(data []byte) error {
	var event IssueEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal issue report status updated event: %v", err)
	}
	log.Printf("Issue report status updated: %d", event.ReportID)
	// Process the event
	return nil
}


// User service event handlers
func HandleUserCreated(data []byte) error {
	var event UserEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal user created event: %v", err)
	}
	log.Printf("User created: %s", event.UserID)
	// Process the event (e.g., update search index, send notifications)
	return nil
}

func HandleUserUpdated(data []byte) error {
	var event UserEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal user updated event: %v", err)
	}
	log.Printf("User updated: %s", event.UserID)
	// Process the event
	return nil
}

func HandleUserDeleted(data []byte) error {
	var event UserEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal user deleted event: %v", err)
	}
	log.Printf("User deleted: %s", event.UserID)
	// Process the event
	return nil
}

func HandleHospitalCreated(data []byte) error {
	var event HospitalEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal hospital created event: %v", err)
	}
	log.Printf("Hospital created: %s", event.HospitalID)
	// Process the event
	return nil
}

func HandleHospitalUpdated(data []byte) error {
	var event HospitalEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal hospital updated event: %v", err)
	}
	log.Printf("Hospital updated: %s", event.HospitalID)
	// Process the event
	return nil
}

func HandleHospitalDeleted(data []byte) error {
	var event HospitalEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal hospital deleted event: %v", err)
	}
	log.Printf("Hospital deleted: %s", event.HospitalID)
	// Process the event
	return nil
}



// Billing event handlers
func HandleInvoiceCreated(data []byte) error {
	var event InvoiceEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal invoice created event: %v", err)
	}
	log.Printf("Invoice created: %s", event.InvoiceID)
	// Process the event (e.g., send notification)
	return nil
}

func HandleInvoiceUpdated(data []byte) error {
	var event InvoiceEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal invoice updated event: %v", err)
	}
	log.Printf("Invoice updated: %s", event.InvoiceID)
	// Process the event
	return nil
}

func HandleInvoiceDeleted(data []byte) error {
	var event InvoiceEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal invoice deleted event: %v", err)
	}
	log.Printf("Invoice deleted: %s", event.InvoiceID)
	// Process the event
	return nil
}

func HandlePaymentProcessed(data []byte) error {
	var event PaymentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal payment processed event: %v", err)
	}
	log.Printf("Payment processed: %s", event.PaymentID)
	// Process the event (e.g., update accounting system)
	return nil
}


// Add these event handlers
func HandleServiceRecordCreated(data []byte) error {
    var event ServiceRecordEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal service record created event: %v", err)
    }
    log.Printf("Service record created: %s for patient %s", event.RecordID, event.PatientID)
    return nil
}

func HandleServiceRecordStatusUpdated(data []byte) error {
    var event ServiceRecordEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal service record status updated event: %v", err)
    }
    log.Printf("Service record status updated: %s", event.RecordID)
    return nil
}



func HandlePermissionAdded(data []byte) error {
    var event PermissionEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal permission added event: %v", err)
    }
    log.Printf("Permission added for route: %s", event.Route)
    // Process the event (e.g., update cache)
    return nil
}

func HandlePermissionUpdated(data []byte) error {
    var event PermissionEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return fmt.Errorf("failed to unmarshal permission updated event: %v", err)
    }
    log.Printf("Permission updated for route: %s", event.Route)
    // Process the event (e.g., update cache)
    return nil
}