<script lang="ts">
  import { onMount } from "svelte";
  import { api } from "$lib/api/client";
  import type { Service, Project, Post } from "$lib/types";
  import Button from "$lib/components/ui/Button.svelte";
  import ServiceCard from "$lib/components/ui/ServiceCard.svelte";
  import WaveDivider from "$lib/components/ui/WaveDivider.svelte";
  import Accordion from "$lib/components/ui/Accordion.svelte";
  import { ButtonVariant } from "$lib/types/enums";
  import {
    ShieldCheck,
    Headphones,
    ArrowRight,
    MapPin,
    Calendar,
  } from "lucide-svelte";

  let services: Service[] = [];
  let projects: Project[] = [];
  let posts: Post[] = [];
  let loading = true;

  onMount(async () => {
    try {
      const [servicesRes, projectsRes, postsRes] = await Promise.all([
        api.getServices().catch(() => []),
        api.getProjects().catch(() => []),
        api.getPosts({ limit: 3 }).catch(() => ({ data: [] })),
      ]);

      services = servicesRes;
      projects = projectsRes.slice(0, 3);
      posts = postsRes.data.slice(0, 3);
    } finally {
      loading = false;
    }
  });

  const propertyTypes = [
    {
      title: "Rumah Tinggal & Apartemen Pribadi",
      desc: "Melindungi kenyamanan keluarga dan kebersihan dapur dari gangguan kecoa, tikus, rayap perabot kayu, dan nyamuk demam berdarah dengan metode tanpa bau residu.",
    },
    {
      title: "Restoran, Kafe & Fasilitas Kuliner",
      desc: "Sanitasi dapur ketat sesuai standar dinas kesehatan dan audit higienitas. Penanganan lalat dan kecoa jerman tanpa mengganggu jam operasional makan tamu.",
    },
    {
      title: "Gudang & Pusat Logistik",
      desc: "Pengendalian tikus pengerat dan rayap tanah untuk mencegah kerusakan kardus kemasan, produk komoditas bernilai tinggi, serta instalasi kelistrikan pabrik.",
    },
    {
      title: "Gedung Perkantoran, Klinik & Perhotelan",
      desc: "Treatment berkala dan sterilisasi ruangan menjaga reputasi bisnis Anda tetap profesional, higienis, dan terbebas dari keluhan penyewa atau tamu.",
    },
  ];
</script>

<svelte:head>
  <title>Jaya Abadi Solusi - Jasa Pengendalian Hama & Sanitasi Surabaya</title>
  <meta
    name="description"
    content="Jasa pembasmi hama tikus, rayap, serangga, fumigasi PE, dan disinfeksi ruangan di Surabaya. Berpengalaman 15+ tahun, bergaransi resmi. Hubungi 0813 9948 4916."
  />
</svelte:head>

<!-- HERO SECTION (Restored to earlier clean design without arch crop or fake logos) -->
<section
  class="relative bg-hero-navy pt-16 sm:pt-24 pb-12 sm:pb-16 text-white overflow-hidden"
>
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-center">
      <!-- Left Column -->
      <div class="lg:col-span-7 flex flex-col items-start z-10">
        <div
          class="inline-flex items-center gap-2 rounded-full bg-white/10 px-4 py-1 text-xs font-semibold text-white mb-5"
        >
          <ShieldCheck size={14} class="text-amber-400" />
          <span>Terpercaya Sejak 2010</span>
        </div>

        <h1
          class="text-3xl sm:text-5xl lg:text-[52px] font-heading font-extrabold text-white leading-[1.15] tracking-tight mb-5"
        >
          Solusi Tuntas Pengendalian Hama & Sanitasi Higienis
        </h1>

        <p
          class="text-base sm:text-lg text-slate-300 leading-relaxed font-body max-w-xl mb-8"
        >
          Kami melindungi hunian, kantor, restoran, dan pergudangan Anda di
          Surabaya dari gangguan rayap, tikus, kecoa, serta bakteri patogen
          dengan metode bersertifikat dan bergaransi.
        </p>

        <div
          class="flex flex-col sm:flex-row items-stretch sm:items-center gap-4 w-full sm:w-auto mb-8"
        >
          <Button href="/kontak" variant={ButtonVariant.PRIMARY} size="default">
            Minta Penawaran Gratis
          </Button>

          <Button
            href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20ingin%20konsultasi%20layanan%20pengendalian%20hama"
            variant={ButtonVariant.WHATSAPP}
            size="default"
          >
            Chat via WhatsApp
          </Button>
        </div>

        <!-- Customer Service Phone Box -->
        <div
          class="flex items-center gap-3.5 rounded-xl bg-white/10 px-4 py-2.5"
        >
          <div
            class="flex h-10 w-10 items-center justify-center rounded-lg bg-amber-400/20 text-amber-400"
          >
            <Headphones size={20} />
          </div>
          <div>
            <div
              class="text-[11px] uppercase tracking-wider text-slate-300 font-semibold"
            >
              Konsultasi Cepat
            </div>
            <a
              href="tel:081399484916"
              class="text-base font-heading font-bold text-white hover:text-amber-400 transition-colors"
            >
              0813 9948 4916
            </a>
          </div>
        </div>
      </div>

      <!-- Right Column: Female Technician Photo -->
      <div class="lg:col-span-5 relative flex justify-center">
        <div
          class="relative w-full max-w-md overflow-hidden rounded-xl shadow-lg bg-slate-800"
        >
          <img
            src="https://images.unsplash.com/photo-1573496359142-b8d87734a5a2?auto=format&fit=crop&w=800&q=80"
            alt="Customer Service dan Konsultan Pengendalian Hama Profesional Jaya Abadi Solusi Surabaya"
            class="h-[420px] sm:h-[460px] w-full object-cover object-top"
            width="800"
            height="1000"
          />

          <div
            class="absolute bottom-4 left-4 right-4 rounded-lg bg-slate-900/85 backdrop-blur-sm p-3.5 flex items-center justify-between text-white"
          >
            <div class="flex items-center gap-3">
              <span
                class="font-heading font-extrabold text-xl text-amber-400 leading-none"
                >15+</span
              >
              <div>
                <div class="text-xs font-bold leading-tight">
                  Tahun Pengalaman
                </div>
                <div class="text-[11px] text-slate-300">
                  Teknisi Berlisensi Resmi
                </div>
              </div>
            </div>
            <span class="text-xs font-semibold text-amber-400"
              >Garansi Tuntas</span
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- LAYANAN UTAMA SECTION (Clean minimal cards, no harsh AI borders) -->
<section class="py-20 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="text-center max-w-2xl mx-auto mb-16">
      <h2
        class="text-3xl sm:text-4xl font-heading font-bold text-slate-900 mb-3"
      >
        Layanan Pengendalian Hama Terpadu
      </h2>
      <p class="text-base text-slate-600 font-body leading-relaxed">
        Peralatan standar industri dan formulasi ramah lingkungan berizin
        kementerian untuk melindungi properti Anda secara menyeluruh.
      </p>
    </div>

    <!-- 5 Clean Service Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mb-12">
      {#each services as item}
        <ServiceCard
          slug={item.slug}
          name={item.name}
          description={item.short_description}
        />
      {/each}
    </div>

    <div class="text-center">
      <Button href="/layanan" variant={ButtonVariant.PRIMARY} size="default">
        <span>Lihat Semua Layanan</span>
        <ArrowRight size={16} class="ml-2" />
      </Button>
    </div>
  </div>
</section>

<!-- KEUNGGULAN SECTION (Clean minimalist layout, no pastel colored icon circles) -->
<section class="py-20 sm:py-24 bg-surface-light border-y border-slate-200/80">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
      <!-- Left Big Metric -->
      <div
        class="lg:col-span-5 bg-navy rounded-xl p-8 sm:p-10 text-white shadow-xs flex flex-col justify-between"
      >
        <div>
          <div
            class="text-amber-400 font-heading font-extrabold text-5xl sm:text-6xl leading-none mb-3"
          >
            15+ Tahun
          </div>
          <div class="text-xl font-heading font-bold text-white mb-3">
            Pengalaman Pengendalian Hama di Surabaya
          </div>
          <p class="text-sm leading-relaxed text-slate-300 font-body mb-6">
            Berdiri sejak tahun 2010, PT Jaya Abadi Solusi dipercaya oleh ribuan
            pemilik rumah, fasilitas pergudangan logistik, dan pengelola bisnis
            di Jawa Timur.
          </p>
        </div>

        <div
          class="pt-6 border-t border-white/10 flex items-center justify-between text-xs"
        >
          <span class="text-slate-300 uppercase tracking-wider font-semibold"
            >Survei Lokasi</span
          >
          <span class="font-bold text-amber-400">Gratis Area Surabaya</span>
        </div>
      </div>

      <!-- Right 4 Points (Minimalist text hierarchy) -->
      <div class="lg:col-span-7 grid grid-cols-1 sm:grid-cols-2 gap-6">
        <div
          class="bg-white rounded-xl p-6 shadow-xs border border-slate-200/70"
        >
          <div
            class="text-xs font-bold uppercase tracking-wider text-slate-500 mb-2"
          >
            01. Kualifikasi
          </div>
          <h3 class="text-base font-heading font-bold text-slate-900 mb-2">
            Teknisi Tersertifikasi Resmi
          </h3>
          <p class="text-sm text-slate-600 font-body leading-relaxed">
            Seluruh teknisi lapangan memiliki sertifikasi kompetensi resmi
            standar Kemenkes dan karantina pertanian.
          </p>
        </div>

        <div
          class="bg-white rounded-xl p-6 shadow-xs border border-slate-200/70"
        >
          <div
            class="text-xs font-bold uppercase tracking-wider text-slate-500 mb-2"
          >
            02. Garansi
          </div>
          <h3 class="text-base font-heading font-bold text-slate-900 mb-2">
            Jaminan Penanganan Ulang
          </h3>
          <p class="text-sm text-slate-600 font-body leading-relaxed">
            Sertifikat garansi resmi dengan kunjungan inspeksi ulang tanpa biaya
            bila hama target masih terdeteksi.
          </p>
        </div>

        <div
          class="bg-white rounded-xl p-6 shadow-xs border border-slate-200/70"
        >
          <div
            class="text-xs font-bold uppercase tracking-wider text-slate-500 mb-2"
          >
            03. Keamanan
          </div>
          <h3 class="text-base font-heading font-bold text-slate-900 mb-2">
            Metode Ramah Lingkungan
          </h3>
          <p class="text-sm text-slate-600 font-body leading-relaxed">
            Bahan aktif berizin resmi Kementerian Pertanian yang aman untuk
            keluarga dan hewan peliharaan.
          </p>
        </div>

        <div
          class="bg-white rounded-xl p-6 shadow-xs border border-slate-200/70"
        >
          <div
            class="text-xs font-bold uppercase tracking-wider text-slate-500 mb-2"
          >
            04. Kecepatan
          </div>
          <h3 class="text-base font-heading font-bold text-slate-900 mb-2">
            Respon Cepat 24 Jam
          </h3>
          <p class="text-sm text-slate-600 font-body leading-relaxed">
            Layanan darurat dan penjadwalan inspeksi fleksibel menyesuaikan jam
            operasional bisnis Anda.
          </p>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- TIPE PROPERTI SECTION -->
<section class="py-20 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-start">
      <div class="lg:col-span-5">
        <h2 class="text-3xl font-heading font-bold text-slate-900 mb-4">
          Perlindungan Tepat untuk Setiap Tipe Properti
        </h2>
        <p class="text-base text-slate-600 font-body leading-relaxed mb-6">
          Karakteristik hama di rumah tinggal sangat berbeda dengan di gudang
          industri atau restoran. Kami menerapkan pendekatan khusus sesuai
          kondisi lapangan.
        </p>

        <div class="p-6 rounded-xl bg-slate-50 border border-slate-200/70">
          <div class="text-sm font-bold text-slate-900 mb-1">
            Butuh Survei Langsung?
          </div>
          <p class="text-xs text-slate-600 font-body mb-4">
            Hubungi tim kami untuk mendiskusikan jadwal kunjungan inspeksi
            teknisi.
          </p>
          <Button href="/kontak" variant={ButtonVariant.PRIMARY} size="sm">
            Konsultasikan Sekarang
          </Button>
        </div>
      </div>

      <div class="lg:col-span-7 space-y-3">
        {#each propertyTypes as prop, i}
          <Accordion title={prop.title} open={i === 0}>
            {prop.desc}
          </Accordion>
        {/each}
      </div>
    </div>
  </div>
</section>

<!-- ALUR KERJA (HOW IT WORKS) SECTION (Clean minimalist process, no exaggerated shapes) -->
<section class="py-20 sm:py-24 bg-surface-light border-y border-slate-200/80">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="text-center max-w-2xl mx-auto mb-16">
      <h2 class="text-3xl font-heading font-bold text-slate-900 mb-3">
        Alur Penanganan Hama
      </h2>
      <p class="text-base text-slate-600 font-body leading-relaxed">
        Ini cara kami membantu menuntaskan masalah anda.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <!-- Step 1 -->
      <div
        class="bg-white rounded-xl p-8 shadow-xs border border-slate-200/70 flex flex-col items-start"
      >
        <span class="text-3xl font-heading font-extrabold text-slate-300 mb-3"
          >01</span
        >
        <h3 class="text-lg font-heading font-bold text-slate-900 mb-2">
          Survei & Analisis Lokasi
        </h3>
        <p class="text-sm text-slate-600 font-body leading-relaxed">
          Hubungi kami via telepon atau WhatsApp. Kami jadwalkan survei ke
          lokasi Anda untuk mengidentifikasi kondisi terkini.
        </p>
      </div>

      <!-- Step 2 -->
      <div
        class="bg-white rounded-xl p-8 shadow-xs border border-slate-200/70 flex flex-col items-start"
      >
        <span class="text-3xl font-heading font-extrabold text-slate-300 mb-3"
          >02</span
        >
        <h3 class="text-lg font-heading font-bold text-slate-900 mb-2">
          Treatment & Penanganan
        </h3>
        <p class="text-sm text-slate-600 font-body leading-relaxed">
          Penyusunan penawaran transparan diikuti eksekusi oleh tim kami
          menggunakan metode yang aman dan terpercaya.
        </p>
      </div>

      <!-- Step 3 -->
      <div
        class="bg-white rounded-xl p-8 shadow-xs border border-slate-200/70 flex flex-col items-start"
      >
        <span class="text-3xl font-heading font-extrabold text-slate-300 mb-3"
          >03</span
        >
        <h3 class="text-lg font-heading font-bold text-slate-900 mb-2">
          Evaluasi dan Monitoring
        </h3>
        <p class="text-sm text-slate-600 font-body leading-relaxed">
          Pekerjaan selesai disertai laporan dan edukasi menjaga kebersihan
          lingkungan properti Anda.
        </p>
      </div>
    </div>
  </div>
</section>

<!-- PROYEK TERBARU -->
<section class="py-20 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="flex items-end justify-between mb-12">
      <div>
        <h2 class="text-2xl sm:text-3xl font-heading font-bold text-slate-900">
          Proyek Terbaru
        </h2>
      </div>
      <a
        href="/project"
        class="text-xs sm:text-sm font-semibold text-cta hover:underline flex items-center gap-1"
      >
        <span>Lihat Semua Proyek</span>
        <ArrowRight size={14} />
      </a>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      {#each projects as pr}
        <a
          href="/project/{pr.slug}"
          class="group block bg-white rounded-xl overflow-hidden shadow-sm hover:shadow-md transition-all duration-200 border border-slate-200/60"
        >
          <div class="aspect-[16/10] w-full overflow-hidden bg-slate-100">
            <img
              src="https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=600&q=80"
              alt={pr.title}
              class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
            />
          </div>
          <div class="p-5">
            <div
              class="flex items-center gap-1.5 text-xs text-slate-500 mb-1.5"
            >
              <MapPin size={13} class="text-brand-blue" />
              <span>{pr.location}</span>
            </div>
            <h3
              class="text-base font-heading font-bold text-slate-900 group-hover:text-brand-blue transition-colors mb-2 line-clamp-1"
            >
              {pr.title}
            </h3>
            <p
              class="text-xs text-slate-600 font-body line-clamp-2 leading-relaxed"
            >
              {pr.description}
            </p>
          </div>
        </a>
      {/each}
    </div>
  </div>
</section>

<!-- BAND CTA NAVY -->
<section class="bg-navy py-16 text-white text-center">
  <div class="mx-auto max-w-container px-5 sm:px-8 max-w-2xl">
    <h2 class="text-2xl sm:text-3xl font-heading font-bold mb-3">
      Hama Bikin Resah? Serahkan pada Ahlinya
    </h2>
    <p class="text-sm text-slate-300 mb-8 leading-relaxed font-body">
      Konsultasikan sekarang, kami siap membantu dan memberikan penanganan
      bergaransi.
    </p>

    <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
      <Button href="/kontak" variant={ButtonVariant.PRIMARY} size="default">
        Dapatkan Penawaran
      </Button>
      <Button
        href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20butuh%20penanganan%20hama"
        variant={ButtonVariant.WHATSAPP}
        size="default"
      >
        Hubungi Kami
      </Button>
    </div>
  </div>
</section>
