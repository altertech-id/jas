<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Project } from '$lib/types';
  import StatusPill from '$lib/components/ui/StatusPill.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { Plus, Edit, Trash2, ArrowLeft, MapPin } from 'lucide-svelte';

  let projects: Project[] = [];
  let loading = true;
  let activeTab: 'list' | 'edit' = 'list';

  // Form edit state
  let editId: string | undefined = undefined;
  let title = '';
  let slug = '';
  let location = '';
  let description = '';
  let completedAt = '';
  let isPublished = true;
  let saving = false;

  onMount(async () => {
    await loadProjects();
  });

  async function loadProjects() {
    loading = true;
    try {
      projects = await api.getProjects();
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
    location = '';
    description = '';
    completedAt = new Date().toISOString().split('T')[0];
    isPublished = true;
    activeTab = 'edit';
  }

  function startEdit(pr: Project) {
    editId = pr.id;
    title = pr.title;
    slug = pr.slug;
    location = pr.location;
    description = pr.description;
    completedAt = pr.completed_at ? pr.completed_at.split('T')[0] : '';
    isPublished = pr.is_published;
    activeTab = 'edit';
  }

  async function handleSave() {
    if (!title.trim() || !location.trim()) {
      toast.show('Judul dan lokasi proyek wajib diisi.', 'error');
      return;
    }
    if (!slug.trim()) {
      slug = title.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/(^-|-$)/g, '');
    }

    saving = true;
    try {
      await api.saveProject({
        id: editId,
        title: title.trim(),
        slug: slug.trim(),
        location: location.trim(),
        description: description.trim(),
        completed_at: completedAt ? completedAt : undefined,
        is_published: isPublished
      });

      toast.show('Proyek berhasil disimpan.', 'success');
      activeTab = 'list';
      await loadProjects();
    } catch (err: any) {
      toast.show(err.message || 'Gagal menyimpan proyek.', 'error');
    } finally {
      saving = false;
    }
  }

  async function handleDelete(id: string, titleStr: string) {
    if (!confirm(`Hapus proyek "${titleStr}"?`)) return;
    try {
      await api.deleteProject(id);
      toast.show('Proyek berhasil dihapus.', 'success');
      await loadProjects();
    } catch (err: any) {
      toast.show(err.message || 'Gagal menghapus proyek.', 'error');
    }
  }
</script>

<svelte:head>
  <title>Kelola Proyek - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  {#if activeTab === 'list'}
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-xl font-heading font-semibold text-slate-900">Manajemen Portofolio Proyek</h2>
        <p class="text-xs text-slate-500 mt-0.5">Kelola dokumentasi proyek pengerjaan lapangan JAS.</p>
      </div>

      <Button variant={ButtonVariant.PRIMARY} size="sm" on:click={startCreate}>
        <Plus size={15} class="mr-1" />
        <span>Tambah Proyek Baru</span>
      </Button>
    </div>

    <!-- Table -->
    <div class="rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
      {#if loading}
        <div class="flex justify-center p-12">
          <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
        </div>
      {:else if projects.length === 0}
        <div class="p-12 text-center text-sm text-slate-500">
          Belum ada proyek terdaftar.
        </div>
      {:else}
        <div class="overflow-x-auto">
          <table class="w-full text-left text-sm">
            <thead class="bg-slate-50/80 text-xs font-medium text-slate-500 border-b border-slate-200/80">
              <tr>
                <th class="py-3 px-6">Nama Proyek</th>
                <th class="py-3 px-6">Lokasi</th>
                <th class="py-3 px-6">Status</th>
                <th class="py-3 px-6 text-right">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100">
              {#each projects as pr}
                <tr class="hover:bg-slate-50/60 transition-colors">
                  <td class="py-3.5 px-6 font-medium text-slate-900">
                    <div>{pr.title}</div>
                    <div class="text-xs text-slate-500 font-mono font-normal">/project/{pr.slug}</div>
                  </td>
                  <td class="py-3.5 px-6 text-xs text-slate-600">
                    <div class="flex items-center gap-1.5">
                      <MapPin size={13} class="text-slate-400" />
                      <span>{pr.location}</span>
                    </div>
                  </td>
                  <td class="py-3.5 px-6">
                    <StatusPill status={pr.is_published ? 'published' : 'draft'} />
                  </td>
                  <td class="py-3.5 px-6 text-right space-x-2">
                    <button
                      type="button"
                      class="text-slate-600 hover:text-slate-900 p-1 rounded transition-colors"
                      on:click={() => startEdit(pr)}
                      title="Edit"
                    >
                      <Edit size={15} />
                    </button>
                    <button
                      type="button"
                      class="text-slate-400 hover:text-red-600 p-1 rounded transition-colors"
                      on:click={() => handleDelete(pr.id, pr.title)}
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
    <!-- Form Create / Edit -->
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
          {editId ? 'Edit Proyek' : 'Tambah Proyek Baru'}
        </h3>
      </div>

      <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4 max-w-2xl">
        <div>
          <label for="project_title" class="block text-xs font-medium text-slate-700 mb-1.5">
            Nama / Judul Proyek <span class="text-red-500">*</span>
          </label>
          <input
            id="project_title"
            type="text"
            placeholder="Contoh: Proteksi Rayap Pergudangan Rungkut"
            bind:value={title}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
          />
        </div>

        <div>
          <label for="project_slug" class="block text-xs font-medium text-slate-700 mb-1.5">
            URL Slug
          </label>
          <input
            id="project_slug"
            type="text"
            placeholder="proteksi-rayap-pergudangan-rungkut"
            bind:value={slug}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-xs font-mono text-slate-700 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
          />
        </div>

        <div>
          <label for="project_location" class="block text-xs font-medium text-slate-700 mb-1.5">
            Lokasi Pengerjaan <span class="text-red-500">*</span>
          </label>
          <input
            id="project_location"
            type="text"
            placeholder="Contoh: Kawasan Industri SIER, Surabaya"
            bind:value={location}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
          />
        </div>

        <div>
          <label for="project_desc" class="block text-xs font-medium text-slate-700 mb-1.5">
            Deskripsi & Ruang Lingkup Pengerjaan
          </label>
          <textarea
            id="project_desc"
            rows={4}
            placeholder="Jelaskan detail perlakuan kimia, luasan area, dan hasil penanganan..."
            bind:value={description}
            class="w-full rounded-lg border border-slate-300 bg-white p-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"
          ></textarea>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label for="project_completed" class="block text-xs font-medium text-slate-700 mb-1.5">
              Tanggal Selesai
            </label>
            <input
              id="project_completed"
              type="date"
              bind:value={completedAt}
              class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none"
            />
          </div>

          <div>
            <label for="project_status" class="block text-xs font-medium text-slate-700 mb-1.5">
              Status Publikasi
            </label>
            <select
              id="project_status"
              bind:value={isPublished}
              class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-xs text-slate-900 focus:border-navy focus:outline-none"
            >
              <option value={true}>Terbit (Tampil di Website)</option>
              <option value={false}>Draft (Disembunyikan)</option>
            </select>
          </div>
        </div>

        <div class="pt-4 flex gap-3">
          <Button variant={ButtonVariant.PRIMARY} size="default" loading={saving} on:click={handleSave}>
            Simpan Proyek
          </Button>
          <Button variant={ButtonVariant.OUTLINE} size="default" on:click={() => activeTab = 'list'}>
            Batal
          </Button>
        </div>
      </div>
    </div>
  {/if}
</div>
