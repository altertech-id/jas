import type { 
  SiteSettings, 
  Service, 
  Project, 
  Category, 
  Post, 
  GalleryItem, 
  PageContent, 
  Lead, 
  DashboardStats, 
  Media,
  Profile
} from '../types';
import type { LeadStatus } from '../types/enums';

const API_BASE_URL = (import.meta.env.VITE_API_URL as string) || 
  (typeof window !== 'undefined' && window.location.hostname !== 'localhost'
    ? '' 
    : 'http://localhost:8080');

class ApiClient {
  private getAuthToken(): string | null {
    if (typeof window === 'undefined') return null;
    return localStorage.getItem('jas_auth_token');
  }

  private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const url = `${API_BASE_URL}${endpoint}`;
    const headers = new Headers(options.headers || {});

    if (!headers.has('Content-Type') && !(options.body instanceof FormData)) {
      headers.set('Content-Type', 'application/json');
    }

    const token = this.getAuthToken();
    if (token && !headers.has('Authorization')) {
      headers.set('Authorization', `Bearer ${token}`);
    }

    const response = await fetch(url, { ...options, headers });

    const contentType = response.headers.get('content-type') || '';
    let data: any;

    if (contentType.includes('application/json')) {
      data = await response.json();
    } else {
      const text = await response.text();
      if (!response.ok) {
        throw new Error(`Server error (${response.status}): ${text.slice(0, 100)}`);
      }
      data = text;
    }

    if (!response.ok) {
      throw new Error(data?.error || `Terjadi kesalahan sistem (${response.status}).`);
    }

    return data;
  }

  // --- PUBLIC ENDPOINTS ---

  async getSiteSettings(): Promise<SiteSettings> {
    const res = await this.request<{ data: Record<string, any> }>('/api/site-settings');
    return res.data as SiteSettings;
  }

  async getServices(): Promise<Service[]> {
    const res = await this.request<{ data: Service[] }>('/api/services');
    return res.data ?? [];
  }

  async getServiceBySlug(slug: string): Promise<Service> {
    const res = await this.request<{ data: Service }>(`/api/services/${slug}`);
    return res.data;
  }

  async getProjects(): Promise<Project[]> {
    const res = await this.request<{ data: Project[] }>('/api/projects');
    return res.data ?? [];
  }

  async getProjectBySlug(slug: string): Promise<Project> {
    const res = await this.request<{ data: Project }>(`/api/projects/${slug}`);
    return res.data;
  }

  async getCategories(): Promise<Category[]> {
    const res = await this.request<{ data: Category[] }>('/api/categories');
    return res.data ?? [];
  }

  async getPosts(params?: { category?: string; page?: number; limit?: number }): Promise<{ data: Post[]; total: number; page: number; limit: number }> {
    const query = new URLSearchParams();
    if (params?.category) query.set('category', params.category);
    if (params?.page) query.set('page', params.page.toString());
    if (params?.limit) query.set('limit', params.limit.toString());

    const qs = query.toString() ? `?${query.toString()}` : '';
    const res = await this.request<{ data: Post[]; total: number; page: number; limit: number }>(`/api/posts${qs}`);
    return { ...res, data: res.data ?? [] };
  }

  async getPostBySlug(slug: string): Promise<Post> {
    const res = await this.request<{ data: Post }>(`/api/posts/${slug}`);
    return res.data;
  }

  async getGallery(category?: string): Promise<GalleryItem[]> {
    const qs = category ? `?category=${encodeURIComponent(category)}` : '';
    const res = await this.request<{ data: GalleryItem[] }>(`/api/gallery${qs}`);
    return res.data ?? [];
  }

  async getPageContent(key: string): Promise<PageContent> {
    const res = await this.request<{ data: PageContent }>(`/api/pages/${key}`);
    return res.data;
  }

  async submitLead(payload: {
    name: string;
    phone: string;
    email?: string;
    service_id?: string;
    message?: string;
    source_page?: string;
    hp_field?: string;
  }): Promise<{ status: string; message: string; whatsapp_url: string }> {
    return this.request<{ status: string; message: string; whatsapp_url: string }>('/api/leads', {
      method: 'POST',
      body: JSON.stringify(payload)
    });
  }

  // --- ADMIN ENDPOINTS ---

  async getDashboardStats(): Promise<DashboardStats> {
    const res = await this.request<{ data: DashboardStats }>('/api/admin/dashboard/stats');
    return res.data;
  }

  async getLeads(params?: { status?: string; limit?: number; offset?: number }): Promise<{ data: Lead[]; total: number }> {
    const query = new URLSearchParams();
    if (params?.status) query.set('status', params.status);
    if (params?.limit) query.set('limit', params.limit.toString());
    if (params?.offset) query.set('offset', params.offset.toString());

    const qs = query.toString() ? `?${query.toString()}` : '';
    const res = await this.request<{ data: Lead[]; total: number }>(`/api/admin/leads${qs}`);
    return { ...res, data: res.data ?? [] };
  }

  async updateLeadStatus(id: string, status: LeadStatus): Promise<void> {
    await this.request(`/api/admin/leads/${id}/status`, {
      method: 'PATCH',
      body: JSON.stringify({ status })
    });
  }

  async uploadMedia(file: File, altText?: string): Promise<Media> {
    const formData = new FormData();
    formData.append('file', file);
    if (altText) formData.append('alt_text', altText);

    const res = await this.request<{ data: Media }>('/api/admin/media', {
      method: 'POST',
      body: formData
    });
    return res.data;
  }

  async getMediaList(): Promise<Media[]> {
    const res = await this.request<{ data: Media[] }>('/api/admin/media');
    return res.data ?? [];
  }

  async deleteMedia(id: string): Promise<void> {
    await this.request(`/api/admin/media/${id}`, { method: 'DELETE' });
  }

  async saveService(service: Partial<Service>): Promise<Service> {
    const method = service.id ? 'PUT' : 'POST';
    const endpoint = service.id ? `/api/admin/services/${service.id}` : '/api/admin/services';
    const res = await this.request<{ data: Service }>(endpoint, {
      method,
      body: JSON.stringify(service)
    });
    return res.data;
  }

  async deleteService(id: string): Promise<void> {
    await this.request(`/api/admin/services/${id}`, { method: 'DELETE' });
  }

  async saveProject(project: Partial<Project>): Promise<Project> {
    const method = project.id ? 'PUT' : 'POST';
    const endpoint = project.id ? `/api/admin/projects/${project.id}` : '/api/admin/projects';
    const res = await this.request<{ data: Project }>(endpoint, {
      method,
      body: JSON.stringify(project)
    });
    return res.data;
  }

  async deleteProject(id: string): Promise<void> {
    await this.request(`/api/admin/projects/${id}`, { method: 'DELETE' });
  }

  async savePost(post: Partial<Post>): Promise<Post> {
    const method = post.id ? 'PUT' : 'POST';
    const endpoint = post.id ? `/api/admin/posts/${post.id}` : '/api/admin/posts';
    const res = await this.request<{ data: Post }>(endpoint, {
      method,
      body: JSON.stringify(post)
    });
    return res.data;
  }

  async deletePost(id: string): Promise<void> {
    await this.request(`/api/admin/posts/${id}`, { method: 'DELETE' });
  }

  async createCategory(cat: { name: string; slug?: string }): Promise<Category> {
    const res = await this.request<{ data: Category }>('/api/admin/categories', {
      method: 'POST',
      body: JSON.stringify(cat)
    });
    return res.data;
  }

  async savePageContent(key: string, data: { title: string; content: any; seo_title: string; seo_description: string }): Promise<void> {
    await this.request(`/api/admin/pages/${key}`, {
      method: 'PUT',
      body: JSON.stringify(data)
    });
  }

  async saveGalleryItem(item: { media_id: string; caption: string; category: string; sort_order: number }): Promise<void> {
    await this.request('/api/admin/gallery', {
      method: 'POST',
      body: JSON.stringify(item)
    });
  }

  async deleteGalleryItem(id: string): Promise<void> {
    await this.request(`/api/admin/gallery/${id}`, { method: 'DELETE' });
  }

  async updateSiteSetting(key: string, value: any): Promise<void> {
    await this.request(`/api/admin/site-settings/${key}`, {
      method: 'PUT',
      body: JSON.stringify(value)
    });
  }

  async getUsers(): Promise<Profile[]> {
    const res = await this.request<{ data: Profile[] }>('/api/admin/users');
    return res.data ?? [];
  }

  async updateUserStatus(id: string, isActive: boolean): Promise<void> {
    await this.request(`/api/admin/users/${id}/status`, {
      method: 'PATCH',
      body: JSON.stringify({ is_active: isActive })
    });
  }
}

export const api = new ApiClient();
