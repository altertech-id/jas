<script lang="ts">
  import Navbar from '$lib/components/public/Navbar.svelte';
  import Footer from '$lib/components/public/Footer.svelte';
  import FloatingWhatsApp from '$lib/components/public/FloatingWhatsApp.svelte';
  import Toast from '$lib/components/ui/Toast.svelte';
  import '../../app.css';
</script>

<div class="flex min-h-screen flex-col bg-white">
  <Navbar />
  <main class="flex-grow">
    <slot />
  </main>
  <Footer />
  <FloatingWhatsApp />
  <Toast />
</div>
