<script lang="ts">
  import { toast } from '$lib/stores/toast';
  import { CheckCircle2, AlertCircle, AlertTriangle, Info, X } from 'lucide-svelte';

  const icons = {
    success: CheckCircle2,
    error: AlertCircle,
    warning: AlertTriangle,
    info: Info
  };

  const borders = {
    success: 'border-l-4 border-l-green-600',
    error: 'border-l-4 border-l-brand-red',
    warning: 'border-l-4 border-l-amber',
    info: 'border-l-4 border-l-brand-blue'
  };
</script>

<div class="fixed bottom-6 right-6 z-50 flex flex-col gap-2 max-w-sm w-full pointer-events-none">
  {#each $toast as item (item.id)}
    <div
      class="pointer-events-auto flex items-center justify-between gap-3 rounded-xl bg-white p-4 shadow-float border border-line-soft {borders[item.type]} transition-all duration-200"
      role="alert"
    >
      <div class="flex items-center gap-3">
        <svelte:component
          this={icons[item.type]}
          size={20}
          strokeWidth={2}
          class={item.type === 'success' ? 'text-green-600' : item.type === 'error' ? 'text-brand-red' : item.type === 'warning' ? 'text-amber' : 'text-brand-blue'}
        />
        <p class="text-sm font-medium text-ink-900 leading-snug">
          {item.message}
        </p>
      </div>

      <button
        type="button"
        class="text-ink-300 hover:text-ink-700 p-1"
        on:click={() => toast.remove(item.id)}
        aria-label="Tutup"
      >
        <X size={16} strokeWidth={2} />
      </button>
    </div>
  {/each}
</div>
