import { writable, derived } from 'svelte/store';
import type { Profile } from '../types';
import { UserRole } from '../types/enums';

interface AuthState {
  isAuthenticated: boolean;
  token: string | null;
  profile: Profile | null;
  loading: boolean;
}

const initialState: AuthState = {
  isAuthenticated: false,
  token: null,
  profile: null,
  loading: true
};

function createAuthStore() {
  const { subscribe, set, update } = writable<AuthState>(initialState);

  return {
    subscribe,
    init: () => {
      if (typeof window === 'undefined') return;
      const token = localStorage.getItem('jas_auth_token');
      const profileStr = localStorage.getItem('jas_auth_profile');

      if (token && profileStr) {
        try {
          const profile = JSON.parse(profileStr);
          set({
            isAuthenticated: true,
            token,
            profile,
            loading: false
          });
          return;
        } catch {
          // invalid storage
        }
      }
      set({ ...initialState, loading: false });
    },
    login: (token: string, profile: Profile) => {
      if (typeof window !== 'undefined') {
        localStorage.setItem('jas_auth_token', token);
        localStorage.setItem('jas_auth_profile', JSON.stringify(profile));
      }
      set({
        isAuthenticated: true,
        token,
        profile,
        loading: false
      });
    },
    logout: () => {
      if (typeof window !== 'undefined') {
        localStorage.removeItem('jas_auth_token');
        localStorage.removeItem('jas_auth_profile');
      }
      set({
        isAuthenticated: false,
        token: null,
        profile: null,
        loading: false
      });
    }
  };
}

export const auth = createAuthStore();

export const isAdmin = derived(auth, ($auth) => {
  return $auth.isAuthenticated && $auth.profile?.role === UserRole.ADMIN;
});

export const isEditor = derived(auth, ($auth) => {
  return $auth.isAuthenticated && ($auth.profile?.role === UserRole.EDITOR || $auth.profile?.role === UserRole.ADMIN);
});
