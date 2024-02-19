# A Simple Go Application to Test Auto Scaling functionalities for Kubernetes Clusters

This application allows you to send compute intensive HTTP requests to a simple webserver.

## Docker Image

`switzerchees/go-scaling:latest`

### Build Image by your Own

```bash
# Simple Build
docker build -t go-scaling .
# Multi Platform Build and Push to Docker Hub
docker buildx build --push --platform=linux/amd64,linux/arm64 -t <docker-hub-username>/go-scaling .
```

### Endpoints

- `/compute/{fibNumber}`: Allows you to calculate the n-th fibonacci number. Produces a lot of load on the http server depends on the number. Don't go higher than 45 if you do you will not get a response until the end of time 😉
- `/escalate`: Behind this endpoint is an empty infinity loop. That means the http server behind will go up to 100% CPU load immediately.
- `/stopescalate`: This endpoint will stop the infinite loop.

## Kubernetes port Forwarding for Service

This command allow you to forward kubernetes ressources running inside your kubernetes cluster and bind it to a port on your local machine for testing purposes.

### Pod Forwarding

```bash
kubectl port-forward pod/[POD_NAME] [LOCAL_PORT]:[POD_PORT] -n [NAMESPACE]
# Example
kubectl port-forward pod/my-pod 8080:80 -n my-namespace
```

### Service Forwarding

```bash
kubectl port-forward service/[SERVICE_NAME] [LOCAL_PORT]:[SERVICE_PORT] -n [NAMESPACE]
# Example
kubectl port-forward service/my-service 8080:80 -n my-namespace
```

#### Example

`curl http://localhost:8080/compute/20` -> This will send a http request to the compute endpoint

## Bechmarking Tools

#### Example - Hey (Windows, Linux, Mac)

https://github.com/rakyll/hey?tab=readme-ov-file#installation

`hey -z 20s -c 10 http://localhost:8080/compute/30` -> Will make 10 concurrent requests for 30 seconds and outputs a summary after

### Benchmarking Tool - Apache Bench (Mac, Linux)

https://www.inmotionhosting.com/support/edu/wordpress/performance/stress-test-with-apachebench/

`ab -n 3000000 -c 10 http://localhost:8080/compute/30` -> Will make 10 concurrent requests for 30 seconds and outputs a summary after

### Benchmarking Tool - Siege (Mac, Linux, Windows)

- Linux: https://linuxhint.com/install-siege-ubuntu/
- Mac: https://formulae.brew.sh/formula/siege
- Windows: https://code.google.com/archive/p/siege-windows/

`siege -c10 -t1M http://example.com/ http://localhost:8080/compute/30` -> Will make 10 concurrent requests for 1 minute with live output.
