#Build the Go server
FROM golang:1.22-alpine AS builder
WORKDIR /go/src/app
COPY server/ .
RUN go mod tidy
RUN go build -o server

#Prepare the Python client
FROM python:3.9-slim AS client-builder
WORKDIR /app
COPY client/requirements.txt .
RUN pip install --no-cache-dir -r ./requirements.txt
COPY client/ .

#Combine both server and client
FROM python:3.9-slim
WORKDIR /app
COPY --from=builder /go/src/app/server /app/server
COPY --from=client-builder /app /app
RUN pip install grpcio grpcio-tools googleapis-common-protos
EXPOSE 50051
CMD ["sh", "-c", "./server & sleep 5 && python3 client.py"]
