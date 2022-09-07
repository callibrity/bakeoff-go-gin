# Build
.PHONY: build
build:
	@echo "Building local image..."
	@docker build --platform amd64 -t callibrity/bakeoff-go-gin:1.0.1 .

# Stack
.PHONY:	stop
stop:
	@docker compose -f stack.yml down -v

.PHONY:	prod
prod:
	@docker compose -f stack.yml down -v
	@docker compose -f stack.yml up --build

.PHONY: dev
dev:
	@docker compose -f stack.yml down -v
	@docker compose -f stack.yml -f stack.dev.yml up