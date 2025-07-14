# SupplyChainGuard Notification Contracts (`scg-notification-contracts`)

[![CI](https://github.com/hbttundar/scg-notification-contracts/actions/workflows/CI.yml/badge.svg)](https://github.com/hbttundar/scg-notification-contracts/actions/workflows/CI.yml)

This repository contains the **canonical, versioned Protobuf contracts** for all notifications (user-facing and system-facing) in the SupplyChainGuard enterprise supply chain SaaS platform.

## ✨ Why use this library?

- **Notification standardization:** Provides a consistent structure for all notifications across the platform.
- **Multi-channel support:** Designed for emails, SMS, push notifications, webhooks, dashboard alerts, in-app banners, escalations, digests, and more.
- **Business context preservation:** Each notification includes references to the triggering business event and entity.
- **User preference awareness:** Supports user-defined notification preferences and delivery channels.
- **Comprehensive coverage:** Includes notifications for all major supply chain workflows (onboarding, compliance, logistics, digital twins, etc.).
- **Plug-and-play for Go:** `option go_package` is set for all `.proto` files, enabling easy Go codegen and direct microservice import.

## 📂 Project Structure

```
scg-notification-contracts/
├── .github/                            # GitHub configuration
│   ├── workflows/                      # GitHub Actions workflows
│   │   ├── ci.yml                      # CI workflow for linting and breaking changes
│   │   └── generate.yml                # Workflow for generating Go code
│   ├── ISSUE_TEMPLATE/                 # Issue templates
│   │   ├── bug_report.md               # Bug report template
│   │   └── feature_request.md          # Feature request template
│   └── pull_request_template.md        # Pull request template
├── proto/                              # Protobuf definitions
│   ├── v1/                             # Version 1 of the API
│   │   ├── enums/
│   │   │   └── notification_enums.proto    # Enums for notification types, channels, severity, etc.
│   │   ├── models/
│   │   │   └── notification_models.proto   # Core notification data models
│   │   └── events/
│   │       └── notification_events.proto   # Notification events for various business workflows
├── scg                                 # SCG bash helper script
├── buf.yaml                            # Buf configuration
├── go.mod                              # Go module definition
└── README.md                           # Project documentation
```

## 📋 Notification Contracts Overview

### Enums

- **DeliveryChannel**: Defines channels for notification delivery (EMAIL, SMS, PUSH, WEBHOOK, etc.)
- **NotificationType**: Defines types of notifications (ALERT, INFORMATION, CONFIRMATION, etc.)
- **NotificationSeverity**: Defines severity levels (LOW, MEDIUM, HIGH, CRITICAL)
- **NotificationAudience**: Defines target audiences (USER, ADMIN, PARTNER, SYSTEM, INTEGRATION)
- **NotificationStatus**: Defines notification statuses (PENDING, SENT, DELIVERED, READ, FAILED, CANCELLED)
- **NotificationCategory**: Defines business domain categories (ONBOARDING, COMPLIANCE, RISK, etc.)

### Models

- **NotificationTemplate**: Templates for notifications with fields for subject, body, metadata, etc.
- **NotificationPreference**: User preferences for receiving notifications
- **NotificationDelivery**: Tracks the delivery status of notifications
- **Notification**: Core notification model with references to triggering event, entity, workflow step, and reason
- **UserNotification**: Links a notification to a specific user and tracks read/archived status
- **NotificationDigest**: Represents a digest of multiple notifications for a user

### Events

- **Onboarding Notifications**: UserInvited, UserRegistered, PartnerOnboardingStarted, PartnerVetted, PartnerOnboardingCompleted
- **Compliance and Risk Notifications**: ComplianceCheckFailed, RiskAlert
- **Logistics and Consignment Notifications**: ConsignmentCreated, ConsignmentStatusChanged, ConsignmentArrived
- **Digital Twin Notifications**: DigitalTwinAlert
- **System Notifications**: SystemMaintenance, IntegrationFailure

## 🚀 Installation and Usage

### Go

Add the library to your Go project:

```bash
go get github.com/hbttundar/scg-notification-contracts
```

Import and use in your Go code:

```go
import (
    notificationEnums "github.com/hbttundar/scg-notification-contracts/proto/enums"
    notificationModels "github.com/hbttundar/scg-notification-contracts/proto/models"
    notificationEvents "github.com/hbttundar/scg-notification-contracts/proto/events"
)

// Create a notification
notification := &notificationModels.Notification{
    Id:          uuid.New().String(),
    EventId:     eventID,
    EventType:   "ConsignmentCreated",
    EntityId:    consignmentID,
    EntityType:  "Consignment",
    WorkflowStep: "ConsignmentRegistration",
    Reason:      "New consignment registered in the system",
    Type:        notificationEnums.NotificationType_NOTIFICATION_TYPE_INFORMATION,
    Category:    notificationEnums.NotificationCategory_NOTIFICATION_CATEGORY_LOGISTICS,
    Severity:    notificationEnums.NotificationSeverity_NOTIFICATION_SEVERITY_LOW,
    Audience:    notificationEnums.NotificationAudience_NOTIFICATION_AUDIENCE_PARTNER,
    Subject:     "New Consignment Registered",
    Body:        "A new consignment has been registered in the system.",
    CreatedAt:   timestamppb.Now(),
}
```

### Code Generation

Generate Go code from Protobuf definitions:

```bash
# Install protoc and Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/enums/*.proto proto/models/*.proto proto/events/*.proto
```

## 🔄 Integration with Core Business Events

This notification contracts library is designed to work seamlessly with the core business events defined in the [scg-contracts](https://github.com/hbttundar/scg-contracts) repository. The notification events reference the triggering business events and entities, providing a complete picture of the notification context.

## 🤝 Contributing

### Contribution Guidelines

1. **Fork the repository**: Create your own fork of the repository to work on changes
2. **Create a feature branch**: Make your changes in a new branch
3. **Follow the naming conventions**: Ensure your changes adhere to the naming conventions
4. **Add documentation**: Include comprehensive documentation for any new messages, fields, or enums
5. **Run linting**: Use `buf lint` to check for any linting issues
6. **Check for breaking changes**: Use `buf breaking` to ensure your changes don't break compatibility
7. **Submit a pull request**: Include a clear description of the changes and any relevant context

### Naming Conventions

- Use **PascalCase** for message and enum names (e.g., `NotificationTemplate`, `DeliveryChannel`)
- Use **snake_case** for field names (e.g., `delivery_channel`, `notification_id`)
- Prefix enum values with the enum name (e.g., `NOTIFICATION_TYPE_ALERT`, `DELIVERY_CHANNEL_EMAIL`)
- Use descriptive names that align with the business domain

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
