<script lang="ts">
  import { auth } from '$lib/stores/auth';
  import { Menu, User } from 'lucide-svelte';

  export let title = 'Admin Panel';
  export let toggleSidebar: () => void = () => {};
</script>

<header class="sticky top-0 z-30 flex h-16 w-full items-center justify-between border-b border-slate-200/80 bg-white px-6 sm:px-8">
  <div class="flex items-center gap-4">
    <button
      type="button"
      class="text-slate-600 hover:text-slate-900 lg:hidden p-1.5 rounded-md border border-slate-200"
      on:click={toggleSidebar}
      aria-label="Buka navigasi"
    >
      <Menu size={18} />
    </button>
    <h1 class="text-base sm:text-lg font-heading font-semibold text-slate-900">
      {title}
    </h1>
  </div>

  <div class="flex items-center gap-3">
    {#if $auth.profile}
      <div class="flex items-center gap-3">
        <div class="hidden sm:flex flex-col text-right">
          <span class="text-xs font-semibold text-slate-900 leading-tight">
            {$auth.profile.name}
          </span>
          <span class="text-[11px] font-medium text-slate-500 capitalize">
            {$auth.profile.role}
          </span>
        </div>
        <div class="flex h-8 w-8 items-center justify-center rounded-full bg-slate-900 text-white font-bold text-xs">
          {$auth.profile.name.charAt(0)}
        </div>
      </div>
    {/if}
  </div>
</header>
