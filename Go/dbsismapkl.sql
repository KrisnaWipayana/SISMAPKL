-- Tabel dospem
create table dospem (
    id uuid primary key not null,
    nama varchar(100) not null,
    nip varchar(30) not null,
    password varchar(255) not null,
    id_mahasiswa uuid, -- Tidak ada foreign key pada tahap ini
    created_at TIMESTAMPTZ default current_timestamp,
    updated_at TIMESTAMPTZ default current_timestamp,
    deleted_at TIMESTAMPTZ
);

-- Tabel perusahaan
create table perusahaan (
    id uuid primary key not null,
    nama varchar(100) not null, 
    koordinat varchar(100) not null,
    alamat varchar(200) not null,
    kabupaten varchar(100) not null,
    provinsi varchar(100) not null,
    id_mahasiswa uuid, -- Tidak ada foreign key pada tahap ini
    created_at TIMESTAMPTZ default current_timestamp,
    updated_at TIMESTAMPTZ default current_timestamp,
    deleted_at TIMESTAMPTZ
);

-- Tabel mahasiswa
create table mahasiswa (
    id uuid primary key not null,
    nim varchar(15) not null,
    password varchar(255) not null,
    nama varchar(100) not null,
    kelas varchar(10) not null,
    prodi varchar(100) not null,
    jurusan varchar(100) not null,
    id_dospem uuid, -- Tidak ada foreign key pada tahap ini
    id_perusahaan uuid, -- Tidak ada foreign key pada tahap ini
    created_at TIMESTAMPTZ default current_timestamp,
    updated_at TIMESTAMPTZ default current_timestamp,
    deleted_at TIMESTAMPTZ
);

-- Tabel mentor
create table mentor (
    id uuid primary key not null,
    nama varchar(100) not null,
    email varchar(100) not null,
    password varchar(255) not null,
    id_perusahaan uuid not null, -- Tidak ada foreign key pada tahap ini
    created_at TIMESTAMPTZ default current_timestamp,
    updated_at TIMESTAMPTZ default current_timestamp,
    deleted_at TIMESTAMPTZ
);

-- Tabel berkas
create table berkas (
    id uuid primary key not null,
    nama varchar(100) not null,
    status varchar(20) not null,
    id_mahasiswa uuid not null, -- Tidak ada foreign key pada tahap ini
    created_at TIMESTAMPTZ default current_timestamp,
    updated_at TIMESTAMPTZ default current_timestamp,
    deleted_at TIMESTAMPTZ
);

-- Tabel laporan
create table laporan (
    id uuid primary key not null,
    laporan varchar(100) not null,
    status_dospem varchar(20) not null,
    status_mentor varchar(20) not null,
    id_mahasiswa uuid not null, -- Tidak ada foreign key pada tahap ini
    created_at TIMESTAMPTZ default current_timestamp,
    updated_at TIMESTAMPTZ default current_timestamp,
    deleted_at TIMESTAMPTZ
);
-- End database

-- Tambahkan foreign key setelah tabel dibuat

-- Tambahkan foreign key di tabel dospem
alter table dospem
add constraint fk_dospem_mahasiswa foreign key (id_mahasiswa) references mahasiswa(id) on delete set null;

-- Tambahkan foreign key di tabel perusahaan
alter table perusahaan
add constraint fk_perusahaan_mahasiswa foreign key (id_mahasiswa) references mahasiswa(id) on delete cascade;

-- Tambahkan foreign key di tabel mahasiswa
alter table mahasiswa
add constraint fk_mahasiswa_dospem foreign key (id_dospem) references dospem(id) on delete set null;
alter table mahasiswa
add constraint fk_mahasiswa_perusahaan foreign key (id_perusahaan) references perusahaan(id) on delete cascade;

-- Tambahkan foreign key di tabel mentor
alter table mentor
add constraint fk_mentor_perusahaan foreign key (id_perusahaan) references perusahaan(id) on delete cascade;

-- Tambahkan foreign key di tabel berkas
alter table berkas
add constraint fk_berkas_mahasiswa foreign key (id_mahasiswa) references mahasiswa(id) on delete cascade;

-- Tambahkan foreign key di tabel laporan
alter table laporan
add constraint fk_laporan_mahasiswa foreign key (id_mahasiswa) references mahasiswa(id) on delete cascade;
