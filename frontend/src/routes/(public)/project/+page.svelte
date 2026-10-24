<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Project } from '$lib/types';
  import Badge from '$lib/components/ui/Badge.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { MapPin, Calendar, ArrowRight } from 'lucide-svelte';

  let projects: Project[] = [];
  let loading = true;

  onMount(async () => {
    try {
      projects = (await api.getProjects()) ?? [];
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  <title>Portofolio Proyek - Jaya Abadi Solusi Surabaya</title>
  <meta name="description" content="Dokumentasi proyek pengerjaan pest control, penanganan rayap pondasi, fumigasi kontainer, dan disinfeksi ruangan di Surabaya." />
</svelte:head>

<!-- Header -->
<section class="bg-hero-navy py-16 sm:py-20 text-white text-center">
  <div class="mx-auto max-w-container px-5 sm:px-8 max-w-3xl">
    <Badge variant="hero" class="mb-3">Rekam Jejak</Badge>
    <h1 class="text-3xl sm:text-5xl font-heading font-extrabold text-white mb-4">
      Portofolio & Pekerjaan Lapangan
    </h1>
    <p class="text-base text-on-navy-soft leading-relaxed">
      Bukti kepercayaan ratusan klien korporasi, fasilitas logistik, restoran, dan pemilik hunian terhadap keandalan penanganan Jaya Abadi Solusi.
    </p>
  </div>
</section>

<!-- Projects Grid -->
<section class="py-16 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    {#if loading}
      <div class="flex justify-center py-16">
        <div class="h-8 w-8 animate-spin rounded-full border-4 border-navy border-t-transparent"></div>
      </div>
    {:else if (projects?.length ?? 0) === 0}
      <div class="text-center py-16 text-ink-500">
        Belum ada dokumentasi proyek yang dipublikasikan.
      </div>
    {:else}
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        {#each projects as pr}
          <a
            href="/project/{pr.slug}"
            class="group block bg-white rounded-xl overflow-hidden shadow-sm hover:shadow-md transition-all duration-200 border border-slate-200/60 flex flex-col justify-between"
          >
            <div>
              <div class="aspect-[16/10] w-full overflow-hidden bg-slate-100">
                <img
                  src="https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=600&q=80"
                  alt={pr.title}
                  class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
                  loading="lazy"
                />
              </div>

              <div class="p-6">
                <div class="flex items-center gap-1.5 text-xs text-slate-500 mb-2">
                  <MapPin size={13} class="text-brand-blue" />
                  <span>{pr.location}</span>
                </div>

                <h3 class="text-lg font-heading font-bold text-slate-900 group-hover:text-brand-blue transition-colors mb-2">
                  {pr.title}
                </h3>

                <p class="text-xs leading-relaxed text-slate-600 font-body line-clamp-3">
                  {pr.description}
                </p>
              </div>
            </div>

            <div class="px-6 pb-6 pt-0">
              <span class="text-xs font-semibold text-brand-blue group-hover:underline flex items-center gap-1">
                <span>Rincian Proyek</span>
                <ArrowRight size={13} />
              </span>
            </div>
          </a>
        {/each}
      </div>
    {/if}
  </div>
</section>
