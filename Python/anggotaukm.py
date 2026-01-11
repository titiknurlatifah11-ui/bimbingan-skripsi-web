selesai = False
anggota = []

while True:
    print("program pencatatan anggota ukm")
    print("1. cetak daftar anggota")
    print("2. tambah anggota")
    print("3. hapus anggota")
    print("4. keluar")

    pilihan = input("masukkan pilihan: ")

    if pilihan == "1":
        print("daftar anggota:")
        for nama in anggota:
            print(nama)

    elif pilihan == "2":
        nama = input("masukkan nama: ")
        anggota.append(nama)
        print("Anggota berhasil ditambahkan")

    elif pilihan == "3":
        nama = input("masukkan nama yang ingin dihapus: ")
        index = anggota.index(nama)
        del anggota[index]
        # del anggota[anggota.index(nama)]

    elif pilihan == "4":
        break #keluar dari while True