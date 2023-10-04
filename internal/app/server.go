package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"kredit-service/internal/repository"
	"net/http"
	"time"
)

type Handler struct {
	User    repository.User
	Product repository.Product
}

func Run(u repository.User, p repository.Product) {
	e := echo.New()

	handler := Handler{
		User:    u,
		Product: p,
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Rate Limiter Configuration
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	e.Use(middleware.RateLimiterWithConfig(config))
	//e.POST("/register", handler.RegisterUser)
	//e.POST("/login", handler.LoginUser)
	//
	//mobile := e.Group("/mobile")
	//{
	//	mobile.Use(JWTMiddleware("secret"))
	//	mobile.GET("/products", handler.GetListProduct)
	//	mobile.GET("/categories", handler.GetListCategories)
	//	mobile.GET("/products/detail", handler.GetProductDetail)
	//	mobile.POST("/cart", handler.AddToCart)
	//	mobile.GET("/cart", handler.GetUserCarts)
	//}
	//
	//admin := e.Group("/admin")
	//{
	//	admin.Use(JWTMiddleware("secret"))
	//	admin.Use(AdminMiddleware)
	//	admin.POST("/product", handler.AddProduct)
	//	admin.PUT("/product", handler.UpdateProduct)
	//}
	//
	//e.Logger.Fatal(e.Start(":8080"))
}
