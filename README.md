# IP service

Golang service, that accept access logs records with IPs values and provide Prometheus metric for count of unique IPs

## How to start service

### Test start:

```shell
go run main.go tree.go
```

### Run as executable:

Build an executable file:

```shell
go build
```

Run executable file:

```shell
./ip-service
```

### Run in docker:

Build docker file:

```shell
docker build . -t ip-service
```

Run docker image:

```shell
docker run -p 5000:5000 -p 9102:9102 ip-service:latest
```

## Testing

Run

```shell
go test
```

## Endpoints

Service provide two enpoints with different ports

### host-url:5000/logs

`/logs` endpoint accept access logs records. Log record should have following format:

```json
{ "timestamp": "2020-06-24T15:27:00.123456Z", "ip": "83.150.59.250", "url": ... }
```

### host-url:9102/metrics

`/metrics` endpoint provide [Prometheus](https://prometheus.io/) metric `unique_ip_addresses`, that contain cuulative count of unique IPs, occured in access logs records.

## Some technical background

Tree structure was used as a storage for unique IPs. All incomming IP values is splitting in symbols, and array of symbols is used as a branch in a tree. 

Here is example of how it's look like.

Let's imagine, that `1.2.3.4` IP came to service. Here is how tree will look like after it will be added:

```
root -> '1' -> '.' -> '2' -> '.' -> '3' -> '.' -> '4'
```

If after that `1.2.5.6` address will came, tree will looks like that:


```
                               ---> '3' -> '.' -> '4'
                              |
root -> '1' -> '.' -> '2' -> '.' 
                              |
                               ---> '5' -> '.' -> '6'
```

Such approach provide a way to use memory more efficient and check IP uniqness much more quick.

## Used libraries

- **github.com/prometheus/client_golang/prometheus** for Prometheus metric generation
- **github.com/prometheus/client_golang/prometheus/promhttp** for Prometheus endpoint handling
- **github.com/sirupsen/logrus** for logging