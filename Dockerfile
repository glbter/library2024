FROM golang:1.23 AS builder

#apt-get install -y --no-install-recommends
RUN apt-get update \
    && apt-get install -y apt-transport-https \
    && go install github.com/a-h/templ/cmd/templ@latest \
    && apt-get install make

RUN wget https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN mkdir /app
RUN mv tailwindcss-linux-x64 /app/tailwindcss
#RUN #alias tailwindcss=/bin/tailwindcss

COPY . /app
WORKDIR /app

RUN make build
#    APP_NAME=ourApp

#RUN ./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
#RUN templ generate
#
## Toggle CGO based on your app requirement. CGO_ENABLED=1 for enabling CGO
#RUN CGO_ENABLED=0 go build -ldflags '-X main.Environment=production -s -w -extldflags "-static"' -o /app/appbin ./cmd/$(APP_NAME)/main.go

FROM scratch

COPY --from=builder /app/static /bin/static
COPY --from=builder /app/bin/ /bin

# Since running as a non-root user, port bindings < 1024 are not possible
# 8000 for HTTP; 8443 for HTTPS;
EXPOSE 8000
EXPOSE 8443

CMD ["./bin/ourApp"]

