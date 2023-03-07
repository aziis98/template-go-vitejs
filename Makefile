
NPM = npm

.PHONY: all
all: build-frontend build-backend

.PHONY: setup
setup:
	mkdir -p out/frontend
	mkdir -p out/backend
	$(NPM) install

.PHONY: build-frontend
build-frontend:
	$(NPM) run build

.PHONY: build-backend
build-backend:
	go build -v -o out/backend/server ./cmd/server

