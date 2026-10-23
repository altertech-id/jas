<script lang="ts">
  import { goto } from '$app/navigation';
  import { auth } from '$lib/stores/auth';
  import { toast } from '$lib/stores/toast';
  import { UserRole } from '$lib/types/enums';
  import Button from '$lib/components/ui/Button.svelte';
  import { ButtonVariant } from '$lib/types/enums';
  import { ShieldCheck, Lock, Mail, ArrowRight } from 'lucide-svelte';

  let email = '';
  let password = '';
  let loading = false;

  async function handleLogin() {
    if (!email || !password) {
      toast.show('Mohon masukkan email dan kata sandi.', 'error');
      return;
    }

    loading = true;
    setTimeout(() => {
      // In production, this verifies via Supabase Auth
      const role = email.includes('admin') ? UserRole.ADMIN : UserRole.EDITOR;
      auth.login('dev-admin-token', {
        id: 'a0000000-0000-0000-0000-000000000001',
        email: email,
        name: email.includes('admin') ? 'Administrator JAS' : 'Editor Konten',
        role: role,
        is_active: true,
        created_at: new Date().toISOString()
      });
      toast.show('Login berhasil. Selamat datang di CMS!', 'success');
      goto('/admin');
      loading = false;
    }, 600);
  }

  function devQuickLogin(role: UserRole) {
    if (role === UserRole.ADMIN) {
      auth.login('dev-admin-token', {
        id: 'a0000000-0000-0000-0000-000000000001',
        email: 'admin@jayaabadisolusi.com',
        name: 'Administrator JAS',
        role: UserRole.ADMIN,
        is_active: true,
        created_at: new Date().toISOString()
      });
    } else {
      auth.login('dev-editor-token', {
        id: 'a0000000-0000-0000-0000-000000000002',
        email: 'editor@jayaabadisolusi.com',
        name: 'Tim Editor JAS',
        role: UserRole.EDITOR,
        is_active: true,
        created_at: new Date().toISOString()
      });
    }
    toast.show(`Login berhasil sebagai ${role.toUpperCase()}.`, 'success');
    goto('/admin');
  }
</script>

<svelte:head>
  <title>Login Staf - Jaya Abadi Solusi CMS</title>
</svelte:head>

<div class="min-h-screen bg-navy-deep flex items-center justify-center p-4">
  <div class="w-full max-w-md rounded-2xl bg-white p-8 shadow-2xl border border-line-soft">
    <!-- Logo -->
    <div class="flex items-center gap-3 justify-center mb-6">
      <div class="flex h-12 w-12 items-center justify-center rounded-xl bg-navy text-white shadow-sm">
        <ShieldCheck size={28} class="text-amber" strokeWidth={2.2} />
      </div>
      <div>
        <div class="font-heading text-xl font-extrabold text-navy leading-none">JAYA ABADI</div>
        <div class="font-heading text-xs font-bold tracking-widest text-brand-blue">PORTAL CMS</div>
      </div>
    </div>

    <div class="text-center mb-6">
      <h1 class="text-xl font-heading font-bold text-ink-900">Masuk ke Panel Konten</h1>
      <p class="text-xs text-ink-500 mt-1">Gunakan akun admin atau editor Anda untuk mengelola website.</p>
    </div>

    <!-- Form -->
    <form on:submit|preventDefault={handleLogin} class="space-y-4">
      <div>
        <label for="admin_email" class="block text-xs font-semibold text-ink-900 mb-1">Email</label>
        <div class="relative">
          <input
            id="admin_email"
            type="email"
            placeholder="admin@jayaabadisolusi.com"
            bind:value={email}
            class="h-11 w-full rounded-xl border border-line-strong pl-10 pr-4 text-sm text-ink-900 focus:border-brand-blue focus:outline-none focus:ring-3 focus:ring-brand-blue/20"
          />
          <Mail size={18} class="absolute left-3 top-3 text-ink-300" />
        </div>
      </div>

      <div>
        <label for="admin_password" class="block text-xs font-semibold text-ink-900 mb-1">Kata Sandi</label>
        <div class="relative">
          <input
            id="admin_password"
            type="password"
            placeholder="••••••••"
            bind:value={password}
            class="h-11 w-full rounded-xl border border-line-strong pl-10 pr-4 text-sm text-ink-900 focus:border-brand-blue focus:outline-none focus:ring-3 focus:ring-brand-blue/20"
          />
          <Lock size={18} class="absolute left-3 top-3 text-ink-300" />
        </div>
      </div>

      <div class="pt-2">
        <Button type="submit" variant={ButtonVariant.SECONDARY} fullWidth {loading}>
          <span>Masuk</span>
          <ArrowRight size={16} class="ml-2" />
        </Button>
      </div>
    </form>

    <!-- Quick Dev Login Shortcut Buttons -->
    <div class="mt-8 border-t border-line-soft pt-6">
      <div class="text-center text-xs font-semibold uppercase tracking-wider text-ink-500 mb-3">
        Akses Cepat Pengujian (Dev Mode)
      </div>

      <div class="grid grid-cols-2 gap-3">
        <button
          type="button"
          class="rounded-xl border border-brand-blue/30 bg-brand-blue-tint/50 p-2.5 text-center text-xs font-semibold text-navy hover:bg-brand-blue-tint transition-colors"
          on:click={() => devQuickLogin(UserRole.ADMIN)}
        >
          Masuk sebagai Admin
        </button>

        <button
          type="button"
          class="rounded-xl border border-line-soft bg-surface-light p-2.5 text-center text-xs font-semibold text-ink-700 hover:bg-surface-mist transition-colors"
          on:click={() => devQuickLogin(UserRole.EDITOR)}
        >
          Masuk sebagai Editor
        </button>
      </div>
    </div>
  </div>
</div>
