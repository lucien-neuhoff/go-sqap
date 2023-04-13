## Project structure

```
.
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── database
│   │   ├── database.go
│   │   └── migrations
│   │       ├── create_user_table.sql
│   │       ├── create_sessions_table.sql
│   │       └── create_public_keys_table.sql
│   ├── handlers
│   │   ├── auth_handler.go
│   │   └── user_handler.go
│   ├── models
│   │   ├── login_request.go
│   │   ├── public_key.go
│   │   ├── session.go
│   │   └── user.go
│   ├── repositories
│   │   ├── session_repository.go
│   │   └── user_repository.go
│   ├── router
│   │   └── router.go
│   ├── services
│   │   ├── auth_service.go
│   │   └── user_service.go
│   └── utils
│       └── logger.go
├── test
│   ├── controllers
│   │   └── auth_test.go
│   ├── services
│   │   └── auth_service_test.go
│   └── test_helper.go
├── .env
└── go.mod
```

## Project structure explained

-   `cmd/`: This directory contains the main application entry point, `main.go`, which initializes the API and starts the HTTP server.

-   `internal/`: This directory contains all of the internal application logic, including controllers, models, repositories, routes, and services.

    -   `controllers/`: This directory contains the API controllers, which handle incoming HTTP requests, perform validation, and call the appropriate service functions.

    -   `models/`: This directory contains the application models, which represent the domain objects being manipulated by the application.

    -   `repositories/`: This directory contains the data access layer, which is responsible for interacting with the database to retrieve and store data.

    -   `routes/`: This directory contains the application routes, which map incoming HTTP requests to the appropriate controller functions.

    -   `services/`: This directory contains the business logic layer, which encapsulates the application's core functionality and is responsible for implementing use cases.

-   `test/`: This directory contains all of the test code for the application.

    -   `controllers/`: This directory contains the test code for the API controllers.

    -   `services/`: This directory contains the test code for the application services.

    -   `test_helper.go`: This file contains utility functions and setup code that can be used across multiple test files.
