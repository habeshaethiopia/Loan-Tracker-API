# Library Management API

API using Golang with the Gin framework and clean architecture principles
Here is a [`README.md`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FREADME.md%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/home/adane/Repository/Loan-Tracker-API/README.md") file with instructions on how to clone and run the project:

```markdown
# Loan Tracker API

API using Golang with the Gin framework and clean architecture principles.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.22.2 or later)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

## Clone the Repository

To clone the repository, run the following command:

```sh
git clone https://github.com/yourusername/LoanTrackerAPI.git
cd LoanTrackerAPI
```

## Setup Environment Variables

Create a [`.env`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2F.env%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/home/adane/Repository/Loan-Tracker-API/.env") file in the root directory of the project and add the following environment variables:

```env
DATABASE_URL=your_mongodb_connection_string
DB_NAME=your_database_name
PORT=8080
EMAIL_FROM=your_email@gmail.com
EMAIL_PASSWORD=your_app_specific_password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=465
SERVER_HOST=http://localhost:8080
TOKEN_TTL=1h
```

## Install Dependencies

To install the required dependencies, run:

```sh
go mod tidy
```

## Run the Application

To run the application, execute the following command:

```sh
go run cmd/main.go
```

The server will start on the port specified in the [`.env`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2F.env%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/home/adane/Repository/Loan-Tracker-API/.env") file (default is 8080).

## API Endpoints

The API provides the following endpoints:

- [`POST /books`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22external%22%3A%22file%3A%2F%2F%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A19%2C%22character%22%3A11%7D%7D%5D%5D "Go to definition") - Create a new book
- [`GET /books`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22external%22%3A%22file%3A%2F%2F%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A22%2C%22character%22%3A15%7D%7D%5D%5D "Go to definition") - Get all books
- [`GET /books/available`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22external%22%3A%22file%3A%2F%2F%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A22%2C%22character%22%3A15%7D%7D%5D%5D "Go to definition") - Get all available books
- [`GET /books/:id`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22external%22%3A%22file%3A%2F%2F%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A22%2C%22character%22%3A15%7D%7D%5D%5D "Go to definition") - Get a book by ID
- [`PUT /books/:id`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22external%22%3A%22file%3A%2F%2F%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A25%2C%22character%22%3A15%7D%7D%5D%5D "Go to definition") - Update a book by ID
- [`DELETE /books/:id`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22external%22%3A%22file%3A%2F%2F%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FDelivery%2Frouters%2Fbook_router.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A26%2C%22character%22%3A15%7D%7D%5D%5D "Go to definition") - Delete a book by ID

For more details, refer to the [API Documentation](doc/doc.md).

## License

This project is licensed under the MIT License. See the [`LICENSE`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2Fadane%2FRepository%2FLoan-Tracker-API%2FLICENSE%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/home/adane/Repository/Loan-Tracker-API/LICENSE") file for details.
```

Make sure to replace `yourusername` with your actual GitHub username and update the environment variables with your actual values.