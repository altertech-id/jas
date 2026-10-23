<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api/client';
  import type { SiteSettings } from '$lib/types';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { toast } from '$lib/stores/toast';
  import { Settings, Save, Phone, Mail, MapPin } from 'lucide-svelte';

  let settings: SiteSettings = {};
  let loading = true;
  let saving = false;
  let activeTab: 'contact' | 'general' | 'social' = 'contact';

  // Fields
  let phone = '081399484916';
  let phoneDisplay = '0813 9948 4916';
  let email = 'altertech.id@gmail.com';
  let address = 'Jl. Rungkut Industri Raya No. 45, Surabaya, Jawa Timur 60293';
  let workingHours = 'Senin - Sabtu: 08.00 - 17.00 WIB';

  let companyName = 'PT Jaya Abadi Solusi';
  let tagline = 'Solusi Profesional Pengendalian Hama & Sanitasi Surabaya';

  let instagram = 'https://instagram.com/jayaabadisolusi';
  let facebook = 'https://facebook.com/jayaabadisolusi';
  let linkedin = 'https://linkedin.com/company/jayaabadisolusi';

  onMount(async () => {
    try {
      settings = await api.getSiteSettings();
      if (settings.contact) {
        phone = settings.contact.phone || phone;
        phoneDisplay = settings.contact.phone_display || phoneDisplay;
        email = settings.contact.email || email;
        address = settings.contact.address || address;
        workingHours = settings.contact.working_hours || workingHours;
      }
      if (settings.general) {
        companyName = settings.general.company_name || companyName;
        tagline = settings.general.tagline || tagline;
      }
      if (settings.social_media) {
        instagram = settings.social_media.instagram || instagram;
        facebook = settings.social_media.facebook || facebook;
        linkedin = settings.social_media.linkedin || linkedin;
      }
    } catch (err: any) {
      toast.show(err.message, 'error');
    } finally {
      loading = false;
    }
  });

  async function handleSaveContact() {
    saving = true;
    try {
      await api.updateSiteSetting('contact', {
        phone: phone.trim(),
        phone_display: phoneDisplay.trim(),
        whatsapp: phone.trim(),
        whatsapp_url: `https://wa.me/62${phone.replace(/^0/, '')}`,
        email: email.trim(),
        address: address.trim(),
        working_hours: workingHours.trim(),
        emergency_service: 'Layanan Darurat & Konsultasi 24 Jam',
        maps_embed_url: settings.contact?.maps_embed_url || ''
      });
      toast.show('Pengaturan kontak berhasil diperbarui.', 'success');
    } catch (err: any) {
      toast.show(err.message || 'Gagal menyimpan pengaturan.', 'error');
    } finally {
      saving = false;
    }
  }

  async function handleSaveGeneral() {
    saving = true;
    try {
      await api.updateSiteSetting('general', {
        company_name: companyName.trim(),
        short_name: 'Jaya Abadi Solusi',
        tagline: tagline.trim(),
        established_year: 2010,
        years_experience: '15+',
        verified_badge: 'Terpercaya Sejak 2010'
      });
      toast.show('Pengaturan umum berhasil diperbarui.', 'success');
    } catch (err: any) {
      toast.show(err.message || 'Gagal menyimpan pengaturan.', 'error');
    } finally {
      saving = false;
    }
  }

  async function handleSaveSocial() {
    saving = true;
    try {
      await api.updateSiteSetting('social_media', {
        instagram: instagram.trim(),
        facebook: facebook.trim(),
        linkedin: linkedin.trim()
      });
      toast.show('Tautan sosial media berhasil diperbarui.', 'success');
    } catch (err: any) {
      toast.show(err.message || 'Gagal menyimpan pengaturan.', 'error');
    } finally {
      saving = false;
    }
  }
</script>

<svelte:head>
  <title>Pengaturan Situs - Jaya Abadi Solusi Admin</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
    <div>
      <h2 class="text-xl font-heading font-semibold text-slate-900">Pengaturan Situs & Kontak</h2>
      <p class="text-xs text-slate-500 mt-0.5">Konfigurasi nomor WhatsApp publik, alamat Surabaya, dan informasi profil.</p>
    </div>

    <!-- Tabs -->
    <div class="inline-flex rounded-lg border border-slate-200 bg-slate-50/80 p-0.5">
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {activeTab === 'contact' ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => activeTab = 'contact'}
      >
        Kontak & WhatsApp
      </button>
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {activeTab === 'general' ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => activeTab = 'general'}
      >
        Identitas Brand
      </button>
      <button
        type="button"
        class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors {activeTab === 'social' ? 'bg-white text-slate-900 shadow-xs' : 'text-slate-600 hover:text-slate-900'}"
        on:click={() => activeTab = 'social'}
      >
        Sosial Media
      </button>
    </div>
  </div>

  {#if loading}
    <div class="flex justify-center p-12">
      <div class="h-7 w-7 animate-spin rounded-full border-2 border-slate-900 border-t-transparent"></div>
    </div>
  {:else if activeTab === 'contact'}
    <!-- Tab Kontak -->
    <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-5 max-w-2xl">
      <h3 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
        Kontak Layanan Publik & WhatsApp
      </h3>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label for="set_phone" class="block text-xs font-medium text-slate-700 mb-1">Nomor WhatsApp & Telepon</label>
          <input id="set_phone" type="text" bind:value={phone} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
        </div>
        <div>
          <label for="set_phonedisp" class="block text-xs font-medium text-slate-700 mb-1">Format Tampilan UI</label>
          <input id="set_phonedisp" type="text" bind:value={phoneDisplay} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
        </div>
      </div>

      <div>
        <label for="set_email" class="block text-xs font-medium text-slate-700 mb-1">Email Resmi Perusahaan</label>
        <input id="set_email" type="email" bind:value={email} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
      </div>

      <div>
        <label for="set_addr" class="block text-xs font-medium text-slate-700 mb-1">Alamat Kantor Surabaya</label>
        <textarea id="set_addr" rows={2} bind:value={address} class="w-full rounded-lg border border-slate-300 bg-white p-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"></textarea>
      </div>

      <div>
        <label for="set_hours" class="block text-xs font-medium text-slate-700 mb-1">Jam Operasional</label>
        <input id="set_hours" type="text" bind:value={workingHours} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
      </div>

      <div class="pt-2 flex justify-end">
        <Button variant={ButtonVariant.PRIMARY} size="default" loading={saving} on:click={handleSaveContact}>
          <Save size={15} class="mr-2" />
          <span>Simpan Pengaturan Kontak</span>
        </Button>
      </div>
    </div>
  {:else if activeTab === 'general'}
    <!-- Tab General -->
    <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-5 max-w-2xl">
      <h3 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
        Identitas & Slogan Perusahaan
      </h3>

      <div>
        <label for="set_cname" class="block text-xs font-medium text-slate-700 mb-1">Nama Perusahaan</label>
        <input id="set_cname" type="text" bind:value={companyName} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
      </div>

      <div>
        <label for="set_tagline" class="block text-xs font-medium text-slate-700 mb-1">Tagline Utama</label>
        <textarea id="set_tagline" rows={2} bind:value={tagline} class="w-full rounded-lg border border-slate-300 bg-white p-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy"></textarea>
      </div>

      <div class="pt-2 flex justify-end">
        <Button variant={ButtonVariant.PRIMARY} size="default" loading={saving} on:click={handleSaveGeneral}>
          <Save size={15} class="mr-2" />
          <span>Simpan Identitas Brand</span>
        </Button>
      </div>
    </div>
  {:else}
    <!-- Tab Social -->
    <div class="rounded-xl border border-slate-200/80 bg-white p-6 shadow-xs space-y-5 max-w-2xl">
      <h3 class="text-sm font-heading font-semibold text-slate-900 border-b border-slate-100 pb-2">
        Tautan Akun Sosial Media
      </h3>

      <div>
        <label for="set_ig" class="block text-xs font-medium text-slate-700 mb-1">Instagram URL</label>
        <input id="set_ig" type="url" bind:value={instagram} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
      </div>

      <div>
        <label for="set_fb" class="block text-xs font-medium text-slate-700 mb-1">Facebook URL</label>
        <input id="set_fb" type="url" bind:value={facebook} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
      </div>

      <div>
        <label for="set_in" class="block text-xs font-medium text-slate-700 mb-1">LinkedIn URL</label>
        <input id="set_in" type="url" bind:value={linkedin} class="h-10 w-full rounded-lg border border-slate-300 bg-white px-3 text-sm text-slate-900 placeholder:text-slate-400 focus:border-navy focus:outline-none focus:ring-1 focus:ring-navy" />
      </div>

      <div class="pt-2 flex justify-end">
        <Button variant={ButtonVariant.PRIMARY} size="default" loading={saving} on:click={handleSaveSocial}>
          <Save size={15} class="mr-2" />
          <span>Simpan Sosial Media</span>
        </Button>
      </div>
    </div>
  {/if}
</div>
