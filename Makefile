# Directories
PROTO_DIR := ./proto
GOOGLEAPIS_DIR := ./third_party/googleapis
THIRD_PARTY_DIR := ./third_party
OUTPUT_BASE := ./shared

# Proto files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Protoc command template
define PROTOC_CMD
protoc -I . -I $(GOOGLEAPIS_DIR) -I $(THIRD_PARTY_DIR) \
    --go_out=$(OUTPUT_BASE)/$(1)/ \
    --go_opt=paths=source_relative \
    --go-grpc_out=$(OUTPUT_BASE)/$(1)/ \
    --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=$(OUTPUT_BASE)/$(1)/ \
    --grpc-gateway_opt=paths=source_relative \
    $(PROTO_DIR)/$(2).proto
endef

# Targets
all: $(patsubst $(PROTO_DIR)/%.proto,%,$(PROTO_FILES))

# Individual proto targets
# appointment_service:
# 	$(call PROTOC_CMD,appointment-service,appointment_service)

# billing_service:
# 	$(call PROTOC_CMD,billing-service,billing_service)

# chat_service:
# 	$(call PROTOC_CMD,chat-service,chat_service)

# human_resource_service:
# 	$(call PROTOC_CMD,human-resource-service,human_resource_service)

# issue_report_service:
# 	$(call PROTOC_CMD,issue-report-service,issue_report_service)

# laboratory_service:
# 	$(call PROTOC_CMD,laboratory-service,laboratory_service)

# medical_records_service:
# 	$(call PROTOC_CMD,medical-records-service,medical_records_service)

# patient_service:
# 	$(call PROTOC_CMD,patient-service,patient_service)

# route_permissions_service:
# 	$(call PROTOC_CMD,route-permissions-service,route_permissions_service)

# pharmacy_service:
# 	$(call PROTOC_CMD,pharmacy-service,pharmacy_service)

# supply_chain_service:
# 	$(call PROTOC_CMD,supply-chain-service,supply_chain_service)

user_service:
	$(call PROTOC_CMD,user-service,user_service)

# ward_service:
# 	$(call PROTOC_CMD,ward-service,ward_service)

# analytics_service:
# 	$(call PROTOC_CMD,analytics-service,analytics_service)

# encounter_service:
# 	$(call PROTOC_CMD,encounter-service,encounter_service)

.PHONY: all $(patsubst $(PROTO_DIR)/%.proto,%,$(PROTO_FILES))