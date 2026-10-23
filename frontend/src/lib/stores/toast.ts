import { writable } from 'svelte/store';

export interface ToastItem {
  id: string;
  type: 'success' | 'error' | 'warning' | 'info';
  message: string;
}

function createToastStore() {
  const { subscribe, update } = writable<ToastItem[]>([]);

  return {
    subscribe,
    show: (message: string, type: 'success' | 'error' | 'warning' | 'info' = 'success', duration = 4000) => {
      const id = Math.random().toString(36).substring(2, 9);
      update((toasts) => [...toasts, { id, type, message }]);

      setTimeout(() => {
        update((toasts) => toasts.filter((t) => t.id !== id));
      }, duration);
    },
    remove: (id: string) => {
      update((toasts) => toasts.filter((t) => t.id !== id));
    }
  };
}

export const toast = createToastStore();
