<script lang="ts">
  import { ButtonVariant } from '$lib/types/enums';

  export let variant: ButtonVariant | string = ButtonVariant.PRIMARY;
  export let size: 'default' | 'sm' = 'default';
  export let href: string | undefined = undefined;
  export let type: 'button' | 'submit' | 'reset' = 'button';
  export let disabled = false;
  export let loading = false;
  export let fullWidth = false;

  const baseStyles = 'inline-flex items-center justify-center font-heading font-semibold rounded-full transition-all duration-150 ease-out focus:outline-none disabled:cursor-not-allowed select-none';

  const sizeStyles = {
    default: 'h-11 sm:h-12 px-7 text-sm tracking-wide',
    sm: 'h-9 sm:h-10 px-5 text-xs sm:text-sm tracking-wide'
  };

  const variantStyles: Record<string, string> = {
    [ButtonVariant.PRIMARY]: 'bg-cta hover:bg-cta-hover active:bg-cta-active text-white shadow-xs',
    [ButtonVariant.SECONDARY]: 'bg-navy hover:bg-navy-deep text-white shadow-xs',
    [ButtonVariant.OUTLINE]: 'bg-transparent border border-navy text-navy hover:bg-navy/5',
    [ButtonVariant.OUTLINE_ON_NAVY]: 'bg-transparent border border-white text-white hover:bg-white/10',
    [ButtonVariant.WHATSAPP]: 'bg-wa hover:bg-wa-hover text-white shadow-xs',
    [ButtonVariant.GHOST]: 'bg-transparent text-navy hover:text-brand-blue p-0 h-auto font-medium'
  };
</script>

{#if href && !disabled}
  <a
    {href}
    class="{baseStyles} {variant === ButtonVariant.GHOST ? '' : sizeStyles[size]} {variantStyles[variant]} {fullWidth ? 'w-full' : ''} {$$props.class || ''}"
    on:click
  >
    {#if loading}
      <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-current" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
      </svg>
    {/if}
    <slot />
  </a>
{:else}
  <button
    {type}
    {disabled}
    class="{baseStyles} {variant === ButtonVariant.GHOST ? '' : sizeStyles[size]} {variantStyles[variant]} {fullWidth ? 'w-full' : ''} {$$props.class || ''}"
    on:click
  >
    {#if loading}
      <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-current" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
      </svg>
    {/if}
    <slot />
  </button>
{/if}
