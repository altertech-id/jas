package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"jayaabadisolusi/backend/internal/domain"
	"jayaabadisolusi/backend/internal/middleware"
	"jayaabadisolusi/backend/internal/repository"
	"jayaabadisolusi/backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	repo      repository.Repository
	storage   *service.StorageService
	sanitizer *service.Sanitizer
}

func NewAdminHandler(repo repository.Repository, storage *service.StorageService, sanitizer *service.Sanitizer) *AdminHandler {
	return &AdminHandler{
		repo:      repo,
		storage:   storage,
		sanitizer: sanitizer,
	}
}

// GetDashboardStats returns high-level stats for dashboard cards
func (h *AdminHandler) GetDashboardStats(c *fiber.Ctx) error {
	stats, err := h.repo.GetDashboardStats(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": stats})
}

// ListLeads lists leads with optional status filter
func (h *AdminHandler) ListLeads(c *fiber.Ctx) error {
	status := c.Query("status")
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	leads, total, err := h.repo.ListLeads(c.Context(), status, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data":   leads,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

// UpdateLeadStatus changes lead status
func (h *AdminHandler) UpdateLeadStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var req domain.UpdateLeadStatusRequest
	if err := c.BodyParser(&req); err != nil || !req.Status.IsValid() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Status lead tidak valid (baru, dihubungi, selesai)."})
	}

	if err := h.repo.UpdateLeadStatus(c.Context(), id, req.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Status lead berhasil diperbarui."})
}

// UploadMedia handles single image upload
func (h *AdminHandler) UploadMedia(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File gambar wajib diunggah."})
	}

	altText := c.FormValue("alt_text", "")
	profile, _ := c.Locals(middleware.UserContextKey).(*domain.Profile)
	var uploaderID *string
	if profile != nil {
		uploaderID = &profile.ID
	}

	media, err := h.storage.UploadImage(c.Context(), file, altText, uploaderID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": media})
}

// ListMedia lists media library files
func (h *AdminHandler) ListMedia(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "30"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	items, err := h.repo.ListMedia(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

// DeleteMedia removes media
func (h *AdminHandler) DeleteMedia(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.repo.DeleteMedia(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Media berhasil dihapus."})
}

// SaveService creates or updates a service
func (h *AdminHandler) SaveService(c *fiber.Ctx) error {
	var req domain.SaveServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid."})
	}

	req.Slug = strings.TrimSpace(req.Slug)
	req.Name = strings.TrimSpace(req.Name)
	if req.Slug == "" || req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nama dan slug layanan wajib diisi."})
	}

	service := &domain.Service{
		ID:               c.Params("id"),
		Slug:             req.Slug,
		Name:             req.Name,
		Icon:             req.Icon,
		ShortDescription: req.ShortDescription,
		Content:          req.Content,
		HeroMediaID:      req.HeroMediaID,
		SEOTitle:         req.SEOTitle,
		SEODescription:   req.SEODescription,
		IsPublished:      req.IsPublished,
		SortOrder:        req.SortOrder,
		UpdatedAt:        time.Now(),
	}

	if err := h.repo.SaveService(c.Context(), service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": service, "message": "Layanan berhasil disimpan."})
}

// DeleteService removes a service
func (h *AdminHandler) DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.repo.DeleteService(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Layanan berhasil dihapus."})
}

// SaveProject creates or updates a project
func (h *AdminHandler) SaveProject(c *fiber.Ctx) error {
	var req domain.SaveProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid."})
	}

	req.Slug = strings.TrimSpace(req.Slug)
	req.Title = strings.TrimSpace(req.Title)
	if req.Slug == "" || req.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Judul dan slug proyek wajib diisi."})
	}

	var completedAt *time.Time
	if req.CompletedAt != nil && *req.CompletedAt != "" {
		parsed, err := time.Parse("2006-01-02", *req.CompletedAt)
		if err == nil {
			completedAt = &parsed
		}
	}

	project := &domain.Project{
		ID:           c.Params("id"),
		Slug:         req.Slug,
		Title:        req.Title,
		Location:     req.Location,
		Description:  req.Description,
		CompletedAt:  completedAt,
		CoverMediaID: req.CoverMediaID,
		ServiceID:    req.ServiceID,
		IsPublished:  req.IsPublished,
	}

	if err := h.repo.SaveProject(c.Context(), project); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": project, "message": "Proyek berhasil disimpan."})
}

// DeleteProject removes a project
func (h *AdminHandler) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.repo.DeleteProject(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Proyek berhasil dihapus."})
}

// SavePost creates or updates a blog post (sanitizes HTML)
func (h *AdminHandler) SavePost(c *fiber.Ctx) error {
	var req domain.SavePostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid."})
	}

	req.Slug = strings.TrimSpace(req.Slug)
	req.Title = strings.TrimSpace(req.Title)
	if req.Slug == "" || req.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Judul dan slug artikel wajib diisi."})
	}

	// Sanitize HTML using bluemonday policy
	cleanHTML := h.sanitizer.SanitizeHTML(req.ContentHTML)

	profile, _ := c.Locals(middleware.UserContextKey).(*domain.Profile)
	var authorID *string
	if profile != nil {
		authorID = &profile.ID
	}

	now := time.Now()
	var publishedAt *time.Time
	if req.Status == domain.PostStatusPublished {
		publishedAt = &now
	}

	post := &domain.Post{
		ID:             c.Params("id"),
		CategoryID:     req.CategoryID,
		AuthorID:       authorID,
		CoverMediaID:   req.CoverMediaID,
		Slug:           req.Slug,
		Title:          req.Title,
		Excerpt:        req.Excerpt,
		ContentHTML:    cleanHTML,
		Status:         req.Status,
		SEOTitle:       req.SEOTitle,
		SEODescription: req.SEODescription,
		PublishedAt:    publishedAt,
		UpdatedAt:      now,
	}

	if err := h.repo.SavePost(c.Context(), post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": post, "message": "Artikel berhasil disimpan."})
}

// DeletePost removes a blog post
func (h *AdminHandler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.repo.DeletePost(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Artikel berhasil dihapus."})
}

// CreateCategory adds a category
func (h *AdminHandler) CreateCategory(c *fiber.Ctx) error {
	var cat domain.Category
	if err := c.BodyParser(&cat); err != nil || strings.TrimSpace(cat.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nama kategori wajib diisi."})
	}
	if strings.TrimSpace(cat.Slug) == "" {
		cat.Slug = strings.ToLower(strings.ReplaceAll(cat.Name, " ", "-"))
	}
	if err := h.repo.CreateCategory(c.Context(), &cat); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": cat})
}

// SavePage updates static CMS page
func (h *AdminHandler) SavePage(c *fiber.Ctx) error {
	key := domain.PageKey(c.Params("key"))
	var req domain.SavePageRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid."})
	}

	page := &domain.Page{
		Key:            key,
		Title:          req.Title,
		Content:        req.Content,
		SEOTitle:       req.SEOTitle,
		SEODescription: req.SEODescription,
		UpdatedAt:      time.Now(),
	}

	if err := h.repo.SavePage(c.Context(), page); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": page, "message": "Konten halaman berhasil disimpan."})
}

// SaveGalleryItem adds gallery item
func (h *AdminHandler) SaveGalleryItem(c *fiber.Ctx) error {
	var item domain.GalleryItem
	if err := c.BodyParser(&item); err != nil || item.MediaID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Media ID wajib diisi."})
	}
	if err := h.repo.SaveGalleryItem(c.Context(), &item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": item, "message": "Foto galeri berhasil disimpan."})
}

// DeleteGalleryItem removes gallery item
func (h *AdminHandler) DeleteGalleryItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.repo.DeleteGalleryItem(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Foto galeri berhasil dihapus."})
}

// UpdateSiteSetting updates global setting
func (h *AdminHandler) UpdateSiteSetting(c *fiber.Ctx) error {
	key := c.Params("key")
	var raw json.RawMessage = c.Body()
	if len(raw) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nilai pengaturan tidak boleh kosong."})
	}

	if err := h.repo.UpdateSiteSetting(c.Context(), key, raw); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": fmt.Sprintf("Pengaturan '%s' berhasil diperbarui.", key)})
}

// ListProfiles lists users (Admin role only)
func (h *AdminHandler) ListProfiles(c *fiber.Ctx) error {
	profiles, err := h.repo.ListProfiles(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": profiles})
}

// UpdateProfileStatus toggles user status (Admin role only)
func (h *AdminHandler) UpdateProfileStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var req struct {
		IsActive bool `json:"is_active"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid."})
	}

	if err := h.repo.UpdateProfileStatus(c.Context(), id, req.IsActive); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Status pengguna berhasil diperbarui."})
}
