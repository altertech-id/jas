<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { GalleryItem } from '$lib/types';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { Plus, Trash2, Upload, Images } from 'lucide-svelte';

  let items: GalleryItem[] = [];
  let loading = true;
  let showUploadModal = false;

  // New item form
  let caption = '';
  let category = 'Pest Control';
  let sortOrder = 1;
  let selectedFile: File | null = null;
  let uploading = false;

  onMount(async () => {
    await loadGallery();
  });

  async function loadGallery() {
    loading = true;
    try {
      items = await api.getGallery();
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  }

  function handleFileChange(e: Event) {
    const input = e.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      selectedFile = input.files[0];
    }
  }

  async function handleUpload() {
    if (!selectedFile) {
      toast.show('Pilih file foto terlebih dahulu.', 'error');
      return;
    }
    if (!caption.trim()) {
      toast.show('Keterangan foto (caption) wajib diisi.', 'error');
      return;
    }

    uploading = true;
    try {
      // 1. Upload media file to Supabase Storage
      const media = await api.uploadMedia(selectedFile, caption);

      // 2. Add record to gallery items
      await api.saveGalleryItem({
        media_id: media.id,
        caption: caption.trim(),
        category: category,
        sort_order: sortOrder
      });

      toast.show('Foto dokumentasi berhasil ditambahkan ke galeri.', 'success');
      showUploadModal = false;
      selectedFile = null;
      caption = '';
      await loadGallery();
    } catch (err: any) {
      toast.show(err.message || 'Gagal mengunggah foto.', 'error');
    } finally {
      uploading = false;
    }
  }

  async function handleDelete(id: string) {
    if (!confirm('Hapus foto ini dari galeri?')) return;
    try {
      await api.deleteGalleryItem(id);
      toast.show('Foto berhasil dihapus.', 'success');
      await loadGallery();
    } catch (err: any) {
      toast.show(err.message || 'Gagal menghapus foto.', 'error');
    }
  }
</script>

<svelte:head>
  <title>Kelola Galeri - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex items-center justify-between">
    <div>
      <h2 class="text-xl font-heading font-semibold text-slate-900">Manajemen Galeri & Dokumentasi</h2>
      <p class="text-xs text-slate-500 mt-0.5">Foto asli pengerjaan lapangan teknisi JAS di website publik.</p>
    </div>

    <Button variant={ButtonVariant.PRIMARY} size="sm" on:click={() => showUploadModal = true}>
      <Plus size={15} class="mr-1" />
      <span>Unggah Foto Dokumentasi</span>
    </Button>
  </div>

  {#if loading}
    <div class="flex justify-center p-12">
      <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
    </div>
  {:else if items.length === 0}
    <div class="rounded-xl border border-slate-200/80 bg-white p-12 text-center text-sm text-slate-500">
      Belum ada foto galeri terdaftar.
    </div>
  {:else}
    <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-4">
      {#each items as item}
        <div class="group relative rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs flex flex-col justify-between">
          <div class="aspect-[3/2] w-full bg-slate-100 overflow-hidden relative">
            <img
              src={item.media?.public_url || 'https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=600&q=80'}
              alt={item.caption}
              class="h-full w-full object-cover"
            />
            <span class="absolute top-2 left-2 rounded bg-slate-900/80 px-2 py-0.5 text-[10px] font-medium text-white">
              {item.category}
            </span>
          </div>

          <div class="p-3">
            <p class="text-xs text-slate-900 line-clamp-2 leading-relaxed mb-3">
              {item.caption}
            </p>

            <div class="flex items-center justify-between pt-2 border-t border-slate-100 text-[11px] text-slate-500">
              <span>Urutan: {item.sort_order}</span>
              <button
                type="button"
                class="text-slate-400 hover:text-red-600 p-1 rounded"
                on:click={() => handleDelete(item.id)}
                title="Hapus foto"
              >
                <Trash2 size={14} />
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<!-- Modal Upload -->
{#if showUploadModal}
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-slate-900/50 backdrop-blur-xs p-4 animate-in fade-in">
    <div class="w-full max-w-lg rounded-xl bg-white p-6 shadow-xl border border-slate-200 space-y-4">
      <h3 class="text-base font-heading font-semibold text-slate-900">
        Unggah Foto Dokumentasi Baru
      </h3>

      <!-- Dropzone -->
      <div>
        <label for="gallery_file" class="block text-xs font-medium text-slate-700 mb-1.5">File Foto (JPG, PNG, WebP max 5 MB)</label>
        <input
          id="gallery_file"
          type="file"
          accept="image/jpeg,image/png,image/webp"
          on:change={handleFileChange}
          class="w-full text-xs text-slate-600 file:mr-4 file:py-1.5 file:px-3 file:rounded-md file:border-0 file:text-xs file:font-medium file:bg-slate-100 file:text-slate-800 hover:file:bg-slate-200"
        />
      </div>

      <div>
        <label for="gallery_caption" class="block text-xs font-medium text-slate-700 mb-1">Keterangan / Caption Foto</label>
        <textarea
          id="gallery_caption"
          rows={2}
          placeholder="Contoh: Teknisi melakukan fogging kecoa di area gudang..."
          bind:value={caption}
          class="w-full rounded-lg border border-slate-300 bg-white p-3 text-xs text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none"
        ></textarea>
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label for="gallery_cat" class="block text-xs font-medium text-slate-700 mb-1">Kategori</label>
          <select
            id="gallery_cat"
            bind:value={category}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none"
          >
            <option value="Pest Control">Pest Control</option>
            <option value="Termite Control">Termite Control</option>
            <option value="Rodent Control">Rodent Control</option>
            <option value="Fumigasi">Fumigasi</option>
            <option value="Disinfeksi">Disinfeksi</option>
          </select>
        </div>

        <div>
          <label for="gallery_order" class="block text-xs font-medium text-slate-700 mb-1">Urutan Tampil</label>
          <input
            id="gallery_order"
            type="number"
            bind:value={sortOrder}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none"
          />
        </div>
      </div>

      <div class="pt-4 flex justify-end gap-3 border-t border-slate-100">
        <button
          type="button"
          class="rounded-md px-4 py-2 text-xs font-medium text-slate-600 hover:text-slate-900"
          on:click={() => showUploadModal = false}
        >
          Batal
        </button>

        <Button
          variant={ButtonVariant.PRIMARY}
          size="sm"
          loading={uploading}
          on:click={handleUpload}
        >
          Unggah ke Galeri
        </Button>
      </div>
    </div>
  </div>
{/if}
