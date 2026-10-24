<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { DashboardStats, Lead } from '$lib/types';
  import StatusPill from '$lib/components/ui/StatusPill.svelte';
  import { ArrowRight, MessageCircle } from 'lucide-svelte';

  let stats: DashboardStats = {
    new_leads_month: 0,
    total_leads_month: 0,
    published_posts: 0,
    published_projects: 0
  };

  let recentLeads: Lead[] = [];
  let loading = true;

  onMount(async () => {
    try {
      const [statsRes, leadsRes] = await Promise.all([
        api.getDashboardStats().catch(() => ({
          new_leads_month: 3,
          total_leads_month: 12,
          published_posts: 3,
          published_projects: 3
        })),
        api.getLeads({ limit: 5 }).catch(() => ({ data: [] }))
      ]);

      stats = statsRes;
      recentLeads = leadsRes.data ?? [];
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  <title>Dashboard - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-8">
  <!-- 4 Clean Metric Cards -->
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
    <!-- Stat 1: Lead Baru -->
    <div class="rounded-xl border border-slate-200/80 bg-white p-5 shadow-xs">
      <div class="text-xs font-medium text-slate-500">Lead Baru (Bulan Ini)</div>
      <div class="mt-2 text-3xl font-heading font-bold text-slate-900 tabular-nums">
        {stats.new_leads_month}
      </div>
      <div class="mt-1 text-xs text-brand-blue font-medium">Perlu ditindaklanjuti</div>
    </div>

    <!-- Stat 2: Total Lead -->
    <div class="rounded-xl border border-slate-200/80 bg-white p-5 shadow-xs">
      <div class="text-xs font-medium text-slate-500">Total Lead Masuk</div>
      <div class="mt-2 text-3xl font-heading font-bold text-slate-900 tabular-nums">
        {stats.total_leads_month}
      </div>
      <div class="mt-1 text-xs text-slate-500">Permintaan penawaran</div>
    </div>

    <!-- Stat 3: Artikel Terbit -->
    <div class="rounded-xl border border-slate-200/80 bg-white p-5 shadow-xs">
      <div class="text-xs font-medium text-slate-500">Artikel Terbit</div>
      <div class="mt-2 text-3xl font-heading font-bold text-slate-900 tabular-nums">
        {stats.published_posts}
      </div>
      <div class="mt-1 text-xs text-slate-500">Tayang di blog publik</div>
    </div>

    <!-- Stat 4: Project Terbit -->
    <div class="rounded-xl border border-slate-200/80 bg-white p-5 shadow-xs">
      <div class="text-xs font-medium text-slate-500">Portofolio Proyek</div>
      <div class="mt-2 text-3xl font-heading font-bold text-slate-900 tabular-nums">
        {stats.published_projects}
      </div>
      <div class="mt-1 text-xs text-slate-500">Dokumentasi lapangan</div>
    </div>
  </div>

  <!-- Content Split: 2/3 Recent Leads & 1/3 Quick Actions -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
    <!-- Recent Leads Table (2/3) -->
    <div class="lg:col-span-8 rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
      <div class="flex items-center justify-between border-b border-slate-200/80 px-6 py-4">
        <div>
          <h2 class="text-base font-heading font-semibold text-slate-900">
            Permintaan Survei & Penawaran Masuk
          </h2>
          <p class="text-xs text-slate-500 mt-0.5">Daftar calon pelanggan terbaru dari formulir website.</p>
        </div>
        <a href="/admin/leads" class="text-xs font-semibold text-brand-blue hover:underline">
          Lihat Semua &rarr;
        </a>
      </div>

      {#if loading}
        <div class="p-8 text-center text-sm text-slate-500">Memuat data lead...</div>
      {:else if (recentLeads?.length ?? 0) === 0}
        <div class="p-8 text-center text-sm text-slate-500">
          Belum ada lead baru yang masuk.
        </div>
      {:else}
        <div class="overflow-x-auto">
          <table class="w-full text-left text-sm">
            <thead class="bg-slate-50/80 text-xs font-medium text-slate-500 border-b border-slate-200/80">
              <tr>
                <th class="py-3 px-6">Nama Pengirim</th>
                <th class="py-3 px-6">Kontak</th>
                <th class="py-3 px-6">Status</th>
                <th class="py-3 px-6 text-right">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100">
              {#each recentLeads as lead}
                <tr class="hover:bg-slate-50/60 transition-colors">
                  <td class="py-3.5 px-6 font-medium text-slate-900">
                    <div>{lead.name}</div>
                    <div class="text-xs text-slate-500 font-normal">{lead.source_page || 'Website'}</div>
                  </td>
                  <td class="py-3.5 px-6">
                    <a href="tel:{lead.phone}" class="font-medium text-slate-900 hover:text-brand-blue">
                      {lead.phone}
                    </a>
                  </td>
                  <td class="py-3.5 px-6">
                    <StatusPill status={lead.status} />
                  </td>
                  <td class="py-3.5 px-6 text-right">
                    <a
                      href="https://wa.me/{lead.phone.replace(/[^0-9]/g, '')}?text=Halo%20{encodeURIComponent(lead.name)}%2C%20saya%20dari%20PT%20Jaya%20Abadi%20Solusi..."
                      target="_blank"
                      rel="noopener noreferrer"
                      class="inline-flex items-center gap-1.5 rounded-md bg-emerald-50 hover:bg-emerald-100 text-emerald-700 px-2.5 py-1 text-xs font-medium transition-colors"
                    >
                      <MessageCircle size={13} />
                      <span>Balas WA</span>
                    </a>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      {/if}
    </div>

    <!-- Quick Actions (1/3) -->
    <div class="lg:col-span-4 rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-4">
      <h2 class="text-base font-heading font-semibold text-slate-900 pb-3 border-b border-slate-200/80">
        Aksi Cepat
      </h2>

      <div class="space-y-2.5">
        <a
          href="/admin/blog"
          class="flex items-center justify-between rounded-lg border border-slate-200/80 p-3 hover:bg-slate-50 hover:border-slate-300 transition-colors group"
        >
          <div>
            <div class="text-xs font-semibold text-slate-900 group-hover:text-navy">Tulis Artikel Baru</div>
            <div class="text-[11px] text-slate-500">Publikasi artikel tips & edukasi</div>
          </div>
          <ArrowRight size={15} class="text-slate-400 group-hover:text-slate-900 transition-colors" />
        </a>

        <a
          href="/admin/project"
          class="flex items-center justify-between rounded-lg border border-slate-200/80 p-3 hover:bg-slate-50 hover:border-slate-300 transition-colors group"
        >
          <div>
            <div class="text-xs font-semibold text-slate-900 group-hover:text-navy">Tambah Proyek Baru</div>
            <div class="text-[11px] text-slate-500">Unggah portofolio & lokasi pengerjaan</div>
          </div>
          <ArrowRight size={15} class="text-slate-400 group-hover:text-slate-900 transition-colors" />
        </a>

        <a
          href="/admin/galeri"
          class="flex items-center justify-between rounded-lg border border-slate-200/80 p-3 hover:bg-slate-50 hover:border-slate-300 transition-colors group"
        >
          <div>
            <div class="text-xs font-semibold text-slate-900 group-hover:text-navy">Unggah Foto Galeri</div>
            <div class="text-[11px] text-slate-500">Dokumentasi hasil pengerjaan lapangan</div>
          </div>
          <ArrowRight size={15} class="text-slate-400 group-hover:text-slate-900 transition-colors" />
        </a>
      </div>
    </div>
  </div>
</div>