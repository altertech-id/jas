<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { api } from '$lib/api/client';
  import type { Post } from '$lib/types';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { Calendar, User, ArrowLeft, MessageCircle } from 'lucide-svelte';

  let post: Post | null = null;
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
      post = await api.getPostBySlug(slug);
    } catch (err: any) {
      errorMsg = err.message || 'Artikel tidak ditemukan.';
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  {#if post}
    <title>{post.seo_title || post.title} - Jaya Abadi Solusi</title>
    <meta name="description" content={post.seo_description || post.excerpt} />
  {:else}
    <title>Artikel - Jaya Abadi Solusi</title>
  {/if}
</svelte:head>

{#if loading}
  <div class="flex min-h-[60vh] items-center justify-center">
    <div class="h-8 w-8 animate-spin rounded-full border-4 border-navy border-t-transparent"></div>
  </div>
{:else if errorMsg || !post}
  <div class="mx-auto max-w-container px-5 py-24 text-center">
    <h1 class="text-2xl font-heading font-bold text-navy mb-2">Artikel Tidak Ditemukan</h1>
    <p class="text-sm text-ink-500 mb-6">{errorMsg}</p>
    <Button href="/blog" variant={ButtonVariant.PRIMARY}>Kembali ke Blog</Button>
  </div>
{:else}
  <article class="py-12 sm:py-20 bg-white">
    <div class="mx-auto max-w-article px-5 sm:px-8">
      <!-- Back Link -->
      <a href="/blog" class="inline-flex items-center gap-1.5 text-xs font-semibold text-brand-blue hover:underline mb-6">
        <ArrowLeft size={16} />
        <span>Kembali ke Kumpulan Artikel</span>
      </a>

      <!-- Category Pill -->
      {#if post.category}
        <div class="mb-3">
          <span class="rounded bg-slate-100 px-2.5 py-1 text-xs font-medium text-slate-700">
            {post.category.name}
          </span>
        </div>
      {/if}

      <!-- H1 Title -->
      <h1 class="text-3xl sm:text-4xl lg:text-[40px] font-heading font-extrabold text-slate-900 leading-tight mb-4">
        {post.title}
      </h1>

      <!-- Meta Bar -->
      <div class="flex items-center gap-4 text-xs text-slate-500 pb-6 border-b border-slate-200/80 mb-8 font-body">
        <div class="flex items-center gap-1.5">
          <Calendar size={13} class="text-slate-400" />
          <span>20 Okt 2026</span>
        </div>
        <div class="flex items-center gap-1.5">
          <User size={13} class="text-slate-400" />
          <span>Tim Ahli Jaya Abadi Solusi</span>
        </div>
      </div>

      <!-- Featured Image (16:10 ratio, clean radius) -->
      <div class="aspect-[16/10] w-full overflow-hidden rounded-xl bg-slate-100 shadow-sm mb-8">
        <img
          src="https://images.unsplash.com/photo-1621905251189-08b45d6a269e?auto=format&fit=crop&w=1200&q=80"
          alt={post.title}
          class="h-full w-full object-cover"
        />
      </div>

      <!-- HTML Content (Sanitized) -->
      <div class="prose max-w-none text-slate-700 leading-relaxed space-y-4 mb-12 font-body">
        {@html post.content_html}
      </div>

      <!-- CTA Box at Bottom of Article -->
      <div class="rounded-xl border border-slate-200/80 bg-slate-50 p-6 sm:p-8 text-center my-12">
        <h3 class="text-xl font-heading font-bold text-slate-900 mb-2">
          Punya Masalah Hama di Rumah atau Tempat Usaha?
        </h3>
        <p class="text-sm text-slate-600 max-w-lg mx-auto mb-6 leading-relaxed font-body">
          Konsultasikan dengan tim ahli PT Jaya Abadi Solusi untuk survei lokasi gratis dan rekomendasi penanganan bergaransi.
        </p>
        <div class="flex flex-col sm:flex-row items-center justify-center gap-3">
          <Button href="/kontak" variant={ButtonVariant.PRIMARY} size="sm">
            Minta Penawaran Gratis
          </Button>
          <Button
            href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20baru%20membaca%20artikel%20{encodeURIComponent(post.title)}"
            variant={ButtonVariant.WHATSAPP}
            size="sm"
          >
            Chat WhatsApp (0813 9948 4916)
          </Button>
        </div>
      </div>
    </div>
  </article>
{/if}
