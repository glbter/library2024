FROM golang:1.22 AS builder

COPY ${PWD} /app
WORKDIR /app


RUN apt-get update  \
    && apt-get install tailwindcss \
    && apt-get install templ

RUN ./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
RUN templ generate

# Toggle CGO based on your app requirement. CGO_ENABLED=1 for enabling CGO
RUN CGO_ENABLED=0 go build -ldflags '-X main.Environment=production -s -w -extldflags "-static"' -o /app/appbin ./cmd/$(APP_NAME)/main.go

FROM scratch

COPY --from=builder /app /app

WORKDIR /app
# Since running as a non-root user, port bindings < 1024 are not possible
# 8000 for HTTP; 8443 for HTTPS;
EXPOSE 8000
EXPOSE 8443

CMD ["go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go"]

