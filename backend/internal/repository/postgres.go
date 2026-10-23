package repository

import (
	"context"
	"encoding/json"

	"jayaabadisolusi/backend/internal/domain"
)

// PostgresRepository method implementations querying Supabase Postgres

func (p *PostgresRepository) GetSiteSettings(ctx context.Context) (map[string]json.RawMessage, error) {
	rows, err := p.pool.Query(ctx, "SELECT key, value FROM site_settings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[string]json.RawMessage)
	for rows.Next() {
		var k string
		var v json.RawMessage
		if err := rows.Scan(&k, &v); err == nil {
			res[k] = v
		}
	}
	return res, nil
}

func (p *PostgresRepository) GetSiteSetting(ctx context.Context, key string) (json.RawMessage, error) {
	var val json.RawMessage
	err := p.pool.QueryRow(ctx, "SELECT value FROM site_settings WHERE key = $1", key).Scan(&val)
	return val, err
}

func (p *PostgresRepository) UpdateSiteSetting(ctx context.Context, key string, value json.RawMessage) error {
	_, err := p.pool.Exec(ctx, "INSERT INTO site_settings (key, value, updated_at) VALUES ($1, $2, now()) ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value, updated_at = now()", key, value)
	return err
}

func (p *PostgresRepository) ListServices(ctx context.Context, onlyPublished bool) ([]domain.Service, error) {
	query := "SELECT id, slug, name, icon, short_description, content, seo_title, seo_description, is_published, sort_order, created_at, updated_at FROM services"
	if onlyPublished {
		query += " WHERE is_published = true ORDER BY sort_order ASC"
	} else {
		query += " ORDER BY sort_order ASC"
	}

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []domain.Service
	for rows.Next() {
		var s domain.Service
		err := rows.Scan(&s.ID, &s.Slug, &s.Name, &s.Icon, &s.ShortDescription, &s.Content, &s.SEOTitle, &s.SEODescription, &s.IsPublished, &s.SortOrder, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

func (p *PostgresRepository) GetServiceBySlug(ctx context.Context, slug string) (*domain.Service, error) {
	var s domain.Service
	err := p.pool.QueryRow(ctx, "SELECT id, slug, name, icon, short_description, content, seo_title, seo_description, is_published, sort_order, created_at, updated_at FROM services WHERE slug = $1", slug).
		Scan(&s.ID, &s.Slug, &s.Name, &s.Icon, &s.ShortDescription, &s.Content, &s.SEOTitle, &s.SEODescription, &s.IsPublished, &s.SortOrder, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (p *PostgresRepository) SaveService(ctx context.Context, service *domain.Service) error {
	if service.ID != "" {
		_, err := p.pool.Exec(ctx, "UPDATE services SET slug=$1, name=$2, icon=$3, short_description=$4, content=$5, seo_title=$6, seo_description=$7, is_published=$8, sort_order=$9, updated_at=now() WHERE id=$10",
			service.Slug, service.Name, service.Icon, service.ShortDescription, service.Content, service.SEOTitle, service.SEODescription, service.IsPublished, service.SortOrder, service.ID)
		return err
	}
	err := p.pool.QueryRow(ctx, "INSERT INTO services (slug, name, icon, short_description, content, seo_title, seo_description, is_published, sort_order, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now()) RETURNING id",
		service.Slug, service.Name, service.Icon, service.ShortDescription, service.Content, service.SEOTitle, service.SEODescription, service.IsPublished, service.SortOrder).Scan(&service.ID)
	return err
}

func (p *PostgresRepository) DeleteService(ctx context.Context, id string) error {
	_, err := p.pool.Exec(ctx, "DELETE FROM services WHERE id = $1", id)
	return err
}

func (p *PostgresRepository) ListProjects(ctx context.Context, onlyPublished bool) ([]domain.Project, error) {
	query := "SELECT id, slug, title, location, description, completed_at, is_published, created_at FROM projects"
	if onlyPublished {
		query += " WHERE is_published = true ORDER BY created_at DESC"
	} else {
		query += " ORDER BY created_at DESC"
	}

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []domain.Project
	for rows.Next() {
		var pr domain.Project
		if err := rows.Scan(&pr.ID, &pr.Slug, &pr.Title, &pr.Location, &pr.Description, &pr.CompletedAt, &pr.IsPublished, &pr.CreatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, pr)
	}
	return projects, nil
}

func (p *PostgresRepository) GetProjectBySlug(ctx context.Context, slug string) (*domain.Project, error) {
	var pr domain.Project
	err := p.pool.QueryRow(ctx, "SELECT id, slug, title, location, description, completed_at, is_published, created_at FROM projects WHERE slug = $1", slug).
		Scan(&pr.ID, &pr.Slug, &pr.Title, &pr.Location, &pr.Description, &pr.CompletedAt, &pr.IsPublished, &pr.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

func (p *PostgresRepository) SaveProject(ctx context.Context, project *domain.Project) error {
	if project.ID != "" {
		_, err := p.pool.Exec(ctx, "UPDATE projects SET slug=$1, title=$2, location=$3, description=$4, completed_at=$5, is_published=$6 WHERE id=$7",
			project.Slug, project.Title, project.Location, project.Description, project.CompletedAt, project.IsPublished, project.ID)
		return err
	}
	err := p.pool.QueryRow(ctx, "INSERT INTO projects (slug, title, location, description, completed_at, is_published) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		project.Slug, project.Title, project.Location, project.Description, project.CompletedAt, project.IsPublished).Scan(&project.ID)
	return err
}

func (p *PostgresRepository) DeleteProject(ctx context.Context, id string) error {
	_, err := p.pool.Exec(ctx, "DELETE FROM projects WHERE id = $1", id)
	return err
}

func (p *PostgresRepository) ListCategories(ctx context.Context) ([]domain.Category, error) {
	rows, err := p.pool.Query(ctx, "SELECT id, name, slug FROM categories ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []domain.Category
	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug); err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, nil
}

func (p *PostgresRepository) CreateCategory(ctx context.Context, cat *domain.Category) error {
	return p.pool.QueryRow(ctx, "INSERT INTO categories (name, slug) VALUES ($1, $2) RETURNING id", cat.Name, cat.Slug).Scan(&cat.ID)
}

func (p *PostgresRepository) ListPosts(ctx context.Context, categorySlug string, status string, page, limit int) ([]domain.Post, int64, error) {
	offset := (page - 1) * limit
	var total int64
	_ = p.pool.QueryRow(ctx, "SELECT count(*) FROM posts").Scan(&total)

	query := "SELECT id, slug, title, excerpt, content_html, status, seo_title, seo_description, published_at, created_at, updated_at FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	rows, err := p.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		if err := rows.Scan(&post.ID, &post.Slug, &post.Title, &post.Excerpt, &post.ContentHTML, &post.Status, &post.SEOTitle, &post.SEODescription, &post.PublishedAt, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, 0, err
		}
		posts = append(posts, post)
	}
	return posts, total, nil
}

func (p *PostgresRepository) GetPostBySlug(ctx context.Context, slug string) (*domain.Post, error) {
	var post domain.Post
	err := p.pool.QueryRow(ctx, "SELECT id, slug, title, excerpt, content_html, status, seo_title, seo_description, published_at, created_at, updated_at FROM posts WHERE slug = $1", slug).
		Scan(&post.ID, &post.Slug, &post.Title, &post.Excerpt, &post.ContentHTML, &post.Status, &post.SEOTitle, &post.SEODescription, &post.PublishedAt, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PostgresRepository) SavePost(ctx context.Context, post *domain.Post) error {
	if post.ID != "" {
		_, err := p.pool.Exec(ctx, "UPDATE posts SET slug=$1, title=$2, excerpt=$3, content_html=$4, status=$5, seo_title=$6, seo_description=$7, published_at=$8, updated_at=now() WHERE id=$9",
			post.Slug, post.Title, post.Excerpt, post.ContentHTML, post.Status, post.SEOTitle, post.SEODescription, post.PublishedAt, post.ID)
		return err
	}
	return p.pool.QueryRow(ctx, "INSERT INTO posts (slug, title, excerpt, content_html, status, seo_title, seo_description, published_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now()) RETURNING id",
		post.Slug, post.Title, post.Excerpt, post.ContentHTML, post.Status, post.SEOTitle, post.SEODescription, post.PublishedAt).Scan(&post.ID)
}

func (p *PostgresRepository) DeletePost(ctx context.Context, id string) error {
	_, err := p.pool.Exec(ctx, "DELETE FROM posts WHERE id = $1", id)
	return err
}

func (p *PostgresRepository) ListGalleryItems(ctx context.Context, category string) ([]domain.GalleryItem, error) {
	query := `SELECT g.id, g.media_id, g.caption, g.category, g.sort_order, g.created_at, 
              m.id, m.public_url, m.alt_text, m.mime_type
              FROM gallery_items g 
              LEFT JOIN media m ON g.media_id = m.id 
              ORDER BY g.sort_order ASC`
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.GalleryItem
	for rows.Next() {
		var item domain.GalleryItem
		var med domain.Media
		if err := rows.Scan(&item.ID, &item.MediaID, &item.Caption, &item.Category, &item.SortOrder, &item.CreatedAt, &med.ID, &med.PublicURL, &med.AltText, &med.MimeType); err == nil {
			item.Media = &med
			items = append(items, item)
		}
	}
	return items, nil
}

func (p *PostgresRepository) SaveGalleryItem(ctx context.Context, item *domain.GalleryItem) error {
	return p.pool.QueryRow(ctx, "INSERT INTO gallery_items (media_id, caption, category, sort_order) VALUES ($1, $2, $3, $4) RETURNING id",
		item.MediaID, item.Caption, item.Category, item.SortOrder).Scan(&item.ID)
}

func (p *PostgresRepository) DeleteGalleryItem(ctx context.Context, id string) error {
	_, err := p.pool.Exec(ctx, "DELETE FROM gallery_items WHERE id = $1", id)
	return err
}

func (p *PostgresRepository) GetPageByKey(ctx context.Context, key domain.PageKey) (*domain.Page, error) {
	var page domain.Page
	err := p.pool.QueryRow(ctx, "SELECT id, key, title, content, seo_title, seo_description, updated_at FROM pages WHERE key = $1", string(key)).
		Scan(&page.ID, &page.Key, &page.Title, &page.Content, &page.SEOTitle, &page.SEODescription, &page.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (p *PostgresRepository) SavePage(ctx context.Context, page *domain.Page) error {
	_, err := p.pool.Exec(ctx, "INSERT INTO pages (key, title, content, seo_title, seo_description, updated_at) VALUES ($1, $2, $3, $4, $5, now()) ON CONFLICT (key) DO UPDATE SET title=EXCLUDED.title, content=EXCLUDED.content, seo_title=EXCLUDED.seo_title, seo_description=EXCLUDED.seo_description, updated_at=now()",
		string(page.Key), page.Title, page.Content, page.SEOTitle, page.SEODescription)
	return err
}

func (p *PostgresRepository) CreateLead(ctx context.Context, lead *domain.Lead) error {
	return p.pool.QueryRow(ctx, "INSERT INTO leads (service_id, name, phone, email, message, source_page, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at",
		lead.ServiceID, lead.Name, lead.Phone, lead.Email, lead.Message, lead.SourcePage, string(domain.LeadStatusBaru)).Scan(&lead.ID, &lead.CreatedAt)
}

func (p *PostgresRepository) ListLeads(ctx context.Context, status string, limit, offset int) ([]domain.Lead, int64, error) {
	var total int64
	_ = p.pool.QueryRow(ctx, "SELECT count(*) FROM leads").Scan(&total)

	query := "SELECT id, service_id, name, phone, email, message, source_page, status, created_at FROM leads"
	var rowsArgs []interface{}
	if status != "" {
		query += " WHERE status = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3"
		rowsArgs = append(rowsArgs, status, limit, offset)
	} else {
		query += " ORDER BY created_at DESC LIMIT $1 OFFSET $2"
		rowsArgs = append(rowsArgs, limit, offset)
	}

	rows, err := p.pool.Query(ctx, query, rowsArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var leads []domain.Lead
	for rows.Next() {
		var l domain.Lead
		var statStr string
		if err := rows.Scan(&l.ID, &l.ServiceID, &l.Name, &l.Phone, &l.Email, &l.Message, &l.SourcePage, &statStr, &l.CreatedAt); err != nil {
			return nil, 0, err
		}
		l.Status = domain.LeadStatus(statStr)
		leads = append(leads, l)
	}
	return leads, total, nil
}

func (p *PostgresRepository) UpdateLeadStatus(ctx context.Context, id string, status domain.LeadStatus) error {
	_, err := p.pool.Exec(ctx, "UPDATE leads SET status = $1 WHERE id = $2", string(status), id)
	return err
}

func (p *PostgresRepository) GetDashboardStats(ctx context.Context) (*domain.DashboardStats, error) {
	var stats domain.DashboardStats
	_ = p.pool.QueryRow(ctx, "SELECT count(*) FROM leads WHERE status = 'baru'").Scan(&stats.NewLeadsMonth)
	_ = p.pool.QueryRow(ctx, "SELECT count(*) FROM leads").Scan(&stats.TotalLeadsMonth)
	_ = p.pool.QueryRow(ctx, "SELECT count(*) FROM posts WHERE status = 'published'").Scan(&stats.PublishedPosts)
	_ = p.pool.QueryRow(ctx, "SELECT count(*) FROM projects WHERE is_published = true").Scan(&stats.PublishedProjects)
	return &stats, nil
}

func (p *PostgresRepository) SaveMedia(ctx context.Context, media *domain.Media) error {
	return p.pool.QueryRow(ctx, "INSERT INTO media (storage_path, public_url, alt_text, mime_type, size_bytes, uploaded_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at",
		media.StoragePath, media.PublicURL, media.AltText, media.MimeType, media.SizeBytes, media.UploadedBy).Scan(&media.ID, &media.CreatedAt)
}

func (p *PostgresRepository) ListMedia(ctx context.Context, limit, offset int) ([]domain.Media, error) {
	rows, err := p.pool.Query(ctx, "SELECT id, storage_path, public_url, alt_text, mime_type, size_bytes, uploaded_by, created_at FROM media ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaList []domain.Media
	for rows.Next() {
		var m domain.Media
		if err := rows.Scan(&m.ID, &m.StoragePath, &m.PublicURL, &m.AltText, &m.MimeType, &m.SizeBytes, &m.UploadedBy, &m.CreatedAt); err == nil {
			mediaList = append(mediaList, m)
		}
	}
	return mediaList, nil
}

func (p *PostgresRepository) DeleteMedia(ctx context.Context, id string) error {
	_, err := p.pool.Exec(ctx, "DELETE FROM media WHERE id = $1", id)
	return err
}

func (p *PostgresRepository) ListProfiles(ctx context.Context) ([]domain.Profile, error) {
	rows, err := p.pool.Query(ctx, "SELECT id, email, name, role, is_active, created_at FROM profiles ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []domain.Profile
	for rows.Next() {
		var prof domain.Profile
		var r string
		if err := rows.Scan(&prof.ID, &prof.Email, &prof.Name, &r, &prof.IsActive, &prof.CreatedAt); err == nil {
			prof.Role = domain.UserRole(r)
			profiles = append(profiles, prof)
		}
	}
	return profiles, nil
}

func (p *PostgresRepository) UpdateProfileStatus(ctx context.Context, id string, isActive bool) error {
	_, err := p.pool.Exec(ctx, "UPDATE profiles SET is_active = $1 WHERE id = $2", isActive, id)
	return err
}

// Compile-time check that PostgresRepository implements Repository
var _ Repository = (*PostgresRepository)(nil)
