package domain

// UserRole represents system roles for authentication and permission
type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleEditor UserRole = "editor"
)

func (r UserRole) IsValid() bool {
	return r == RoleAdmin || r == RoleEditor
}

// LeadStatus represents the follow-up status of lead submissions
type LeadStatus string

const (
	LeadStatusBaru      LeadStatus = "baru"
	LeadStatusDihubungi LeadStatus = "dihubungi"
	LeadStatusSelesai   LeadStatus = "selesai"
)

func (s LeadStatus) IsValid() bool {
	return s == LeadStatusBaru || s == LeadStatusDihubungi || s == LeadStatusSelesai
}

// PostStatus represents publication status of blog articles
type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
)

func (s PostStatus) IsValid() bool {
	return s == PostStatusDraft || s == PostStatusPublished
}

// ServiceSlug represents the standardized slugs for the 5 primary services
type ServiceSlug string

const (
	ServicePestControl    ServiceSlug = "pest-control"
	ServiceRodentControl  ServiceSlug = "rodent-control"
	ServiceTermiteControl ServiceSlug = "termite-control"
	ServiceFumigasiPE     ServiceSlug = "fumigasi-pe"
	ServiceDisinfeksi     ServiceSlug = "disinfeksi"
)

// PageKey represents unique keys for static CMS pages
type PageKey string

const (
	PageKeyHome    PageKey = "home"
	PageKeyAboutUs PageKey = "about-us"
	PageKeyContact PageKey = "contact"
)

// SettingKey represents the primary keys in site_settings table
type SettingKey string

const (
	SettingKeyGeneral     SettingKey = "general"
	SettingKeyContact     SettingKey = "contact"
	SettingKeySocialMedia SettingKey = "social_media"
	SettingKeySEODefault  SettingKey = "seo_default"
)
