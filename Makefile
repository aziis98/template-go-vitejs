
.PHONY: all
all: build-frontend build-backend

.PHONY: build-backend
build-backend:
	go build -v -o out/backend/server ./cmd/backend

.PHONY: build-frontend
build-frontend:
	npm run build

