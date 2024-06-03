import sqlite3

# Koneksi ke database
conn = sqlite3.connect('data_kampus.db')
cursor = conn.cursor()
 

# Buat tabel users jika belum ada
cursor.execute('''CREATE TABLE IF NOT EXISTS users (username text, password text)''')


# Buat tabel mahasiswa dan mata_kuliah jika belum ada
cursor.execute('''CREATE TABLE IF NOT EXISTS mahasiswa (
  nim text PRIMARY KEY,
  nama text,
  jurusan text
)''')
cursor.execute('''CREATE TABLE IF NOT EXISTS mata_kuliah (
  kode text PRIMARY KEY,
  nama text,
  sks int
)''')
cursor.execute('''CREATE TABLE IF NOT EXISTS nilai (
  nim text,
  kode text,
  nilai real,
  PRIMARY KEY (nim, kode),
  FOREIGN KEY (nim) REFERENCES mahasiswa (nim),
  FOREIGN KEY (kode) REFERENCES mata_kuliah (kode)
)''')

cursor.execute('''CREATE TABLE IF NOT EXISTS DOSEN
			(NIDN TEXT PRIMARY KEY,
			NAMA TEXT NOT NULL,
			MATKUL TEXT NOT NULL);''')

# Fungsi untuk membuat akun baru
def create_account():
  username = input('Masukkan username: ')
  password = input('Masukkan password: ')

  # Cek apakah username sudah ada di database
  cursor.execute('''SELECT * FROM users WHERE username=?''', (username,))
  user = cursor.fetchone()
  if user:
    print('Username sudah ada, silakan pilih username lain.')
  else:
    # Tambahkan akun baru ke database
    cursor.execute('''INSERT INTO users VALUES (?, ?)''', (username, password))
    conn.commit()
    print('Akun berhasil dibuat.')


# Fungsi untuk login sebagai admin
def login():
  username = input('Masukkan username: ')
  password = input('Masukkan password: ')

  # Cek apakah username dan password cocok dengan yang ada di database
  cursor.execute('''SELECT * FROM users WHERE username=? AND password=?''', (username, password))
  user = cursor.fetchone()
  if user:
      print('Log-in Sukses!')
      menu_data()
  else:
    print('Username atau password salah, silakan coba lagi.')


# Menu Login
def menu_login():
    while True:
        print('==========|| Student Center Learning - University kelompok 3 ||==========')
        print('                        // Selamat datang \\\      ')
        print('Silahkan buat akun jika belum punya akun,')
        print('1. Buat akun')
        print('2. Log in sebagai admin')
        print('3. Keluar')
        pilihan = input('Masukkan pilihan Anda (1/2/3): ')

        if pilihan == '1':
            create_account()
        elif pilihan == '2':
            login()
        elif pilihan == '3':
            break
        else:
            print('Pilihan tidak valid, silakan coba lagi.')


# Fungsi untuk menambahkan data mahasiswa baru
def tambah_mahasiswa():
  nim = input('Masukkan NIM: ')
  nama = input('Masukkan nama: ')
  jurusan = input('Masukkan jurusan: ')

  # Cek apakah NIM sudah ada di database
  cursor.execute('''SELECT * FROM mahasiswa WHERE nim=?''', (nim,))
  mahasiswa = cursor.fetchone()
  if mahasiswa:
    print('NIM sudah ada, silakan masukkan NIM lain.')
  else:
    # Tambahkan data mahasiswa baru ke database
    cursor.execute('''INSERT INTO mahasiswa VALUES (?, ?, ?)''', (nim, nama, jurusan))
    conn.commit()
    print('Data mahasiswa berhasil ditambahkan.')


# Fungsi untuk menambahkan data mata kuliah baru
def tambah_mata_kuliah():
  kode = input('Masukkan kode mata kuliah: ')
  nama = input('Masukkan nama mata kuliah: ')
  sks = int(input('Masukkan jumlah SKS: '))

  # Cek apakah kode sudah ada di database
  cursor.execute('''SELECT * FROM mata_kuliah WHERE kode=?''', (kode,))
  mata_kuliah = cursor.fetchone()
  if mata_kuliah:
    print('Kode sudah ada, silakan masukkan kode lain.')
  else:
    # Tambahkan data mata kuliah baru ke database
    cursor.execute('''INSERT INTO mata_kuliah VALUES (?, ?, ?)''', (kode, nama, sks))
    conn.commit()
    print('Data mata kuliah berhasil ditambahkan.')


# Fungsi untuk menampilkan data mahasiswa
def lihat_mahasiswa():
  # Ambil semua data mahasiswa dari database
  cursor.execute('''SELECT * FROM mahasiswa''')
  mahasiswa = cursor.fetchall()
  if mahasiswa:
    # Print header
    print('|{:^10s}|{:^30s}|{:^20s}|'.format('NIM', 'Nama', 'Jurusan'))
    print('-' * 62)
    # Print semua data mahasiswa
    for m in mahasiswa:
      print('|{:^10s}|{:^30s}|{:^20s}|'.format(m[0], m[1], m[2]))
  else:
    print('Tidak ada data mahasiswa.')


# Fungsi untuk menampilkan data mata kuliah
def lihat_mata_kuliah():
  # Ambil semua data mata kuliah dari database
  cursor.execute('''SELECT * FROM mata_kuliah''')
  mata_kuliah = cursor.fetchall()
  if mata_kuliah:
    # Print header
    print('|{:^10s}|{:^30s}|{:^5s}|'.format('Kode', 'Nama', 'SKS'))
    print('-' * 47)
    # Print semua data mata kuliah
    for mk in mata_kuliah:
      print('|{:^10s}|{:^30s}|{:^5d}|'.format(mk[0], mk[1], mk[2]))
  else:
    print('Tidak ada data mata kuliah.')


# Fungsi untuk mengedit data mahasiswa
def edit_mahasiswa():
  nim = input('Masukkan NIM mahasiswa yang akan diedit: ')

  # Ambil data mahasiswa yang akan diedit dari database
  cursor.execute('''SELECT * FROM mahasiswa WHERE nim=?''', (nim,))
  mahasiswa = cursor.fetchone()
  if mahasiswa:
    # Tampilkan data mahasiswa yang akan diedit
    print('Data mahasiswa sebelum diedit:')
    print('NIM:', mahasiswa[0])
    print('Nama:', mahasiswa[1])
    print('Jurusan:', mahasiswa[2])

    # Minta inputan baru dari pengguna
    nama = input('Masukkan nama baru (kosongkan jika tidak ingin mengubah): ')
    jurusan = input('Masukkan jurusan baru (kosongkan jika tidak ingin mengubah): ')

    # Update data mahasiswa di database
    if nama:
      cursor.execute('''UPDATE mahasiswa SET nama=? WHERE nim=?''', (nama, nim))
    if jurusan:
      cursor.execute('''UPDATE mahasiswa SET jurusan=? WHERE nim=?''', (jurusan, nim))
    conn.commit()
    print('Data mahasiswa berhasil diperbarui.')
  else:
    print('NIM tidak ditemukan.')


# Fungsi untuk mengedit data mata kuliah
def edit_mata_kuliah():
  kode = input('Masukkan kode mata kuliah yang akan diedit: ')

  # Ambil data mata kuliah yang akan diedit dari database
  cursor.execute('''SELECT * FROM mata_kuliah WHERE kode=?''', (kode,))
  mata_kuliah = cursor.fetchone()
  if mata_kuliah:
    # Tampilkan data mata kuliah yang akan diedit
    print('Data mata kuliah sebelum diedit:')
    print('Kode:', mata_kuliah[0])
    print('Nama:', mata_kuliah[1])
    print('SKS:', mata_kuliah[2])

    # Minta inputan baru dari pengguna
    nama = input('Masukkan nama baru (kosongkan jika tidak ingin mengubah): ')
    sks = input('Masukkan jumlah SKS baru (kosongkan jika tidak ingin mengubah): ')

    # Update data mata kuliah di database
    if nama:
      cursor.execute('''UPDATE mata_kuliah SET nama=? WHERE kode=?''', (nama, kode))
    if sks:
      cursor.execute('''UPDATE mata_kuliah SET sks=? WHERE kode=?''', (sks, kode))
    conn.commit()
    print('Data mata kuliah berhasil diperbarui.')
  else:
    print('Kode mata kuliah tidak ditemukan.')


# Fungsi untuk menghapus data mahasiswa
def hapus_mahasiswa():
  nim = input('Masukkan NIM mahasiswa yang akan dihapus: ')

  # Ambil data mahasiswa yang akan dihapus dari database
  cursor.execute('''SELECT * FROM mahasiswa WHERE nim=?''', (nim,))
  mahasiswa = cursor.fetchone()
  if mahasiswa:
    # Tampilkan data mahasiswa yang akan dihapus
    print('Data mahasiswa yang akan dihapus:')
    print('NIM:', mahasiswa[0])
    print('Nama:', mahasiswa[1])
    print('Jurusan:', mahasiswa[2])

    # Minta konfirmasi dari pengguna
    yakin = input('Apakah Anda yakin ingin menghapus data ini? (y/t) ')
    if yakin == 'y':
      # Hapus data mahasiswa dari database
      cursor.execute('''DELETE FROM mahasiswa WHERE nim=?''', (nim,))
      conn.commit()
      print('Data mahasiswa berhasil dihapus.')
    else:
      print('Data mahasiswa tidak dihapus.')
  else:
    print('NIM tidak ditemukan.')


# Fungsi untuk menghapus data mata kuliah
def hapus_mata_kuliah():
  kode = input('Masukkan kode mata kuliah yang akan dihapus: ')

  # Ambil data mata kuliah yang akan dihapus dari database
  cursor.execute('''SELECT * FROM mata_kuliah WHERE kode=?''', (kode,))
  mata_kuliah = cursor.fetchone()
  if mata_kuliah:
    # Tampilkan data mata kuliah yang akan dihapus
    print('Data mata kuliah yang akan dihapus:')
    print('Kode:', mata_kuliah[0])
    print('Nama:', mata_kuliah[1])
    print('SKS:', mata_kuliah[2])

    # Minta konfirmasi dari pengguna
    yakin = input('Apakah Anda yakin ingin menghapus data ini? (y/t) ')
    if yakin == 'y':
      # Hapus data mata kuliah dari database
      cursor.execute('''DELETE FROM mata_kuliah WHERE kode=?''', (kode,))
      conn.commit()
      print('Data mata kuliah berhasil dihapus.')
    else:
      print('Data mata kuliah tidak dihapus.')
  else:
    print('Kode mata kuliah tidak ditemukan.')


# Fungsi untuk menambah nilai
def tambah_nilai():
  # Minta inputan dari pengguna
  nim = input('Masukkan NIM mahasiswa: ')
  kode = input('Masukkan kode mata kuliah: ')
  kehadiran = float(input('Masukan Kehadiran: '))
  praktikum = float(input('Masukan Nilai Praktikum: '))
  tugas = float(input('Masukan Nilai Tugas: '))
  uts = float(input('Masukan Nilai UTS: '))
  uas = float(input('Masukan Nilai UAS: '))
  

  #komposisi nilai
  kehadiran *= 0.2
  praktikum *= 0.1
  tugas *= 0.2
  uts *= 0.25
  uas *= 0.25
  nilai = kehadiran + praktikum + tugas + uts + uas

  if nilai >= 80:
    print("Predikat: A")
  elif nilai >= 70 and nilai < 80:
    print("Predikat: B")
  elif nilai >= 60 and nilai < 70:
    print("Predikat: C")
  elif nilai >= 50 and nilai < 60:
    print("Predikat: D")
  elif nilai < 50:
     print("Predikat: E")
  else:
     print("Nilai Tidak Ditemukan")

  # Tambah data nilai ke database
  cursor.execute('''INSERT INTO nilai (nim, kode, nilai) VALUES (?, ?, ?)''', (nim, kode, nilai))
  conn.commit()
  print('Nilai berhasil ditambahkan.')


# Fungsi untuk melihat nilai
def lihat_nilai():
  # Minta inputan dari pengguna
  nim = input('Masukkan NIM mahasiswa: ')

  # Ambil data nilai dari database
  cursor.execute('''SELECT m.nama, n.kode, mk.nama, n.nilai
                    FROM nilai n
                    INNER JOIN mahasiswa m ON n.nim = m.nim
                    INNER JOIN mata_kuliah mk ON n.kode = mk.kode
                    WHERE n.nim=?''', (nim,))
  nilai = cursor.fetchall()

  if nilai:
    # Tampilkan data nilai
    print('Nilai mahasiswa:')
    for n in nilai:
      print('NIM:', nim)
      print('Nama:', n[0])
      print('Kode mata kuliah:', n[1])
      print('Nama mata kuliah:', n[2])
      print('Nilai:', n[3])
  else:
    print('NIM tidak ditemukan atau belum memiliki nilai.')


# Fungsi untuk mengedit nilai
def edit_nilai():
  # Minta inputan dari pengguna
  nim = input('Masukkan NIM mahasiswa: ')
  kode = input('Masukkan kode mata kuliah: ')

  # Ambil data nilai yang akan diedit dari database
  cursor.execute('''SELECT * FROM nilai WHERE nim=? AND kode=?''', (nim, kode))
  nilai = cursor.fetchone()
  if nilai:
    # Tampilkan data nilai yang akan diedit
    print('Data nilai yang akan diedit:')
    print('NIM:', nilai[0])
    print('Kode mata kuliah:', nilai[1])
    print('Nilai:', nilai[2])

    # Minta inputan baru dari pengguna
    kehadiran = float(input('Masukan Kehadiran: '))
    praktikum = float(input('Masukan Nilai Praktikum: '))
    tugas = float(input('Masukan Nilai Tugas: '))
    uts = float(input('Masukan Nilai UTS: '))
    uas = float(input('Masukan Nilai UAS: '))
     #komposisi nilai
    kehadiran *= 0.2
    praktikum *= 0.1
    tugas *= 0.2
    uts *= 0.25
    uas *= 0.25
    nilai_baru= kehadiran + praktikum + tugas + uts + uas

    if nilai_baru >= 80:
      print("Predikat: A")
    elif nilai_baru >= 70 and nilai < 80:
      print("Predikat: B")
    elif nilai_baru >= 60 and nilai < 70:
      print("Predikat: C")
    elif nilai_baru >= 50 and nilai < 60:
      print("Predikat: D")
    elif nilai_baru < 50:
     print("Predikat: E")
    else:
     print("Nilai Tidak Ditemukan")

    # Update data nilai di database
    cursor.execute('''UPDATE nilai SET nilai=? WHERE nim=? AND kode=?''', (nilai_baru, nim, kode))
    conn.commit()
    print('Nilai berhasil diedit.')
  else:
    print('NIM atau kode mata kuliah tidak ditemukan.')


# Fungsi untuk menghapus nilai
def hapus_nilai():
  # Minta inputan dari pengguna
  nim = input('Masukkan NIM mahasiswa: ')
  kode = input('Masukkan kode mata kuliah: ')

  # Ambil data nilai yang akan dihapus dari database
  cursor.execute('''SELECT * FROM nilai WHERE nim=? AND kode=?''', (nim, kode))
  nilai = cursor.fetchone()
  if nilai:
    # Tampilkan data nilai yang akan dihapus
    print('Data nilai yang akan dihapus:')
    print('NIM:', nilai[0])
    print('Kode mata kuliah:', nilai[1])
    print('Nilai:', nilai[2])

    # Minta konfirmasi dari pengguna
    yakin = input('Apakah Anda yakin ingin menghapus data ini? (y/t) ')
    if yakin == 'y':
      # Hapus data nilai dari database
      cursor.execute('''DELETE FROM nilai WHERE nim=? AND kode=?''', (nim, kode))
      conn.commit()
      print('Nilai berhasil dihapus.')
    else:
      print('Nilai tidak dihapus.')
  else:
    print('NIM atau kode mata kuliah tidak ditemukan.')

def buat_dosen():
	while True:
		NIDN = input("Masukkan NIDN dosen: ")
		NAMA = input("Masukkan nama dosen: ")
		MATKUL = input("Masukkan mata kuliah dosen: ")
		
		if NIDN == "" or NAMA == "" or MATKUL == "":
			print("Data tidak boleh kosong")
			continue
		
		conn.execute("INSERT INTO DOSEN (NIDN, NAMA, MATKUL) VALUES (?, ?, ?)", (NIDN, NAMA, MATKUL))
		conn.commit()
		break

def tampilkan_dosen():
	kumpulan_dosen = conn.execute("SELECT * FROM DOSEN")
	for dosen in kumpulan_dosen:
		print("NIDN :", dosen[0])
		print("Nama :", dosen[1])
		print("Mata Kuliah :", dosen[2])
		print("====================================")

def update_dosen():
	while True:
		NIDN = input("Masukkan NIDN dosen yang ingin di update: ")
		NAMA = input("Masukkan nama dosen: ")
		MATKUL = input("Masukkan mata kuliah dosen: ")

		if NIDN == "" or NAMA == "" or MATKUL == "":
			print("Data tidak boleh kosong")
			continue

		conn.execute("UPDATE DOSEN SET NAMA = ?, MATKUL = ? WHERE NIDN = ?", (NAMA, MATKUL, NIDN))
		conn.commit()
		break

def hapus_dosen():
	while True:
		NIDN = input("Masukkan NIDN dosen yang ingin di update: ")

		if NIDN == "":
			print("Data tidak boleh kosong")
			continue

		conn.execute("DELETE FROM DOSEN WHERE NIDN = ?", (NIDN,))
		conn.commit()
		break


def menu_c():
  while True:
    print('Silahkan anda pilih data yang ingin anda buat')
    print('1. mahasiswa')
    print('2. dosen')
    print('3. mata kuliah')
    print('4. kembali')
    pilihan = input("masukan data yang ingin anda buat (1,2,3,4): ")

    if pilihan == '1':
        tambah_mahasiswa()
    elif pilihan == '2':
        buat_dosen()
    elif pilihan == '3':
       tambah_mata_kuliah()
    elif pilihan == '4':
        break
    else:
       print("tidak tersedia")


def menu_r():
   while True:
    print('Silahkan anda pilih data yang ingin anda lihat')
    print('1. mahasiswa')
    print('2. mata kuliah')
    print('3. kembali')
    pilihan = input("masukan data yang ingin anda lihat (1,2,3): ")

    if pilihan == '1':
        lihat_mahasiswa()
    elif pilihan == '2':
        lihat_mata_kuliah()
    elif pilihan == '3':
        break
    else:
       print("tidak tersedia")

def menu_u():
   while True:
    print('Silahkan anda pilih data yang ingin anda edit')
    print('1. mahasiswa')
    print('2. mata kuliah')
    print('3. kembali')
    pilihan = input("masukan data yang ingin anda edit (1,2,3): ")

    if pilihan == '1':
        edit_mahasiswa()
    elif pilihan == '2':
        edit_mata_kuliah()
    elif pilihan == '3':
        break
    else:
       print("tidak tersedia")


def menu_d():
   while True:
    print('Silahkan anda pilih data yang ingin anda hapus')
    print('1. mahasiswa')
    print('2. mata kuliah')
    print('3. dosen')
    print('4. kembali')
    pilihan = input("masukan data yang ingin anda hapus (1/2/3): ")

    if pilihan == '1':
        hapus_mahasiswa()
    elif pilihan == '2':
        hapus_mata_kuliah()
    elif pilihan == '3':
        break
    else:
       print("tidak tersedia")


def menu_n():
    while True:
        print('Pilih menu yang ingin anda pilih')
        print('1. Tambah Nilai')
        print('2. Lihat Nilai')
        print('3. Update Nilai')
        print('4. Delete Nilai')
        print('5. Kembali ke menu data')
        pilihan = input("masukan data yang ingin anda buat (1/2/3/4/5): ")

        if pilihan == '1':
            tambah_nilai()
        elif pilihan == '2':
            lihat_nilai()
        elif pilihan == '3':
            edit_nilai()
        elif pilihan == '4':
            hapus_nilai()
        elif pilihan == '5':
            break
        else:
            print("tidak tersedia")
   

def menu_data():
  while True:
        print('        !!!Selamat datang di menu utama, admin!!!')
        print('=====|| Silahkan Pilih Menu Yang Ingin anda gunakan ||=====')
        print('C. Create / Buat / Tambah data mahasiswa dan mata kuliah / dosen')
        print('R. Read / Lihat data mahasiswa dan mata kuliah / dosen')
        print('U. Update / Rename data mahasiswa dan mata kuliah / dosen')
        print('D. Delete / Hapus data mahasiswa dan mata kuliah / dosen')
        print('N. Nilai Mata Kuliah')
        print('B. Back / Kembali ke menu log in')
        pilihan = input('Masukkan pilihan Anda (C/R/U/D/N/B): ')

        if pilihan == 'C':
            menu_c()
        elif pilihan == 'R':
            menu_r()
        elif pilihan == 'U':
            menu_u()
        elif pilihan == 'D':
            menu_d()
        elif pilihan == 'N':
            menu_n()
        elif pilihan == 'B':
            break
        else:
            print('Pilihan tidak valid, silakan coba lagi.')

#Menjalankan Program
print("""
        ------------------------------------------
        |  Creator: KELOMPOK 2                   |
        |  Jurusan: Teknik Informatika           |
        |  Tahun Angkatan: 2023                  |
        |  Institut Teknologi Indonesia          |
        ------------------------------------------
        """)
menu_login()
# Tutup koneksi ke database
conn.close()