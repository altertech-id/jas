package domain

import (
	"encoding/json"
	"time"
)

// Profile represents an admin or editor account
type Profile struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      UserRole  `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

// Media represents an uploaded file metadata in Supabase Storage
type Media struct {
	ID          string    `json:"id"`
	StoragePath string    `json:"storage_path"`
	PublicURL   string    `json:"public_url"`
	AltText     string    `json:"alt_text"`
	MimeType    string    `json:"mime_type"`
	SizeBytes   int64     `json:"size_bytes"`
	UploadedBy  *string   `json:"uploaded_by,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// Service represents a pest control service offering and landing page
type Service struct {
	ID               string          `json:"id"`
	Slug             string          `json:"slug"`
	Name             string          `json:"name"`
	Icon             string          `json:"icon"`
	ShortDescription string          `json:"short_description"`
	Content          json.RawMessage `json:"content"`
	HeroMediaID      *string         `json:"hero_media_id,omitempty"`
	HeroMedia        *Media          `json:"hero_media,omitempty"`
	SEOTitle         string          `json:"seo_title"`
	SEODescription   string          `json:"seo_description"`
	IsPublished      bool            `json:"is_published"`
	SortOrder        int             `json:"sort_order"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

// Project represents a completed field project case study
type Project struct {
	ID           string     `json:"id"`
	Slug         string     `json:"slug"`
	Title        string     `json:"title"`
	Location     string     `json:"location"`
	Description  string     `json:"description"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
	CoverMediaID *string    `json:"cover_media_id,omitempty"`
	CoverMedia   *Media     `json:"cover_media,omitempty"`
	ServiceID    *string    `json:"service_id,omitempty"`
	IsPublished  bool       `json:"is_published"`
	CreatedAt    time.Time  `json:"created_at"`
}

// GalleryItem represents an attachment documentation photo
type GalleryItem struct {
	ID        string    `json:"id"`
	MediaID   string    `json:"media_id"`
	Media     *Media    `json:"media,omitempty"`
	ProjectID *string   `json:"project_id,omitempty"`
	Caption   string    `json:"caption"`
	Category  string    `json:"category"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}

// Category represents a blog category
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Post represents a blog post
type Post struct {
	ID             string     `json:"id"`
	CategoryID     *string    `json:"category_id,omitempty"`
	Category       *Category  `json:"category,omitempty"`
	AuthorID       *string    `json:"author_id,omitempty"`
	Author         *Profile   `json:"author,omitempty"`
	CoverMediaID   *string    `json:"cover_media_id,omitempty"`
	CoverMedia     *Media     `json:"cover_media,omitempty"`
	Slug           string     `json:"slug"`
	Title          string     `json:"title"`
	Excerpt        string     `json:"excerpt"`
	ContentHTML    string     `json:"content_html"`
	Status         PostStatus `json:"status"`
	SEOTitle       string     `json:"seo_title"`
	SEODescription string     `json:"seo_description"`
	PublishedAt    *time.Time `json:"published_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// Page represents a CMS managed static page
type Page struct {
	ID             string          `json:"id"`
	Key            PageKey         `json:"key"`
	Title          string          `json:"title"`
	Content        json.RawMessage `json:"content"`
	SEOTitle       string          `json:"seo_title"`
	SEODescription string          `json:"seo_description"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// Lead represents a quotation submission
type Lead struct {
	ID         string     `json:"id"`
	ServiceID  *string    `json:"service_id,omitempty"`
	Service    *Service   `json:"service,omitempty"`
	Name       string     `json:"name"`
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	Message    string     `json:"message"`
	SourcePage string     `json:"source_page"`
	Status     LeadStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
}

// SiteSetting represents global key-value configuration
type SiteSetting struct {
	Key       SettingKey      `json:"key"`
	Value     json.RawMessage `json:"value"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// Request / Response DTOs

type CreateLeadRequest struct {
	Name       string  `json:"name"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	ServiceID  *string `json:"service_id"`
	Message    string  `json:"message"`
	SourcePage string  `json:"source_page"`
	Honeypot   string  `json:"hp_field,omitempty"` // Honeypot anti-spam
}

type UpdateLeadStatusRequest struct {
	Status LeadStatus `json:"status"`
}

type SavePostRequest struct {
	CategoryID     *string    `json:"category_id"`
	CoverMediaID   *string    `json:"cover_media_id"`
	Slug           string     `json:"slug"`
	Title          string     `json:"title"`
	Excerpt        string     `json:"excerpt"`
	ContentHTML    string     `json:"content_html"`
	Status         PostStatus `json:"status"`
	SEOTitle       string     `json:"seo_title"`
	SEODescription string     `json:"seo_description"`
}

type SaveServiceRequest struct {
	Slug             string          `json:"slug"`
	Name             string          `json:"name"`
	Icon             string          `json:"icon"`
	ShortDescription string          `json:"short_description"`
	Content          json.RawMessage `json:"content"`
	HeroMediaID      *string         `json:"hero_media_id"`
	SEOTitle         string          `json:"seo_title"`
	SEODescription   string          `json:"seo_description"`
	IsPublished      bool            `json:"is_published"`
	SortOrder        int             `json:"sort_order"`
}

type SaveProjectRequest struct {
	Slug         string  `json:"slug"`
	Title        string  `json:"title"`
	Location     string  `json:"location"`
	Description  string  `json:"description"`
	CompletedAt  *string `json:"completed_at"`
	CoverMediaID *string `json:"cover_media_id"`
	ServiceID    *string `json:"service_id"`
	IsPublished  bool    `json:"is_published"`
}

type SavePageRequest struct {
	Title          string          `json:"title"`
	Content        json.RawMessage `json:"content"`
	SEOTitle       string          `json:"seo_title"`
	SEODescription string          `json:"seo_description"`
}

type DashboardStats struct {
	NewLeadsMonth   int64 `json:"new_leads_month"`
	TotalLeadsMonth int64 `json:"total_leads_month"`
	PublishedPosts  int64 `json:"published_posts"`
	PublishedProjects int64 `json:"published_projects"`
}
