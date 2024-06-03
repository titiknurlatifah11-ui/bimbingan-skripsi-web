import sqlite3

conn = sqlite3.connect('dosen.db')
print("database berhasil dibuat")


conn.execute('''CREATE TABLE IF NOT EXISTS DOSEN_INFORMATIKA
            (NIDN TEXT PRIMARY KEY,
            NAMA TEXT NOT NULL,
            MATKUL TEXT NOT NULL);''')
conn.commit()
print("Table berhasil dibuat")

# Menambah data dengan input secara berulang dengan batasan nya di inputkan oleh user
def tambah_dosen(conn):
    print("\n===== Tambah Dosen =====")
    banyak_data = int(input("Masukkan banyak data yang ingin di inputkan: "))
    for i in range(banyak_data):
        NIDN = input("Masukkan NIDN: ")
        NAMA = input("Masukkan NAMA: ").lower()
        MATKUL = input("Masukkan MATKUL: ").lower()
        if NIDN == "" or NAMA == "" or MATKUL:
            print("Data tidak boleh kosong")
        conn.execute("INSERT INTO DOSEN INFORMATIKA (NIDN, NAMA, MATKUL) VALUES (?, ?, ?)", (NIDN, NAMA, MATKUL))
        conn.commit()
        print("Data berhasil di inputkan")
        
# Menampilkan data
def tampilkan_dosen():
    cursor = conn.execute("SELECT * FROM DOSEN INFORMATIKA")
    for row in cursor:
        print("NIDN :", row[0])
        print("NAMA :", row[1])
        print("MATKUL :", row[2])
        print("====================================")

# Update data
def update_dosen(conn):
    print("\n===== Update Karakter =====")
    NIDN = input("Masukkan NIDN yang ingin di update: ")
    NAMA = input("Masukkan NAMA  dosen yang ingin di update: ").lower()
    MATKUL = input("Masukkan MATKUL dosen: ").lower()
    if NIDN == "" or NAMA == "" or MATKUL:
        print("Data tidak boleh kosong")
        

    conn.execute("UPDATE DOSEN INFORMATIKA SET NAMA = ?, MATKUL = ?", (NIDN, NAMA, MATKUL))
    conn.commit()
    print("\nData berhasil di update")

 # Hapus data
def hapus_dosen(conn):
     print("\n===== Hapus Dosen =====")
     id = input("Masukkan Nidn yang ingin di hapus: ")
     conn.execute("DELETE FROM DOSEN INFORMATIKA WHERE NIDN = ?", (NIDN))
     conn.commit()
     print("\nData berhasil di hapus")

# Cari data
def cari_dosen(conn):
    print("\n===== Cari Dosen =====")
    nama = input("Masukkan nama dosen yang ingin di cari: ").lower()
    cursor = conn.execute("SELECT * FROM DOSEN INFORMATIKA WHERE NAMA = ?", (nama,))
    print("\nHasil pencarian: ", nama)
    for row in cursor:
        print("NIDN :", row[0])
        print("NAMA :", row[1])
        print("MATKUL :", row[2])
        print("====================================")

    
# Menu
def menu():
    while True:
        print("\n===== Selamat Datang Di DATA DOSEN =====")
        print("1. Tambah dosen")
        print("2. Tampilkan dosen")
        print("3. Update dosen")
        print("4. Hapus dosen")
        print("5. Cari dosen")
        print("6. Keluar")
        pilihan = input("Masukkan pilihan: ")
        if pilihan == "1":
            tambah_dosen()
        elif pilihan == "2":
            tampilkan_dosen()
        elif pilihan == "3":
            update_dosen()
        elif pilihan == "4":
            hapus_dosen()
        elif pilihan == "5":
            cari_dosen()
        elif pilihan == "6":
            print("Terima kasih mengakses data dosen")
            conn.close()
            break
        else:
            print("Pilihan tidak tersedia")
        
menu()