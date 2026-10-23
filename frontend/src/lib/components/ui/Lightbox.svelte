<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte';
  import { X, ChevronLeft, ChevronRight } from 'lucide-svelte';

  export let show = false;
  export let imageUrl = '';
  export let imageAlt = '';
  export let caption = '';
  export let hasPrev = false;
  export let hasNext = false;

  const dispatch = createEventDispatcher<{
    close: void;
    prev: void;
    next: void;
  }>();

  function handleKeydown(e: KeyboardEvent) {
    if (!show) return;
    if (e.key === 'Escape') dispatch('close');
    if (e.key === 'ArrowLeft' && hasPrev) dispatch('prev');
    if (e.key === 'ArrowRight' && hasNext) dispatch('next');
  }

  onMount(() => {
    window.addEventListener('keydown', handleKeydown);
    return () => window.removeEventListener('keydown', handleKeydown);
  });
</script>

{#if show}
  <div 
    class="fixed inset-0 z-50 flex items-center justify-center bg-navy-deep/95 p-4 sm:p-6 backdrop-blur-sm transition-opacity duration-200"
    role="dialog"
    aria-modal="true"
  >
    <!-- Close button -->
    <button
      type="button"
      class="absolute top-4 right-4 z-10 flex h-11 w-11 items-center justify-center rounded-full bg-white/10 text-white hover:bg-white/20 transition-colors"
      on:click={() => dispatch('close')}
      aria-label="Tutup pratinjau"
    >
      <X size={24} strokeWidth={2} />
    </button>

    <!-- Prev button -->
    {#if hasPrev}
      <button
        type="button"
        class="absolute left-4 top-1/2 -translate-y-1/2 z-10 flex h-11 w-11 items-center justify-center rounded-full bg-white/10 text-white hover:bg-white/20 transition-colors"
        on:click={() => dispatch('prev')}
        aria-label="Foto sebelumnya"
      >
        <ChevronLeft size={24} strokeWidth={2} />
      </button>
    {/if}

    <!-- Image container -->
    <div class="relative max-h-[85vh] max-w-5xl flex flex-col items-center">
      <img
        src={imageUrl}
        alt={imageAlt || caption || 'Dokumentasi Jaya Abadi Solusi'}
        class="max-h-[75vh] w-auto max-w-full rounded-xl object-contain shadow-2xl"
      />
      {#if caption}
        <p class="mt-3 text-center text-sm font-medium text-white/90 max-w-xl">
          {caption}
        </p>
      {/if}
    </div>

    <!-- Next button -->
    {#if hasNext}
      <button
        type="button"
        class="absolute right-4 top-1/2 -translate-y-1/2 z-10 flex h-11 w-11 items-center justify-center rounded-full bg-white/10 text-white hover:bg-white/20 transition-colors"
        on:click={() => dispatch('next')}
        aria-label="Foto berikutnya"
      >
        <ChevronRight size={24} strokeWidth={2} />
      </button>
    {/if}
  </div>
{/if}
