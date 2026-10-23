<script lang="ts">
  import { page } from '$app/stores';
  import { auth, isAdmin } from '$lib/stores/auth';
  import { 
    LayoutDashboard, 
    Newspaper, 
    Briefcase, 
    Image, 
    ShieldCheck, 
    FileText, 
    Images, 
    Inbox, 
    Settings, 
    Users, 
    LogOut,
    ExternalLink
  } from 'lucide-svelte';

  export let mobileOpen = false;
  export let closeMobile: () => void = () => {};

  $: currentPath = $page.url.pathname;
</script>

<aside class="fixed inset-y-0 left-0 z-40 flex w-64 flex-col bg-navy-deep text-white transition-transform duration-200 lg:translate-x-0 {mobileOpen ? 'translate-x-0' : '-translate-x-full'} shadow-xl lg:shadow-none">
  <!-- Brand Top -->
  <div class="flex h-16 items-center justify-between px-6 border-b border-white/10">
    <a href="/admin" class="flex items-center gap-2.5">
      <div class="flex h-8 w-8 items-center justify-center rounded-md bg-white text-navy font-extrabold text-xs">
        JAS
      </div>
      <div class="flex flex-col">
        <span class="font-heading text-sm font-bold tracking-tight text-white leading-none">
          JAYA ABADI SOLUSI
        </span>
        <span class="text-[10px] uppercase font-bold tracking-widest text-slate-400 mt-0.5">
          Backoffice CMS
        </span>
      </div>
    </a>
  </div>

  <!-- Navigation Groups -->
  <div class="flex-1 overflow-y-auto px-4 py-5 space-y-6">
    <!-- Group: Utama -->
    <div>
      <div class="px-3 mb-2 text-[10px] font-bold uppercase tracking-wider text-slate-400">
        Utama
      </div>
      <nav class="space-y-1">
        <a
          href="/admin"
          class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath === '/admin' ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
          on:click={closeMobile}
        >
          <LayoutDashboard size={17} />
          <span>Dashboard</span>
        </a>
      </nav>
    </div>

    <!-- Group: Konten -->
    <div>
      <div class="px-3 mb-2 text-[10px] font-bold uppercase tracking-wider text-slate-400">
        Konten Publik
      </div>
      <nav class="space-y-1">
        <a
          href="/admin/blog"
          class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/blog') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
          on:click={closeMobile}
        >
          <Newspaper size={17} />
          <span>Blog Artikel</span>
        </a>

        <a
          href="/admin/project"
          class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/project') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
          on:click={closeMobile}
        >
          <Briefcase size={17} />
          <span>Portofolio Proyek</span>
        </a>

        <a
          href="/admin/media"
          class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/media') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
          on:click={closeMobile}
        >
          <Image size={17} />
          <span>Media Library</span>
        </a>

        <!-- Admin Only Content Menus -->
        {#if $isAdmin}
          <a
            href="/admin/layanan"
            class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/layanan') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
            on:click={closeMobile}
          >
            <ShieldCheck size={17} />
            <span>Layanan & Landing</span>
          </a>

          <a
            href="/admin/pages"
            class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/pages') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
            on:click={closeMobile}
          >
            <FileText size={17} />
            <span>Halaman Statis</span>
          </a>

          <a
            href="/admin/galeri"
            class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/galeri') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
            on:click={closeMobile}
          >
            <Images size={17} />
            <span>Dokumentasi Galeri</span>
          </a>
        {/if}
      </nav>
    </div>

    <!-- Group: Operasional (Admin Only) -->
    {#if $isAdmin}
      <div>
        <div class="px-3 mb-2 text-[10px] font-bold uppercase tracking-wider text-slate-400">
          Operasional
        </div>
        <nav class="space-y-1">
          <a
            href="/admin/leads"
            class="flex items-center justify-between rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/leads') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
            on:click={closeMobile}
          >
            <div class="flex items-center gap-3">
              <Inbox size={17} />
              <span>Inbox Lead</span>
            </div>
            <span class="inline-flex items-center px-1.5 py-0.5 rounded text-[10px] font-bold bg-brand-red text-white leading-none">
              Baru
            </span>
          </a>
        </nav>
      </div>

      <!-- Group: Sistem (Admin Only) -->
      <div>
        <div class="px-3 mb-2 text-[10px] font-bold uppercase tracking-wider text-slate-400">
          Sistem & Konfigurasi
        </div>
        <nav class="space-y-1">
          <a
            href="/admin/pengaturan"
            class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/pengaturan') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
            on:click={closeMobile}
          >
            <Settings size={17} />
            <span>Pengaturan Situs</span>
          </a>

          <a
            href="/admin/pengguna"
            class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors {currentPath.startsWith('/admin/pengguna') ? 'bg-white/10 text-white font-semibold' : 'text-slate-300 hover:bg-white/5 hover:text-white'}"
            on:click={closeMobile}
          >
            <Users size={17} />
            <span>Manajemen Pengguna</span>
          </a>
        </nav>
      </div>
    {/if}
  </div>

  <!-- Bottom User Profile & Links -->
  <div class="p-4 border-t border-white/10 space-y-1.5">
    <a
      href="/"
      target="_blank"
      class="flex items-center gap-2 rounded-lg px-3 py-2 text-xs font-medium text-slate-300 hover:bg-white/5 hover:text-white transition-colors"
    >
      <ExternalLink size={15} />
      <span>Buka Website Publik</span>
    </a>

    <button
      type="button"
      class="flex items-center gap-2 w-full rounded-lg px-3 py-2 text-xs font-medium text-red-300 hover:bg-red-500/10 hover:text-red-200 transition-colors"
      on:click={() => auth.logout()}
    >
      <LogOut size={15} />
      <span>Keluar Akun</span>
    </button>
  </div>
</aside>
