# Create Project Structure

Go to project:

```bash
cd ~/alert_and_notification
```

---

## Create Directories

```bash
mkdir -p \
cmd/api \
configs \
internal/{config,http/{routes,handlers,middleware},models,storage/{mongo,redis},rabbitmq,notifications,services,utils} \
deployments/docker
```

---

## Create Files

```bash
touch \
cmd/api/main.go \
configs/config.dev.yaml \
configs/config.prod.yaml \
internal/config/config.go \
internal/http/routes/v1.go \
internal/http/handlers/alert_publish_handler.go \
internal/http/middleware/error_handling.go \
internal/models/alert.go \
internal/models/audit_log.go \
internal/storage/mongo/client.go \
internal/storage/mongo/audit_repo.go \
internal/storage/redis/conn_store.go \
internal/rabbitmq/client.go \
internal/rabbitmq/publisher.go \
internal/rabbitmq/consumer_audit.go \
internal/rabbitmq/consumer_email.go \
internal/notifications/email_sender.go \
internal/notifications/websocket_hub.go \
internal/services/audit_service.go \
internal/services/email_service.go \
internal/utils/logger.go \
deployments/docker/Dockerfile \
README.md
```

---

## Verify Structure

Install tree if needed:

```bash
sudo apt install tree
```

Check:

```bash
tree -L 4
```

Expected:

```text
alert_and_notification/
├── cmd
│   └── api
│       └── main.go
├── configs
│   ├── config.dev.yaml
│   └── config.prod.yaml
├── internal
│   ├── config
│   │   └── config.go
│   ├── http
│   │   ├── handlers
│   │   │   └── alert_publish_handler.go
│   │   ├── middleware
│   │   │   └── error_handling.go
│   │   └── routes
│   │       └── v1.go
│   ├── models
│   ├── notifications
│   ├── rabbitmq
│   ├── services
│   ├── storage
│   └── utils
├── deployments
│   └── docker
│       └── Dockerfile
└── README.md
```

---

## Purpose of Each Folder

```text
cmd/api                → application entry point
configs                → yaml configs
internal/config        → config loader
internal/http          → routes + handlers + middleware
internal/models        → request/response/domain structs
internal/storage       → mongo + redis
internal/rabbitmq      → publisher + consumers
internal/notifications → email + websocket
internal/services      → business logic
internal/utils         → logger/helpers
deployments/docker     → docker build
```
