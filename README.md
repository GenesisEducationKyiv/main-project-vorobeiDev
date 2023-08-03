
## Running the Project

### Without Docker

To run the project without Docker, follow these steps:

1. Make sure you have Go installed on your machine.
2. Navigate to the root directory of the project.
3. Run the following command to build the project:

```shell
    go build -o main ./cmd/main.go
```


4. After the build is successful, run the following command to start the application:

```shell
    ./main
```

The application will start running on http://localhost:5000.

### With Docker

To run the project with Docker, make sure you have Docker installed on your machine. Then, follow these steps:

1. Navigate to the root directory of the project.
2. Build the Docker image using the provided Dockerfile:

```shell
    docker build -t crypto-client .
```

3. After the build is complete, run the Docker container:

```shell
    docker run -p 5000:5000 crypto-client
```

The application will start running inside the Docker container on http://localhost:5000.

## Endpoints

The project provides the following endpoints:

- **GET /rate**: Retrieves the current Bitcoin exchange rate in UAH.
- **POST /subscribe**: Subscribes an email address to receive updates.
- **POST /sendEmails**: Sends emails to all subscribed email addresses.


Please refer to the respective handler files for more details on how each endpoint is implemented.
