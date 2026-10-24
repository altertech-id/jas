<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { api } from '$lib/api/client';
  import type { Service, GalleryItem } from '$lib/types';
  import Button from '$lib/components/ui/Button.svelte';
  import Badge from '$lib/components/ui/Badge.svelte';
  import Accordion from '$lib/components/ui/Accordion.svelte';
  import QuotationForm from '$lib/components/public/QuotationForm.svelte';
  import MobileStickyCTA from '$lib/components/public/MobileStickyCTA.svelte';
  import Lightbox from '$lib/components/ui/Lightbox.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { 
    ShieldCheck, 
    AlertCircle, 
    CheckCircle2, 
    ArrowRight, 
    MessageCircle,
    PhoneCall,
    Award,
    Sparkles,
    ZoomIn
  } from 'lucide-svelte';

  let service: Service | null = null;
  let galleryItems: GalleryItem[] = [];
  let loading = true;
  let errorMsg = '';

  // Lightbox state
  let lightboxOpen = false;
  let activeImageIndex = 0;

  $: currentSlug = $page.params.slug;

  onMount(async () => {
    if (currentSlug) {
      await loadData(currentSlug);
    } else {
      loading = false;
      errorMsg = 'Layanan tidak ditemukan.';
    }
  });

  async function loadData(slug: string) {
    loading = true;
    errorMsg = '';
    try {
      const [serviceData, galleryData] = await Promise.all([
        api.getServiceBySlug(slug),
        api.getGallery().catch(() => [])
      ]);
      service = serviceData;
      galleryItems = (galleryData ?? []).slice(0, 4);
    } catch (err: any) {
      errorMsg = err.message || 'Layanan tidak ditemukan.';
    } finally {
      loading = false;
    }
  }

  function openLightbox(index: number) {
    activeImageIndex = index;
    lightboxOpen = true;
  }
</script>

<svelte:head>
  {#if service}
    <title>{service.seo_title || `${service.name} - Jaya Abadi Solusi Surabaya`}</title>
    <meta name="description" content={service.seo_description || service.short_description} />
  {:else}
    <title>Layanan - Jaya Abadi Solusi</title>
  {/if}
</svelte:head>

{#if loading}
  <div class="flex min-h-[60vh] items-center justify-center">
    <div class="h-10 w-10 animate-spin rounded-full border-4 border-navy border-t-transparent"></div>
  </div>
{:else if errorMsg || !service}
  <div class="mx-auto max-w-container px-5 py-24 text-center">
    <div class="mb-4 inline-flex h-16 w-16 items-center justify-center rounded-full bg-red-50 text-cta">
      <AlertCircle size={32} />
    </div>
    <h1 class="text-2xl font-heading font-bold text-navy mb-2">Layanan Tidak Ditemukan</h1>
    <p class="text-sm text-ink-500 mb-6">{errorMsg || 'Halaman layanan yang Anda cari tidak tersedia.'}</p>
    <Button href="/layanan" variant={ButtonVariant.PRIMARY}>Kembali ke Katalog Layanan</Button>
  </div>
{:else}
  <!-- HERO SECTION -->
  <section class="bg-hero-navy py-16 sm:py-20 text-white relative">
    <div class="mx-auto max-w-container px-5 sm:px-8">
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-center">
        <!-- Left Column -->
        <div class="lg:col-span-7 flex flex-col items-start">
          <Badge variant="hero" class="mb-4">
            <ShieldCheck slot="icon" size={16} class="text-amber" />
            <span>{service.content.hero_badge || 'Layanan Spesialis Bergaransi'}</span>
          </Badge>

          <h1 class="text-3xl sm:text-4xl lg:text-5xl font-heading font-extrabold text-white leading-tight mb-4">
            {service.content.headline || service.name}
          </h1>

          <p class="text-base sm:text-lg text-on-navy-soft leading-relaxed max-w-xl mb-8">
            {service.content.tagline || service.short_description}
          </p>

          <div class="flex flex-col sm:flex-row items-stretch sm:items-center gap-4 w-full sm:w-auto">
            <Button href="#form-quotation-section" variant={ButtonVariant.PRIMARY} size="default">
              Minta Penawaran Gratis
            </Button>

            <Button
              href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20tertarik%20dengan%20layanan%20{encodeURIComponent(service.name)}"
              variant={ButtonVariant.WHATSAPP}
              size="default"
            >
              Chat via WhatsApp
            </Button>
          </div>
        </div>

        <!-- Right Column (Feature Box) -->
        <div class="lg:col-span-5">
          <div class="rounded-xl border border-white/15 bg-white/10 p-6 sm:p-8 backdrop-blur-sm shadow-lg">
            <div class="text-xs font-semibold uppercase tracking-wider text-amber mb-2">
              Standar Layanan JAS
            </div>
            <h3 class="text-xl font-heading font-bold text-white mb-4">
              Jaminan Kepuasan & Keselamatan
            </h3>
            <ul class="space-y-3 text-sm text-slate-200 mb-6 font-body">
              <li class="flex items-start gap-2.5">
                <CheckCircle2 size={17} class="text-amber mt-0.5 flex-shrink-0" />
                <span>Teknisi bersertifikasi ASPHAMI dan Kemenkes</span>
              </li>
              <li class="flex items-start gap-2.5">
                <CheckCircle2 size={17} class="text-amber mt-0.5 flex-shrink-0" />
                <span>Bahan ramah lingkungan dan terdaftar di Kementan</span>
              </li>
              <li class="flex items-start gap-2.5">
                <CheckCircle2 size={17} class="text-amber mt-0.5 flex-shrink-0" />
                <span>Laporan inspeksi digital dan sertifikat garansi resmi</span>
              </li>
              <li class="flex items-start gap-2.5">
                <CheckCircle2 size={17} class="text-amber mt-0.5 flex-shrink-0" />
                <span>Survei lokasi wilayah Surabaya tanpa dipungut biaya</span>
              </li>
            </ul>

            <div class="pt-4 border-t border-white/10 flex items-center justify-between text-xs text-slate-300">
              <span>Konsultasi Darurat 24 Jam:</span>
              <a href="tel:081399484916" class="font-bold text-amber hover:underline">0813 9948 4916</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- BAR KEPERCAYAAN -->
  <section class="bg-surface-light border-b border-slate-200/80 py-6">
    <div class="mx-auto max-w-container px-5 sm:px-8">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-center">
        <div>
          <div class="font-heading font-extrabold text-2xl sm:text-3xl text-slate-900">15+ Tahun</div>
          <div class="text-xs text-slate-500 font-medium mt-0.5">Pengalaman Pengendalian Hama</div>
        </div>
        <div>
          <div class="font-heading font-extrabold text-2xl sm:text-3xl text-slate-900">100%</div>
          <div class="text-xs text-slate-500 font-medium mt-0.5">Bahan Kimia Berizin Resmi</div>
        </div>
        <div>
          <div class="font-heading font-extrabold text-2xl sm:text-3xl text-slate-900">Garansi</div>
          <div class="text-xs text-slate-500 font-medium mt-0.5">Penanganan Ulang Gratis</div>
        </div>
        <div>
          <div class="font-heading font-extrabold text-2xl sm:text-3xl text-slate-900">24 Jam</div>
          <div class="text-xs text-slate-500 font-medium mt-0.5">Respon Layanan Konsultasi</div>
        </div>
      </div>
    </div>
  </section>

  <!-- MASALAH DAN SOLUSI SECTION (Clean minimal comparison) -->
  {#if service.content.problems && service.content.solutions}
    <section class="py-16 sm:py-24 bg-white">
      <div class="mx-auto max-w-container px-5 sm:px-8">
        <div class="text-center max-w-2xl mx-auto mb-14">
          <Badge variant="outline" class="mb-3">Analisis Kebutuhan</Badge>
          <h2 class="text-3xl sm:text-4xl font-heading font-bold text-slate-900 mb-3">
            Tantangan di Lapangan & Pendekatan Kami
          </h2>
          <p class="text-base text-slate-600 font-body leading-relaxed">
            Hama yang tidak ditangani dengan metode tepat akan terus kembali dan menyebar ke seluruh ruangan.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <!-- Problems Card -->
          <div class="rounded-xl border border-slate-200/80 bg-slate-50/80 p-6 sm:p-8">
            <div class="text-xs font-bold uppercase tracking-wider text-slate-500 mb-2">01. Identifikasi Masalah</div>
            <h3 class="text-lg font-heading font-bold text-slate-900 mb-4">
              Tantangan yang Sering Dihadapi
            </h3>
            <ul class="space-y-3.5 text-sm leading-relaxed text-slate-700 font-body">
              {#each service.content.problems as prob}
                <li class="flex items-start gap-3">
                  <span class="mt-2 h-1.5 w-1.5 rounded-full bg-slate-400 flex-shrink-0"></span>
                  <span>{prob}</span>
                </li>
              {/each}
            </ul>
          </div>

          <!-- Solutions Card -->
          <div class="rounded-xl bg-navy text-white p-6 sm:p-8 shadow-xs">
            <div class="text-xs font-bold uppercase tracking-wider text-amber mb-2">02. Penanganan Terpadu</div>
            <h3 class="text-lg font-heading font-bold text-white mb-4">
              Solusi Profesional dari Tim JAS
            </h3>
            <ul class="space-y-3.5 text-sm leading-relaxed text-slate-200 font-body">
              {#each service.content.solutions as sol}
                <li class="flex items-start gap-3">
                  <CheckCircle2 size={16} class="text-amber mt-0.5 flex-shrink-0" />
                  <span>{sol}</span>
                </li>
              {/each}
            </ul>
          </div>
        </div>
      </div>
    </section>
  {/if}

  <!-- METODE TREATMENT SECTION -->
  {#if service.content.methods && service.content.methods.length > 0}
    <section class="py-16 sm:py-24 bg-surface-light border-y border-slate-200/80">
      <div class="mx-auto max-w-container px-5 sm:px-8">
        <div class="text-center max-w-2xl mx-auto mb-14">
          <Badge variant="outline" class="mb-3">Teknologi & Pendekatan</Badge>
          <h2 class="text-3xl sm:text-4xl font-heading font-bold text-slate-900 mb-3">
            Metode Treatment Terpadu
          </h2>
          <p class="text-base text-slate-600 font-body leading-relaxed">
            Kombinasi teknik modern yang aman, efektif, dan tidak meninggalkan residu berbahaya.
          </p>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {#each service.content.methods as method, i}
            <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs">
              <div class="font-heading font-extrabold text-2xl text-slate-300 mb-2">
                0{i + 1}
              </div>
              <h3 class="text-base font-heading font-semibold text-slate-900 mb-2">
                {method.title}
              </h3>
              <p class="text-xs leading-relaxed text-slate-600 font-body">
                {method.desc}
              </p>
            </div>
          {/each}
        </div>
      </div>
    </section>
  {/if}

  <!-- DOKUMENTASI HASIL PEKERJAAN -->
  {#if (galleryItems?.length ?? 0) > 0}
    <section class="py-16 sm:py-20 bg-white">
      <div class="mx-auto max-w-container px-5 sm:px-8">
        <div class="text-center max-w-2xl mx-auto mb-12">
          <Badge variant="outline" class="mb-3">Bukti Nyata</Badge>
          <h2 class="text-3xl font-heading font-bold text-slate-900 mb-2">
            Dokumentasi Pengerjaan Lapangan
          </h2>
          <p class="text-sm text-slate-600 font-body">
            Foto asli teknisi kami saat melakukan inspeksi dan penanganan di berbagai properti klien.
          </p>
        </div>

        <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
          {#each galleryItems as item, idx}
            <button
              type="button"
              class="group relative aspect-[3/2] w-full overflow-hidden rounded-lg border border-slate-200/80 bg-slate-100 focus:outline-none"
              on:click={() => openLightbox(idx)}
            >
              <img
                src={item.media?.public_url || 'https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=600&q=80'}
                alt={item.caption || 'Dokumentasi JAS'}
                class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
                loading="lazy"
                width="600"
                height="400"
              />
              <div class="absolute inset-0 flex items-center justify-center bg-slate-900/40 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                <ZoomIn size={22} class="text-white" />
              </div>
            </button>
          {/each}
        </div>
      </div>
    </section>
  {/if}

  <!-- FORM QUOTATION & CONTACT SECTION -->
  <section id="form-quotation-section" class="py-16 sm:py-24 bg-surface-light border-t border-slate-200/80 scroll-mt-20">
    <div class="mx-auto max-w-container px-5 sm:px-8">
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-start">
        <div class="lg:col-span-5">
          <Badge variant="outline" class="mb-3">Estimasi Biaya</Badge>
          <h2 class="text-3xl font-heading font-bold text-slate-900 mb-4">
            Minta Penawaran Layanan {service.name}
          </h2>
          <p class="text-sm leading-relaxed text-slate-600 font-body mb-6">
            Lengkapi formulir di samping untuk mendapatkan estimasi penanganan dan penjadwalan survei lokasi gratis oleh teknisi kami.
          </p>

          <div class="rounded-xl border border-slate-200/80 bg-white p-6 space-y-4 mb-6 shadow-xs">
            <div class="flex items-center gap-3">
              <PhoneCall size={18} class="text-slate-500" />
              <div>
                <div class="text-xs text-slate-500 font-medium">Telepon Langsung</div>
                <a href="tel:081399484916" class="text-base font-bold text-slate-900 hover:text-brand-blue transition-colors">
                  0813 9948 4916
                </a>
              </div>
            </div>

            <div class="flex items-center gap-3 pt-3 border-t border-slate-100">
              <MessageCircle size={18} class="text-emerald-600" />
              <div>
                <div class="text-xs text-slate-500 font-medium">Konsultasi WhatsApp</div>
                <a 
                  href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20ingin%20tanya%20layanan%20{encodeURIComponent(service.name)}" 
                  target="_blank" 
                  rel="noopener noreferrer" 
                  class="text-base font-bold text-slate-900 hover:text-emerald-600 transition-colors"
                >
                  0813 9948 4916
                </a>
              </div>
            </div>
          </div>
        </div>

        <div class="lg:col-span-7">
          <QuotationForm
            selectedServiceSlug={service.slug}
            sourcePage={`Landing Page: ${service.name}`}
          />
        </div>
      </div>
    </div>
  </section>

  <!-- FAQ SECTION (IF ANY) -->
  {#if service.content.faq && service.content.faq.length > 0}
    <section class="py-16 sm:py-20 bg-white border-t border-line-soft">
      <div class="mx-auto max-w-3xl px-5 sm:px-8">
        <div class="text-center mb-10">
          <Badge variant="outline" class="mb-3">Tanya Jawab</Badge>
          <h2 class="text-2xl sm:text-3xl font-heading font-bold text-navy">
            Pertanyaan Seputar Layanan {service.name}
          </h2>
        </div>

        <div class="space-y-3">
          {#each service.content.faq as f, idx}
            <Accordion title={f.q} open={idx === 0}>
              {f.a}
            </Accordion>
          {/each}
        </div>
      </div>
    </section>
  {/if}

  <!-- Mobile Sticky CTA Bar -->
  <MobileStickyCTA serviceName={service.name} />

  <!-- Lightbox Modal -->
  {#if (galleryItems?.length ?? 0) > 0}
    <Lightbox
      show={lightboxOpen}
      imageUrl={galleryItems[activeImageIndex]?.media?.public_url || ''}
      caption={galleryItems[activeImageIndex]?.caption || ''}
      hasPrev={activeImageIndex > 0}
      hasNext={activeImageIndex < galleryItems.length - 1}
      on:close={() => lightboxOpen = false}
      on:prev={() => activeImageIndex--}
      on:next={() => activeImageIndex++}
    />
  {/if}
{/if}
