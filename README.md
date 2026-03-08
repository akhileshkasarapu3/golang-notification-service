# Golang Notification Service

A beginner-to-production notification service built with Go.

## MVP Goal

The first version will support:

- REST API to create notification requests
- PostgreSQL to store notifications and templates
- Kafka to publish notification events
- Worker to consume events
- Basic template rendering
- Email sending
- Docker Compose for local setup

## Planned Flow

Client -> REST API -> PostgreSQL -> Kafka -> Worker -> Template Render -> Email Send -> Status Update