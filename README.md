
# Virtual ATM

Virtual ATM adalah aplikasi ATM sederhana berbasis CLI yang memungkinkan pengguna melakukan operasi seperti cek saldo, top up, tarik tunai, dan mengubah PIN. Aplikasi ini terhubung ke database PostgreSQL untuk menyimpan data pengguna dan transaksi.

## Fitur

- **Cek Saldo**: Lihat saldo akun.
- **Top Up**: Tambahkan saldo ke akun.
- **Tarik Tunai**: Tarik saldo dari akun.
- **Ubah PIN**: Ubah PIN akun.
- **Catatan Transaksi**: Semua transaksi (top up dan tarik tunai) tercatat di database.

## Persyaratan

- Go 1.16 atau lebih baru
- PostgreSQL
- Go library: `github.com/lib/pq`

## Instalasi

### 1. Clone Repository

Clone repository ini atau salin kode ke direktori lokal.

```bash
git clone https://github.com/username/repo.git
cd repo
```

### 2. Instal Dependency

Instal library Go untuk PostgreSQL dengan perintah berikut:

```bash
go get github.com/lib/pq
```

### 3. Buat Database di PostgreSQL

Jalankan perintah SQL berikut untuk membuat tabel `users` dan `transactions`:

```sql
CREATE TABLE users (
    username VARCHAR(50) PRIMARY KEY,
    pin VARCHAR(6),
    saldo INT
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    amount INT,
    transaction_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 4. Set Environment Variable

Set environment variable `DATABASE_URL` dengan connection string PostgreSQL kamu.

**Linux/MacOS:**
```bash
export DATABASE_URL="postgresql://<user>:<password>@<host>/<dbname>?sslmode=require"
```

**Windows (PowerShell):**
```powershell
$env:DATABASE_URL = "postgresql://<user>:<password>@<host>/<dbname>?sslmode=require"
```

### 5. Jalankan Aplikasi

Jalankan aplikasi dengan perintah:

```bash
go run main.go
```

## Penggunaan

1. **Login**: Masukkan username dan PIN yang valid untuk mengakses menu utama.
2. **Menu Utama**:
   - `1`: Cek Saldo.
   - `2`: Top Up.
   - `3`: Tarik Tunai.
   - `4`: Ubah PIN.
   - `5`: Keluar dari aplikasi.
3. **Transaksi**: Setiap kali melakukan top up atau tarik tunai, saldo pengguna akan diperbarui dan transaksi tercatat di database.

## Struktur File

- `atm.go`: File utama yang berisi logika aplikasi Virtual ATM.

## Masalah Umum

- **Error: `DATABASE_URL` environment variable not set**  
  Pastikan `DATABASE_URL` sudah di-set dengan benar.
  
- **Error connecting to database**  
  Cek connection string PostgreSQL dan pastikan server database sudah berjalan.
