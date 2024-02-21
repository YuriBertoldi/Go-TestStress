FROM golang:1.21.5 AS builder

WORKDIR /app

RUN mkdir -p cmd
RUN mkdir -p internal/model
RUN mkdir -p internal/usecase

COPY go.mod ./
COPY cmd/main.go ./cmd
COPY internal/model/report.go ./internal/model
COPY internal/model/config.go ./internal/model
COPY internal/usecase/testStress.go ./internal/usecase
COPY internal/usecase/testStress_test.go ./internal/usecase


RUN CGO_ENABLED=0 GOOS=linux go build -o go-teststress cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-teststress .

ENTRYPOINT ["./go-teststress"]
