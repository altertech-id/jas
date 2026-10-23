<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { CheckCircle2, Send, MessageCircle } from 'lucide-svelte';

  export let sourcePage = 'Website Form';
  export let selectedServiceSlug: string | undefined = undefined;

  let name = '';
  let phone = '';
  let email = '';
  let serviceId = '';
  let message = '';
  let hpField = ''; // Honeypot anti-spam
  let loading = false;
  let submittedSuccess = false;
  let whatsappRedirectUrl = '';

  const serviceOptions = [
    { id: 'b0000000-0000-0000-0000-000000000001', slug: 'pest-control', name: 'Pest Control Umum (Kecoa, Semut, Lalat)' },
    { id: 'b0000000-0000-0000-0000-000000000002', slug: 'rodent-control', name: 'Rodent Control (Tikus & Got)' },
    { id: 'b0000000-0000-0000-0000-000000000003', slug: 'termite-control', name: 'Termite Control (Anti Rayap Pondasi)' },
    { id: 'b0000000-0000-0000-0000-000000000004', slug: 'fumigasi-pe', name: 'Fumigasi PE & Komoditas Ekspor' },
    { id: 'b0000000-0000-0000-0000-000000000005', slug: 'disinfeksi', name: 'Disinfeksi & Sterilisasi Ruangan' }
  ];

  onMount(() => {
    if (selectedServiceSlug) {
      const match = serviceOptions.find(s => s.slug === selectedServiceSlug);
      if (match) serviceId = match.id;
    }
  });

  async function handleSubmit() {
    if (!name.trim()) {
      toast.show('Mohon lengkapi nama Anda.', 'error');
      return;
    }
    if (!phone.trim() || phone.trim().length < 9) {
      toast.show('Nomor telepon / WhatsApp minimal 9 digit.', 'error');
      return;
    }

    loading = true;
    try {
      const res = await api.submitLead({
        name: name.trim(),
        phone: phone.trim(),
        email: email.trim(),
        service_id: serviceId || undefined,
        message: message.trim(),
        source_page: sourcePage,
        hp_field: hpField
      });

      submittedSuccess = true;
      whatsappRedirectUrl = res.whatsapp_url;
      toast.show(res.message, 'success');
    } catch (err: any) {
      toast.show(err.message || 'Gagal mengirim formulir. Coba lagi atau hubungi via WhatsApp.', 'error');
    } finally {
      loading = false;
    }
  }

  function resetForm() {
    name = '';
    phone = '';
    email = '';
    message = '';
    submittedSuccess = false;
  }
</script>

<div class="rounded-xl border border-slate-200/80 bg-white p-6 sm:p-8 shadow-xs">
  {#if submittedSuccess}
    <!-- Success State -->
    <div class="flex flex-col items-center text-center py-6 animate-in fade-in duration-300">
      <div class="mb-3 flex h-12 w-12 items-center justify-center rounded-full bg-emerald-50 text-emerald-600">
        <CheckCircle2 size={28} />
      </div>

      <h3 class="mb-2 text-xl font-heading font-bold text-slate-900">
        Permintaan Terkirim
      </h3>

      <p class="mb-6 max-w-md text-sm leading-relaxed text-slate-600 font-body">
        Terima kasih, <strong class="text-slate-900">{name}</strong>. Tim teknisi Jaya Abadi Solusi akan segera menghubungi nomor telepon <strong class="text-slate-900">{phone}</strong> untuk konfirmasi jadwal survei.
      </p>

      <div class="flex flex-col sm:flex-row gap-3 w-full max-w-sm">
        {#if whatsappRedirectUrl}
          <a
            href={whatsappRedirectUrl}
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center justify-center gap-2 rounded-lg bg-emerald-600 hover:bg-emerald-700 py-2.5 px-4 text-sm font-medium text-white transition-colors w-full"
          >
            <MessageCircle size={16} />
            <span>Konfirmasi via WhatsApp</span>
          </a>
        {/if}

        <Button variant={ButtonVariant.OUTLINE} size="default" fullWidth on:click={resetForm}>
          Kirim Form Baru
        </Button>
      </div>
    </div>
  {:else}
    <!-- Form State -->
    <div class="mb-6">
      <h3 class="text-xl font-heading font-bold text-slate-900">
        Minta Penawaran & Jadwal Survei
      </h3>
      <p class="mt-1 text-sm text-slate-500 font-body">
        Isi formulir di bawah ini. Tim kami akan memberikan estimasi penanganan dan survei lokasi gratis.
      </p>
    </div>

    <form on:submit|preventDefault={handleSubmit} class="space-y-4">
      <!-- Honeypot (hidden from real users) -->
      <div class="hidden" aria-hidden="true">
        <label for="hp_field">Jangan diisi:</label>
        <input type="text" id="hp_field" name="hp_field" bind:value={hpField} tabindex="-1" autocomplete="off" />
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <!-- Nama -->
        <div>
          <label for="lead_name" class="block text-xs font-medium text-slate-700 mb-1.5">
            Nama Lengkap <span class="text-red-500">*</span>
          </label>
          <input
            id="lead_name"
            type="text"
            required
            placeholder="Contoh: Budi Santoso"
            bind:value={name}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy transition-colors"
          />
        </div>

        <!-- Nomor Telepon / WA -->
        <div>
          <label for="lead_phone" class="block text-xs font-medium text-slate-700 mb-1.5">
            Nomor Telepon / WhatsApp <span class="text-red-500">*</span>
          </label>
          <input
            id="lead_phone"
            type="tel"
            required
            placeholder="Contoh: 0812 3456 7890"
            bind:value={phone}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy transition-colors"
          />
        </div>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <!-- Email -->
        <div>
          <label for="lead_email" class="block text-xs font-medium text-slate-700 mb-1.5">
            Alamat Email (Opsional)
          </label>
          <input
            id="lead_email"
            type="email"
            placeholder="nama@perusahaan.com"
            bind:value={email}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy transition-colors"
          />
        </div>

        <!-- Pilihan Layanan -->
        <div>
          <label for="lead_service" class="block text-xs font-medium text-slate-700 mb-1.5">
            Jenis Layanan
          </label>
          <select
            id="lead_service"
            bind:value={serviceId}
            class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3.5 text-sm text-slate-900 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy transition-colors"
          >
            <option value="">Pilih Kebutuhan Layanan...</option>
            {#each serviceOptions as s}
              <option value={s.id}>{s.name}</option>
            {/each}
          </select>
        </div>
      </div>

      <!-- Pesan / Catatan -->
      <div>
        <label for="lead_message" class="block text-xs font-medium text-slate-700 mb-1.5">
          Ceritakan Masalah Hama Anda
        </label>
        <textarea
          id="lead_message"
          rows={3}
          placeholder="Tuliskan lokasi, jenis hama yang terlihat, atau perkiraan luas bangunan..."
          bind:value={message}
          class="w-full rounded-lg border border-slate-300 bg-white p-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy transition-colors"
        ></textarea>
      </div>

      <!-- Submit Button -->
      <div class="pt-2">
        <Button
          type="submit"
          variant={ButtonVariant.PRIMARY}
          size="default"
          fullWidth
          {loading}
        >
          <Send size={15} class="mr-2" />
          <span>Kirim Permintaan Survei & Penawaran</span>
        </Button>
      </div>

      <p class="text-center text-xs text-slate-400">
        Data Anda terjamin kerahasiaannya dan hanya digunakan untuk keperluan survei & penawaran resmi PT Jaya Abadi Solusi.
      </p>
    </form>
  {/if}
</div>