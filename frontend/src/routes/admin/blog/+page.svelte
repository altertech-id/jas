<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Post, Category } from '$lib/types';
  import { PostStatus } from '$lib/types/enums';
  import TipTapEditor from '$lib/components/admin/TipTapEditor.svelte';
  import StatusPill from '$lib/components/ui/StatusPill.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { Plus, Edit, Trash2, Calendar, ArrowLeft } from 'lucide-svelte';

  let posts: Post[] = [];
  let categories: Category[] = [];
  let loading = true;
  let activeTab: 'list' | 'edit' = 'list';

  // Form edit state
  let editId: string | undefined = undefined;
  let title = '';
  let slug = '';
  let excerpt = '';
  let contentHtml = '';
  let categoryId = '';
  let status = PostStatus.DRAFT;
  let seoTitle = '';
  let seoDescription = '';
  let saving = false;

  onMount(async () => {
    await loadPosts();
  });

  async function loadPosts() {
    loading = true;
    try {
      const [catsRes, postsRes] = await Promise.all([
        api.getCategories().catch(() => []),
        api.getPosts({ limit: 50 })
      ]);
      categories = catsRes;
      posts = postsRes.data;
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  }

  function startCreate() {
    editId = undefined;
    title = '';
    slug = '';
    excerpt = '';
    contentHtml = '';
    categoryId = categories[0]?.id || '';
    status = PostStatus.DRAFT;
    seoTitle = '';
    seoDescription = '';
    activeTab = 'edit';
  }

  function startEdit(post: Post) {
    editId = post.id;
    title = post.title;
    slug = post.slug;
    excerpt = post.excerpt;
    contentHtml = post.content_html;
    categoryId = post.category_id || '';
    status = post.status;
    seoTitle = post.seo_title;
    seoDescription = post.seo_description;
    activeTab = 'edit';
  }

  async function handleSave(newStatus?: PostStatus) {
    if (!title.trim()) {
      toast.show('Judul artikel wajib diisi.', 'error');
      return;
    }
    if (!slug.trim()) {
      slug = title.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/(^-|-$)/g, '');
    }

    saving = true;
    try {
      await api.savePost({
        id: editId,
        title: title.trim(),
        slug: slug.trim(),
        excerpt: excerpt.trim(),
        content_html: contentHtml,
        category_id: categoryId || undefined,
        status: newStatus || status,
        seo_title: seoTitle.trim(),
        seo_description: seoDescription.trim()
      });

      toast.show('Artikel berhasil disimpan.', 'success');
      activeTab = 'list';
      await loadPosts();
    } catch (err: any) {
      toast.show(err.message || 'Gagal menyimpan artikel.', 'error');
    } finally {
      saving = false;
    }
  }

  async function handleDelete(id: string, titleStr: string) {
    if (!confirm(`Hapus artikel "${titleStr}"?`)) return;
    try {
      await api.deletePost(id);
      toast.show('Artikel berhasil dihapus.', 'success');
      await loadPosts();
    } catch (err: any) {
      toast.show(err.message || 'Gagal menghapus artikel.', 'error');
    }
  }
</script>

<svelte:head>
  <title>Kelola Blog - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  {#if activeTab === 'list'}
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-xl font-heading font-semibold text-slate-900">Manajemen Artikel Blog</h2>
        <p class="text-xs text-slate-500 mt-0.5">Tulis panduan tips dan edukasi penanganan hama untuk publik.</p>
      </div>

      <Button variant={ButtonVariant.PRIMARY} size="sm" on:click={startCreate}>
        <Plus size={15} class="mr-1" />
        <span>Tulis Artikel Baru</span>
      </Button>
    </div>

    <!-- Table -->
    <div class="rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
      {#if loading}
        <div class="flex justify-center p-12">
          <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
        </div>
      {:else if posts.length === 0}
        <div class="p-12 text-center text-sm text-slate-500">
          Belum ada artikel. Klik tombol di atas untuk membuat artikel pertama.
        </div>
      {:else}
        <div class="overflow-x-auto">
          <table class="w-full text-left text-sm">
            <thead class="bg-slate-50/80 text-xs font-medium text-slate-500 border-b border-slate-200/80">
              <tr>
                <th class="py-3 px-6">Judul Artikel</th>
                <th class="py-3 px-6">Kategori</th>
                <th class="py-3 px-6">Status</th>
                <th class="py-3 px-6">Tanggal</th>
                <th class="py-3 px-6 text-right">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100">
              {#each posts as post}
                <tr class="hover:bg-slate-50/60 transition-colors">
                  <td class="py-3.5 px-6 font-medium text-slate-900">
                    <div>{post.title}</div>
                    <div class="text-xs text-slate-500 font-mono font-normal">/blog/{post.slug}</div>
                  </td>
                  <td class="py-3.5 px-6 text-xs text-slate-600">
                    {post.category?.name || 'Umum'}
                  </td>
                  <td class="py-3.5 px-6">
                    <StatusPill status={post.status} />
                  </td>
                  <td class="py-3.5 px-6 text-xs text-slate-500">
                    20 Okt 2026
                  </td>
                  <td class="py-3.5 px-6 text-right space-x-2">
                    <button
                      type="button"
                      class="text-slate-600 hover:text-slate-900 p-1 rounded transition-colors"
                      on:click={() => startEdit(post)}
                      title="Edit"
                    >
                      <Edit size={15} />
                    </button>
                    <button
                      type="button"
                      class="text-slate-400 hover:text-red-600 p-1 rounded transition-colors"
                      on:click={() => handleDelete(post.id, post.title)}
                      title="Hapus"
                    >
                      <Trash2 size={15} />
                    </button>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      {/if}
    </div>
  {:else}
    <!-- Edit Form Tab -->
    <div class="space-y-6">
      <div class="flex items-center justify-between border-b border-slate-200/80 pb-4">
        <button
          type="button"
          class="flex items-center gap-1.5 text-xs font-medium text-slate-600 hover:text-slate-900"
          on:click={() => activeTab = 'list'}
        >
          <ArrowLeft size={15} />
          <span>Kembali ke Daftar</span>
        </button>
        <h3 class="text-lg font-heading font-semibold text-slate-900">
          {editId ? 'Edit Artikel' : 'Tulis Artikel Baru'}
        </h3>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
        <!-- Main Form (2/3) -->
        <div class="lg:col-span-8 space-y-5">
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <div>
              <label for="post_title" class="block text-xs font-medium text-slate-700 mb-1.5">
                Judul Artikel <span class="text-red-500">*</span>
              </label>
              <input
                id="post_title"
                type="text"
                placeholder="Masukkan judul artikel yang menarik..."
                bind:value={title}
                class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-sm font-medium text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
              />
            </div>

            <div>
              <label for="post_slug" class="block text-xs font-medium text-slate-700 mb-1.5">
                URL Slug
              </label>
              <input
                id="post_slug"
                type="text"
                placeholder="contoh-judul-artikel-anda"
                bind:value={slug}
                class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-xs font-mono text-slate-700 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
              />
            </div>

            <div>
              <label for="post_excerpt" class="block text-xs font-medium text-slate-700 mb-1.5">
                Ringkasan Singkat (Excerpt)
              </label>
              <textarea
                id="post_excerpt"
                rows={2}
                placeholder="Deskripsi 1-2 kalimat untuk preview di kartu blog..."
                bind:value={excerpt}
                class="w-full rounded-lg border border-slate-300 bg-white p-3 text-xs leading-relaxed text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
              ></textarea>
            </div>
          </div>

          <!-- TipTap Rich Text Editor -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-2">
            <label class="block text-xs font-medium text-slate-700 mb-1.5">
              Konten Lengkap Artikel
            </label>
            <TipTapEditor bind:content={contentHtml} />
          </div>
        </div>

        <!-- Sidebar Panel (1/3) -->
        <div class="lg:col-span-4 space-y-6">
          <!-- Publish Settings Card -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <h4 class="text-sm font-heading font-semibold text-slate-900 pb-2 border-b border-slate-100">
              Status & Publikasi
            </h4>

            <div>
              <label for="post_status" class="block text-xs font-medium text-slate-700 mb-1">Status</label>
              <select
                id="post_status"
                bind:value={status}
                class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs font-medium text-slate-900 focus:border-navy focus:outline-none"
              >
                <option value={PostStatus.DRAFT}>Draft (Belum Tayang)</option>
                <option value={PostStatus.PUBLISHED}>Published (Tayang Publik)</option>
              </select>
            </div>

            <div>
              <label for="post_category" class="block text-xs font-medium text-slate-700 mb-1">Kategori</label>
              <select
                id="post_category"
                bind:value={categoryId}
                class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs font-medium text-slate-900 focus:border-navy focus:outline-none"
              >
                {#each categories as cat}
                  <option value={cat.id}>{cat.name}</option>
                {/each}
              </select>
            </div>
          </div>

          <!-- SEO Panel -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <h4 class="text-sm font-heading font-semibold text-slate-900 pb-2 border-b border-slate-100">
              Optimasi SEO (Google)
            </h4>

            <div>
              <div class="flex justify-between text-xs font-medium text-slate-700 mb-1">
                <span>SEO Title</span>
                <span class="text-slate-400">{seoTitle.length}/60</span>
              </div>
              <input
                type="text"
                maxlength="60"
                placeholder="Judul optimal di hasil Google..."
                bind:value={seoTitle}
                class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none"
              />
            </div>

            <div>
              <div class="flex justify-between text-xs font-medium text-slate-700 mb-1">
                <span>Meta Description</span>
                <span class="text-slate-400">{seoDescription.length}/160</span>
              </div>
              <textarea
                rows={3}
                maxlength="160"
                placeholder="Ringkasan 1-2 kalimat untuk Google snippet..."
                bind:value={seoDescription}
                class="w-full rounded-lg border border-slate-300 bg-white p-3 text-xs text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none"
              ></textarea>
            </div>
          </div>
        </div>
      </div>

      <!-- Sticky Save Bar -->
      <div class="sticky bottom-4 z-20 flex items-center justify-end gap-3 rounded-xl bg-slate-900/90 backdrop-blur-md p-3.5 text-white shadow-lg">
        <button
          type="button"
          class="rounded-md px-4 py-2 text-xs font-medium text-slate-300 hover:text-white"
          on:click={() => activeTab = 'list'}
        >
          Batal
        </button>

        <button
          type="button"
          class="rounded-md border border-white/20 bg-white/10 px-4 py-2 text-xs font-medium text-white hover:bg-white/20 transition-colors"
          on:click={() => handleSave(PostStatus.DRAFT)}
          disabled={saving}
        >
          Simpan sebagai Draft
        </button>

        <Button
          variant={ButtonVariant.PRIMARY}
          size="sm"
          loading={saving}
          on:click={() => handleSave(PostStatus.PUBLISHED)}
        >
          Simpan & Terbitkan
        </Button>
      </div>
    </div>
  {/if}
</div>
