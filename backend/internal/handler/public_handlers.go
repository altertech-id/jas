package handler

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"jayaabadisolusi/backend/internal/domain"
	"jayaabadisolusi/backend/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type PublicHandler struct {
	repo repository.Repository
}

func NewPublicHandler(repo repository.Repository) *PublicHandler {
	return &PublicHandler{repo: repo}
}

// GetSiteSettings returns public global settings (contact, general, seo)
func (h *PublicHandler) GetSiteSettings(c *fiber.Ctx) error {
	settings, err := h.repo.GetSiteSettings(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": settings})
}

// ListServices returns all published services
func (h *PublicHandler) ListServices(c *fiber.Ctx) error {
	services, err := h.repo.ListServices(c.Context(), true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": services})
}

// GetServiceBySlug returns a single service by slug
func (h *PublicHandler) GetServiceBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	service, err := h.repo.GetServiceBySlug(c.Context(), slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Layanan tidak ditemukan."})
	}
	return c.JSON(fiber.Map{"data": service})
}

// ListProjects returns all published projects
func (h *PublicHandler) ListProjects(c *fiber.Ctx) error {
	projects, err := h.repo.ListProjects(c.Context(), true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": projects})
}

// GetProjectBySlug returns a single project
func (h *PublicHandler) GetProjectBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	project, err := h.repo.GetProjectBySlug(c.Context(), slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proyek tidak ditemukan."})
	}
	return c.JSON(fiber.Map{"data": project})
}

// ListCategories returns blog categories
func (h *PublicHandler) ListCategories(c *fiber.Ctx) error {
	categories, err := h.repo.ListCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": categories})
}

// ListPosts returns published posts with category and pagination
func (h *PublicHandler) ListPosts(c *fiber.Ctx) error {
	category := c.Query("category")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit", "9"))
	if limit < 1 || limit > 50 {
		limit = 9
	}

	posts, total, err := h.repo.ListPosts(c.Context(), category, string(domain.PostStatusPublished), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data":  posts,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// GetPostBySlug returns a single published post
func (h *PublicHandler) GetPostBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	post, err := h.repo.GetPostBySlug(c.Context(), slug)
	if err != nil || post.Status != domain.PostStatusPublished {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artikel tidak ditemukan."})
	}
	return c.JSON(fiber.Map{"data": post})
}

// ListGallery returns gallery items
func (h *PublicHandler) ListGallery(c *fiber.Ctx) error {
	cat := c.Query("category")
	items, err := h.repo.ListGalleryItems(c.Context(), cat)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

// GetPageContent returns CMS block data for home, about-us, or contact
func (h *PublicHandler) GetPageContent(c *fiber.Ctx) error {
	key := domain.PageKey(c.Params("key"))
	page, err := h.repo.GetPageByKey(c.Context(), key)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Halaman tidak ditemukan."})
	}
	return c.JSON(fiber.Map{"data": page})
}

// SubmitLead handles quotation submission
func (h *PublicHandler) SubmitLead(c *fiber.Ctx) error {
	var req domain.CreateLeadRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid."})
	}

	// Honeypot check: bot prevention
	if req.Honeypot != "" {
		// Silently drop bot submission
		return c.JSON(fiber.Map{"status": "success", "message": "Permintaan berhasil terkirim."})
	}

	// Validate required fields
	req.Name = strings.TrimSpace(req.Name)
	req.Phone = strings.TrimSpace(req.Phone)

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nama lengkap wajib diisi."})
	}
	if req.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nomor telepon/WhatsApp wajib diisi."})
	}
	if len(req.Phone) < 9 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nomor telepon minimal 9 digit."})
	}

	lead := &domain.Lead{
		ServiceID:  req.ServiceID,
		Name:       req.Name,
		Phone:      req.Phone,
		Email:      strings.TrimSpace(req.Email),
		Message:    strings.TrimSpace(req.Message),
		SourcePage: strings.TrimSpace(req.SourcePage),
	}

	if err := h.repo.CreateLead(c.Context(), lead); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan permintaan. Silakan hubungi kami via WhatsApp."})
	}

	// Generate WhatsApp click-to-chat URL with prefilled text for instant follow-up
	waPhone := "6281399484916"
	presetMsg := fmt.Sprintf("Halo Jaya Abadi Solusi, saya %s baru saja mengirimkan permintaan survei/penawaran melalui website. Mohon informasinya.", req.Name)
	waURL := fmt.Sprintf("https://wa.me/%s?text=%s", waPhone, url.QueryEscape(presetMsg))

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":       "success",
		"message":      "Terima kasih! Permintaan Anda telah diterima oleh tim Jaya Abadi Solusi.",
		"whatsapp_url": waURL,
	})
}
