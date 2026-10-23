-- Migration 002: Seed Initial Data for Jaya Abadi Solusi (JAS)

-- 1. PROFILES
INSERT INTO profiles (id, email, name, role, is_active)
VALUES 
    ('a0000000-0000-0000-0000-000000000001', 'admin@jayaabadisolusi.com', 'Administrator JAS', 'admin', true),
    ('a0000000-0000-0000-0000-000000000002', 'editor@jayaabadisolusi.com', 'Tim Konten JAS', 'editor', true)
ON CONFLICT (id) DO NOTHING;

-- 2. SITE SETTINGS (Using phone 081399484916 & email altertech.id@gmail.com)
INSERT INTO site_settings (key, value)
VALUES 
('general', '{
    "company_name": "PT Jaya Abadi Solusi",
    "short_name": "Jaya Abadi Solusi",
    "tagline": "Solusi Profesional Pengendalian Hama & Sanitasi Surabaya",
    "established_year": 2010,
    "years_experience": "15+",
    "verified_badge": "Terpercaya Sejak 2010"
}'::jsonb),
('contact', '{
    "phone": "081399484916",
    "phone_display": "0813 9948 4916",
    "whatsapp": "081399484916",
    "whatsapp_url": "https://wa.me/6281399484916",
    "email": "altertech.id@gmail.com",
    "address": "Jl. Rungkut Industri Raya No. 45, Surabaya, Jawa Timur 60293",
    "working_hours": "Senin - Sabtu: 08.00 - 17.00 WIB",
    "emergency_service": "Layanan Darurat & Konsultasi 24 Jam",
    "maps_embed_url": "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3957.2723789423985!2d112.75382587574577!3d-7.323288472017329!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd7facdb18e59e3%3A0x6b2b289069d5ceb8!2sKawasan%20Industri%20Rungkut!5e0!3m2!1sid!2sid!4v1700000000000!5m2!1sid!2sid"
}'::jsonb),
('social_media', '{
    "instagram": "https://instagram.com/jayaabadisolusi",
    "facebook": "https://facebook.com/jayaabadisolusi",
    "linkedin": "https://linkedin.com/company/jayaabadisolusi"
}'::jsonb),
('seo_default', '{
    "title": "Jaya Abadi Solusi - Jasa Pest Control & Pengendalian Hama Surabaya",
    "description": "Jasa pembasmi hama tikus, rayap, serangga, fumigasi PE dan disinfeksi di Surabaya. Berpengalaman 15+ tahun, garansi tuntas dan aman.",
    "og_image": "https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=1200&q=80"
}'::jsonb)
ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value, updated_at = now();

-- 3. 5 SERVICES (Pest Control, Rodent Control, Termite Control, Fumigasi PE, Disinfeksi)
INSERT INTO services (id, slug, name, icon, short_description, content, seo_title, seo_description, is_published, sort_order)
VALUES
(
    'b0000000-0000-0000-0000-000000000001',
    'pest-control',
    'Pest Control Umum',
    'ShieldCheck',
    'Pengendalian serangga terbang dan merayap seperti kecoa, semut, dan lalat menggunakan metode terpadu yang higienis serta aman.',
    '{
        "hero_badge": "Layanan Paling Diminati",
        "headline": "Basmi Kecoa, Semut, dan Lalat Sampai ke Sarangnya",
        "tagline": "Metode Integrated Pest Management (IPM) ramah lingkungan untuk rumah tangga, restoran, kantor, dan gudang.",
        "problems": [
            "Kecoa berseliweran di dapur dan area saji, mencemari makanan.",
            "Lalat beterbangan menurunkan citra higienitas tempat usaha Anda.",
            "Semut merusak instalasi kabel dan mengganggu kenyamanan keluarga.",
            "Penggunaan insektisida semprot biasa hanya membunuh sementara."
        ],
        "solutions": [
            "Inspeksi mendalam ke celah dan sarang tersembunyi.",
            "Pemberian gel umpan khusus (baiting) yang bekerja berantai.",
            "Cold fogging (ULV) untuk area indoor tanpa residu minyak.",
            "Residual spraying pada titik jalan masuk hama secara terukur."
        ],
        "methods": [
            {"title": "Cold Fogging (ULV)", "desc": "Pengabutan partikel mikron halus untuk membasmi serangga terbang dengan cepat tanpa noda."},
            {"title": "Baiting Gel Khusus", "desc": "Umpan gel atraktif berdaya bunuh berantai hingga membasmi koloni sarang kecoa."},
            {"title": "Residual Spraying", "desc": "Penyemprotan bahan aktif berizin kementerian di jalur dinding dan celah tersembunyi."},
            {"title": "Monitoring Rutin", "desc": "Pemasangan lem perangkap monitoring untuk evaluasi populasi hama secara berkala."}
        ],
        "benefits": [
            "Bahan kimia berizin resmi Kementerian Pertanian & Kemenkes",
            "Aman untuk anak-anak dan hewan peliharaan setelah treatment",
            "Laporan inspeksi digital lengkap dan rekomendasi sanitasi",
            "Jadwal fleksibel termasuk malam hari untuk restoran dan kantor"
        ],
        "faq": [
            {"q": "Apakah makanan harus dikosongkan saat treatment?", "a": "Untuk area dapur, bahan makanan terbuka sebaiknya ditutup rapat atau disimpan dalam kulkas selama proses spraying dan fogging."},
            {"q": "Berapa lama ruangan boleh digunakan kembali?", "a": "Ruangan dapat digunakan kembali setelah ventilasi udara dibuka sekitar 30 sampai 60 menit setelah treatment selesai."}
        ]
    }'::jsonb,
    'Jasa Pest Control Surabaya - Basmi Serangga Terpercaya | JAS',
    'Jasa pembasmi kecoa, semut, lalat di Surabaya. Pengendalian hama terpadu bergaransi untuk hunian dan tempat usaha. Hubungi 0813 9948 4916.',
    true,
    1
),
(
    'b0000000-0000-0000-0000-000000000002',
    'rodent-control',
    'Rodent Control',
    'ShieldAlert',
    'Pembasmian dan pencegahan tikus got, tikus atap, dan mencit dengan metode trapping higienis dan baiting aman tanpa bau bangkai.',
    '{
        "hero_badge": "Spesialis Masalah Tikus",
        "headline": "Tuntas Basmi Tikus Tanpa Bau Bangkai Membusuk",
        "tagline": "Kombinasi bait station berkunci, perangkap mekanik modern, dan penutupan jalur masuk (proofing).",
        "problems": [
            "Kabel listrik, plafon, dan arsip dokumen habis digerogoti tikus.",
            "Kotoran dan urin tikus mencemari area makan serta gudang penyimpanan.",
            "Kekhawatiran bangkai tikus mati di tempat tersembunyi dan berbau busuk.",
            "Tikus terus datang kembali karena jalur masuk luar belum ditutup."
        ],
        "solutions": [
            "Identifikasi jalur jelajah (runway) dan lubang sarang utama.",
            "Penggunaan rodentisida antikoagulan yang membuat tikus mencari cahaya dan air ke luar ruangan.",
            "Penempatan tamper-resistant bait box yang aman dari jangkauan anak dan hewan peliharaan.",
            "Rekomendasi dan tindakan rodent proofing untuk menutup akses masuk."
        ],
        "methods": [
            {"title": "Tamper-Proof Bait Station", "desc": "Kotak umpan terkunci berbahan kokoh dipasang di sekeliling perimeter luar dan dalam."},
            {"title": "Mechanical & Glue Trapping", "desc": "Perangkap mekanik higienis ditempatkan di jalur jelajah tanpa risiko racun terbuka."},
            {"title": "Tracking & Runway Detection", "desc": "Pemeriksaan jejak rub mark dan urine dengan lampu UV khusus untuk presisi aksi."},
            {"title": "Rodent Proofing", "desc": "Penutupan celah lubang pipa, ventilasi, dan bawah pintu dengan material tahan gigitan."}
        ],
        "benefits": [
            "Formula antikoagulan mencegah tikus mati mendadak di dalam plafon",
            "Bait box aman terkunci untuk pabrik makanan berstandar HACCP",
            "Garansi penanganan ulang jika populasi tikus masih terdeteksi",
            "Bebas bau dan aman untuk lingkungan kerja"
        ],
        "faq": [
            {"q": "Apakah racun tikus berbahaya bagi hewan peliharaan?", "a": "Kami menempatkan umpan di dalam Bait Box terkunci khusus yang hanya bisa diakses tikus, sehingga kucing atau anjing tidak bisa menjangkaunya."},
            {"q": "Bagaimana bila ada tikus mati di dalam?", "a": "Tim teknisi kami siap membantu evakuasi bangkai dan melakukan disinfeksi area agar bebas dari bau dan bakteri."}
        ]
    }'::jsonb,
    'Jasa Pembasmi Tikus Surabaya - Bebas Tikus & Bau | JAS',
    'Solusi tuntas pembasmian tikus di rumah, gudang, pabrik, dan restoran di Surabaya. Higienis, tanpa bau bangkai, garansi tuntas.',
    true,
    2
),
(
    'b0000000-0000-0000-0000-000000000003',
    'termite-control',
    'Termite Control',
    'Home',
    'Perlindungan struktur bangunan dari rayap kayu dan rayap tanah melalui metode injeksi kimia tanah (soil barrier) serta sistem umpan.',
    '{
        "hero_badge": "Garansi Hingga 5 Tahun",
        "headline": "Lindungi Bangunan dan Furnitur Kayu dari Bahaya Rayap",
        "tagline": "Pencegahan pra-konstruksi dan penanganan pasca-konstruksi bergaransi resmi untuk hunian dan properti komersial.",
        "problems": [
            "Kusen pintu, kitchen set, dan plafon keropos secara tiba-tiba.",
            "Dinding muncul jalur tanah rayap yang merusak estetika dan struktur.",
            "Nilai properti anjlok akibat kerusakan struktural yang tak terdeteksi.",
            "Semprotan serangga biasa tidak mampu menyentuh ratu rayap di bawah tanah."
        ],
        "solutions": [
            "Deteksi kelembapan dan koloni aktif menggunakan alat pendeteksi modern.",
            "Metode Termite Baiting System untuk mengeliminasi koloni hingga ke ratu.",
            "Chemical Barrier System: injeksi termitisida di sekeliling pondasi.",
            "Sertifikat garansi berkala dengan inspeksi kontrol berkala tanpa biaya tambahan."
        ],
        "methods": [
            {"title": "Soil Treatment / Injeksi Pondasi", "desc": "Pengeboran presisi dan injeksi termitisida non-repellent pembentuk benteng perlindungan pondasi."},
            {"title": "Termite Baiting System", "desc": "Pemasangan stasiun umpan berinhibitor sintesis kitin yang dibagikan pekerja ke ratu rayap."},
            {"title": "Wood Spraying & Injection", "desc": "Penyuntikan bahan pengawet kayu khusus pada kayu kusen dan atap rangka."},
            {"title": "Inspeksi Berkala Tahunan", "desc": "Kunjungan rutin teknisi senior untuk memastikan properti Anda tetap bebas rayap."}
        ],
        "benefits": [
            "Sertifikat garansi resmi hingga 3 sampai 5 tahun",
            "Termitisida non-repellent ramah lingkungan bersertifikasi",
            "Peralatan injeksi modern dengan kerapihan pengerjaan tinggi",
            "Survei lokasi dan konsultasi teknis gratis di wilayah Surabaya"
        ],
        "faq": [
            {"q": "Apakah lantai keramik akan rusak saat diinjeksi?", "a": "Teknisi kami menggunakan mata bor khusus berdiameter kecil pada nat keramik, lalu menutupnya kembali dengan semen warna senada sehingga tampak rapi seperti semula."},
            {"q": "Berapa lama garansi rayap berlaku?", "a": "Kami memberikan garansi 3 tahun untuk bangunan pasca-konstruksi dan hingga 5 tahun untuk perlakuan pra-konstruksi."}
        ]
    }'::jsonb,
    'Jasa Anti Rayap Surabaya - Bergaransi Resmi Hingga 5 Tahun | JAS',
    'Jasa pembasmi rayap dan perlindungan bangunan di Surabaya. Metode injeksi soil barrier dan termite baiting sistem bergaransi.',
    true,
    3
),
(
    'b0000000-0000-0000-0000-000000000004',
    'fumigasi-pe',
    'Fumigasi PE & Komoditas',
    'Package',
    'Fumigasi standar karantina komoditas ekspor, pergudangan arsip, dan kontainer menggunakan fosfin (PH3) bersertifikat resmi.',
    '{
        "hero_badge": "Standar Karantina Resmi",
        "headline": "Fumigasi Kontainer & Komoditas Ekspor Bersertifikat",
        "tagline": "Layanan fumigasi fosfin (PH3) sesuai standar ISPM 15 dan Barantan untuk kelancaran logistik ekspor Anda.",
        "problems": [
            "Risiko penolakan kontainer di pelabuhan tujuan akibat hama kutu hidup.",
            "Kerusakan komoditas pangan, tembakau, rempah, atau kayu di gudang penyimpanan.",
            "Persyaratan karantina ekspor yang ketat dan membutuhkan sertifikat resmi.",
            "Proses fumigasi mandiri yang berisiko tinggi terhadap keselamatan pekerja."
        ],
        "solutions": [
            "Fumigasi sungkup gas-tight menggunakan terpal PVC standar kedap udara.",
            "Penggunaan fumigan Phostoxin / Fosfin (PH3) murni dengan takaran presisi.",
            "Monitoring konsentrasi gas secara berkala selama masa aerasi.",
            "Penerbitan Fumigation Certificate resmi yang diakui otoritas karantina internasional."
        ],
        "methods": [
            {"title": "Container Fumigation", "desc": "Perlakuan fumigasi langsung di dalam kontainer 20ft dan 40ft di depo atau pelabuhan."},
            {"title": "Sheet Fumigation (Sungkup)", "desc": "Fumigasi tumpukan komoditas gudang menggunakan terpal khusus kedap gas berstandar tinggi."},
            {"title": "Silo & Bulk Fumigation", "desc": "Pengendalian hama kutu gudang pada komoditas biji-bijian volume besar dalam silo tertutup."},
            {"title": "Aerasi & Clearance Test", "desc": "Pengukuran sisa residu gas dengan detektor gas akurat sebelum kontainer diberangkatkan."}
        ],
        "benefits": [
            "Teknisi fumigator bersertifikat resmi Badan Karantina Pertanian",
            "Penerbitan dokumen sertifikat fumigasi cepat dan valid",
            "Dukungan tim darurat di area pelabuhan Tanjung Perak Surabaya",
            "Jaminan mortalitas hama target mencapai 100%"
        ],
        "faq": [
            {"q": "Berapa lama proses fumigasi kontainer berlangsung?", "a": "Masa eksposur gas umumnya berlangsung antara 3 sampai 5 hari tergantung komoditas dan suhu, diikuti proses aerasi udara selama beberapa jam."},
            {"q": "Apakah sertifikat fumigasi diakui untuk ekspor?", "a": "Ya, sertifikat yang kami terbitkan memenuhi regulasi karantina internasional dan diterima oleh pembeli maupun otoritas bea cukai negara tujuan."}
        ]
    }'::jsonb,
    'Jasa Fumigasi Surabaya - Standar Karantina Ekspor & Gudang | JAS',
    'Layanan fumigasi fosfin (PH3) profesional di Surabaya untuk kontainer ekspor, arsip, dan komoditas pangan bersertifikat karantina.',
    true,
    4
),
(
    'b0000000-0000-0000-0000-000000000005',
    'disinfeksi',
    'Disinfeksi & Sterilisasi',
    'Sparkles',
    'Sterilisasi ruangan kantor, pabrik, klinik, dan hunian dari virus, bakteri, serta jamur patogen menggunakan cold fogging ramah lingkungan.',
    '{
        "hero_badge": "Perlindungan Higienis",
        "headline": "Sterilisasi Ruangan dari Virus, Bakteri, dan Jamur",
        "tagline": "Penyemprotan disinfektan medical-grade dengan Ultra Low Volume (ULV) cold fogger tanpa merusak alat elektronik.",
        "problems": [
            "Penyebaran penyakit musiman di lingkungan kantor dan sekolah.",
            "Bau apek dan spora jamur membandel pada ruangan ber-AC tertutup.",
            "Kekhawatiran residu kimia basah yang dapat merusak komputer atau arsip.",
            "Pembersihan manual permukaan saja tidak mampu menjangkau sirkulasi udara."
        ],
        "solutions": [
            "Teknologi ULV menghasilkan kabut kering mikro berukuran 5 sampai 20 mikron.",
            "Disinfektan berspektrum luas yang membunuh 99.9% patogen dan virus.",
            "Bahan non-korosif, tidak meninggalkan noda minyak, dan tidak beracun.",
            "Pengerjaan cepat dengan waktu tunggu minimal sebelum ruangan dapat dipakai."
        ],
        "methods": [
            {"title": "Dry Mist Cold Fogging", "desc": "Pengabutan kering ke seluruh sudut ruangan, plafon, dan celah sirkulasi pendingin ruangan."},
            {"title": "High-Touch Surface Wipe", "desc": "Penyekaan detail pada gagang pintu, sakelar, meja kerja, dan tombol lift dengan larutan disinfektan."},
            {"title": "Air Quality Neutralization", "desc": "Penyemprotan bahan penetralisir bau dan pemurni udara ramah lingkungan."},
            {"title": "Sanitation Certification", "desc": "Pemberian stiker bukti sterilisasi ruangan higienis untuk transparansi pengunjung."}
        ],
        "benefits": [
            "Bahan disinfektan bersertifikat Kemenkes dan aman untuk kontak makanan",
            "Tidak basah, tidak lengket, dan aman untuk perangkat elektronik kantor",
            "Aroma segar menyegarkan tanpa bau kimia menyengat",
            "Mendukung protokol audit kebersihan industri dan sertifikasi ISO"
        ],
        "faq": [
            {"q": "Apakah barang elektronik perlu ditutup plastik?", "a": "Karena menggunakan teknologi dry fogging mikron sangat halus, barang elektronik tidak akan basah. Namun dokumen kertas penting sebaiknya dimasukkan ke dalam laci."},
            {"q": "Kapan ruangan bisa ditempati kembali?", "a": "Cukup tunggu 30 menit setelah proses disinfeksi selesai dan sirkulasi udara dihidupkan."}
        ]
    }'::jsonb,
    'Jasa Disinfeksi Ruangan Surabaya - Steril & Higienis | JAS',
    'Jasa penyemprotan disinfektan ruangan di Surabaya untuk kantor, pabrik, klinik, dan rumah tinggal. Cepat, aman untuk elektronik, higienis.',
    true,
    5
)
ON CONFLICT (slug) DO UPDATE SET 
    name = EXCLUDED.name,
    icon = EXCLUDED.icon,
    short_description = EXCLUDED.short_description,
    content = EXCLUDED.content,
    seo_title = EXCLUDED.seo_title,
    seo_description = EXCLUDED.seo_description,
    is_published = EXCLUDED.is_published,
    sort_order = EXCLUDED.sort_order,
    updated_at = now();

-- 4. CATEGORIES (Blog)
INSERT INTO categories (id, name, slug)
VALUES
    ('c0000000-0000-0000-0000-000000000001', 'Tips & Pencegahan', 'tips-pencegahan'),
    ('c0000000-0000-0000-0000-000000000002', 'Sanitasi Bisnis', 'sanitasi-bisnis'),
    ('c0000000-0000-0000-0000-000000000003', 'Standar & Regulasi', 'standar-regulasi')
ON CONFLICT (slug) DO NOTHING;

-- 5. POSTS (Blog Sample Posts with Rich Content)
INSERT INTO posts (id, category_id, author_id, slug, title, excerpt, content_html, status, seo_title, seo_description, published_at)
VALUES
(
    'd0000000-0000-0000-0000-000000000001',
    'c0000000-0000-0000-0000-000000000001',
    'a0000000-0000-0000-0000-000000000001',
    'cara-mencegah-serangan-rayap-di-musim-hujan',
    'Cara Ampuh Mencegah Serangan Rayap di Musim Hujan',
    'Musim hujan meningkatkan kelembapan tanah yang memicu laron keluar mencari sarang baru. Pelajari cara melindungi hunian Anda sedini mungkin.',
    '<h2>Mengapa Rayap Begitu Aktif Saat Musim Hujan?</h2><p>Saat musim hujan tiba, tanah menjadi lembap dan suhu lingkungan menjadi ideal bagi koloni rayap untuk berkembang biak. Fenomena laron yang sering mengerubungi lampu rumah sebenarnya adalah kasta reproduksi rayap yang sedang mencari pasangan untuk membangun koloni baru di hunian Anda.</p><h3>Tanda Awal Kehadiran Rayap</h3><ul><li>Adanya sayap laron berserakan di dekat pintu dan jendela.</li><li>Muncul jalur terowongan tanah cokelat di sepanjang dinding atau kusen.</li><li>Kayu kusen atau perabot terdengar kopong saat diketuk perlahan.</li></ul><p>Langkah pencegahan terbaik adalah memastikan tidak ada rembesan air di sekitar dinding pondasi serta melakukan inspeksi rutin pada area furnitur kayu yang menempel langsung ke dinding.</p>',
    'published',
    'Cara Mencegah Serangan Rayap di Musim Hujan | JAS',
    'Tips praktis melindungi rumah dari bahaya rayap saat musim hujan. Kenali tanda awal rayap dan solusi proteksi terbaik.',
    now()
),
(
    'd0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000002',
    'a0000000-0000-0000-0000-000000000001',
    'bahaya-tikus-di-gudang-bisnis-dan-cara-mengatasinya',
    'Bahaya Tikus di Gudang Bisnis dan Cara Mengatasinya',
    'Tikus bukan hanya merusak barang dagangan, tetapi juga menimbulkan ancaman kebakaran dan penyakit. Simak strategi pengendalian terpadu.',
    '<h2>Dampak Serius Hama Tikus pada Bisnis dan Pergudangan</h2><p>Bagi pemilik fasilitas pergudangan dan pabrik, kehadiran tikus adalah ancaman kerugian finansial yang nyata. Selain mengerat kemasan kardus dan produk bernilai tinggi, tikus kerap menggerogoti kabel listrik yang berpotensi memicu korsleting listrik dan bahaya kebakaran besar.</p><h3>Metode Pengendalian Tikus Terpadu (Rodent Management)</h3><p>Pendekatan modern tidak sekadar memasang racun sembarangan. Metode Integrated Pest Management (IPM) mengedepankan kombinasi sanitasi lingkungan, penutupan celah masuk (proofing), serta penempatan bait station higienis.</p><blockquote>Keberhasilan pembasmian tikus terletak pada konsistensi monitoring perimeter luar sebelum hama sempat masuk ke area penyimpanan utama.</blockquote>',
    'published',
    'Bahaya Tikus di Gudang Bisnis & Solusinya | JAS',
    'Strategi efektif membasmi hama tikus di gudang dan pabrik tanpa bau bangkai. Lindungi aset usaha Anda bersama teknisi JAS.',
    now()
),
(
    'd0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000003',
    'a0000000-0000-0000-0000-000000000001',
    'standar-fumigasi-kontainer-ekspor-bebas-penolakan',
    'Standar Fumigasi Kontainer Ekspor agar Bebas Penolakan',
    'Memahami regulasi karantina internasional ISPM 15 dan prosedur perlakuan fumigasi resmi untuk kelancaran arus barang ekspor.',
    '<h2>Pentingnya Sertifikat Fumigasi Resmi untuk Eksportir</h2><p>Setiap negara tujuan ekspor memiliki standar karantina ketat untuk mencegah masuknya organisme pengganggu tumbuhan karantina (OPTK). Ketidaksesuaian perlakuan fumigasi dapat berakibat fatal: kontainer ditolak, dikembalikan, atau bahkan dimusnahkan di pelabuhan tujuan dengan beban biaya yang sangat tinggi pada eksportir.</p><h3>Prosedur Standar Fumigasi Komoditas</h3><p>Fumigasi yang benar mensyaratkan perhitungan dosis gas fosfin atau methyl bromide secara cermat, kerapatan sungkup terpal standar karantina, pemantauan konsentrasi gas selama masa retensi, hingga proses aerasi tuntas sebelum penerbitan sertifikat resmi.</p>',
    'published',
    'Standar Fumigasi Kontainer Ekspor | Jaya Abadi Solusi',
    'Panduan lengkap prosedur fumigasi kontainer ekspor sesuai standar karantina internasional. Pastikan pengiriman kargo Anda aman.',
    now()
)
ON CONFLICT (slug) DO NOTHING;

-- 6. PROJECTS (Sample Portofolio)
INSERT INTO projects (id, slug, title, location, description, completed_at, is_published)
VALUES
(
    'e0000000-0000-0000-0000-000000000001',
    'treatment-rayap-pergudangan-rungkut',
    'Proteksi Rayap Bangunan Pergudangan Logistik Rungkut',
    'Kawasan Industri Rungkut (SIER), Surabaya',
    'Pengerjaan soil treatment injeksi pondasi keliling sepanjang 450 meter dan wood treatment pada area pallet kayu gudang arsip seluas 3.500 m2. Bangunan terlindungi dengan garansi resmi 5 tahun.',
    '2026-08-15',
    true
),
(
    'e0000000-0000-0000-0000-000000000002',
    'pengendalian-hama-terpadu-restoran-surabaya-barat',
    'Integrated Pest Management Jaringan Restoran Cepat Saji',
    'Pakuwon City, Surabaya Barat',
    'Penerapan kontrak sanitasi rutin berkala untuk 4 gerai resto, meliputi penanganan kecoa jerman di dapur dengan baiting gel, monitoring lalat dengan fly catcher UV, dan pencegahan tikus.',
    '2026-09-02',
    true
),
(
    'e0000000-0000-0000-0000-000000000003',
    'fumigasi-kontainer-ekspor-tanjung-perak',
    'Fumigasi 12 Kontainer Komoditas Ekspor Kayu Olahan',
    'Depo Pelabuhan Tanjung Perak, Surabaya',
    'Perlakuan fumigasi fosfin pada 12 unit kontainer 40 feet muatan kayu olahan tujuan Eropa. Proses aerasi aman, lolos uji residu, dan sertifikat resmi diterbitkan tepat waktu.',
    '2026-10-10',
    true
)
ON CONFLICT (slug) DO NOTHING;

-- 7. PAGES (Static Pages Block Data)
-- Notice: Hero photo explicitly features a professional female specialist / customer service representative as requested!
INSERT INTO pages (key, title, content, seo_title, seo_description)
VALUES
('home', 'Beranda', '{
    "hero": {
        "badge": "Terpercaya Sejak 2010",
        "display_heading": "Solusi Tuntas Pengendalian Hama & Sanitasi Higienis",
        "paragraph": "Kami melindungi hunian, kantor, restoran, dan pergudangan Anda dari ancaman tikus, rayap, serangga, dan kuman patogen dengan metode profesional bergaransi di Surabaya.",
        "cta_primary_text": "Minta Penawaran Gratis",
        "cta_secondary_text": "Chat via WhatsApp",
        "cs_label": "Konsultasi Cepat",
        "cs_phone": "0813 9948 4916",
        "cs_phone_raw": "081399484916",
        "hero_image_url": "https://images.unsplash.com/photo-1573496359142-b8d87734a5a2?auto=format&fit=crop&w=800&q=80",
        "hero_image_alt": "Customer Service dan Konsultan Pengendalian Hama Profesional Jaya Abadi Solusi Surabaya"
    },
    "stats": [
        {"value": "15+", "label": "Tahun Pengalaman", "desc": "Melayani ribuan klien di Jawa Timur sejak 2010"},
        {"value": "2.500+", "label": "Properti Tertangani", "desc": "Hunian pribadi hingga kawasan industri skala besar"},
        {"value": "100%", "label": "Teknisi Tersertifikasi", "desc": "Pelatihan resmi standar Kemenkes dan Karantina"},
        {"value": "24 Jam", "label": "Respon Konsultasi", "desc": "Siap membantu kebutuhan inspeksi darurat Anda"}
    ],
    "workflows": [
        {"step": "01", "title": "Konsultasi & Survei", "desc": "Hubungi tim kami via telepon atau WhatsApp. Kami jadwalkan survei lokasi untuk menganalisis tingkat infestasi hama."},
        {"step": "02", "title": "Rencana & Treatment", "desc": "Kami menyusun penawaran transparan dan mengeksekusi treatment dengan metode ramah lingkungan berstandar tinggi."},
        {"step": "03", "title": "Garansi & Monitoring", "desc": "Pekerjaan tuntas disertai laporan inspeksi resmi, jaminan garansi penanganan ulang, dan jadwal kontrol berkala."}
    ],
    "property_types": [
        {"title": "Rumah Tinggal & Apartemen", "desc": "Perlindungan kenyamanan dan kesehatan keluarga dari serbuan kecoa, nyamuk, semut, tikus, dan rayap perabot."},
        {"title": "Restoran, Kafe & Kuliner", "desc": "Sanitasi dapur ketat sesuai standar dinas kesehatan tanpa mengganggu operasional jam saji tamu Anda."},
        {"title": "Gudang & Fasilitas Logistik", "desc": "Pencegahan kerusakan kemasan kardus, stok bahan baku, dan kabel mesin dari gigitan tikus dan rayap tanah."},
        {"title": "Gedung Perkantoran & Hotel", "desc": "Layanan berkala menjaga reputasi profesional properti komersial Anda tetap bersih, steril, dan bebas komplain tamu."}
    ]
}'::jsonb, 'Jaya Abadi Solusi - Jasa Pengendalian Hama Surabaya Terpercaya', 'Jasa pembasmi hama profesional di Surabaya: pest control, tikus, rayap, fumigasi dan disinfeksi bergaransi. Hubungi 0813 9948 4916.'),

('about-us', 'Tentang Kami', '{
    "hero": {
        "title": "Tentang Jaya Abadi Solusi",
        "subtitle": "Mitra Terpercaya Pengendalian Hama dan Sanitasi Lingkungan Sejak 2010"
    },
    "profile": {
        "story": "PT Jaya Abadi Solusi (JAS) berdiri di Surabaya dengan satu komitmen mendasar: memberikan perlindungan menyeluruh bagi properti dan kesehatan masyarakat dari gangguan hama. Berawal dari layanan pest control skala residensial, kini JAS telah dipercaya oleh berbagai perusahaan logistik, industri manufaktur, hotel, restoran, dan ribuan rumah tangga di Jawa Timur.",
        "mission": [
            "Menyediakan solusi pengendalian hama yang tuntas, aman, dan bertanggung jawab terhadap kelestarian lingkungan.",
            "Mengembangkan sumber daya teknisi profesional yang memiliki sertifikasi kompetensi resmi berkelanjutan.",
            "Menjaga standar kualitas pengerjaan dengan peralatan termutakhir dan respon cepat 24 jam bagi pelanggan."
        ],
        "vision": "Menjadi perusahaan pengendalian hama dan sanitasi higienis terdepan di Indonesia yang dikenal karena keandalan teknis, kepuasan pelanggan, dan integritas kerja tinggi."
    },
    "legalities": [
        {"name": "Izin Operasional Pengendalian Hama", "number": "No. 503/442.1/IPPH/Surabaya", "issuer": "Dinas Kesehatan Kota Surabaya"},
        {"name": "Sertifikat Standar Fumigasi Karantina", "number": "Reg. Barantan No. ID-0042-MB/PH3", "issuer": "Badan Karantina Pertanian"},
        {"name": "Sertifikasi Asosiasi ASPHAMI", "number": "Anggota Resmi DPD Jawa Timur", "issuer": "Asosiasi Perusahaan Pengendalian Hama Indonesia"}
    ]
}'::jsonb, 'Tentang Kami - Profil PT Jaya Abadi Solusi Surabaya', 'Profil perusahaan, visi misi, dan legalitas resmi PT Jaya Abadi Solusi, penyedia jasa pengendalian hama terpercaya di Surabaya sejak 2010.'),

('contact', 'Kontak Kami', '{
    "hero": {
        "title": "Hubungi Tim Ahli Hama Kami",
        "subtitle": "Konsultasikan keluhan hama Anda sekarang. Tim kami siap melakukan survei lokasi dan memberikan estimasi penawaran terbaik."
    },
    "direct_contacts": {
        "phone": "0813 9948 4916",
        "phone_raw": "081399484916",
        "whatsapp": "0813 9948 4916",
        "whatsapp_link": "https://wa.me/6281399484916?text=Halo%20Jaya%20Abadi%20Solusi%2C%20saya%20ingin%20konsultasi%20layanan%20hama",
        "email": "altertech.id@gmail.com",
        "office_address": "Jl. Rungkut Industri Raya No. 45, Surabaya, Jawa Timur 60293",
        "working_hours": "Senin - Sabtu: 08.00 - 17.00 WIB (Layanan Darurat 24 Jam)"
    }
}'::jsonb, 'Kontak Jaya Abadi Solusi Surabaya - Telepon & WhatsApp 0813 9948 4916', 'Hubungi PT Jaya Abadi Solusi di Surabaya. Konsultasi gratis via WhatsApp 0813 9948 4916 atau email altertech.id@gmail.com.')
ON CONFLICT (key) DO UPDATE SET 
    title = EXCLUDED.title,
    content = EXCLUDED.content,
    seo_title = EXCLUDED.seo_title,
    seo_description = EXCLUDED.seo_description,
    updated_at = now();
