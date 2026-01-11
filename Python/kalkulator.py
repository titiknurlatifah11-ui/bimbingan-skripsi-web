# Fungsi untuk penambahan
def tambah(x, y):
    return x + y

# Fungsi untuk pengurangan
def kurang(x, y):
    return x - y

# Fungsi untuk perkalian
def kali(x, y):
    return x * y

# Fungsi untuk pembagian
def bagi(x, y):
    if y == 0:
        return "Error: Tidak bisa membagi dengan nol"
    return x / y

print("Pilih operasi:")
print("1. Penambahan")
print("2. Pengurangan")
print("3. Perkalian")
print("4. Pembagian")

# Meminta input dari pengguna
pilihan = input("Masukkan nomor operasi (1/2/3/4): ")

angka1 = float(input("Masukkan angka pertama: "))
angka2 = float(input("Masukkan angka kedua: "))

if pilihan == '1':
    print(angka1, "+", angka2, "=", tambah(angka1, angka2))
elif pilihan == '2':
    print(angka1, "-", angka2, "=", kurang(angka1, angka2))
elif pilihan == '3':
    print(angka1, "*", angka2, "=", kali(angka1, angka2))
elif pilihan == '4':
    print(angka1, "/", angka2, "=", bagi(angka1, angka2))
else:
    print("Pilihan operasi tidak valid")
