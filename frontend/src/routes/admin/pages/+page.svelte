<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { PageContent } from '$lib/types';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { FileText, Save } from 'lucide-svelte';

  let selectedPage = 'home';
  let pageData: PageContent | null = null;
  let loading = true;
  let saving = false;

  // Edit fields
  let title = '';
  let seoTitle = '';
  let seoDescription = '';
  let contentJson = '';

  onMount(async () => {
    await loadPage(selectedPage);
  });

  async function loadPage(key: string) {
    loading = true;
    selectedPage = key;
    try {
      pageData = await api.getPageContent(key);
      if (pageData) {
        title = pageData.title;
        seoTitle = pageData.seo_title || '';
        seoDescription = pageData.seo_description || '';
        contentJson = JSON.stringify(pageData.content, null, 2);
      }
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  }

  async function handleSave() {
    let parsedContent = {};
    try {
      parsedContent = JSON.parse(contentJson);
    } catch {
      toast.show('Format JSON konten tidak valid. Mohon periksa kembali.', 'error');
      return;
    }

    saving = true;
    try {
      await api.savePageContent(selectedPage, {
        title: title.trim(),
        content: parsedContent,
        seo_title: seoTitle.trim(),
        seo_description: seoDescription.trim()
      });
      toast.show(`Konten halaman ${title} berhasil diperbarui.`, 'success');
    } catch (err: any) {
      toast.show(err.message || 'Gagal menyimpan konten halaman.', 'error');
    } finally {
      saving = false;
    }
  }
</script>

<svelte:head>
  <title>Kelola Halaman Statis - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  <!-- Header -->
  <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
    <div>
      <h2 class="text-xl font-heading font-semibold text-slate-900">Manajemen Halaman Utama (CMS)</h2>
      <p class="text-xs text-slate-500 mt-0.5">Edit teks hero, statistik, profil, visi misi, dan informasi kontak.</p>
    </div>

    <!-- Page Switcher Tabs -->
    <div class="inline-flex rounded-lg border border-slate-200 bg-slate-50/80 p-0.5">
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {selectedPage === 'home' ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => loadPage('home')}
      >
        Beranda
      </button>
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {selectedPage === 'about-us' ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => loadPage('about-us')}
      >
        Tentang Kami
      </button>
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {selectedPage === 'contact' ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => loadPage('contact')}
      >
        Kontak
      </button>
    </div>
  </div>

  {#if loading}
    <div class="flex justify-center p-12">
      <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
    </div>
  {:else}
    <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-5">
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div>
          <label for="page_title" class="block text-xs font-medium text-slate-700 mb-1">Judul Halaman</label>
          <input id="page_title" type="text" bind:value={title} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
        </div>
        <div>
          <label for="page_seotitle" class="block text-xs font-medium text-slate-700 mb-1">SEO Title Tag</label>
          <input id="page_seotitle" type="text" bind:value={seoTitle} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
        </div>
      </div>

      <div>
        <label for="page_seodesc" class="block text-xs font-medium text-slate-700 mb-1">Meta Description</label>
        <textarea id="page_seodesc" rows={2} bind:value={seoDescription} class="w-full rounded-lg border border-slate-300 bg-white p-3 text-xs text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"></textarea>
      </div>

      <div>
        <label for="page_json" class="block text-xs font-medium text-slate-700 mb-1">
          Struktur Konten Blok Halaman (JSON)
        </label>
        <textarea
          id="page_json"
          rows={14}
          bind:value={contentJson}
          class="w-full rounded-lg border border-slate-300 p-4 font-mono text-xs text-slate-900 leading-relaxed bg-slate-50 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
        ></textarea>
      </div>

      <div class="pt-2 flex justify-end">
        <Button variant={ButtonVariant.PRIMARY} size="default" loading={saving} on:click={handleSave}>
          <Save size={15} class="mr-2" />
          <span>Simpan Konten Halaman</span>
        </Button>
      </div>
    </div>
  {/if}
</div>
