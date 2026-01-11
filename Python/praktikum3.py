# def printquotes():
#     print("may your journey be overflowing with curses and blessings")

# printquotes()

# def penjumlahan():
#     userinput1 = int(input("angka 1 :"))
#     userinput2 = int(input("angka 2 :"))
#     hasil = userinput1 + userinput2
#     print(f"hasil = {hasil}")
    
# penjumlahan()

# def printNama(namadepan = "Titik", namabelakang = "Titik"):
#     print("nama depan saya adalah : ", namadepan, "nama belakang saya adalah :", namabelakang)

# # printNama("Titik")
# printNama()

# def testing(a, b ):
#     print("ini adalah :", a, "ini adalah:", b)

# testing("aku", "saya")

# def perhitungan2bilangan(bil1, bil2):
#    print( bil1 + bil2 )

# perhitungan2bilangan(10, 30)


# def perhitungan2bilangan(bil1, bil2):
#    a = bil1 + bil2
#    print(a)
# perhitungan2bilangan(10, 30)

# def returnhasil2bilangan(bil1, bil2):
#    return bil1 + bil2
# print(returnhasil2bilangan(10, 40))

# 

# 
# import math
# def mencariluaslingkaran():
#     radius = float(input('masukkan radius lingkaran : '))
#     luas = math.pi * radius**2
#     print(luas)

# mencariluaslingkaran()


# def penjumlahan3bil():
#      a = int(input("masukkan bil pertama : "))
#      b = int(input("masukkan bil kedua :"))
#      c = int(input("masukkan bil ketiga:"))
#      print(a+b+c)

# def pengurangan3bil():
#      a = int(input("masukkan bil pertama : "))
#      b = int(input("masukkan bil kedua :"))
#      c = int(input("masukkan bil ketiga:"))
#      print(a-b-c)

#      print("menu")
#      print("1. penjumlahan")
#      print("2. pengurangan")
# userinput = input("pilih menu : ")
# if userinput =='1':
#      penjumlahan3bil()
# elif userinput =='2':
#      pengurangan3bil()

data_bank = []

jumlah_data = int(input("Masukkan jumlah data bank yang ingin diinput: "))

for i in range(jumlah_data):
    nama_bank = input("Masukkan Nama Bank ke-{}: ".format(i+1))
    kode_bank = input("Masukkan Kode Bank ke-{}: ".format(i+1))
    variabel_bank = input("Masukkan Variabel Bank ke-{}: ".format(i+1))

    # Membuat dictionary untuk menyimpan data bank
    bank = {
        'Nama Bank': nama_bank,
        'Kode Bank': kode_bank,
        'Variabel Bank': variabel_bank
    }
# Menambahkan data bank ke dalam list data_bank
    data_bank.append(bank)

# Menampilkan hasil data_bank
print("\nData Bank yang telah diinput:")
for bank in data_bank:
    print("Nama Bank: {BRI}, Kode Bank: {001}, Variabel Bank: {5T}".format(bank['BNI'], bank['002'], bank['4T']))
