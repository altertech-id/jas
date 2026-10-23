<script lang="ts">
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  import Button from "$lib/components/ui/Button.svelte";
  import { ButtonVariant } from "$lib/types/enums";
  import { Menu, X, ChevronDown, Phone } from "lucide-svelte";

  let scrolled = false;
  let mobileMenuOpen = false;
  let servicesDropdownOpen = false;

  const services = [
    { name: "Pest Control Umum", slug: "pest-control" },
    { name: "Rodent Control (Tikus)", slug: "rodent-control" },
    { name: "Termite Control (Rayap)", slug: "termite-control" },
    { name: "Fumigasi PE & Komoditas", slug: "fumigasi-pe" },
    { name: "Disinfeksi & Sterilisasi", slug: "disinfeksi" },
  ];

  function handleScroll() {
    scrolled = window.scrollY > 15;
  }

  onMount(() => {
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  });

  function closeMobileMenu() {
    mobileMenuOpen = false;
    servicesDropdownOpen = false;
  }
</script>

<header
  class="sticky top-0 z-40 w-full bg-white transition-all duration-200 {scrolled
    ? 'shadow-xs border-b border-line-soft'
    : ''}"
>
  <div
    class="mx-auto flex h-16 sm:h-[72px] max-w-container items-center justify-between px-5 sm:px-8"
  >
    <!-- Brand Logo (Clean & Professional) -->
    <a href="/" class="flex items-center gap-2" on:click={closeMobileMenu}>
      <div class="flex flex-col">
        <span
          class="font-heading text-xl font-extrabold tracking-tight text-navy leading-none"
        >
          JAYA ABADI
        </span>
        <span
          class="font-heading text-[11px] font-bold tracking-widest text-cta leading-tight"
        >
          SOLUSI
        </span>
      </div>
    </a>

    <!-- Desktop Navigation Links -->
    <nav class="hidden md:flex items-center gap-8">
      <a
        href="/"
        class="text-xs sm:text-sm font-heading font-semibold uppercase tracking-wider transition-colors hover:text-cta {$page
          .url.pathname === '/'
          ? 'text-cta font-bold'
          : 'text-ink-700'}"
      >
        Beranda
      </a>

      <!-- Services Dropdown -->
      <div
        class="relative group"
        on:mouseleave={() => (servicesDropdownOpen = false)}
      >
        <button
          type="button"
          class="flex items-center gap-1 text-xs sm:text-sm font-heading font-semibold uppercase tracking-wider text-ink-700 transition-colors hover:text-cta py-2 {$page.url.pathname.startsWith(
            '/layanan',
          )
            ? 'text-cta'
            : ''}"
          on:mouseenter={() => (servicesDropdownOpen = true)}
          on:click={() => (servicesDropdownOpen = !servicesDropdownOpen)}
        >
          <span>Layanan</span>
          <ChevronDown
            size={14}
            class="transition-transform duration-200 {servicesDropdownOpen
              ? 'rotate-180'
              : ''}"
          />
        </button>

        {#if servicesDropdownOpen}
          <div
            class="absolute left-0 top-full w-60 rounded-xl bg-white p-2 shadow-float border border-slate-100 transition-all duration-150 animate-in fade-in"
          >
            {#each services as item}
              <a
                href="/layanan/{item.slug}"
                class="block rounded-lg px-3 py-2 text-xs sm:text-sm font-medium text-ink-700 transition-colors hover:bg-surface-mist hover:text-navy"
                on:click={() => (servicesDropdownOpen = false)}
              >
                {item.name}
              </a>
            {/each}
            <div class="mt-1 border-t border-slate-100 pt-1">
              <a
                href="/layanan"
                class="block rounded-lg px-3 py-2 text-xs font-semibold text-cta hover:bg-surface-mist"
                on:click={() => (servicesDropdownOpen = false)}
              >
                Lihat Semua Layanan &rarr;
              </a>
            </div>
          </div>
        {/if}
      </div>

      <a
        href="/project"
        class="text-xs sm:text-sm font-heading font-semibold uppercase tracking-wider transition-colors hover:text-cta {$page.url.pathname.startsWith(
          '/project',
        )
          ? 'text-cta font-bold'
          : 'text-ink-700'}"
      >
        Project
      </a>

      <a
        href="/blog"
        class="text-xs sm:text-sm font-heading font-semibold uppercase tracking-wider transition-colors hover:text-cta {$page.url.pathname.startsWith(
          '/blog',
        )
          ? 'text-cta font-bold'
          : 'text-ink-700'}"
      >
        Blog
      </a>

      <a
        href="/galeri"
        class="text-xs sm:text-sm font-heading font-semibold uppercase tracking-wider transition-colors hover:text-cta {$page.url.pathname.startsWith(
          '/galeri',
        )
          ? 'text-cta font-bold'
          : 'text-ink-700'}"
      >
        Galeri
      </a>

      <a
        href="/tentang-kami"
        class="text-xs sm:text-sm font-heading font-semibold uppercase tracking-wider transition-colors hover:text-cta {$page
          .url.pathname === '/tentang-kami'
          ? 'text-cta font-bold'
          : 'text-ink-700'}"
      >
        Tentang Kami
      </a>

      <a
        href="/kontak"
        class="text-xs sm:text-sm font-heading font-semibold uppercase tracking-wider transition-colors hover:text-cta {$page
          .url.pathname === '/kontak'
          ? 'text-cta font-bold'
          : 'text-ink-700'}"
      >
        Kontak
      </a>
    </nav>

    <!-- Desktop Action Button -->
    <div class="hidden md:flex items-center gap-4">
      <Button href="/kontak" variant={ButtonVariant.PRIMARY} size="sm">
        Hubungi Kami
      </Button>
    </div>

    <!-- Mobile Hamburger Toggle -->
    <button
      type="button"
      class="flex h-10 w-10 items-center justify-center rounded-lg text-ink-700 md:hidden"
      on:click={() => (mobileMenuOpen = !mobileMenuOpen)}
      aria-label="Toggle menu"
    >
      {#if mobileMenuOpen}
        <X size={24} />
      {:else}
        <Menu size={24} />
      {/if}
    </button>
  </div>

  <!-- Mobile Drawer -->
  {#if mobileMenuOpen}
    <div
      class="fixed inset-x-0 top-16 bottom-0 z-50 bg-white p-6 shadow-xl overflow-y-auto md:hidden flex flex-col justify-between"
    >
      <div class="flex flex-col space-y-4">
        <a
          href="/"
          class="text-base font-heading font-semibold text-ink-900 py-2 border-b border-slate-100"
          on:click={closeMobileMenu}
        >
          Beranda
        </a>

        <div class="py-2 border-b border-slate-100">
          <div
            class="text-xs font-semibold uppercase tracking-wider text-ink-500 mb-2"
          >
            Layanan Kami
          </div>
          <div class="flex flex-col space-y-2 pl-2">
            {#each services as item}
              <a
                href="/layanan/{item.slug}"
                class="text-sm text-ink-700 py-1"
                on:click={closeMobileMenu}
              >
                {item.name}
              </a>
            {/each}
            <a
              href="/layanan"
              class="text-xs font-semibold text-cta py-1"
              on:click={closeMobileMenu}
            >
              Lihat Katalog Semua Layanan &rarr;
            </a>
          </div>
        </div>

        <a
          href="/project"
          class="text-base font-heading font-semibold text-ink-900 py-2 border-b border-slate-100"
          on:click={closeMobileMenu}
        >
          Project Lapangan
        </a>
        <a
          href="/blog"
          class="text-base font-heading font-semibold text-ink-900 py-2 border-b border-slate-100"
          on:click={closeMobileMenu}
        >
          Artikel & Tips
        </a>
        <a
          href="/galeri"
          class="text-base font-heading font-semibold text-ink-900 py-2 border-b border-slate-100"
          on:click={closeMobileMenu}
        >
          Galeri Dokumentasi
        </a>
        <a
          href="/tentang-kami"
          class="text-base font-heading font-semibold text-ink-900 py-2 border-b border-slate-100"
          on:click={closeMobileMenu}
        >
          Tentang Kami
        </a>
        <a
          href="/kontak"
          class="text-base font-heading font-semibold text-ink-900 py-2 border-b border-slate-100"
          on:click={closeMobileMenu}
        >
          Kontak
        </a>
      </div>

      <div class="pt-6 border-t border-slate-100 flex flex-col gap-3">
        <a
          href="https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20ingin%20konsultasi%20layanan%20hama"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center justify-center gap-2 rounded-full bg-wa py-3 text-sm font-semibold text-white shadow-sm"
        >
          Chat WhatsApp (0813 9948 4916)
        </a>

        <Button
          href="/kontak"
          variant={ButtonVariant.PRIMARY}
          fullWidth
          on:click={closeMobileMenu}
        >
          Minta Penawaran Gratis
        </Button>
      </div>
    </div>
  {/if}
</header>
