<script lang="ts">
  import { Plus } from 'lucide-svelte';

  export let title: string;
  export let open = false;

  function toggle() {
    open = !open;
  }
</script>

<div class="rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
  <button
    type="button"
    class="flex w-full items-center justify-between px-6 py-4 text-left transition-colors hover:bg-slate-50/50"
    on:click={toggle}
    aria-expanded={open}
  >
    <span class="text-base font-heading font-semibold text-slate-900 pr-4">
      {title}
    </span>
    <span class="flex-shrink-0 text-slate-400 transition-transform duration-200 {open ? 'rotate-45 text-slate-900' : ''}">
      <Plus size={18} />
    </span>
  </button>

  {#if open}
    <div class="border-t border-slate-100 px-6 py-4 text-sm leading-relaxed text-slate-600 font-body">
      <slot />
    </div>
  {/if}
</div>
