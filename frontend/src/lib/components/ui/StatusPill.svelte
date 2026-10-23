<script lang="ts">
  export let status: string;

  const statusConfigs: Record<string, { label: string; bg: string; text: string; dot: string }> = {
    baru: { label: 'Baru', bg: 'bg-blue-50', text: 'text-blue-700', dot: 'bg-blue-600' },
    dihubungi: { label: 'Dihubungi', bg: 'bg-amber-50', text: 'text-amber-800', dot: 'bg-amber-500' },
    selesai: { label: 'Selesai', bg: 'bg-green-50', text: 'text-green-800', dot: 'bg-green-600' },
    draft: { label: 'Draft', bg: 'bg-slate-100', text: 'text-slate-700', dot: 'bg-slate-400' },
    published: { label: 'Published', bg: 'bg-green-50', text: 'text-green-800', dot: 'bg-green-600' },
    aktif: { label: 'Aktif', bg: 'bg-green-50', text: 'text-green-800', dot: 'bg-green-600' },
    nonaktif: { label: 'Nonaktif', bg: 'bg-slate-100', text: 'text-slate-600', dot: 'bg-slate-400' }
  };

  $: config = statusConfigs[status.toLowerCase()] || {
    label: status,
    bg: 'bg-slate-100',
    text: 'text-slate-700',
    dot: 'bg-slate-400'
  };
</script>

<span class="inline-flex items-center gap-1.5 rounded-full px-2.5 py-0.5 text-xs font-semibold {config.bg} {config.text}">
  <span class="h-1.5 w-1.5 rounded-full {config.dot}"></span>
  {config.label}
</span>
