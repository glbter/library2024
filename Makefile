.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: tailwind-build-dev
tailwind-build-dev:
	pnpm exec tailwindcss -i ./static/css/input.css -o ./static/css/style.css

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: dev
dev:
	air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	CGO_ENABLED=0 go build -ldflags "-X main.Environment=production" -o ./bin/ourApp ./cmd/main.go

.PHONY: build-dev
build-dev:
	make tailwind-build-dev
	make templ-generate
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	go test -race -v -timeout 30s ./...