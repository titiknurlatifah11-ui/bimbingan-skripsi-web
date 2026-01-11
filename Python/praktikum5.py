# warna = ["merah", "hitam", "kuning", "biru"]

# #mengakses warna index ke 0
# print(warna[0])

# #mengakses warna index ke-1  sampai sebelum index ke-3
# print(warna[1:3])

# a = [1,2,3]
# b = [True, False, a]

# print(b)

# kosong = []
# print(kosong)

#Akses List
def cetak_daftar(daftar):
     for nama in daftar:
        # if nama == "Titik":
        if "i" in nama:
          print("Nama saya: ", nama)

daftar_nama = ["Titik", "eha", "vina"]
daftar_nama.append("natan")

del daftar_nama[0]

# print(daftar_nama.index("natan"))

# print(daftar_nama)

def cek_status(nama, daftar):
    if nama in daftar:
        print(nama, "ada didalam list")
    else:
        print(nama, "tidak ada didalam list")

user_input = input("masukkan nama: ")

cek_status(user_input, daftar_nama)

def cetak_semua(daftar):
    for item in daftar:
        print(item)






