# Summary

This app contains a POC health risk api is used to register and login users with the healthfusiongenai platform.

## Overview

The simulator can:

- create users and their profiles
- simulate energy usages based on profiles

## Setting up Postgres

Postgres can be run within Docker, I am using Rancher to do so.

- helm install gj-postgresql bitnami/postgresql
- Setup port forwarding in rancher dashboard on port 5432 (matches hlth-api.env)
-- click on port forwarding menu,
-- in Rancher udpate local port to 5432
- Get the password from k8s secret
-- ```export POSTGRES_PASSWORD=$(kubectl get secret --namespace default gj-postgresql -o jsonpath="{.data.postgres-password}" | base64 -d)```
- test connection to database: psql
-- ```psql --host 127.0.0.1 -U postgres -d gj  -p 5432```
-- There can be issues on a macos, where listen only happens with ipv6 address.  
-- ```netstat -an  | grep LISTEN | grep 5432```
-- ```tcp6       0      0  ::1.5432               *.*                    LISTEN```

## Running the Simulator

- Run the simulator
-- ```go run ./main.go```
- Check Healthchecker from browser and/or curl
-- go to ```http://localhost:8000/api/healthchecker```
-- ```➜  ~ curl http://localhost:8000/api/healthchecker```
--- response is ```{"message":"Welcome to GJ Simulator","status":"success"}```

## Debugging

You can check the logs of the container by running

```shell
➜  health-risk-api git:(main) ✗ docker container logs health-risk-api
2024/10/10 03:37:17 Failed to connect to the Database
```

### Check to see what the service is listening on

You will need to run the simulator with the make command, to get the IP of the postgres service.

-- ```make run```

- Attach to the container

-- ```make exec```

```shell
➜  health-risk-api git:(main) ✗ make exec
docker exec -it health-risk-api /bin/sh
~ # netstat 
Active Internet connections (w/o servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       
tcp        0      0 0f3250667a85:39960      10.42.0.31:postgresql   ESTABLISHED 
Active UNIX domain sockets (w/o servers)
Proto RefCnt Flags       Type       State         I-Node Path
~ # netstat -l
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       
tcp        0      0 :::8000                 :::*                    LISTEN      
Active UNIX domain sockets (only servers)
Proto RefCnt Flags       Type       State         I-Node Path
~ # 
```

### Understanding what IP to use for postgres

You need to check the pods, and describe them

```shell
➜  health-risk-api git:(main) ✗ k get pods
NAME                     READY   STATUS    RESTARTS          AGE
gj-postgresql-0          1/1     Running   16 (34h ago)      246d
gj-mariadb-0             1/1     Running   23 (34h ago)      228d
spark-cluster-master-0   1/1     Running   17 (34h ago)      318d
spark-cluster-worker-1   1/1     Running   181 (6m41s ago)   318d
spark-cluster-worker-0   1/1     Running   187 (4m21s ago)   318d
➜  health-risk-api git:(main) ✗ k describe pod gj-postgresql-0
```

Then describe the pod

```shell
➜  health-risk-api git:(main) ✗ k describe pod gj-postgresql-0
Name:         gj-postgresql-0
Namespace:    default
Priority:     0
Node:         ip-10-42-0-31.eu-west-1.compute.internal/10.42.0.31
Start Time:   Tue, 03 Jul 2024 14:46:26 +0200
Labels:       app.kubernetes.io/instance-type=db.t3.small
...
...
...
Annotations:      <none>
Status:           Running
IP:               10.42.0.31
IPs:
IP:           10.42.0.31
```

Use that IP to connet to database in the hlth-api.env file

```shell
➜  health-risk-api git:(main) ✗ cat hlth-api.env
POSTGRES_HOST=10.42.0.31
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=123
POSTGRES_DB=gj
```

- Run the simulator
-- ```make run```
- Attach to the container
-- ```make exec```
