# Project - Load Balancer

## Set-up Instructions & App Details

For this project, we will be running API servers. Primarily, we will need to run 5 backend servers from ports 8081 to 8085. To do so, you will need to run 5 seperate terminals, each running 1 of the backend server, which you can do by running the respective commands:

```bash
go run backend/main.go -port=8081
```

```bash
go run backend/main.go -port=8082
```

```bash
go run backend/main.go -port=8083
```

```bash
go run backend/main.go -port=8084
```

```bash
go run backend/main.go -port=8085
```

After running the commands in the respective terminals, we will now need to run the load balancer. To do so, we can simply run:

```bash
scripts/run_load_balancer.sh
```

Running this `run_load_balancer.sh` script will start up the load balancer server at Port 8000 (set in environmental variables)

## Triggering of Load Balancer

To see how the Load Balancer is balancing the loads across the backends, we will need to send HTTP requests to the Load Balancer. To do so, you may simply run:

```bash
scripts/hit_load_balancer.sh
```

This will do 10 cURL requests to the Load Balancer. In the respective terminals, you can easily see how the requests are being allocated to each backend server.
