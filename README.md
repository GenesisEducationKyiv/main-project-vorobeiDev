## Project Structure
    ├── dockerfile
    ├── go.mod
    ├── go.sum
    ├── .gitignore
    ├── cmd
    │   └──main.go
    └── pkg
        ├── service
        │   ├── currency_service.go
        │   ├── email_service.go
        │   └── services.go
        ├── repository
        │   ├── email_repository.go
        │   └── email_repository_test.go
        └── handler
            ├── email_handler.go
            ├── handlers.go
            ├── rate_handler.go
            ├── subscribe_handler.go
            └── subscribe_handler_integration_test.go


- **dockerfile**: Contains the Dockerfile used to build the project's Docker image.
- **go.mod** and **go.sum**: Go module files that manage the project's dependencies.
- **.gitignore**: Specifies files and directories to be ignored by Git version control.
- **cmd**: Directory that holds the main Go file of the project.
- **pkg**: Directory that contains packages with specific functionalities.
  - **handler**: Contains handler packages responsible for handling HTTP requests.
    - **handlers.go**: General handler. And route registration method.
    - **email_handler.go**: Handles sending emails to a list of email addresses.
    - **rate_handler.go**: Handles retrieving the Bitcoin exchange rate in UAH.
    - **subscribe_handler.go**: Handles subscribing email addresses and writing them to a file.
    - **subscribe_handler_integration_test.go**: Integration test for handles subscribing email addresses and writing them to a file.

  - **service**: Contains service packages responsible for providing various functionalities.
    - **currency_service.go**: Implements functionality related to fetching the Bitcoin exchange rate in UAH (Ukrainian Hryvnia) using CoinGecko API.
    - **email_service.go**: Implements functionality for sending emails using SMTP.
    - **services**: General file for methods implementing.
    
  - **repository**
    - **email_repository.go**: Implements file-related operations such as writing and reading email addresses to/from a file. Implements email address validation.
    - **email_repository_test.go**: Unit test for implements file-related operations such as writing and reading email addresses to/from a file. Implements email address validation.


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
