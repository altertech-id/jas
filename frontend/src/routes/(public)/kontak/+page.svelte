<script lang="ts">
  import Badge from '$lib/components/ui/Badge.svelte';
  import QuotationForm from '$lib/components/public/QuotationForm.svelte';
  import { MapPin, Phone, Mail, Clock, MessageCircle } from 'lucide-svelte';
</script>

<svelte:head>
  <title>Hubungi Kami - PT Jaya Abadi Solusi Surabaya</title>
  <meta name="description" content="Kontak langsung tim Jaya Abadi Solusi Surabaya. Telepon & WhatsApp 0813 9948 4916, email altertech.id@gmail.com. Konsultasi dan survei gratis." />
</svelte:head>

<!-- Header -->
<section class="bg-hero-navy py-16 sm:py-20 text-white text-center">
  <div class="mx-auto max-w-container px-5 sm:px-8 max-w-3xl">
    <Badge variant="hero" class="mb-3">Layanan Pelanggan</Badge>
    <h1 class="text-3xl sm:text-5xl font-heading font-extrabold text-white mb-4">
      Hubungi Tim Ahli Hama Kami
    </h1>
    <p class="text-base text-slate-300 leading-relaxed font-body">
      Konsultasikan masalah hama pada properti Anda. Kami siap menjadwalkan survei lokasi dan menyusun estimasi penawaran terbaik.
    </p>
  </div>
</section>

<!-- Content: 2 Columns -->
<section class="py-16 sm:py-24 bg-surface-light">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-start">
      <!-- Left Column: Info & Map -->
      <div class="lg:col-span-5 space-y-6">
        <div class="rounded-xl border border-slate-200/80 bg-white p-6 sm:p-8 shadow-xs space-y-6">
          <h2 class="text-xl font-heading font-bold text-slate-900">
            Kantor & Kontak Resmi
          </h2>

          <div class="space-y-4 text-sm text-slate-700 font-body">
            <div class="flex items-start gap-3.5">
              <MapPin size={18} class="text-slate-400 flex-shrink-0 mt-0.5" />
              <div>
                <div class="font-medium text-slate-900">Alamat Operasional</div>
                <div class="text-slate-500 leading-relaxed mt-0.5 text-xs">
                  Jl. Rungkut Industri Raya No. 45, Surabaya, Jawa Timur 60293
                </div>
              </div>
            </div>

            <div class="flex items-start gap-3.5 pt-3 border-t border-slate-100">
              <Phone size={18} class="text-slate-400 flex-shrink-0 mt-0.5" />
              <div>
                <div class="font-medium text-slate-900">Telepon & Konsultasi</div>
                <a href="tel:081399484916" class="text-base font-bold text-slate-900 hover:text-brand-blue transition-colors">
                  0813 9948 4916
                </a>
              </div>
            </div>

            <div class="flex items-start gap-3.5 pt-3 border-t border-slate-100">
              <Mail size={18} class="text-slate-400 flex-shrink-0 mt-0.5" />
              <div>
                <div class="font-medium text-slate-900">Alamat Email</div>
                <a href="mailto:altertech.id@gmail.com" class="text-slate-600 hover:text-brand-blue transition-colors text-xs">
                  altertech.id@gmail.com
                </a>
              </div>
            </div>

            <div class="flex items-start gap-3.5 pt-3 border-t border-slate-100">
              <Clock size={18} class="text-slate-400 flex-shrink-0 mt-0.5" />
              <div>
                <div class="font-medium text-slate-900">Jam Layanan</div>
                <div class="text-slate-500 text-xs leading-relaxed mt-0.5">
                  Senin - Sabtu: 08.00 - 17.00 WIB<br>
                  <span class="font-medium text-emerald-700">Layanan Darurat 24 Jam via WhatsApp</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Big WhatsApp Button -->
          <div class="pt-2">
            <a
              href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20ingin%20konsultasi%20layanan%20hama"
              target="_blank"
              rel="noopener noreferrer"
              class="flex items-center justify-center gap-2 rounded-lg bg-emerald-600 hover:bg-emerald-700 py-3 px-4 text-sm font-medium text-white shadow-xs transition-colors w-full"
            >
              <MessageCircle size={18} />
              <span>Chat WhatsApp (0813 9948 4916)</span>
            </a>
          </div>
        </div>

        <!-- Google Maps Embed Card -->
        <div class="rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
          <div class="p-4 border-b border-slate-100">
            <span class="text-xs font-semibold uppercase tracking-wider text-slate-500">Lokasi Kantor di Google Maps</span>
          </div>
          <div class="aspect-[16/9] w-full bg-slate-100">
            <iframe
              title="Lokasi Kantor PT Jaya Abadi Solusi Surabaya"
              src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3957.3887222306766!2d112.75338167499997!3d-7.310156992697843!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd7fb10b27b5003%3A0x6b29d919ff7d7045!2sJl.%20Rungkut%20Industri%20Raya%2C%20Surabaya!5e0!3m2!1sid!2sid!4v1700000000000"
              width="100%"
              height="100%"
              style="border:0;"
              allowfullscreen={false}
              loading="lazy"
              referrerpolicy="no-referrer-when-downgrade"
            ></iframe>
          </div>
        </div>
      </div>

      <!-- Right Column: Form Quotation -->
      <div class="lg:col-span-7">
        <QuotationForm sourcePage="Halaman Kontak Resmi" />
      </div>
    </div>
  </div>
</section>