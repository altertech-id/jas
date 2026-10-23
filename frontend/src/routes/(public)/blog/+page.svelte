<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Post, Category } from '$lib/types';
  import Badge from '$lib/components/ui/Badge.svelte';
  import { Calendar, ArrowRight } from 'lucide-svelte';

  let posts: Post[] = [];
  let categories: Category[] = [];
  let selectedCategory = '';
  let loading = true;

  onMount(async () => {
    try {
      const [catsRes, postsRes] = await Promise.all([
        api.getCategories().catch(() => []),
        api.getPosts()
      ]);
      categories = catsRes;
      posts = postsRes.data;
    } finally {
      loading = false;
    }
  });

  async function filterCategory(slug: string) {
    selectedCategory = slug;
    loading = true;
    try {
      const res = await api.getPosts({ category: slug });
      posts = res.data;
    } finally {
      loading = false;
    }
  }
</script>

<svelte:head>
  <title>Blog & Panduan Pengendalian Hama - Jaya Abadi Solusi</title>
  <meta name="description" content="Kumpulan tips pencegahan hama rumah dan industri, sanitasi bisnis, panduan karantina fumigasi dan standar kebersihan higienis di Surabaya." />
</svelte:head>

<!-- Header -->
<section class="bg-hero-navy py-16 sm:py-20 text-white text-center">
  <div class="mx-auto max-w-container px-5 sm:px-8 max-w-3xl">
    <Badge variant="hero" class="mb-3">Edukasi & Tips</Badge>
    <h1 class="text-3xl sm:text-5xl font-heading font-extrabold text-white mb-4">
      Artikel & Panduan Bebas Hama
    </h1>
    <p class="text-base text-on-navy-soft leading-relaxed">
      Informasi praktis, tips pencegahan hama harian, dan standar sanitasi untuk menjaga properti Anda tetap aman dan higienis.
    </p>
  </div>
</section>

<!-- Filter Tabs & Posts -->
<section class="py-16 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <!-- Category Tabs -->
    <div class="flex flex-wrap items-center justify-center gap-1.5 mb-12">
      <button
        type="button"
        class="rounded-lg px-4 py-2 text-xs font-medium transition-colors {selectedCategory === '' ? 'bg-navy text-white shadow-xs' : 'bg-slate-100 text-slate-700 hover:bg-slate-200'}"
        on:click={() => filterCategory('')}
      >
        Semua Kategori
      </button>

      {#each categories as cat}
        <button
          type="button"
          class="rounded-lg px-4 py-2 text-xs font-medium transition-colors {selectedCategory === cat.slug ? 'bg-navy text-white shadow-xs' : 'bg-slate-100 text-slate-700 hover:bg-slate-200'}"
          on:click={() => filterCategory(cat.slug)}
        >
          {cat.name}
        </button>
      {/each}
    </div>

    <!-- Posts Grid -->
    {#if loading}
      <div class="flex justify-center py-16">
        <div class="h-8 w-8 animate-spin rounded-full border-4 border-navy border-t-transparent"></div>
      </div>
    {:else if posts.length === 0}
      <div class="text-center py-16 text-slate-500 font-body">
        Belum ada artikel pada kategori ini.
      </div>
    {:else}
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        {#each posts as post}
          <a
            href="/blog/{post.slug}"
            class="group block bg-white rounded-xl overflow-hidden shadow-sm hover:shadow-md transition-all duration-200 border border-slate-200/60 flex flex-col justify-between"
          >
            <div>
              <div class="aspect-[16/10] w-full overflow-hidden bg-slate-100">
                <img
                  src="https://images.unsplash.com/photo-1621905251189-08b45d6a269e?auto=format&fit=crop&w=600&q=80"
                  alt={post.title}
                  class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
                  loading="lazy"
                />
              </div>

              <div class="p-6">
                <div class="flex items-center justify-between text-xs text-slate-500 mb-2">
                  <div class="flex items-center gap-1.5">
                    <Calendar size={13} class="text-slate-400" />
                    <span>20 Okt 2026</span>
                  </div>
                  {#if post.category}
                    <span class="rounded bg-slate-100 px-2 py-0.5 text-[11px] font-medium text-slate-700">
                      {post.category.name}
                    </span>
                  {/if}
                </div>

                <h3 class="text-lg font-heading font-bold text-slate-900 group-hover:text-brand-blue transition-colors mb-2 line-clamp-2">
                  {post.title}
                </h3>

                <p class="text-xs leading-relaxed text-slate-600 font-body line-clamp-3">
                  {post.excerpt}
                </p>
              </div>
            </div>

            <div class="px-6 pb-6 pt-0">
              <span class="text-xs font-semibold text-brand-blue group-hover:underline flex items-center gap-1">
                <span>Baca Selengkapnya</span>
                <ArrowRight size={13} />
              </span>
            </div>
          </a>
        {/each}
      </div>
    {/if}
  </div>
</section>
