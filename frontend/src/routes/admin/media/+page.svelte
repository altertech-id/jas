<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Media } from '$lib/types';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { UploadCloud, Copy, Trash2, Check, Image as ImageIcon } from 'lucide-svelte';

  let mediaList: Media[] = [];
  let loading = true;
  let uploading = false;
  let copiedId: string | null = null;

  onMount(async () => {
    await loadMedia();
  });

  async function loadMedia() {
    loading = true;
    try {
      mediaList = await api.getMediaList();
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  }

  async function handleFileUpload(e: Event) {
    const input = e.target as HTMLInputElement;
    if (!input.files || input.files.length === 0) return;

    const file = input.files[0];
    uploading = true;
    try {
      await api.uploadMedia(file, file.name);
      toast.show('Gambar berhasil diunggah ke storage.', 'success');
      await loadMedia();
    } catch (err: any) {
      toast.show(err.message || 'Gagal mengunggah gambar.', 'error');
    } finally {
      uploading = false;
      input.value = '';
    }
  }

  function copyUrl(url: string, id: string) {
    navigator.clipboard.writeText(url);
    copiedId = id;
    toast.show('URL gambar berhasil disalin ke clipboard.', 'info');
    setTimeout(() => {
      copiedId = null;
    }, 2000);
  }

  async function handleDelete(id: string) {
    if (!confirm('Hapus file media ini?')) return;
    try {
      await api.deleteMedia(id);
      toast.show('Media berhasil dihapus.', 'success');
      await loadMedia();
    } catch (err: any) {
      toast.show(err.message || 'Gagal menghapus media.', 'error');
    }
  }
</script>

<svelte:head>
  <title>Media Library - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  <div>
    <h2 class="text-xl font-heading font-semibold text-slate-900">Media Library</h2>
    <p class="text-xs text-slate-500 mt-0.5">Penyimpanan aset gambar terunggah di Supabase Storage bucket media.</p>
  </div>

  <!-- Dropzone Card -->
  <div class="rounded-xl border border-dashed border-slate-300 bg-white p-8 text-center transition-colors hover:border-slate-400">
    <div class="mx-auto flex h-12 w-12 items-center justify-center rounded-lg bg-slate-50 text-slate-700 mb-3">
      <UploadCloud size={24} />
    </div>
    <div class="text-sm font-heading font-semibold text-slate-900 mb-1">
      Tarik file gambar ke sini atau klik untuk memilih
    </div>
    <p class="text-xs text-slate-500 mb-4">
      Mendukung format JPG, PNG, atau WebP (Maksimal 5 MB)
    </p>

    <label class="inline-flex cursor-pointer">
      <input
        type="file"
        accept="image/jpeg,image/png,image/webp"
        on:change={handleFileUpload}
        class="hidden"
        disabled={uploading}
      />
      <span class="inline-flex items-center justify-center rounded-md bg-slate-900 hover:bg-slate-800 text-white px-4 py-2 text-xs font-medium transition-colors">
        {uploading ? 'Mengunggah...' : 'Pilih File dari Komputer'}
      </span>
    </label>
  </div>

  <!-- Media Grid -->
  {#if loading}
    <div class="flex justify-center p-12">
      <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
    </div>
  {:else if mediaList.length === 0}
    <div class="rounded-xl border border-slate-200/80 bg-white p-12 text-center text-sm text-slate-500">
      Belum ada aset gambar terunggah.
    </div>
  {:else}
    <div class="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-6 gap-4">
      {#each mediaList as media}
        <div class="group relative rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs flex flex-col justify-between">
          <div class="aspect-square w-full bg-slate-100 overflow-hidden relative">
            <img
              src={media.public_url}
              alt={media.alt_text || 'Media JAS'}
              class="h-full w-full object-cover"
            />
          </div>

          <div class="p-2 border-t border-slate-100 flex items-center justify-between text-xs bg-white">
            <button
              type="button"
              class="flex items-center gap-1 text-[11px] font-medium text-brand-blue hover:underline"
              on:click={() => copyUrl(media.public_url, media.id)}
            >
              {#if copiedId === media.id}
                <Check size={12} class="text-emerald-600" />
                <span class="text-emerald-600">Disalin</span>
              {:else}
                <Copy size={12} />
                <span>Salin URL</span>
              {/if}
            </button>

            <button
              type="button"
              class="text-slate-400 hover:text-red-600 p-1"
              on:click={() => handleDelete(media.id)}
              title="Hapus media"
            >
              <Trash2 size={13} />
            </button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>
