## Project structure

```
.
├── cmd
│   └── api
│       └── main.go
├── encryption
│   └── encryption.go
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
│   │   ├── keys_handler.go
│   │   ├── session_handler.go
│   │   └── validate_create_user_request.go
│   ├── models
│   │   ├── login_request.go
│   │   ├── public_key.go
│   │   ├── session.go
│   │   └── user.go
│   ├── repositories
│   │   ├── keys_repository.go
│   │   ├── session_repository.go
│   │   └── user_repository.go
│   ├── router
│   │   └── router.go
│   ├── services
│   │   ├── auth_service.go
│   │   ├── keys_service.go
│   │   └── session_service.go
│   └── utils
│       ├── utils.go
│       └── logger.go
├── test
│   ├── controllers
│   │   └── auth_test.go
│   ├── services
│   │   ├── auth_test.go
│   │   ├── keys_test.go
│   │   └── session_test.go
│   └── test_helper.go
├── .env
├── go.mod
├── private.pem
├── public.pem
└── go.mod
```
