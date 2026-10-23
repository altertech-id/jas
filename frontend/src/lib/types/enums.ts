/**
 * Jaya Abadi Solusi (JAS) - Standardized Enums
 * Reference & Standardization across the frontend application
 */

export enum UserRole {
  ADMIN = 'admin',
  EDITOR = 'editor'
}

export enum LeadStatus {
  BARU = 'baru',
  DIHUBUNGI = 'dihubungi',
  SELESAI = 'selesai'
}

export enum PostStatus {
  DRAFT = 'draft',
  PUBLISHED = 'published'
}

export enum ServiceSlug {
  PEST_CONTROL = 'pest-control',
  RODENT_CONTROL = 'rodent-control',
  TERMITE_CONTROL = 'termite-control',
  FUMIGASI_PE = 'fumigasi-pe',
  DISINFEKSI = 'disinfeksi'
}

export enum PageKey {
  HOME = 'home',
  ABOUT_US = 'about-us',
  CONTACT = 'contact'
}

export enum SettingKey {
  GENERAL = 'general',
  CONTACT = 'contact',
  SOCIAL_MEDIA = 'social_media',
  SEO_DEFAULT = 'seo_default'
}

export enum ButtonVariant {
  PRIMARY = 'primary',
  SECONDARY = 'secondary',
  OUTLINE = 'outline',
  OUTLINE_ON_NAVY = 'outline-on-navy',
  WHATSAPP = 'whatsapp',
  GHOST = 'ghost'
}
