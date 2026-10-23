<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { GalleryItem } from '$lib/types';
  import Badge from '$lib/components/ui/Badge.svelte';
  import Lightbox from '$lib/components/ui/Lightbox.svelte';
  import { ZoomIn } from 'lucide-svelte';

  let items: GalleryItem[] = [];
  let selectedCategory = '';
  let loading = true;

  // Lightbox
  let lightboxOpen = false;
  let activeIndex = 0;

  const categories = ['Pest Control', 'Termite Control', 'Rodent Control', 'Fumigasi'];

  onMount(async () => {
    await fetchGallery('');
  });

  async function fetchGallery(cat: string) {
    selectedCategory = cat;
    loading = true;
    try {
      items = await api.getGallery(cat);
    } finally {
      loading = false;
    }
  }

  function openImage(idx: number) {
    activeIndex = idx;
    lightboxOpen = true;
  }
</script>

<svelte:head>
  <title>Galeri Dokumentasi Lapangan - Jaya Abadi Solusi Surabaya</title>
  <meta name="description" content="Koleksi foto dokumentasi pengerjaan pest control, penyemprotan kecoa restoran, injeksi anti rayap pondasi, dan fumigasi kontainer di Surabaya." />
</svelte:head>

<!-- Header -->
<section class="bg-hero-navy py-16 sm:py-20 text-white text-center">
  <div class="mx-auto max-w-container px-5 sm:px-8 max-w-3xl">
    <Badge variant="hero" class="mb-3">Dokumentasi Asli</Badge>
    <h1 class="text-3xl sm:text-5xl font-heading font-extrabold text-white mb-4">
      Galeri Dokumentasi Lapangan
    </h1>
    <p class="text-base text-on-navy-soft leading-relaxed">
      Kumpulan foto pengerjaan langsung oleh teknisi bersertifikasi PT Jaya Abadi Solusi di berbagai lokasi proyek di Jawa Timur.
    </p>
  </div>
</section>

<!-- Gallery Filter & Grid -->
<section class="py-16 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <!-- Category Tabs -->
    <div class="flex flex-wrap items-center justify-center gap-1.5 mb-12">
      <button
        type="button"
        class="rounded-lg px-4 py-2 text-xs font-medium transition-colors {selectedCategory === '' ? 'bg-navy text-white shadow-xs' : 'bg-slate-100 text-slate-700 hover:bg-slate-200'}"
        on:click={() => fetchGallery('')}
      >
        Semua Kategori
      </button>

      {#each categories as cat}
        <button
          type="button"
          class="rounded-lg px-4 py-2 text-xs font-medium transition-colors {selectedCategory === cat ? 'bg-navy text-white shadow-xs' : 'bg-slate-100 text-slate-700 hover:bg-slate-200'}"
          on:click={() => fetchGallery(cat)}
        >
          {cat}
        </button>
      {/each}
    </div>

    <!-- Images Grid (Ratio 3:2, gap 12px, 2-5 columns) -->
    {#if loading}
      <div class="flex justify-center py-16">
        <div class="h-8 w-8 animate-spin rounded-full border-4 border-navy border-t-transparent"></div>
      </div>
    {:else if items.length === 0}
      <div class="text-center py-16 text-slate-500 font-body">
        Belum ada foto dokumentasi pada kategori ini.
      </div>
    {:else}
      <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-3">
        {#each items as item, idx}
          <button
            type="button"
            class="group relative aspect-[3/2] w-full overflow-hidden rounded-lg border border-slate-200/80 bg-slate-100 focus:outline-none"
            on:click={() => openImage(idx)}
          >
            <img
              src={item.media?.public_url || 'https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=600&q=80'}
              alt={item.caption || 'Dokumentasi Jaya Abadi Solusi'}
              class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
              loading="lazy"
              width="600"
              height="400"
            />

            <!-- Hover Overlay -->
            <div class="absolute inset-0 flex flex-col items-center justify-center bg-slate-900/60 p-3 text-center opacity-0 transition-opacity duration-200 group-hover:opacity-100">
              <ZoomIn size={20} class="text-white mb-1.5" />
              <p class="line-clamp-2 text-[11px] font-medium text-white/90">
                {item.caption}
              </p>
            </div>
          </button>
        {/each}
      </div>
    {/if}
  </div>
</section>

<!-- Lightbox Modal -->
{#if items.length > 0}
  <Lightbox
    show={lightboxOpen}
    imageUrl={items[activeIndex]?.media?.public_url || ''}
    caption={items[activeIndex]?.caption || ''}
    hasPrev={activeIndex > 0}
    hasNext={activeIndex < items.length - 1}
    on:close={() => lightboxOpen = false}
    on:prev={() => activeIndex--}
    on:next={() => activeIndex++}
  />
{/if}
