<script lang="ts">
  import { MessageCircle, FileText } from 'lucide-svelte';

  export let serviceName = '';
  const phone = '6281399484916';
  $: waMessage = `Halo Jaya Abadi Solusi, saya ingin penawaran untuk layanan ${serviceName || 'pengendalian hama'}.`;
  $: waUrl = `https://wa.me/${phone}?text=${encodeURIComponent(waMessage)}`;

  function scrollToForm() {
    const el = document.getElementById('form-quotation-section');
    if (el) {
      el.scrollIntoView({ behavior: 'smooth' });
    }
  }
</script>

<div class="fixed bottom-0 inset-x-0 z-30 bg-white/95 backdrop-blur-md border-t border-line-soft p-3 sm:hidden shadow-lg">
  <div class="flex items-center gap-2">
    <a
      href={waUrl}
      target="_blank"
      rel="noopener noreferrer"
      class="flex-1 flex items-center justify-center gap-1.5 rounded-full bg-wa py-2.5 px-3 text-xs font-semibold text-white shadow-sm active:scale-95 transition-transform"
    >
      <MessageCircle size={16} />
      <span>Chat WA</span>
    </a>

    <button
      type="button"
      on:click={scrollToForm}
      class="flex-1 flex items-center justify-center gap-1.5 rounded-full bg-cta py-2.5 px-3 text-xs font-semibold text-white shadow-sm active:scale-95 transition-transform"
    >
      <FileText size={16} />
      <span>Minta Survei</span>
    </button>
  </div>
</div>
