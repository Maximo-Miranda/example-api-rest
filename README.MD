## Example APIRest

Golang api rest Service

### For run service with docker-compose:

```

Requirements:

- Docker >= v2.1.0.5 and docker-compose

1. cd example-api-rest
2. docker-compose up --build

```

### For run service in local:

```

Requirements:

- Go >= v1.12
- Postgres >= v10

Steps:

1. cd example-api-rest
2. cp .env.example .env.local
3. validate enviroment data in .env.local file
    Example:

        export EXAMPLE_API_REST_APP_PORT=8090
        export EXAMPLE_API_REST_DB_HOST=localhost
        export EXAMPLE_API_REST_DB_PORT=5432
        export EXAMPLE_API_REST_DB_NAME=postgres
        export EXAMPLE_API_REST_DB_USER=postgres
        export EXAMPLE_API_REST_DB_PASSWORD=secret

4. go build ./...
5. go run main.go

```

### API Example:

```

Endpoint: "http://localhost:8090/api/v1/user"
Method: "POST"
Header: "Content-Type": "application/json"
Example Body Request:

{
	"full_name": "Maximo Miranda",
	"dni": "123456789",
	"date_of_birth": "2019-11-22"
}

```


## Authors

* Maximo Miranda <maximomirandah@gmail.com>