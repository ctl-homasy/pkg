// pkg/middleware/validator.go

package middleware

import (
	"context"
	"regexp"

	appointment_pb "github.com/homasy/pkg/shared/appointment-service/proto"
	patient_pb "github.com/homasy/pkg/shared/patient-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidationInterceptor creates a gRPC interceptor for request validation
func ValidationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validateRequest(req); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func validateRequest(req interface{}) error {
	switch r := req.(type) {
	// Appointment service requests
	case *appointment_pb.CreateAppointmentRequest:
		return validateCreateAppointmentRequest(r)
	case *appointment_pb.UpdateAppointmentRequest:
		return validateUpdateAppointmentRequest(r)
	case *appointment_pb.CreateQueueRequest:
		return validateCreateQueueRequest(r)
	case *appointment_pb.UpdateQueueStatusRequest:
		return validateUpdateQueueStatusRequest(r)
	
	// Patient service requests
	case *patient_pb.AddPatientRequest:
		return validateAddPatientRequest(r)
	case *patient_pb.UpdatePatientRequest:
		return validateUpdatePatientRequest(r)
	}
	return nil
}

// Appointment service validators
func validateCreateAppointmentRequest(req *appointment_pb.CreateAppointmentRequest) error {
	// Validate required fields
	if req.GetPatientId() == "" {
		return status.Error(codes.InvalidArgument, "patient ID is required")
	}
	if req.GetDoctorId() == "" {
		return status.Error(codes.InvalidArgument, "doctor ID is required")
	}
	if req.GetPreferredDate() == "" {
		return status.Error(codes.InvalidArgument, "preferred date is required")
	}
	if req.GetPreferredTime() == "" {
		return status.Error(codes.InvalidArgument, "preferred time is required")
	}
	if req.GetAppointmentType() == "" {
		return status.Error(codes.InvalidArgument, "appointment type is required")
	}
	if req.GetSpecialistCategory() == "" {
		return status.Error(codes.InvalidArgument, "specialist category is required")
	}
	if req.GetStatus() == "" {
		return status.Error(codes.InvalidArgument, "status is required")
	}

	return nil
}

func validateUpdateAppointmentRequest(req *appointment_pb.UpdateAppointmentRequest) error {
	// Validate ID
	if req.GetAppointmentId() == "" {
		return status.Error(codes.InvalidArgument, "appointment ID is required")
	}

	// Other validations similar to create
	if req.GetPatientId() == "" {
		return status.Error(codes.InvalidArgument, "patient ID is required")
	}
	if req.GetDoctorId() == "" {
		return status.Error(codes.InvalidArgument, "doctor ID is required")
	}

	return nil
}

func validateCreateQueueRequest(req *appointment_pb.CreateQueueRequest) error {
	// Validate required fields
	if req.GetPatientId() == "" {
		return status.Error(codes.InvalidArgument, "patient ID is required")
	}
	if req.GetDoctorId() == "" {
		return status.Error(codes.InvalidArgument, "doctor ID is required")
	}
	if req.GetArrivalDate() == "" {
		return status.Error(codes.InvalidArgument, "arrival date is required")
	}
	if req.GetArrivalTime() == "" {
		return status.Error(codes.InvalidArgument, "arrival time is required")
	}
	if req.GetStatus() == "" {
		return status.Error(codes.InvalidArgument, "status is required")
	}

	return nil
}

func validateUpdateQueueStatusRequest(req *appointment_pb.UpdateQueueStatusRequest) error {
	// Validate ID
	if req.GetQueueId() == "" {
		return status.Error(codes.InvalidArgument, "queue ID is required")
	}

	// Validate status
	if req.GetNewStatus() == "" {
		return status.Error(codes.InvalidArgument, "new status is required")
	}

	// Check if status is valid
	validStatuses := map[string]bool{
		"pending":   true,
		"completed": true,
	}

	if !validStatuses[req.GetNewStatus()] {
		return status.Error(codes.InvalidArgument, "invalid status: must be 'pending' or 'completed'")
	}

	return nil
}

// Patient service validators
func validateAddPatientRequest(req *patient_pb.AddPatientRequest) error {
	// Validate required fields
	if req.GetFirstName() == "" {
		return status.Error(codes.InvalidArgument, "first name is required")
	}
	if req.GetLastName() == "" {
		return status.Error(codes.InvalidArgument, "last name is required")
	}
	if req.GetDateOfBirth() == "" {
		return status.Error(codes.InvalidArgument, "date of birth is required")
	}
	if req.GetGender() == "" {
		return status.Error(codes.InvalidArgument, "gender is required")
	}
	if req.GetPhoneNumber() == "" {
		return status.Error(codes.InvalidArgument, "phone number is required")
	}
	if req.GetAddress() == "" {
		return status.Error(codes.InvalidArgument, "address is required")
	}

	// Validate email format if provided
	if req.GetEmail() != "" {
		if !isValidEmail(req.GetEmail()) {
			return status.Error(codes.InvalidArgument, "invalid email format")
		}
	}

	// Validate phone number format
	if !isValidPhoneNumber(req.GetPhoneNumber()) {
		return status.Error(codes.InvalidArgument, "invalid phone number format")
	}

	return nil
}

func validateUpdatePatientRequest(req *patient_pb.UpdatePatientRequest) error {
	// Validate ID
	if req.GetId() == "" {
		return status.Error(codes.InvalidArgument, "patient ID is required")
	}

	// Validate email format if provided
	if req.GetEmail() != "" {
		if !isValidEmail(req.GetEmail()) {
			return status.Error(codes.InvalidArgument, "invalid email format")
		}
	}

	// Validate phone number format if provided
	if req.GetPhoneNumber() != "" {
		if !isValidPhoneNumber(req.GetPhoneNumber()) {
			return status.Error(codes.InvalidArgument, "invalid phone number format")
		}
	}

	return nil
}

// Shared validation functions
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func isValidPhoneNumber(phone string) bool {
	pattern := `^\+?[0-9]{10,15}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}