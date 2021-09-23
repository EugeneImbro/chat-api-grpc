FROM golang:1.16-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

# Build
RUN go build -o chat-backend ./cmd/chat-backend/main.go

RUN chmod +x chat-backend

# Run
CMD [ "./chat-backend" ]