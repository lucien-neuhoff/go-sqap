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