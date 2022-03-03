package todo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	sessionStore SessionStore
	userStore    UserStore
	todoStore    TodoStore
}

func NewService(redisURI string, pgURI string) *Service {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURI,
	})

	pgdb, err := pgxpool.Connect(context.Background(), pgURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	} else { fmt.Println(pgURI, " Connected to database !") }

	if !checkIfDatabaseAlreadyExist(pgdb) { initDatabase(pgdb) }

	_ = rdb
	_ = pgdb

	return &Service{
		// sessionStore: NewSessionStoreRedis(rdb),
		userStore: NewUserStorePG(pgdb),
		todoStore: NewTodoStorePG(pgdb),
	}
}

func checkIfDatabaseAlreadyExist(pgdb *pgxpool.Pool) bool {
	_, err := pgdb.Query(context.Background(), "SELECT * FROM \"User\";")
	if err != nil { return false }
	return true
}

func initDatabase(dbPool *pgxpool.Pool) {
	fmt.Println("init database -- gonna read file")

	file, err := ioutil.ReadFile("./database.sql")
	if err != nil { log.Fatal(err) }

	requests := strings.Split(string(file), ";\n")
	for _, request := range requests {
		_, err := dbPool.Exec(context.Background(), request)
		if err != nil { log.Fatal(err) }
	}
}

func (s *Service) SetupRoute(r gin.IRouter) {
	// Account --
	r.POST("/user/signup", s.Signup)
	r.POST("/user/login", s.Login)

	// User --
	r.GET("/user/:id", s.GetUserById)

	// Todo --
	r.POST("/todos", s.CreateTodo)

	// GET  /user/logout

	// Must be protected by a middleware checking the validity of the token in the cookie
	// 401 if not valid

	// POST    /todos
	// GET     /todos
	// GET     /todos/:id
	// UPDATE  /todos/:id
	// DELETE  /todos/:id
}
