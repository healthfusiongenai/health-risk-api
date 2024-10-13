FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o energy-usage-simulator .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/energy-usage-simulator .
COPY gj-sim.env ./

EXPOSE 8000

CMD ["./energy-usage-simulator"]
#CMD [ "tail", "-f", "/dev/null" ]