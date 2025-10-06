package routes

import (
	"pos_api_go/internal/handlers"
	"pos_api_go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *handlers.Handler
}

func NewRouter(h *handlers.Handler) *Router {
	return &Router{
		handler: h,
	}
}

func (rt *Router) SetupAPIRoutes() *gin.Engine {
	engine := gin.New()

	// Security configuration - Trust only specific proxies
	engine.SetTrustedProxies([]string{"127.0.0.1", "::1"}) // Only localhost

	// For production, you might want to set specific proxy IPs:
	// engine.SetTrustedProxies([]string{"192.168.1.0/24", "10.0.0.0/8"})

	// Or disable proxy trust entirely for maximum security:
	// engine.SetTrustedProxies(nil)

	// Global middleware
	engine.Use(middleware.RequestLogger())
	engine.Use(middleware.CrossOriginHandler())
	engine.Use(gin.Recovery())
	engine.Use(middleware.ErrorLogger())

	// Health check
	engine.GET("/health", rt.handler.HealthCheck)

	// API version 1
	v1 := engine.Group("/api/v1")
	rt.setupAuthRoutes(v1)
	rt.setupProtectedRoutes(v1)

	return engine
}

func (rt *Router) setupAuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/signin", rt.handler.SignIn)
		authGroup.POST("/signup", rt.handler.SignUp)
		authGroup.POST("/refresh", rt.handler.RefreshToken)
	}
}

func (rt *Router) setupProtectedRoutes(rg *gin.RouterGroup) {
	protected := rg.Group("/")
	protected.Use(middleware.TokenValidator())

	// User management
	userGroup := protected.Group("/users")
	{
		userGroup.GET("/", rt.handler.ListUsers)
		userGroup.POST("/", rt.handler.CreateUser)
		userGroup.GET("/:id", rt.handler.GetUserByID)
		userGroup.PUT("/:id", rt.handler.UpdateUser)
		userGroup.DELETE("/:id", rt.handler.RemoveUser)
	}

	// Product management
	productGroup := protected.Group("/products")
	{
		productGroup.GET("/", rt.handler.ListProducts)
		productGroup.POST("/", rt.handler.AddProduct)
		productGroup.GET("/:id", rt.handler.GetProductByID)
		productGroup.PUT("/:id", rt.handler.ModifyProduct)
		productGroup.DELETE("/:id", rt.handler.RemoveProduct)
	}

	// Category management
	categoryGroup := protected.Group("/categories")
	{
		categoryGroup.GET("/", rt.handler.ListCategories)
		categoryGroup.POST("/", rt.handler.CreateCategory)
	}

	// Customer management
	customerGroup := protected.Group("/customers")
	{
		customerGroup.GET("/", rt.handler.ListCustomers)
		customerGroup.POST("/", rt.handler.CreateCustomer)
	}

	// Sales transactions
	saleGroup := protected.Group("/sales")
	{
		saleGroup.GET("/", rt.handler.ListSales)
		saleGroup.POST("/", rt.handler.ProcessSale)
		saleGroup.GET("/:id", rt.handler.GetSaleByID)
	}

	// Dashboard/reports (require higher privileges)
	dashboardGroup := protected.Group("/dashboard")
	dashboardGroup.Use(middleware.RequireRole("Manager")) // Example role-based access
	{
		// TODO: Add dashboard endpoints
	}
}
