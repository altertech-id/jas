<script lang="ts">
  export let variant: 'hero' | 'light' | 'outline' = 'light';
</script>

{#if variant === 'hero'}
  <div class="inline-flex items-center gap-2 rounded-full bg-white/10 backdrop-blur-sm border border-white/20 px-4 py-1.5 text-xs sm:text-sm font-medium text-white shadow-sm {$$props.class || ''}">
    <slot name="icon" />
    <slot />
  </div>
{:else if variant === 'outline'}
  <div class="inline-flex items-center gap-1.5 rounded-full border border-slate-200 bg-white px-3 py-1 text-xs font-semibold uppercase tracking-wider text-slate-800 {$$props.class || ''}">
    <slot />
  </div>
{:else}
  <div class="inline-flex items-center gap-1.5 rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-800 {$$props.class || ''}">
    <slot />
  </div>
{/if}
