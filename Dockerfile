FROM golang:alpine3.17 as builder

WORKDIR /app
COPY . .
RUN go build -o ci-cd-workshop-app .

FROM alpine:3.17

WORKDIR /bin
COPY --from=builder /app/ci-cd-workshop-app .
CMD ["./ci-cd-workshop-app"]