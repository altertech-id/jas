<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { Profile } from '$lib/types';
  import StatusPill from '$lib/components/ui/StatusPill.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { Users, UserCheck, UserX, Shield } from 'lucide-svelte';

  let profiles: Profile[] = [];
  let loading = true;

  onMount(async () => {
    await loadUsers();
  });

  async function loadUsers() {
    loading = true;
    try {
      profiles = await api.getUsers();
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  }

  async function toggleStatus(p: Profile) {
    const newStatus = !p.is_active;
    try {
      await api.updateUserStatus(p.id, newStatus);
      p.is_active = newStatus;
      profiles = [...profiles];
      toast.show(`Status pengguna ${p.name} berhasil diubah.`, 'success');
    } catch (err: any) {
      toast.show(err.message || 'Gagal mengubah status pengguna.', 'error');
    }
  }
</script>

<svelte:head>
  <title>Manajemen Pengguna - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  <div>
    <h2 class="text-xl font-heading font-semibold text-slate-900">Manajemen Pengguna (Staf & Editor)</h2>
    <p class="text-xs text-slate-500 mt-0.5">Kelola hak akses pengguna internal panel CMS.</p>
  </div>

  <div class="rounded-xl border border-slate-200/80 bg-white overflow-hidden shadow-xs">
    {#if loading}
      <div class="flex justify-center p-12">
        <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
      </div>
    {:else}
      <div class="overflow-x-auto">
        <table class="w-full text-left text-sm">
          <thead class="bg-slate-50/80 text-xs font-medium text-slate-500 border-b border-slate-200/80">
            <tr>
              <th class="py-3 px-6">Nama Pengguna</th>
              <th class="py-3 px-6">Email Akun</th>
              <th class="py-3 px-6">Hak Akses (Role)</th>
              <th class="py-3 px-6">Status Akun</th>
              <th class="py-3 px-6 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100">
            {#each profiles as p}
              <tr class="hover:bg-slate-50/60 transition-colors">
                <td class="py-3.5 px-6 font-medium text-slate-900">
                  <div class="flex items-center gap-3">
                    <div class="flex h-7 w-7 items-center justify-center rounded-full bg-slate-900 text-white text-xs font-bold">
                      {p.name.charAt(0)}
                    </div>
                    <span>{p.name}</span>
                  </div>
                </td>
                <td class="py-3.5 px-6 text-xs text-slate-600">
                  {p.email}
                </td>
                <td class="py-3.5 px-6">
                  <span class="inline-flex items-center gap-1 rounded px-2 py-0.5 text-xs font-medium uppercase {p.role === 'admin' ? 'bg-slate-100 text-slate-800' : 'bg-slate-100 text-slate-600'}">
                    <Shield size={11} />
                    <span>{p.role}</span>
                  </span>
                </td>
                <td class="py-3.5 px-6">
                  <StatusPill status={p.is_active ? 'aktif' : 'nonaktif'} />
                </td>
                <td class="py-3.5 px-6 text-right">
                  <button
                    type="button"
                    class="inline-flex items-center gap-1 rounded-md border border-slate-200 px-2.5 py-1 text-xs font-medium transition-colors {p.is_active ? 'text-red-600 hover:bg-red-50' : 'text-emerald-700 hover:bg-emerald-50'}"
                    on:click={() => toggleStatus(p)}
                  >
                    {#if p.is_active}
                      <UserX size={13} />
                      <span>Nonaktifkan</span>
                    {:else}
                      <UserCheck size={13} />
                      <span>Aktifkan</span>
                    {/if}
                  </button>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  </div>
</div>
