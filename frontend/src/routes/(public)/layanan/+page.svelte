<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Service } from '$lib/types';
  import ServiceCard from '$lib/components/ui/ServiceCard.svelte';
  import Badge from '$lib/components/ui/Badge.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';

  let services: Service[] = [];
  let loading = true;

  onMount(async () => {
    try {
      services = await api.getServices();
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  <title>Layanan Pengendalian Hama & Sanitasi - Jaya Abadi Solusi Surabaya</title>
  <meta name="description" content="Katalog lengkap jasa pest control, rodent control tikus, anti rayap, fumigasi PE standar karantina, dan disinfeksi di Surabaya bergaransi resmi." />
</svelte:head>

<!-- Header -->
<section class="bg-hero-navy py-16 sm:py-20 text-white">
  <div class="mx-auto max-w-container px-5 sm:px-8 text-center max-w-3xl">
    <div class="mb-3">
      <Badge variant="hero">Solusi Terpadu</Badge>
    </div>
    <h1 class="text-3xl sm:text-5xl font-heading font-extrabold text-white mb-4">
      Layanan Profesional Jaya Abadi Solusi
    </h1>
    <p class="text-base sm:text-lg text-on-navy-soft leading-relaxed">
      Metode terpadu Integrated Pest Management (IPM) dengan bahan kimia berizin resmi dan teknisi bersertifikasi untuk melindungi aset Anda secara menyeluruh.
    </p>
  </div>
</section>

<!-- Services Grid -->
<section class="py-16 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      {#each services as item}
        <ServiceCard
          slug={item.slug}
          name={item.name}
          description={item.short_description}
          imageUrl={item.hero_media?.public_url}
        />
      {/each}
    </div>
  </div>
</section>

<!-- Band CTA -->
<section class="bg-surface-light border-t border-slate-200/80 py-16 text-center">
  <div class="mx-auto max-w-2xl px-5">
    <h2 class="text-2xl sm:text-3xl font-heading font-bold text-slate-900 mb-3">
      Butuh Rekomendasi Penanganan Spesifik?
    </h2>
    <p class="text-sm text-slate-600 font-body mb-6 leading-relaxed">
      Konsultasikan masalah hama pada properti Anda langsung dengan konsultan teknis kami via WhatsApp atau jadwalkan survei gratis.
    </p>
    <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
      <Button href="/kontak" variant={ButtonVariant.PRIMARY}>
        Minta Penawaran Gratis
      </Button>
      <Button
        href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20ingin%20tanya%20layanan"
        variant={ButtonVariant.WHATSAPP}
      >
        Chat WhatsApp 0813 9948 4916
      </Button>
    </div>
  </div>
</section>
