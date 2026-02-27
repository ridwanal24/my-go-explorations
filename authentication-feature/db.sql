-- 1. Hapus semua tabel yang ada (Manual DROP berdasarkan urutan)
-- Catatan: SQL tidak memiliki satu perintah "DROP ALL", jadi harus sebutkan tabelnya
-- atau hapus dan buat ulang database-nya.
DROP TABLE IF EXISTS users;

-- 2. Buat tabel user
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- 3. Masukkan 3 data dummy
INSERT INTO users (name, email, password) VALUES 
('Budi Santoso', 'budi@example.com', 'hashed_password_123'),
('Siti Aminah', 'siti@example.com', 'hashed_password_456'),
('Rian Wijaya', 'rian@example.com', 'hashed_password_789');
