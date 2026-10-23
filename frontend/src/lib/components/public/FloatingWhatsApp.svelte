<script lang="ts">
  import { onMount } from 'svelte';
  import { X, Send } from 'lucide-svelte';

  const phone = '6281399484916';
  const defaultMessage = 'Halo Jaya Abadi Solusi, saya ingin konsultasi dan tanya penawaran layanan pengendalian hama.';
  const waUrl = `https://wa.me/${phone}?text=${encodeURIComponent(defaultMessage)}`;

  let showGreeting = false;

  onMount(() => {
    const dismissed = sessionStorage.getItem('jas_wa_greeting_dismissed');
    if (!dismissed) {
      const timer = setTimeout(() => {
        showGreeting = true;
      }, 5000);
      return () => clearTimeout(timer);
    }
  });

  function closeGreeting() {
    showGreeting = false;
    sessionStorage.setItem('jas_wa_greeting_dismissed', 'true');
  }
</script>

<div class="fixed bottom-5 right-5 z-40 flex flex-col items-end gap-3 pointer-events-none">
  <!-- Greeting Box -->
  {#if showGreeting}
    <div 
      class="pointer-events-auto relative w-[calc(100vw-40px)] sm:w-[330px] rounded-2xl bg-[#ECE5DD] shadow-2xl border border-line-soft overflow-hidden animate-in fade-in slide-in-from-bottom-3 duration-300"
    >
      <!-- Header -->
      <div class="flex items-center justify-between bg-[#075E54] p-3 text-white">
        <div class="flex items-center gap-3">
          <div class="relative flex h-10 w-10 items-center justify-center rounded-full bg-white text-navy font-bold text-sm">
            JAS
            <span class="absolute bottom-0 right-0 h-3 w-3 rounded-full bg-[#25D366] border-2 border-[#075E54]"></span>
          </div>
          <div>
            <div class="text-sm font-heading font-semibold leading-tight">Customer Service JAS</div>
            <div class="text-[11px] text-white/80">Online &bull; Respon Cepat</div>
          </div>
        </div>

        <button
          type="button"
          class="text-white/70 hover:text-white p-1"
          on:click={closeGreeting}
          aria-label="Tutup sapaan"
        >
          <X size={18} />
        </button>
      </div>

      <!-- Chat Bubble Body -->
      <div class="p-4 space-y-3">
        <div class="max-w-[85%] rounded-2xl rounded-tl-sm bg-[#DCF8C6] p-3 text-xs leading-relaxed text-ink-900 shadow-sm">
          <p class="font-medium mb-1">Halo! Butuh penanganan hama hari ini?</p>
          <p class="text-ink-700">Tim kami siap memberikan estimasi biaya dan menjadwalkan survei ke lokasi Anda di Surabaya dan sekitarnya.</p>
          <div class="mt-1 text-right text-[10px] text-ink-500">Sekarang</div>
        </div>

        <a
          href={waUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center justify-center gap-2 w-full rounded-full bg-[#25D366] hover:bg-[#00D34D] py-2.5 px-4 text-xs font-semibold text-white shadow-sm transition-colors"
          on:click={closeGreeting}
        >
          <Send size={14} />
          <span>Mulai Chat WhatsApp</span>
        </a>
      </div>
    </div>
  {/if}

  <!-- Floating Green WhatsApp Button -->
  <a
    href={waUrl}
    target="_blank"
    rel="noopener noreferrer"
    class="pointer-events-auto flex h-14 w-14 sm:h-15 sm:w-15 items-center justify-center rounded-full bg-[#25D366] hover:bg-[#00D34D] text-white shadow-float transition-transform duration-200 hover:scale-105 active:scale-95"
    aria-label="Chat via WhatsApp 0813 9948 4916"
  >
    <svg class="h-8 w-8 fill-current" viewBox="0 0 24 24">
      <path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-.999 3.648 3.742-.981zm11.387-5.464c-.074-.124-.272-.198-.57-.347-.297-.149-1.758-.868-2.031-.967-.272-.099-.47-.149-.669.149-.198.297-.768.967-.941 1.165-.173.198-.347.223-.644.074-.297-.149-1.255-.462-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.297-.347.446-.521.151-.172.2-.296.3-.495.099-.198.05-.372-.025-.521-.075-.148-.669-1.611-.916-2.206-.242-.579-.487-.501-.669-.51l-.57-.01c-.198 0-.52.074-.792.372s-1.04 1.016-1.04 2.479 1.065 2.876 1.213 3.074c.149.198 2.095 3.2 5.076 4.487.709.306 1.263.489 1.694.626.712.226 1.36.194 1.872.118.571-.085 1.758-.719 2.006-1.413.248-.695.248-1.29.173-1.414z"/>
    </svg>
  </a>
</div>
