import { createClient } from '@supabase/supabase-js';

const supabaseUrl = (import.meta.env.PUBLIC_SUPABASE_URL as string) || 'https://glqtzcmzkkhqcgyjesxt.supabase.co';
const supabaseKey = (import.meta.env.PUBLIC_SUPABASE_PUBLISHABLE_KEY as string) || 
  (import.meta.env.PUBLIC_SUPABASE_ANON_KEY as string) || 
  'sb_publishable_Xet_GY8ON937gExWjKBsDg_d5JZaKyE';

export const supabase = createClient(supabaseUrl, supabaseKey);
