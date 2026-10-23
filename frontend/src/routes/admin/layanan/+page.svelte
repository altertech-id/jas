<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Service } from '$lib/types';
  import StatusPill from '$lib/components/ui/StatusPill.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { Edit, ShieldCheck, ArrowLeft, Plus, Trash2 } from 'lucide-svelte';

  let services: Service[] = [];
  let loading = true;
  let activeTab: 'list' | 'edit' = 'list';

  // Form edit state
  let editId: string | undefined = undefined;
  let name = '';
  let slug = '';
  let icon = 'ShieldCheck';
  let shortDescription = '';
  let isPublished = true;
  let sortOrder = 1;
  let heroBadge = '';
  let headline = '';
  let tagline = '';
  let problems: string[] = [];
  let solutions: string[] = [];
  let methods: { title: string; desc: string }[] = [];
  let seoTitle = '';
  let seoDescription = '';
  let saving = false;

  onMount(async () => {
    await loadServices();
  });

  async function loadServices() {
    loading = true;
    try {
      services = await api.getServices();
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  }

  function startEdit(s: Service) {
    editId = s.id;
    name = s.name;
    slug = s.slug;
    icon = s.icon;
    shortDescription = s.short_description;
    isPublished = s.is_published;
    sortOrder = s.sort_order;
    heroBadge = s.content?.hero_badge || '';
    headline = s.content?.headline || '';
    tagline = s.content?.tagline || '';
    problems = s.content?.problems ? [...s.content.problems] : [];
    solutions = s.content?.solutions ? [...s.content.solutions] : [];
    methods = s.content?.methods ? [...s.content.methods] : [];
    seoTitle = s.seo_title || '';
    seoDescription = s.seo_description || '';
    activeTab = 'edit';
  }

  function addProblem() {
    problems = [...problems, ''];
  }

  function removeProblem(idx: number) {
    problems = problems.filter((_, i) => i !== idx);
  }

  function addSolution() {
    solutions = [...solutions, ''];
  }

  function removeSolution(idx: number) {
    solutions = solutions.filter((_, i) => i !== idx);
  }

  function addMethod() {
    methods = [...methods, { title: '', desc: '' }];
  }

  function removeMethod(idx: number) {
    methods = methods.filter((_, i) => i !== idx);
  }

  async function handleSave() {
    if (!name.trim() || !slug.trim()) {
      toast.show('Nama dan slug layanan wajib diisi.', 'error');
      return;
    }

    saving = true;
    try {
      await api.saveService({
        id: editId,
        name: name.trim(),
        slug: slug.trim(),
        icon: icon.trim(),
        short_description: shortDescription.trim(),
        is_published: isPublished,
        sort_order: sortOrder,
        content: {
          hero_badge: heroBadge.trim(),
          headline: headline.trim(),
          tagline: tagline.trim(),
          problems: problems.filter(p => p.trim() !== ''),
          solutions: solutions.filter(s => s.trim() !== ''),
          methods: methods.filter(m => m.title.trim() !== ''),
          benefits: []
        },
        seo_title: seoTitle.trim(),
        seo_description: seoDescription.trim()
      });

      toast.show('Layanan & landing page berhasil diperbarui.', 'success');
      activeTab = 'list';
      await loadServices();
    } catch (err: any) {
      toast.show(err.message || 'Gagal menyimpan layanan.', 'error');
    } finally {
      saving = false;
    }
  }
</script>

<svelte:head>
  <title>Kelola Layanan - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  {#if activeTab === 'list'}
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-xl font-heading font-semibold text-slate-900">Manajemen Layanan & Landing Page</h2>
        <p class="text-xs text-slate-500 mt-0.5">Edit isi landing page 5 layanan utama untuk optimasi konversi trafik iklan.</p>
      </div>
    </div>

    <div class="rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
      {#if loading}
        <div class="flex justify-center p-12">
          <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
        </div>
      {:else}
        <div class="overflow-x-auto">
          <table class="w-full text-left text-sm">
            <thead class="bg-slate-50/80 text-xs font-medium text-slate-500 border-b border-slate-200/80">
              <tr>
                <th class="py-3 px-6">Urutan</th>
                <th class="py-3 px-6">Nama Layanan</th>
                <th class="py-3 px-6">URL Landing</th>
                <th class="py-3 px-6">Status</th>
                <th class="py-3 px-6 text-right">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100">
              {#each services as item}
                <tr class="hover:bg-slate-50/60 transition-colors">
                  <td class="py-3.5 px-6 font-mono text-xs text-slate-500">
                    {item.sort_order}
                  </td>
                  <td class="py-3.5 px-6 font-medium text-slate-900">
                    {item.name}
                  </td>
                  <td class="py-3.5 px-6 font-mono text-xs text-brand-blue">
                    /layanan/{item.slug}
                  </td>
                  <td class="py-3.5 px-6">
                    <StatusPill status={item.is_published ? 'published' : 'draft'} />
                  </td>
                  <td class="py-3.5 px-6 text-right">
                    <button
                      type="button"
                      class="inline-flex items-center gap-1.5 rounded-md border border-slate-200 bg-white px-3 py-1 text-xs font-medium text-slate-700 hover:bg-slate-50 transition-colors"
                      on:click={() => startEdit(item)}
                    >
                      <Edit size={13} />
                      <span>Edit Konten</span>
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
    <!-- Form Edit Landing Page Content -->
    <div class="space-y-6">
      <div class="flex items-center justify-between border-b border-slate-200/80 pb-4">
        <button
          type="button"
          class="flex items-center gap-1.5 text-xs font-medium text-slate-600 hover:text-slate-900"
          on:click={() => activeTab = 'list'}
        >
          <ArrowLeft size={15} />
          <span>Kembali ke Daftar Layanan</span>
        </button>
        <h3 class="text-lg font-heading font-semibold text-slate-900">
          Edit Landing Page: {name}
        </h3>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
        <!-- Main Form (8 cols) -->
        <div class="lg:col-span-8 space-y-6">
          <!-- Basic Info -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <h4 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
              Informasi Umum
            </h4>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label for="srv_name" class="block text-xs font-medium text-slate-700 mb-1">Nama Layanan</label>
                <input id="srv_name" type="text" bind:value={name} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
              </div>
              <div>
                <label for="srv_slug" class="block text-xs font-medium text-slate-700 mb-1">Slug URL</label>
                <input id="srv_slug" type="text" bind:value={slug} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs font-mono text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
              </div>
            </div>

            <div>
              <label for="srv_desc" class="block text-xs font-medium text-slate-700 mb-1">Deskripsi Singkat (Kartu)</label>
              <textarea id="srv_desc" rows={2} bind:value={shortDescription} class="w-full rounded-lg border border-slate-300 bg-white p-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"></textarea>
            </div>
          </div>

          <!-- Hero Landing Page Blocks -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <h4 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
              Hero Section (Landing Page)
            </h4>

            <div>
              <label for="hero_badge" class="block text-xs font-medium text-slate-700 mb-1">Badge Hero Pill</label>
              <input id="hero_badge" type="text" placeholder="Contoh: Layanan Paling Diminati" bind:value={heroBadge} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
            </div>

            <div>
              <label for="hero_head" class="block text-xs font-medium text-slate-700 mb-1">Headline Utama (H1)</label>
              <input id="hero_head" type="text" placeholder="Contoh: Basmi Kecoa dan Semut Sampai ke Sarangnya" bind:value={headline} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
            </div>

            <div>
              <label for="hero_tag" class="block text-xs font-medium text-slate-700 mb-1">Subheadline / Tagline</label>
              <textarea id="hero_tag" rows={2} bind:value={tagline} class="w-full rounded-lg border border-slate-300 bg-white p-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"></textarea>
            </div>
          </div>

          <!-- Masalah & Solusi (2 Columns Block) -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-6">
            <h4 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
              Blok Analisis: Masalah vs Solusi JAS
            </h4>

            <!-- Problems list -->
            <div>
              <div class="flex items-center justify-between mb-2">
                <span class="text-xs font-semibold uppercase text-slate-600">Daftar Masalah Klien:</span>
                <button type="button" class="text-xs text-brand-blue hover:underline" on:click={addProblem}>+ Tambah Poin</button>
              </div>
              <div class="space-y-2">
                {#each problems as prob, idx}
                  <div class="flex items-center gap-2">
                    <input type="text" bind:value={problems[idx]} class="h-9 flex-1 rounded-md border border-slate-300 px-3 text-xs text-slate-900 focus:border-navy focus:outline-none" />
                    <button type="button" class="text-slate-400 hover:text-red-600 p-1" on:click={() => removeProblem(idx)}>✕</button>
                  </div>
                {/each}
              </div>
            </div>

            <!-- Solutions list -->
            <div>
              <div class="flex items-center justify-between mb-2">
                <span class="text-xs font-semibold uppercase text-slate-600">Solusi Profesional JAS:</span>
                <button type="button" class="text-xs text-brand-blue hover:underline" on:click={addSolution}>+ Tambah Solusi</button>
              </div>
              <div class="space-y-2">
                {#each solutions as sol, idx}
                  <div class="flex items-center gap-2">
                    <input type="text" bind:value={solutions[idx]} class="h-9 flex-1 rounded-md border border-slate-300 px-3 text-xs text-slate-900 focus:border-navy focus:outline-none" />
                    <button type="button" class="text-slate-400 hover:text-red-600 p-1" on:click={() => removeSolution(idx)}>✕</button>
                  </div>
                {/each}
              </div>
            </div>
          </div>

          <!-- Metode Treatment -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <div class="flex items-center justify-between border-b border-slate-100 pb-2">
              <h4 class="text-sm font-heading font-semibold text-slate-900">
                Metode Treatment
              </h4>
              <button type="button" class="text-xs font-medium text-brand-blue hover:underline" on:click={addMethod}>
                + Tambah Metode
              </button>
            </div>

            <div class="space-y-3">
              {#each methods as method, idx}
                <div class="p-3.5 rounded-lg border border-slate-200 bg-slate-50/60 space-y-2">
                  <div class="flex items-center justify-between">
                    <span class="text-xs font-medium text-slate-600">Metode #{idx + 1}</span>
                    <button type="button" class="text-xs text-red-500 hover:underline" on:click={() => removeMethod(idx)}>Hapus</button>
                  </div>
                  <input type="text" placeholder="Judul Metode (Contoh: Cold Fogging ULV)" bind:value={methods[idx].title} class="h-9 w-full rounded-md border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none" />
                  <textarea rows={2} placeholder="Penjelasan metode treatment..." bind:value={methods[idx].desc} class="w-full rounded-md border border-slate-300 bg-white p-2 text-xs text-slate-900 focus:border-navy focus:outline-none"></textarea>
                </div>
              {/each}
            </div>
          </div>
        </div>

        <!-- Side Panel (4 cols) -->
        <div class="lg:col-span-4 space-y-6">
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <h4 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
              Pengaturan Tampil
            </h4>

            <div>
              <label for="srv_status" class="block text-xs font-medium text-slate-700 mb-1">Status Publikasi</label>
              <select id="srv_status" bind:value={isPublished} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none">
                <option value={true}>Terbit (Aktif)</option>
                <option value={false}>Draft (Nonaktif)</option>
              </select>
            </div>

            <div>
              <label for="srv_order" class="block text-xs font-medium text-slate-700 mb-1">Nomor Urutan Tampil</label>
              <input id="srv_order" type="number" bind:value={sortOrder} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none" />
            </div>

            <div>
              <label for="srv_icon" class="block text-xs font-medium text-slate-700 mb-1">Nama Ikon Lucide</label>
              <input id="srv_icon" type="text" bind:value={icon} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none" />
            </div>
          </div>

          <!-- SEO Panel -->
          <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
            <h4 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
              SEO Google (Landing Iklan)
            </h4>

            <div>
              <label for="seo_title" class="block text-xs font-medium text-slate-700 mb-1">SEO Title</label>
              <input id="seo_title" type="text" bind:value={seoTitle} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none" />
            </div>

            <div>
              <label for="seo_desc" class="block text-xs font-medium text-slate-700 mb-1">Meta Description</label>
              <textarea id="seo_desc" rows={3} bind:value={seoDescription} class="w-full rounded-lg border border-slate-300 bg-white p-3 text-xs text-slate-900 focus:border-navy focus:outline-none"></textarea>
            </div>
          </div>
        </div>
      </div>

      <!-- Sticky Save Bar -->
      <div class="sticky bottom-4 z-20 flex items-center justify-end gap-3 rounded-xl bg-slate-900/90 backdrop-blur-md p-3.5 text-white shadow-lg">
        <button type="button" class="rounded-md px-4 py-2 text-xs font-medium text-slate-300 hover:text-white" on:click={() => activeTab = 'list'}>
          Batal
        </button>
        <Button variant={ButtonVariant.PRIMARY} size="sm" loading={saving} on:click={handleSave}>
          Simpan Seluruh Perubahan
        </Button>
      </div>
    </div>
  {/if}
</div>
