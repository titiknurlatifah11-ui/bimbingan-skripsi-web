#perulangan dan kamus

#list
dompet = [5000,10000,20000,50000]
#akses element list: positif 0 kiri
#akses element list: negatif

#perulangan
#for variabel in sequence
#sequence: list, string, tuple
#ciri sequence: elemennya bisa diakses menggunakan indeks

for i in dompet:
    print(i)

#string
nama = "Titik nur latifah"
print(nama[0]) #T

for ch in nama:
    print(ch)

#while ->(menyebabkan infinite looping)
#inisialisasi variabel n
n = 0
while (n<5):
    print(n)
    #tambah nilai n nya
    n += 1

#cara membedakan looping for(hrs pny sequence) dan while(incr)

#menggabungkan looping dan percabangan
print(range(0,9))

for i in range(0, 9):
    if i == 4:
        continue
    elif i == 7:
        break
    
    print(i)

    #kombinasi looping,percabangan, dan fungsi input

jumlahdata = int(input("masukkan jumlah data yang akan diinput:"))

#list untuk menampung data
#append = untuk menambahkan data ke dalam list
nama = []

for i in range(jumlahdata):
    tambahnama = input("masukkan nama yang akan ditambahkan:")
    nama.append(tambahnama)
    if tambahnama == "nurul":
        break

print(nama)

user = "Ahmad"
passw = 1122


print("Daftar Menu:")
print("1. log in")
print("2. input data")
print("3. keluar")

pilihanUser = input("silahkan masukkan menu:")
while True:
    print("Test")
    if pilihanUser == "1":
        u = input("masukkan username anda:")
        p = int(input("masukkan password anda:"))
        if user == u and passw == p:
            print("berhasil masuk ke system")
            break
        else:
            print("masukkan akun dengan benar")


