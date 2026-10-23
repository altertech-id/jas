import type { UserRole, LeadStatus, PostStatus, ServiceSlug, PageKey, SettingKey } from './enums';

export interface Profile {
  id: string;
  email: string;
  name: string;
  role: UserRole;
  is_active: boolean;
  created_at: string;
}

export interface Media {
  id: string;
  storage_path: string;
  public_url: string;
  alt_text: string;
  mime_type: string;
  size_bytes: number;
  uploaded_by?: string;
  created_at: string;
}

export interface ServiceProblemSolution {
  problems: string[];
  solutions: string[];
  methods: { title: string; desc: string }[];
  benefits: string[];
  faq?: { q: string; a: string }[];
  hero_badge?: string;
  headline?: string;
  tagline?: string;
}

export interface Service {
  id: string;
  slug: string;
  name: string;
  icon: string;
  short_description: string;
  content: ServiceProblemSolution;
  hero_media_id?: string;
  hero_media?: Media;
  seo_title: string;
  seo_description: string;
  is_published: boolean;
  sort_order: number;
  created_at: string;
  updated_at: string;
}

export interface Project {
  id: string;
  slug: string;
  title: string;
  location: string;
  description: string;
  completed_at?: string;
  cover_media_id?: string;
  cover_media?: Media;
  service_id?: string;
  is_published: boolean;
  created_at: string;
}

export interface Category {
  id: string;
  name: string;
  slug: string;
}

export interface Post {
  id: string;
  category_id?: string;
  category?: Category;
  author_id?: string;
  author?: Profile;
  cover_media_id?: string;
  cover_media?: Media;
  slug: string;
  title: string;
  excerpt: string;
  content_html: string;
  status: PostStatus;
  seo_title: string;
  seo_description: string;
  published_at?: string;
  created_at: string;
  updated_at: string;
}

export interface GalleryItem {
  id: string;
  media_id: string;
  media?: Media;
  project_id?: string;
  caption: string;
  category: string;
  sort_order: number;
  created_at: string;
}

export interface PageContent {
  id: string;
  key: PageKey;
  title: string;
  content: Record<string, any>;
  seo_title: string;
  seo_description: string;
  updated_at: string;
}

export interface Lead {
  id: string;
  service_id?: string;
  service?: Service;
  name: string;
  phone: string;
  email: string;
  message: string;
  source_page: string;
  status: LeadStatus;
  created_at: string;
}

export interface SiteSettings {
  general?: {
    company_name: string;
    short_name: string;
    tagline: string;
    established_year: number;
    years_experience: string;
    verified_badge: string;
  };
  contact?: {
    phone: string;
    phone_display: string;
    whatsapp: string;
    whatsapp_url: string;
    email: string;
    address: string;
    working_hours: string;
    emergency_service: string;
    maps_embed_url: string;
  };
  social_media?: {
    instagram: string;
    facebook: string;
    linkedin: string;
  };
  seo_default?: {
    title: string;
    description: string;
    og_image: string;
  };
}

export interface DashboardStats {
  new_leads_month: number;
  total_leads_month: number;
  published_posts: number;
  published_projects: number;
}
