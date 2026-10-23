<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Lead } from '$lib/types';
  import { LeadStatus } from '$lib/types/enums';
  import StatusPill from '$lib/components/ui/StatusPill.svelte';
  import { toast } from '$lib/stores/toast';
  import { MessageCircle, X, Calendar, Phone, Mail } from 'lucide-svelte';

  let leads: Lead[] = [];
  let filterStatus = '';
  let loading = true;

  // Drawer state
  let selectedLead: Lead | null = null;
  let updatingStatus = false;

  onMount(async () => {
    await fetchLeads();
  });

  async function fetchLeads() {
    loading = true;
    try {
      const res = await api.getLeads({ status: filterStatus });
      leads = res.data;
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  }

  async function changeStatus(lead: Lead, newStatus: LeadStatus) {
    updatingStatus = true;
    try {
      await api.updateLeadStatus(lead.id, newStatus);
      lead.status = newStatus;
      if (selectedLead && selectedLead.id === lead.id) {
        selectedLead.status = newStatus;
      }
      toast.show('Status lead berhasil diperbarui.', 'success');
    } catch (err: any) {
      toast.show(err.message || 'Gagal mengubah status.', 'error');
    } finally {
      updatingStatus = false;
    }
  }

  function getWhatsAppUrl(phone: string, name: string) {
    const cleanPhone = phone.replace(/[^0-9]/g, '');
    const intlPhone = cleanPhone.startsWith('0') ? '62' + cleanPhone.slice(1) : cleanPhone;
    const msg = `Halo ${name}, terima kasih telah menghubungi PT Jaya Abadi Solusi. Kami ingin menindaklanjuti permintaan survei/penawaran Anda.`;
    return `https://wa.me/${intlPhone}?text=${encodeURIComponent(msg)}`;
  }
</script>

<svelte:head>
  <title>Inbox Lead - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  <!-- Header & Filter -->
  <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
    <div>
      <h2 class="text-xl font-heading font-semibold text-slate-900">
        Inbox Permintaan & Quotation Lead
      </h2>
      <p class="text-xs text-slate-500 mt-0.5">
        Daftar seluruh calon pelanggan yang mengirimkan formulir dari website.
      </p>
    </div>

    <!-- Status Filter -->
    <div class="inline-flex rounded-lg border border-slate-200 bg-slate-50/80 p-0.5">
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {filterStatus === '' ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => { filterStatus = ''; fetchLeads(); }}
      >
        Semua
      </button>
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {filterStatus === LeadStatus.BARU ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => { filterStatus = LeadStatus.BARU; fetchLeads(); }}
      >
        Baru
      </button>
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {filterStatus === LeadStatus.DIHUBUNGI ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => { filterStatus = LeadStatus.DIHUBUNGI; fetchLeads(); }}
      >
        Dihubungi
      </button>
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {filterStatus === LeadStatus.SELESAI ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => { filterStatus = LeadStatus.SELESAI; fetchLeads(); }}
      >
        Selesai
      </button>
    </div>
  </div>

  <!-- Table Card -->
  <div class="rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
    {#if loading}
      <div class="flex justify-center p-12">
        <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
      </div>
    {:else if leads.length === 0}
      <div class="p-12 text-center text-sm text-slate-500">
        Tidak ada data lead pada status ini.
      </div>
    {:else}
      <div class="overflow-x-auto">
        <table class="w-full text-left text-sm">
          <thead class="bg-slate-50/80 text-xs font-medium text-slate-500 border-b border-slate-200/80">
            <tr>
              <th class="py-3 px-6">Tanggal</th>
              <th class="py-3 px-6">Nama Pengirim</th>
              <th class="py-3 px-6">Kontak</th>
              <th class="py-3 px-6">Sumber Halaman</th>
              <th class="py-3 px-6">Status</th>
              <th class="py-3 px-6 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100">
            {#each leads as lead}
              <tr 
                class="hover:bg-slate-50/60 transition-colors cursor-pointer"
                on:click={() => selectedLead = lead}
              >
                <td class="py-3.5 px-6 text-xs text-slate-500">
                  {new Date(lead.created_at).toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' })}
                </td>
                <td class="py-3.5 px-6 font-medium text-slate-900">
                  {lead.name}
                </td>
                <td class="py-3.5 px-6 text-slate-900 font-medium">
                  {lead.phone}
                </td>
                <td class="py-3.5 px-6 text-xs text-slate-500">
                  {lead.source_page || 'Form Website'}
                </td>
                <td class="py-3.5 px-6">
                  <StatusPill status={lead.status} />
                </td>
                <td class="py-3.5 px-6 text-right" on:click|stopPropagation>
                  <a
                    href={getWhatsAppUrl(lead.phone, lead.name)}
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
</div>

<!-- Slide-over Drawer for Lead Details (480px width) -->
{#if selectedLead}
  <div class="fixed inset-0 z-50 overflow-hidden bg-slate-900/40 backdrop-blur-xs flex justify-end animate-in fade-in duration-200">
    <div class="w-full max-w-[480px] bg-white h-full shadow-2xl flex flex-col justify-between overflow-y-auto">
      <div>
        <!-- Drawer Header -->
        <div class="flex items-center justify-between border-b border-slate-200 px-6 py-5">
          <div>
            <div class="text-xs uppercase tracking-wider text-slate-400 font-semibold">Rincian Lead Masuk</div>
            <h3 class="text-lg font-heading font-semibold text-slate-900 mt-0.5">{selectedLead.name}</h3>
          </div>
          <button
            type="button"
            class="text-slate-400 hover:text-slate-700 p-1.5 rounded-md"
            on:click={() => selectedLead = null}
          >
            <X size={18} />
          </button>
        </div>

        <!-- Drawer Content -->
        <div class="p-6 space-y-6">
          <!-- Status Selector -->
          <div class="rounded-lg border border-slate-200 bg-slate-50/70 p-4">
            <label class="block text-xs font-medium text-slate-500 mb-2">
              Ubah Status Tindak Lanjut:
            </label>
            <div class="flex gap-2">
              {#each [LeadStatus.BARU, LeadStatus.DIHUBUNGI, LeadStatus.SELESAI] as st}
                <button
                  type="button"
                  class="flex-1 py-1.5 px-3 rounded-md text-xs font-medium capitalize border transition-colors {selectedLead.status === st ? 'bg-slate-900 text-white border-slate-900' : 'bg-white text-slate-700 border-slate-200 hover:bg-slate-50'}"
                  on:click={() => changeStatus(selectedLead, st)}
                  disabled={updatingStatus}
                >
                  {st}
                </button>
              {/each}
            </div>
          </div>

          <!-- Contact Info -->
          <div class="space-y-3">
            <div class="text-xs font-medium text-slate-500">
              Informasi Kontak
            </div>
            <div class="flex items-center gap-3 text-sm text-slate-900">
              <Phone size={15} class="text-slate-400" />
              <a href="tel:{selectedLead.phone}" class="font-medium hover:underline">
                {selectedLead.phone}
              </a>
            </div>
            {#if selectedLead.email}
              <div class="flex items-center gap-3 text-sm text-slate-700">
                <Mail size={15} class="text-slate-400" />
                <span>{selectedLead.email}</span>
              </div>
            {/if}
            <div class="flex items-center gap-3 text-sm text-slate-500">
              <Calendar size={15} class="text-slate-400" />
              <span>Diterima: {new Date(selectedLead.created_at).toLocaleString('id-ID')}</span>
            </div>
          </div>

          <!-- Message -->
          <div>
            <div class="text-xs font-medium text-slate-500 mb-2">
              Pesan / Keluhan Hama:
            </div>
            <div class="rounded-lg border border-slate-200 bg-slate-50/70 p-4 text-sm leading-relaxed text-slate-800 whitespace-pre-line">
              {selectedLead.message || 'Tidak ada catatan tambahan.'}
            </div>
          </div>
        </div>
      </div>

      <!-- Drawer Footer -->
      <div class="border-t border-slate-200 p-6 bg-slate-50/50 flex gap-3">
        <a
          href={getWhatsAppUrl(selectedLead.phone, selectedLead.name)}
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center justify-center gap-2 rounded-lg bg-emerald-600 hover:bg-emerald-700 py-2.5 px-4 text-sm font-medium text-white transition-colors w-full"
        >
          <MessageCircle size={17} />
          <span>Buka Percakapan WhatsApp</span>
        </a>
      </div>
    </div>
  </div>
{/if}