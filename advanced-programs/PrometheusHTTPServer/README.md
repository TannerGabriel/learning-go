# Prometheus HTTP Server

An example application that shows how to implement Prometheus Metrics into your Golang application and visualize the data using Grafana.

## Requirements

- Golang
- Docker

## Getting started

### Starting the application

You can start all services by running the docker-compose file:

```
docker-compose up -d
```

### Generating load on HTTP Server

You can either generate load manually by visiting the HTTP page on localhost:8080 or generate load using some load generator. I use [Hey](https://github.com/rakyll/hey), which is an easy to use command-line tool.

```
hey -z 5m -q 5 -m GET -H "Accept: text/html" http://127.0.0.1:8080

hey -n 10000 -c 100 -m GET -H "Accept: text/html" http://127.0.0.1:8080
```

You can now see the metrics of your application on http://localhost:8080/metrics.

### Importing Grafana dashboard

Lastly, you can visualize the metrics data using Grafana, which should be running on localhost:3000 after deploying the docker-compose file.

1. You need to set Prometheus as your data source
2. Visualize the exposed Prometheus metrics in a Dashboard (I have also included a dashboard file in the Grafana folder that you can import)
