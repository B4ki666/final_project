FROM golang:1.25.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /final_project ./cmd/app

FROM alpine:3.23.4
WORKDIR /app
ENV TODO_PORT=7540
ENV TODO_DBFILE=scheduler.db
EXPOSE 7540
COPY --from=builder /final_project .
COPY --from=builder /app/web ./web
CMD [ "./final_project" ]