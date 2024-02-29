# Stori MVP

## Description

The summary email contains information on:

1. Total balance is 39.74
2. Number of transactions in July: 2
3. Number of transactions in August: 2
4. Average debit amount: -15.38
5. Average credit amount: 35.25

### Follow the next steps to install the api

Run the containers:

```sh
docker-compose up --build
```

Stop the containers:

```sh
docker-compose down
```

Open the database:

```sh
docker exec -it docker_db mysql -u storiuser -pasdf -h localhost storidb
```
