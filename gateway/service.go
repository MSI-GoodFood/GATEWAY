package gateway

import (
	"context"
	"fmt"
	"github.com/MSI-GoodFood/GATEWAY/gateway/interface"
	"github.com/MSI-GoodFood/GATEWAY/gateway/store"
	service "github.com/MSI-GoodFood/GATEWAY/proto"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Service struct {
	sessionStore  _interface.SessionStore
	userStore     _interface.UserStore
	userRoleStore _interface.UserRoleStore

	grpcTry 	  _interface.GrpcTryEndpoint
	country 	  _interface.CountryEndpoint
	recipe 		  _interface.RecipeEndpoint
	recipeType    _interface.RecipeTypeEndpoint
	service       _interface.ServiceEndpoint
	orderStatus   _interface.OrderStatusEndpoint
	shopType      _interface.ShopTypeEndpoint
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
		sessionStore:  store.NewSessionStoreRedis(rdb),
		userStore:     store.NewUserStorePG(pgdb),
		userRoleStore: store.NewUserRoleStorePG(pgdb),
		grpcTry:       store.NewGrpcTryStore(),
		country:       store.NewCountryStore(),
		recipe:        store.NewRecipeStore(),
		recipeType:    store.NewRecipeTypeStore(),
		service:       store.NewServiceStore(),
		orderStatus:   store.NewOrderStatusStore(),
		shopType:      store.NewShopTypeStore(),
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
	if err != nil { log.Println(err) }

	requests := strings.Split(string(file), ";\n")
	for _, request := range requests {
		_, err := dbPool.Exec(context.Background(), request)
		if err != nil { log.Fatal(err) }
	}
}

func (s *Service) SetupRoute(r gin.IRouter, serverGestion *grpc.ClientConn) {

	ApiTest := service.NewTestClient(serverGestion)

	r.Use(s.TokenMiddleware())
	r.Use(s.CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{

		base := v1.Group("/")
		{
			base.GET("logout", s.Logout)
			base.POST("signup", s.Signup)
			base.POST("login", s.Login)
		}

		test := v1.Group("/test")
		{
			test.GET("", func(c *gin.Context) { s.Name(c, ApiTest) })
		}

		users := v1.Group("/user")
		{
			users.GET("", s.GetCurrentUser)
			users.PUT("", s.UpdateCurrentUser)
		}

		userRoles := v1.Group("/userRoles")
		{
			userRoles.GET("", s.GetAllUserRole)
			userRoles.POST("", s.CreateUserRole)
			userRoles.PUT("", s.UpdateUserRole)
			userRoles.DELETE("", s.DeleteUserRole)
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
			recipes.POST("shops/:id", s.AddRecipeToShopById)
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

		services := v1.Group("/services")
		{
			services.GET("", s.GetAllService)
			services.POST("", s.CreateService)
			services.PUT(":id", s.UpdateService)
			services.DELETE(":id", s.DeleteService)
		}

		orderStatus := v1.Group("/orderStatus")
		{
			orderStatus.GET("", s.GetAllOrderStatus)
			orderStatus.POST("", s.CreateOrderStatus)
			orderStatus.PUT(":id", s.UpdateOrderStatus)
			orderStatus.DELETE(":id", s.DeleteOrderStatus)
		}

		shopType := v1.Group("/shopTypes")
		{
			shopType.GET("", s.GetAllShopType)
			shopType.POST("", s.CreateShopType)
			shopType.PUT(":id", s.UpdateShopType)
			shopType.DELETE(":id", s.DeleteShopType)
		}
	}
}
