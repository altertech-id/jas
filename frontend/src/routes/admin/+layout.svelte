<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { auth } from '$lib/stores/auth';
  import Sidebar from '$lib/components/admin/Sidebar.svelte';
  import Topbar from '$lib/components/admin/Topbar.svelte';
  import Toast from '$lib/components/ui/Toast.svelte';
  import '../../app.css';

  let mobileSidebarOpen = false;

  $: isLoginPage = $page.url.pathname === '/admin/login';

  onMount(() => {
    auth.init();
    if (!$auth.isAuthenticated && !isLoginPage) {
      goto('/admin/login');
    }
  });

  $: if (!$auth.loading && !$auth.isAuthenticated && !isLoginPage) {
    goto('/admin/login');
  }

  // Determine page title based on path
  $: pageTitle = getTitle($page.url.pathname);

  function getTitle(path: string): string {
    if (path === '/admin') return 'Dashboard Ringkasan';
    if (path.startsWith('/admin/leads')) return 'Inbox Lead Pelanggan';
    if (path.startsWith('/admin/blog')) return 'Manajemen Artikel Blog';
    if (path.startsWith('/admin/project')) return 'Manajemen Portofolio Proyek';
    if (path.startsWith('/admin/layanan')) return 'Manajemen Layanan & Landing Page';
    if (path.startsWith('/admin/pages')) return 'Manajemen Konten Halaman';
    if (path.startsWith('/admin/galeri')) return 'Manajemen Galeri Dokumentasi';
    if (path.startsWith('/admin/media')) return 'Media Library';
    if (path.startsWith('/admin/pengaturan')) return 'Pengaturan Situs & Kontak';
    if (path.startsWith('/admin/pengguna')) return 'Manajemen Pengguna';
    return 'Admin Panel';
  }
</script>

{#if isLoginPage}
  <slot />
  <Toast />
{:else}
  <div class="min-h-screen bg-surface-light">
    <!-- Sidebar -->
    <Sidebar
      mobileOpen={mobileSidebarOpen}
      closeMobile={() => mobileSidebarOpen = false}
    />

    <!-- Main Content Area -->
    <div class="lg:pl-64 flex flex-col min-h-screen">
      <Topbar
        title={pageTitle}
        toggleSidebar={() => mobileSidebarOpen = !mobileSidebarOpen}
      />

      <main class="flex-1 p-6 sm:p-8 max-w-[1280px] w-full mx-auto">
        <slot />
      </main>
    </div>

    <Toast />
  </div>
{/if}
