module github.com/MSI-GoodFood/GATEWAY

go 1.14

replace github.com/MSI-GoodFood/GATEWAY => github.com/MSI-GoodFood/GATEWAY latest

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.4
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/jackc/pgx/v4 v4.15.0
	github.com/joho/godotenv v1.4.0
)

require (
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.4.1
	github.com/swaggo/swag v1.8.1
	golang.org/x/net v0.0.0-20220325170049-de3da57026de // indirect
	golang.org/x/sys v0.0.0-20220330033206-e17cdc41300f // indirect
	golang.org/x/tools v0.1.10 // indirect
)
