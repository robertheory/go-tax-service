# Tax Service

This service is responsible for calculating the tax for a given products.

## Description

This service is a concept project to demonstrate how Go can handle a high number of requests and tasks in parallel with low resource consumption.

## Concepts and Technologies

1. **Go**: The service is written in Go, a statically typed, compiled language that is designed for simplicity and efficiency.

   1. **Goroutines**: The service uses goroutines to handle the requests concurrently and efficiently.

   2. **Channels**: The service uses channels to communicate between the goroutines and to handle the concurrency.

   3. **WaitGroup**: The service uses the WaitGroup to wait for all the goroutines to finish before returning the response.

   4. **Testing**: The service uses the testing package to write unit tests for the service.

2. **RabbitMQ**: The service uses RabbitMQ as a message broker to handle the communication between the services.

3. **Prometheus**: The service uses Prometheus to collect and store the metrics from the service.

4. **Grafana**: The service uses Grafana to visualize the metrics collected by Prometheus.

5. **Docker**: The service is containerized with Docker to ensure that the service runs in the same environment as the production environment.

6. **Docker Compose**: The service uses Docker Compose to manage the services and dependencies in the development environment.

## Setup

1. Install [Docker](https://docs.docker.com/get-docker/)
2. Install [Docker Compose](https://docs.docker.com/compose/install/)
3. Clone this repository
4. Run the following command to start the services:

```bash
docker-compose up
```

By now the services should be running and you can access the following services:

- RabbitMQ: `http://localhost:15672`
- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000`

Now you should be able to access the RabbitMQ GUI and see the queues and exchanges created by the services.

With the provided `prometheus.yml` configuration file, the Prometheus service will scrape the metrics from RabbitMQ.

#### Grafana Dashboard

1. Access the Grafana dashboard at `http://localhost:3000`
2. Login with the default credentials: `admin` and `admin`
3. Add a new data source with the following configurations:
   - Name: `Prometheus`
   - Type: `Prometheus`
   - URL: `http://prometheus:9090`
   - Access: `Server`
4. Import the dashboard from the id `10991`.

#### RabbitMQ Configuration

Before running the services, you need to configure the RabbitMQ server with the following:

1. Create a new queue named `orders`
2. Bind the `orders` queue to the `amq.direct` exchange.
