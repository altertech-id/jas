# Jaya Abadi Solusi (JAS) - Web & Backoffice CMS

Sistem website company profile, 5 landing page spesialis konversi iklan, dan backoffice admin CMS untuk **PT Jaya Abadi Solusi (JAS)** Surabaya.

---

## 🛠️ Tech Stack & Arsitektur

- **Frontend:** [SvelteKit](https://kit.svelte.dev/) (SSR untuk performa & SEO Lighthouse > 90), [Tailwind CSS](https://tailwindcss.com/) dengan pemetaan token design system, [TipTap](https://tiptap.dev/) rich text editor, dan [Lucide Svelte](https://lucide.dev/).
- **Backend:** [Go](https://go.dev/) (Framework: **Fiber v2**) sebagai REST API performa tinggi, sanitasi HTML dengan `bluemonday`, middleware JWT, dan rate limiting.
- **Database, Auth & Storage:** Supabase Cloud (PostgreSQL dengan 10 tabel, Supabase Auth, dan Supabase Storage bucket `media`).
- **Dev Fallback:** Dilengkapi dengan *In-Memory Seed Repository* sehingga frontend dan API dapat langsung dijalankan dan diuji secara lokal tanpa memerlukan instance database eksternal.

---

## 📂 Struktur Monorepo

```
jayaabadisolusi/
├── backend/                    # Go Fiber REST API
│   ├── cmd/api/main.go         # Entry point server
│   ├── internal/
│   │   ├── config/             # Environment loader
│   │   ├── domain/             # Enums & model data
│   │   ├── handler/            # Public & Admin handlers
│   │   ├── middleware/         # Auth JWT, Role Guard & Rate Limiter
│   │   ├── repository/         # Postgres + In-memory seed fallback
│   │   └── service/            # Supabase Storage & HTML Sanitizer
│   ├── Dockerfile
│   └── go.mod
├── frontend/                   # SvelteKit SSR App
│   ├── src/
│   │   ├── app.css             # Design tokens & CSS variables
│   │   ├── lib/
│   │   │   ├── api/client.ts   # Client API terintegrasi
│   │   │   ├── components/     # UI, Public, dan Admin components
│   │   │   ├── stores/         # Auth & Toast notification stores
│   │   │   └── types/enums.ts  # Standardized TypeScript Enums
│   │   └── routes/
│   │       ├── (public)/       # Beranda, 5 Landing Layanan, Project, Blog, Tentang, Kontak, Galeri
│   │       └── admin/          # Dashboard, Inbox Lead, Layanan, Project, Blog, Pages, Galeri, Media, Pengaturan, Pengguna
│   ├── Dockerfile
│   └── package.json
├── supabase/
│   └── migrations/
│       ├── 001_initial_schema.sql    # 10 tabel, enums, RLS & indeks
│       └── 002_seed_initial_data.sql # Seed data 5 layanan, kontak, dan halaman
├── docker-compose.yml          # Orkestrasi Docker untuk produksi
└── .env.example                # Template konfigurasi environment
```

---

## 🚀 Panduan Menjalankan Secara Lokal

### 1. Menjalankan Backend (Go API)

```bash
cd backend
go run ./cmd/api
```
*Server API aktif di `http://localhost:8080` (port default).*

### 2. Menjalankan Frontend (SvelteKit)

Buka terminal baru:
```bash
cd frontend
npm install
npm run dev
```
*Buka browser di `http://localhost:5173`.*

---

## 🔑 Akses Portal Admin

- Buka URL: `http://localhost:5173/admin/login`
- Anda dapat mengklik tombol **"Masuk sebagai Admin"** atau **"Masuk sebagai Editor"** untuk pengujian langsung tanpa perlu konfigurasi Supabase awal.

---

## 📞 Informasi Kontak & Konfigurasi Default

- **Nomor Telepon & WhatsApp:** `0813 9948 4916` (`081399484916`)
- **Email:** `altertech.id@gmail.com`
- **Alamat:** Jl. Rungkut Industri Raya No. 45, Surabaya, Jawa Timur 60293
