package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"jayaabadisolusi/backend/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository interface defines all data operations required by JAS
type Repository interface {
	// Site Settings
	GetSiteSettings(ctx context.Context) (map[string]json.RawMessage, error)
	GetSiteSetting(ctx context.Context, key string) (json.RawMessage, error)
	UpdateSiteSetting(ctx context.Context, key string, value json.RawMessage) error

	// Services
	ListServices(ctx context.Context, onlyPublished bool) ([]domain.Service, error)
	GetServiceBySlug(ctx context.Context, slug string) (*domain.Service, error)
	SaveService(ctx context.Context, service *domain.Service) error
	DeleteService(ctx context.Context, id string) error

	// Projects
	ListProjects(ctx context.Context, onlyPublished bool) ([]domain.Project, error)
	GetProjectBySlug(ctx context.Context, slug string) (*domain.Project, error)
	SaveProject(ctx context.Context, project *domain.Project) error
	DeleteProject(ctx context.Context, id string) error

	// Blog & Categories
	ListCategories(ctx context.Context) ([]domain.Category, error)
	CreateCategory(ctx context.Context, cat *domain.Category) error
	ListPosts(ctx context.Context, categorySlug string, status string, page, limit int) ([]domain.Post, int64, error)
	GetPostBySlug(ctx context.Context, slug string) (*domain.Post, error)
	SavePost(ctx context.Context, post *domain.Post) error
	DeletePost(ctx context.Context, id string) error

	// Gallery
	ListGalleryItems(ctx context.Context, category string) ([]domain.GalleryItem, error)
	SaveGalleryItem(ctx context.Context, item *domain.GalleryItem) error
	DeleteGalleryItem(ctx context.Context, id string) error

	// Pages
	GetPageByKey(ctx context.Context, key domain.PageKey) (*domain.Page, error)
	SavePage(ctx context.Context, page *domain.Page) error

	// Leads
	CreateLead(ctx context.Context, lead *domain.Lead) error
	ListLeads(ctx context.Context, status string, limit, offset int) ([]domain.Lead, int64, error)
	UpdateLeadStatus(ctx context.Context, id string, status domain.LeadStatus) error
	GetDashboardStats(ctx context.Context) (*domain.DashboardStats, error)

	// Media
	SaveMedia(ctx context.Context, media *domain.Media) error
	ListMedia(ctx context.Context, limit, offset int) ([]domain.Media, error)
	DeleteMedia(ctx context.Context, id string) error

	// Users / Profiles
	ListProfiles(ctx context.Context) ([]domain.Profile, error)
	UpdateProfileStatus(ctx context.Context, id string, isActive bool) error
}

// MemoryRepository provides robust fallback storage initialized with official seed data
type MemoryRepository struct {
	mu           sync.RWMutex
	settings     map[string]json.RawMessage
	services     []domain.Service
	projects     []domain.Project
	categories   []domain.Category
	posts        []domain.Post
	galleryItems []domain.GalleryItem
	pages        map[domain.PageKey]*domain.Page
	leads        []domain.Lead
	mediaList    []domain.Media
	profiles     []domain.Profile
}

func NewMemoryRepository() *MemoryRepository {
	repo := &MemoryRepository{
		settings:     make(map[string]json.RawMessage),
		pages:        make(map[domain.PageKey]*domain.Page),
		services:     make([]domain.Service, 0),
		projects:     make([]domain.Project, 0),
		categories:   make([]domain.Category, 0),
		posts:        make([]domain.Post, 0),
		galleryItems: make([]domain.GalleryItem, 0),
		leads:        make([]domain.Lead, 0),
		mediaList:    make([]domain.Media, 0),
		profiles:     make([]domain.Profile, 0),
	}
	repo.seedDefaults()
	return repo
}

// PostgresRepository executes SQL queries on Supabase Postgres
type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

// ConnectDB attempts to connect to Postgres, returns nil if connection cannot be established
func ConnectDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("gagal parse konfigurasi db: %w", err)
	}

	// Supabase PgBouncer (Transaction Pooler) compatibility:
	// Disable prepared statement cache to prevent "prepared statement already exists" (SQLSTATE 42P05)
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec
	config.ConnConfig.StatementCacheCapacity = 0
	config.ConnConfig.DescriptionCacheCapacity = 0

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat pool db: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("db ping gagal: %w", err)
	}

	log.Println("[DATABASE] Terhubung ke PostgreSQL / Supabase Database!")
	return pool, nil
}
