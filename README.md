## Project Structure
    ├── dockerfile
    ├── go.mod
    ├── go.sum
    ├── .gitignore
    ├── cmd
    │   └──main.go
    └── pkg
        ├── service
        │   ├── currencyService.go
        │   ├── emailService.go
        │   ├── fileService.go
        │   └── validationService.go
        └── handler
            ├── emailHandler.go
            ├── rateHandler.go
            └── subscribeHandler.go


- **dockerfile**: Contains the Dockerfile used to build the project's Docker image.
- **go.mod** and **go.sum**: Go module files that manage the project's dependencies.
- **.gitignore**: Specifies files and directories to be ignored by Git version control.
- **cmd**: Directory that holds the main Go file of the project.
- **pkg**: Directory that contains packages with specific functionalities.
    - **service**: Contains service packages responsible for providing various functionalities.
        - **currencyService.go**: Implements functionality related to fetching the Bitcoin exchange rate in UAH (Ukrainian Hryvnia) using CoinGecko API.
        - **emailService.go**: Implements functionality for sending emails using SMTP.
        - **fileService.go**: Implements file-related operations such as writing and reading email addresses to/from a file.
        - **validationService.go**: Implements email address validation.
    - **handler**: Contains handler packages responsible for handling HTTP requests.
        - **emailHandler.go**: Handles sending emails to a list of email addresses.
        - **rateHandler.go**: Handles retrieving the Bitcoin exchange rate in UAH.
        - **subscribeHandler.go**: Handles subscribing email addresses and writing them to a file.

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
