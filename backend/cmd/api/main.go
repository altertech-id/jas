package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"jayaabadisolusi/backend/internal/config"
	"jayaabadisolusi/backend/internal/domain"
	"jayaabadisolusi/backend/internal/handler"
	"jayaabadisolusi/backend/internal/middleware"
	"jayaabadisolusi/backend/internal/repository"
	"jayaabadisolusi/backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()

	// 1. Initialize Repository (Postgres with Memory Fallback)
	var repo repository.Repository
	dbPool, err := repository.ConnectDB(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Printf("[WARNING] Database PostgreSQL tidak terhubung (%v). Beralih ke in-memory repository dengan data seed resmi.", err)
		repo = repository.NewMemoryRepository()
	} else {
		defer dbPool.Close()
		repo = repository.NewPostgresRepository(dbPool)
	}

	// 2. Initialize Services & Handlers
	storageService := service.NewStorageService(cfg, repo)
	sanitizerService := service.NewSanitizer()

	publicHandler := handler.NewPublicHandler(repo)
	adminHandler := handler.NewAdminHandler(repo, storageService, sanitizerService)

	// 3. Initialize Fiber App
	app := fiber.New(fiber.Config{
		AppName:      "Jaya Abadi Solusi REST API",
		BodyLimit:    10 * 1024 * 1024, // 10MB limit for multipart image uploads
		ServerHeader: "Fiber",
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("%s, http://localhost:5173, http://127.0.0.1:5173, http://localhost:4173", cfg.FrontendURL),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"company": "PT Jaya Abadi Solusi",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// Public Routes
	api := app.Group("/api")
	{
		api.Get("/site-settings", publicHandler.GetSiteSettings)
		api.Get("/services", publicHandler.ListServices)
		api.Get("/services/:slug", publicHandler.GetServiceBySlug)
		api.Get("/projects", publicHandler.ListProjects)
		api.Get("/projects/:slug", publicHandler.GetProjectBySlug)
		api.Get("/categories", publicHandler.ListCategories)
		api.Get("/posts", publicHandler.ListPosts)
		api.Get("/posts/:slug", publicHandler.GetPostBySlug)
		api.Get("/gallery", publicHandler.ListGallery)
		api.Get("/pages/:key", publicHandler.GetPageContent)

		// Rate limit on leads: max 5 submissions per 5 minutes per IP
		leadRateLimiter := middleware.NewRateLimiter(5, 5*time.Minute)
		api.Post("/leads", leadRateLimiter.Middleware(), publicHandler.SubmitLead)
	}

	// Admin Routes (Protected by AuthMiddleware)
	adminAuth := middleware.AuthMiddleware(cfg)
	registerAdminRoutes := func(router fiber.Router) {
		// Dashboard & Common
		router.Get("/dashboard/stats", adminHandler.GetDashboardStats)

		// Media (Admin + Editor)
		router.Post("/media", adminHandler.UploadMedia)
		router.Get("/media", adminHandler.ListMedia)
		router.Delete("/media/:id", adminHandler.DeleteMedia)

		// Blog (Admin + Editor)
		router.Post("/posts", adminHandler.SavePost)
		router.Put("/posts/:id", adminHandler.SavePost)
		router.Delete("/posts/:id", adminHandler.DeletePost)
		router.Post("/categories", adminHandler.CreateCategory)

		// Projects (Admin + Editor)
		router.Post("/projects", adminHandler.SaveProject)
		router.Put("/projects/:id", adminHandler.SaveProject)
		router.Delete("/projects/:id", adminHandler.DeleteProject)

		// Strict Admin-Only Routes
		adminOnly := middleware.RequireRole(domain.RoleAdmin)

		// Services
		router.Post("/services", adminOnly, adminHandler.SaveService)
		router.Put("/services/:id", adminOnly, adminHandler.SaveService)
		router.Delete("/services/:id", adminOnly, adminHandler.DeleteService)

		// Pages (CMS)
		router.Put("/pages/:key", adminOnly, adminHandler.SavePage)

		// Gallery
		router.Post("/gallery", adminOnly, adminHandler.SaveGalleryItem)
		router.Delete("/gallery/:id", adminOnly, adminHandler.DeleteGalleryItem)

		// Leads Management
		router.Get("/leads", adminOnly, adminHandler.ListLeads)
		router.Patch("/leads/:id/status", adminOnly, adminHandler.UpdateLeadStatus)

		// Site Settings
		router.Put("/site-settings/:key", adminOnly, adminHandler.UpdateSiteSetting)

		// User Management
		router.Get("/users", adminOnly, adminHandler.ListProfiles)
		router.Patch("/users/:id/status", adminOnly, adminHandler.UpdateProfileStatus)
	}

	// Mount under /api/admin (primary for reverse proxy) and /admin (compatibility)
	registerAdminRoutes(app.Group("/api/admin", adminAuth))
	registerAdminRoutes(app.Group("/admin", adminAuth))

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("[SERVER] Mematikan server secara aman...")
		_ = app.Shutdown()
	}()

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("[SERVER] Jaya Abadi Solusi API aktif di port %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Printf("[SERVER] Server terhenti: %v", err)
	}
}
