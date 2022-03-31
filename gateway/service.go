package gateway

import (
	"context"
	"fmt"
	"gateway/gateway/interface"
	"gateway/gateway/store"
	swaggerFiles "github.com/swaggo/files"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"

	ginSwagger "github.com/swaggo/gin-swagger"
)

type Service struct {
	sessionStore _interface.SessionStore
	userStore    _interface.UserStore

	country 	 _interface.CountryEndpoint
	recipe 		 _interface.RecipeEndpoint
	recipeType   _interface.RecipeTypeEndpoint
}


func NewService(redisURI string, pgURI string) *Service {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURI,
		DB:   0,
	})

	pgdb, err := pgxpool.Connect(context.Background(), pgURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	} else {
		fmt.Println(pgURI, " Connected to database !")
	}

	if !checkIfDatabaseAlreadyExist(pgdb) {
		initDatabase(pgdb)
	}

	_ = rdb
	_ = pgdb

	return &Service{
		sessionStore: store.NewSessionStoreRedis(rdb),
		userStore:    store.NewUserStorePG(pgdb),
		country:      store.NewCountryStore(),
		recipe:       store.NewRecipeStore(),
		recipeType:   store.NewRecipeTypeStore(),
	}
}

func checkIfDatabaseAlreadyExist(pgdb *pgxpool.Pool) bool {
	_, err := pgdb.Query(context.Background(), "SELECT * FROM \"User\";")
	if err != nil {
		return false
	}
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

	r.Use(s.TokenMiddleware())

	// Account --
	v1 := r.Group("/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		base := v1.Group("/")
		{
			base.GET("logout", s.Logout)
			base.POST("signup", s.Signup)
			base.POST("login", s.Login)
		}

		users := v1.Group("/user")
		{
			users.GET("", s.GetCurrentUser)
			users.PUT("", s.UpdateCurrentUser)
		}

		countries := v1.Group("/countries")
		{
			countries.GET("", s.GetAllCountries)
			countries.POST("", s.CreateCountry)
			countries.DELETE(":id", s.DeleteCountry)
			countries.PUT("", s.UpdateCountry)
		}

		recipes := v1.Group("/recipes")
		{
			recipes.GET("shops/:id", s.GetAllRecipesForShopById)
			recipes.GET(":id", s.GetRecipeById)
			recipes.POST("", s.CreateRecipe)
			recipes.POST("shops", s.AddRecipeToShopById)
			recipes.PUT(":id", s.UpdateRecipe)
			recipes.DELETE(":id", s.DeleteRecipe)
		}

		recipeTypes := v1.Group("/recipeTypes")
		{
			recipeTypes.GET("", s.GetAllRecipesType)
			recipeTypes.POST("", s.CreateRecipeType)
			recipeTypes.PUT(":id", s.UpdateRecipeType)
			recipeTypes.DELETE(":id", s.DeleteRecipeType)
		}
	}
}
