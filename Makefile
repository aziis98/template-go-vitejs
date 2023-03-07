
.PHONY: all
all: build-frontend build-backend

.PHONY: build-frontend
build-frontend:
	npm run build

.PHONY: build-backend
build-backend:
	go build -v -o out/backend/server ./cmd/backend

