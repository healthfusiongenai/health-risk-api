FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o health-risk-api .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/health-risk-api .
COPY hlth-api.env ./

EXPOSE 8000

#CMD ["./health-risk-api"]
CMD [ "tail", "-f", "/dev/null" ]