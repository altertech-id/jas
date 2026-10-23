<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { api } from '$lib/api/client';
  import type { Project } from '$lib/types';
  import Badge from '$lib/components/ui/Badge.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { MapPin, Calendar, ArrowLeft, ShieldCheck } from 'lucide-svelte';

  let project: Project | null = null;
  let loading = true;
  let errorMsg = '';

  $: slug = $page.params.slug;

  onMount(async () => {
    if (!slug) {
      errorMsg = 'Slug tidak valid.';
      loading = false;
      return;
    }
    try {
      project = await api.getProjectBySlug(slug);
    } catch (err: any) {
      errorMsg = err.message || 'Proyek tidak ditemukan.';
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  {#if project}
    <title>{project.title} - Jaya Abadi Solusi</title>
    <meta name="description" content={project.description} />
  {:else}
    <title>Detail Proyek - Jaya Abadi Solusi</title>
  {/if}
</svelte:head>

{#if loading}
  <div class="flex min-h-[60vh] items-center justify-center">
    <div class="h-8 w-8 animate-spin rounded-full border-4 border-navy border-t-transparent"></div>
  </div>
{:else if errorMsg || !project}
  <div class="mx-auto max-w-container px-5 py-24 text-center">
    <h1 class="text-2xl font-heading font-bold text-navy mb-2">Proyek Tidak Ditemukan</h1>
    <p class="text-sm text-ink-500 mb-6">{errorMsg}</p>
    <Button href="/project" variant={ButtonVariant.PRIMARY}>Kembali ke Daftar Proyek</Button>
  </div>
{:else}
  <article class="py-12 sm:py-20 bg-white">
    <div class="mx-auto max-w-4xl px-5 sm:px-8">
      <!-- Back Link -->
      <a href="/project" class="inline-flex items-center gap-1.5 text-xs font-semibold text-brand-blue hover:underline mb-6">
        <ArrowLeft size={16} />
        <span>Kembali ke Portofolio Proyek</span>
      </a>

      <!-- Header -->
      <h1 class="text-3xl sm:text-4xl lg:text-5xl font-heading font-extrabold text-navy leading-tight mb-4">
        {project.title}
      </h1>

      <div class="flex flex-wrap items-center gap-4 text-xs sm:text-sm text-ink-500 pb-6 border-b border-line-soft mb-8">
        <div class="flex items-center gap-1.5">
          <MapPin size={16} class="text-brand-blue" />
          <span>{project.location}</span>
        </div>
        <div class="flex items-center gap-1.5">
          <Calendar size={16} class="text-amber" />
          <span>Selesai: {project.completed_at ? new Date(project.completed_at).toLocaleDateString('id-ID', { year: 'numeric', month: 'short' }) : '2026'}</span>
        </div>
      </div>

      <!-- Cover Image -->
      <div class="aspect-[16/9] w-full overflow-hidden rounded-xl bg-slate-100 shadow-sm mb-8">
        <img
          src="https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=1200&q=80"
          alt={project.title}
          class="h-full w-full object-cover"
        />
      </div>

      <!-- Content -->
      <div class="prose max-w-none text-slate-700 leading-relaxed mb-12 font-body">
        <h3 class="text-xl font-heading font-bold text-slate-900 mb-3">Deskripsi Proyek & Ruang Lingkup Pengerjaan</h3>
        <p class="text-base text-slate-600 leading-relaxed whitespace-pre-line mb-6 font-body">
          {project.description}
        </p>

        <div class="rounded-xl border border-slate-200/80 bg-slate-50 p-6 sm:p-8">
          <div class="flex items-center gap-2 font-heading font-bold text-slate-900 text-lg mb-2">
            <ShieldCheck class="text-amber" size={20} />
            <span>Garansi & Standar Kualitas</span>
          </div>
          <p class="text-sm text-slate-600 leading-relaxed font-body">
            Seluruh pengerjaan proyek oleh PT Jaya Abadi Solusi dilengkapi dengan sertifikat garansi resmi, pemantauan berkala, dan penanganan tindak lanjut secara proaktif demi kepuasan klien.
          </p>
        </div>
      </div>

      <!-- CTA Box -->
      <div class="rounded-xl bg-hero-navy p-8 text-center text-white">
        <h3 class="text-2xl font-heading font-bold mb-2">Punya Kebutuhan Pengendalian Serupa?</h3>
        <p class="text-sm text-slate-300 mb-6 max-w-xl mx-auto font-body">
          Konsultasikan masalah hama pada fasilitas Anda dengan tim ahli kami untuk mendapatkan survei lokasi dan estimasi penawaran resmi.
        </p>
        <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
          <Button href="/kontak" variant={ButtonVariant.PRIMARY}>Minta Penawaran Gratis</Button>
          <Button href="https://wa.me/6281399484916?text=Halo%20JAS%2C%20saya%20tertarik%20dengan%20proyek%20{encodeURIComponent(project.title)}" variant={ButtonVariant.WHATSAPP}>
            Chat via WhatsApp
          </Button>
        </div>
      </div>
    </div>
  </article>
{/if}
