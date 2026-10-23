<script lang="ts">
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { MapPin, Phone, Mail, Clock } from 'lucide-svelte';
</script>

<footer class="bg-navy-deep text-slate-300 pt-16 pb-8">
  <div class="mx-auto max-w-container px-5 sm:px-8">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-10 pb-12">
      <!-- Col 1: Logo & Company Info -->
      <div class="flex flex-col">
        <a href="/" class="flex flex-col mb-4">
          <span class="font-heading text-xl font-extrabold tracking-tight text-white leading-none">
            JAYA ABADI
          </span>
          <span class="font-heading text-xs font-bold tracking-widest text-cta leading-tight">
            SOLUSI
          </span>
        </a>

        <p class="text-xs sm:text-sm leading-relaxed text-slate-400 mb-6 font-body">
          Penyedia jasa pest control, anti rayap bergaransi, fumigasi karantina, dan disinfeksi ruangan terpercaya di Surabaya dan sekitarnya sejak 2010.
        </p>

        <div class="flex items-center gap-3 text-slate-400">
          <a href="https://instagram.com/jayaabadisolusi" target="_blank" rel="noopener noreferrer" class="hover:text-amber-400 transition-colors" aria-label="Instagram">
            <svg class="h-5 w-5 fill-current" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/></svg>
          </a>
          <a href="https://facebook.com/jayaabadisolusi" target="_blank" rel="noopener noreferrer" class="hover:text-amber-400 transition-colors" aria-label="Facebook">
            <svg class="h-5 w-5 fill-current" viewBox="0 0 24 24"><path d="M9 8H6v4h3v12h5V12h3.642L18 8h-4V6.333C14 5.374 14.5 5 15.667 5H18V0h-3.808C10.597 0 9 1.583 9 4.615V8z"/></svg>
          </a>
          <a href="https://linkedin.com/company/jayaabadisolusi" target="_blank" rel="noopener noreferrer" class="hover:text-amber-400 transition-colors" aria-label="LinkedIn">
            <svg class="h-5 w-5 fill-current" viewBox="0 0 24 24"><path d="M4.98 3.5c0 1.381-1.11 2.5-2.48 2.5s-2.48-1.119-2.48-2.5c0-1.38 1.11-2.5 2.48-2.5s2.48 1.12 2.48 2.5zm.02 4.5h-5v16h5v-16zm7.982 0h-4.968v16h4.969v-8.399c0-4.67 6.029-5.052 6.029 0v8.399h4.988v-10.131c0-7.88-8.922-7.593-11.018-3.714v-2.155z"/></svg>
          </a>
        </div>
      </div>

      <!-- Col 2: Contact Info -->
      <div>
        <h4 class="text-xs font-heading font-bold uppercase tracking-wider text-white mb-4">
          Kontak Kami
        </h4>
        <ul class="space-y-3 text-xs sm:text-sm text-slate-400 font-body">
          <li class="flex items-start gap-2.5">
            <MapPin size={16} class="text-amber-400 mt-0.5 flex-shrink-0" />
            <span>Jl. Rungkut Industri Raya No. 45, Surabaya, Jawa Timur 60293</span>
          </li>
          <li class="flex items-center gap-2.5">
            <Phone size={16} class="text-amber-400 flex-shrink-0" />
            <a href="tel:081399484916" class="hover:text-white transition-colors text-white font-medium">
              0813 9948 4916
            </a>
          </li>
          <li class="flex items-center gap-2.5">
            <Mail size={16} class="text-amber-400 flex-shrink-0" />
            <a href="mailto:altertech.id@gmail.com" class="hover:text-white transition-colors">
              altertech.id@gmail.com
            </a>
          </li>
          <li class="flex items-start gap-2.5">
            <Clock size={16} class="text-amber-400 mt-0.5 flex-shrink-0" />
            <span>Senin - Sabtu: 08.00 - 17.00 WIB (Layanan Darurat 24 Jam)</span>
          </li>
        </ul>
      </div>

      <!-- Col 3: Layanan Kami -->
      <div>
        <h4 class="text-xs font-heading font-bold uppercase tracking-wider text-white mb-4">
          Layanan Unggulan
        </h4>
        <ul class="space-y-2 text-xs sm:text-sm text-slate-400 font-body">
          <li>
            <a href="/layanan/termite-control" class="hover:text-amber-400 transition-colors">
              Anti Rayap (Termite Control)
            </a>
          </li>
          <li>
            <a href="/layanan/rodent-control" class="hover:text-amber-400 transition-colors">
              Pengendalian Tikus (Rodent Control)
            </a>
          </li>
          <li>
            <a href="/layanan/pest-control" class="hover:text-amber-400 transition-colors">
              Pengendalian Nyamuk & Serangga
            </a>
          </li>
          <li>
            <a href="/layanan/fumigasi-pe" class="hover:text-amber-400 transition-colors">
              Fumigasi Kontainer & Gudang
            </a>
          </li>
          <li>
            <a href="/layanan/disinfeksi" class="hover:text-amber-400 transition-colors">
              Disinfeksi & Sterilisasi Ruangan
            </a>
          </li>
        </ul>
      </div>

      <!-- Col 4: Request Survey Box -->
      <div class="rounded-xl bg-white/5 p-6 flex flex-col justify-between border border-white/5">
        <div>
          <h4 class="text-sm font-heading font-bold text-white mb-2">
            Konsultasi Gratis
          </h4>
          <p class="text-xs leading-relaxed text-slate-400 mb-5 font-body">
            Jadwalkan survei ke lokasi Anda di Surabaya. Tim kami siap menganalisis tingkat infestasi hama secara gratis.
          </p>
        </div>

        <Button href="/kontak" variant={ButtonVariant.PRIMARY} size="sm" fullWidth>
          Minta Penawaran
        </Button>
      </div>
    </div>

    <!-- Bottom Bar -->
    <div class="border-t border-slate-800 pt-6 flex flex-col sm:flex-row items-center justify-between text-[11px] text-slate-500 gap-3 font-body">
      <p>&copy; 2026 PT Jaya Abadi Solusi. Seluruh hak cipta dilindungi undang-undang.</p>
      <div class="flex gap-4">
        <a href="/tentang-kami" class="hover:text-slate-300 transition-colors">Tentang Kami</a>
        <a href="/kontak" class="hover:text-slate-300 transition-colors">Kontak</a>
        <a href="/admin/login" class="hover:text-slate-300 transition-colors">Portal Staf</a>
      </div>
    </div>
  </div>
</footer>
