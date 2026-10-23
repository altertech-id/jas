<script lang="ts">
  import Badge from '$lib/components/ui/Badge.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { ShieldCheck, Award, FileCheck } from 'lucide-svelte';
</script>

<svelte:head>
  <title>Tentang Kami - Profil PT Jaya Abadi Solusi Surabaya</title>
  <meta name="description" content="Profil perusahaan, legalitas resmi, visi misi, dan komitmen PT Jaya Abadi Solusi dalam menyediakan jasa pest control terpercaya di Surabaya sejak 2010." />
</svelte:head>

<!-- Header -->
<section class="bg-hero-navy py-16 sm:py-20 text-white text-center">
  <div class="mx-auto max-w-container px-5 sm:px-8 max-w-3xl">
    <Badge variant="hero" class="mb-3">Tentang Kami</Badge>
    <h1 class="text-3xl sm:text-5xl font-heading font-extrabold text-white mb-4">
      Profil PT Jaya Abadi Solusi
    </h1>
    <p class="text-base text-slate-300 leading-relaxed font-body">
      Mitra terpercaya pengendalian hama dan sanitasi lingkungan sejak 2010 di Surabaya dan Jawa Timur.
    </p>
  </div>
</section>

<!-- Story Section -->
<section class="py-16 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
      <div class="lg:col-span-6">
        <Badge variant="outline" class="mb-3">Sejarah Singkat</Badge>
        <h2 class="text-3xl sm:text-4xl font-heading font-bold text-slate-900 mb-4">
          Lebih dari 15 Tahun Memberikan Kenyamanan Bebas Hama
        </h2>
        <div class="space-y-4 text-base text-slate-600 font-body leading-relaxed">
          <p>
            PT Jaya Abadi Solusi (JAS) berdiri di Surabaya dengan satu komitmen mendasar: memberikan perlindungan menyeluruh bagi properti dan kesehatan masyarakat dari gangguan hama.
          </p>
          <p>
            Berawal dari layanan skala hunian pribadi, kini JAS telah dipercaya oleh berbagai fasilitas pergudangan logistik di kawasan Rungkut dan Tanjung Perak, jaringan restoran, pabrik manufaktur, hotel, hingga ribuan rumah tangga di Jawa Timur.
          </p>
          <p>
            Kami menerapkan sistem Integrated Pest Management (IPM) yang berkelanjutan, ramah lingkungan, dan aman bagi keluarga maupun aktivitas operasional bisnis Anda.
          </p>
        </div>
      </div>

      <div class="lg:col-span-6">
        <div class="relative overflow-hidden rounded-xl border border-slate-200/80 bg-slate-100 shadow-sm">
          <img
            src="https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=800&q=80"
            alt="Tim Profesional Jaya Abadi Solusi"
            class="h-[380px] w-full object-cover"
          />
        </div>
      </div>
    </div>
  </div>
</section>

<!-- Visi & Misi (Clean, minimal typography) -->
<section class="py-16 sm:py-24 bg-surface-light border-y border-slate-200/80">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <!-- Visi Card -->
      <div class="rounded-xl border border-slate-200/80 bg-white p-8 shadow-xs flex flex-col">
        <div class="text-xs font-bold uppercase tracking-wider text-slate-500 mb-2">01. Visi Perusahaan</div>
        <h3 class="text-2xl font-heading font-bold text-slate-900 mb-3">Visi Kami</h3>
        <p class="text-base text-slate-600 font-body leading-relaxed">
          Menjadi perusahaan penyedia jasa pengendalian hama dan sanitasi higienis terdepan di Indonesia yang dikenal karena keandalan teknis berstandar tinggi, kepuasan pelanggan, dan integritas kerja yang kokoh.
        </p>
      </div>

      <!-- Misi Card -->
      <div class="rounded-xl border border-slate-200/80 bg-white p-8 shadow-xs flex flex-col">
        <div class="text-xs font-bold uppercase tracking-wider text-slate-500 mb-2">02. Misi Perusahaan</div>
        <h3 class="text-2xl font-heading font-bold text-slate-900 mb-3">Misi Kami</h3>
        <ul class="space-y-3 text-sm text-slate-600 font-body leading-relaxed">
          <li class="flex items-start gap-2.5">
            <span class="mt-2 h-1.5 w-1.5 rounded-full bg-slate-400 flex-shrink-0"></span>
            <span>Menyediakan solusi pembasmian hama yang tuntas, aman, dan bertanggung jawab terhadap kelestarian lingkungan.</span>
          </li>
          <li class="flex items-start gap-2.5">
            <span class="mt-2 h-1.5 w-1.5 rounded-full bg-slate-400 flex-shrink-0"></span>
            <span>Mengembangkan sumber daya teknisi profesional yang tersertifikasi kompetensi resmi secara berkelanjutan.</span>
          </li>
          <li class="flex items-start gap-2.5">
            <span class="mt-2 h-1.5 w-1.5 rounded-full bg-slate-400 flex-shrink-0"></span>
            <span>Menjaga standar kualitas pengerjaan dengan peralatan modern dan respon cepat 24 jam bagi pelanggan.</span>
          </li>
        </ul>
      </div>
    </div>
  </div>
</section>

<!-- Legalitas & Sertifikasi -->
<section class="py-16 sm:py-24 bg-white">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="text-center max-w-2xl mx-auto mb-14">
      <Badge variant="outline" class="mb-3">Kepatuhan Regulasi</Badge>
      <h2 class="text-3xl sm:text-4xl font-heading font-bold text-slate-900 mb-3">
        Legalitas Resmi & Sertifikasi Industri
      </h2>
      <p class="text-base text-slate-600 font-body leading-relaxed">
        Seluruh operasional PT Jaya Abadi Solusi berlandaskan izin resmi dari instansi pemerintah dan asosiasi profesi.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs">
        <div class="text-slate-900 font-heading font-bold text-base mb-1.5">
          Izin Operasional IPPH
        </div>
        <div class="text-xs text-slate-500 mb-2 font-mono">No. 503/442.1/IPPH/Surabaya</div>
        <div class="text-xs font-medium text-slate-600">Dinas Kesehatan Kota Surabaya</div>
      </div>

      <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs">
        <div class="text-slate-900 font-heading font-bold text-base mb-1.5">
          Sertifikasi Standar ISPM 15
        </div>
        <div class="text-xs text-slate-500 mb-2 font-mono">Reg. Barantan ID-0042-MB/PH3</div>
        <div class="text-xs font-medium text-slate-600">Badan Karantina Pertanian RI</div>
      </div>

      <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs">
        <div class="text-slate-900 font-heading font-bold text-base mb-1.5">
          Keanggotaan ASPHAMI
        </div>
        <div class="text-xs text-slate-500 mb-2 font-mono">Anggota Resmi DPD Jawa Timur</div>
        <div class="text-xs font-medium text-slate-600">Asosiasi Pengendalian Hama Indonesia</div>
      </div>
    </div>
  </div>
</section>