<script lang="ts">
  import { ArrowRight } from 'lucide-svelte';

  export let slug: string;
  export let name: string;
  export let description: string;
  export let imageUrl: string = '';

  const defaultImages: Record<string, string> = {
    'termite-control': 'https://images.unsplash.com/photo-1548767797-d8c844163c4c?auto=format&fit=crop&w=600&q=80',
    'rodent-control': 'https://images.unsplash.com/photo-1579202673506-ca3ce28943ef?auto=format&fit=crop&w=600&q=80',
    'pest-control': 'https://images.unsplash.com/photo-1563245372-f21724e3856d?auto=format&fit=crop&w=600&q=80',
    'fumigasi-pe': 'https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?auto=format&fit=crop&w=600&q=80',
    'disinfeksi': 'https://images.unsplash.com/photo-1584744982491-665216d95f8b?auto=format&fit=crop&w=600&q=80'
  };

  $: displayImage = imageUrl || defaultImages[slug] || 'https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=600&q=80';
</script>

<a
  href="/layanan/{slug}"
  class="group block bg-white rounded-xl overflow-hidden shadow-sm hover:shadow-md transition-all duration-200"
>
  <div class="aspect-[16/10] w-full overflow-hidden bg-slate-100">
    <img
      src={displayImage}
      alt={name}
      class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
      loading="lazy"
    />
  </div>

  <div class="p-6">
    <h3 class="text-lg font-heading font-bold text-slate-900 group-hover:text-brand-blue transition-colors mb-2">
      {name}
    </h3>

    <p class="text-sm leading-relaxed text-slate-600 line-clamp-3 mb-4 font-body">
      {description}
    </p>

    <div class="inline-flex items-center gap-1.5 text-xs font-semibold text-cta group-hover:translate-x-1 transition-transform">
      <span>Lihat Layanan</span>
      <ArrowRight size={14} />
    </div>
  </div>
</a>
