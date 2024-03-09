# GOAPI

## Description

This project is a Go API that allows users to check their coin balance. It is a simple API that demonstrates the basic concepts of building a Go API.

## Usage

To start the server, run the following command:

```sh
go run ./cmd/api/
```

The server will start on localhost:8000.

## API Endpoints

-   **GET /account/coins/{username}**: Returns the coin balance of a user. The request body should be a JSON object with a username field.

## Contributing

Contributions are welcome. Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
