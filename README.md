# Stori MVP

## Description

The summary email contains information on:

1. Total balance is 39.74
2. Number of transactions in July: 2
3. Number of transactions in August: 2
4. Average debit amount: -15.38
5. Average credit amount: 35.25

### Requeriments:

You need Docker Engine, to use Docker Compose on your computer

### Follow the next steps to install the api

In a terminal:

```sh
git clone https://github.com/hectorgool/storimvp.git
```

go to the project folder

```sh
cd storimvp.git
```

Run the containers:

```sh
docker-compose up --build
```

Please, be patient, wait for a response like this:

```sh
docker_api | 2024/02/29 23:31:14 Build ok.
docker_api | 2024/02/29 23:31:14 Restarting the given command.
docker_api | 2024/02/29 23:31:15 stdout: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
docker_api | 2024/02/29 23:31:15 stdout:
docker_api | 2024/02/29 23:31:15 stdout: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
docker_api | 2024/02/29 23:31:15 stdout:  - using env:	export GIN_MODE=release
docker_api | 2024/02/29 23:31:15 stdout:  - using code:	gin.SetMode(gin.ReleaseMode)
docker_api | 2024/02/29 23:31:15 stdout:
docker_api | 2024/02/29 23:31:15 stdout: [GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
docker_api | 2024/02/29 23:31:15 stdout: [GIN-debug] GET    /sendmail/:userEmail      --> storimvp/controller.SendMail (3 handlers)
docker_api | 2024/02/29 23:31:15 stdout: [GIN-debug] DELETE /reset                    --> storimvp/controller.Reset (3 handlers)
docker_api | 2024/02/29 23:31:15 stdout: [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
docker_api | 2024/02/29 23:31:15 stdout: Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
docker_api | 2024/02/29 23:31:15 stdout: [GIN-debug] Listening and serving HTTP on :8080
```

This api have 3 endpoits:

For test

```sh
GET http://localhost:8080
```

For this endpoint:
In the email account that you put on this endpoint(for example: hectorgool@gmail.com), you will receive an email with the following information:
Total balance is 39.74
Average debit amount: -15.38
Number of transactions in July: 2
Average credit amount: 35.25
Number of transactions in August: 2
The file information: txns.cvs, is saved in a MySQL database in the db_documents table: db_documents.
Check your email to validate the report, it is important to mention that the email account where you check the email with the report is the same one you put on the endpoint

```sh
GET http://localhost:8080/sendmail/hectorgool@gmail.com
```

This endpoint, deletes all records from the db_docuements table

```sh
DELETE http://localhost:8080/reset
```

In another terminal you can, stop the containers:

```sh
docker-compose down
```

In another terminal you can open the database with:

```sh
docker exec -it docker_db mysql -u storiuser -pasdf -h localhost storidb
```

Use gmail credentials for sending mail to work, edit the variables.env file to put your credentials

```sh
SMTP_SENDER=
SMTP_PASSWD=""
```

For support contact Santo at hectorgool@gmail.com
