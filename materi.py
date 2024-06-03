#komentar dalam bahasa pemrograman python
'''
Ini adalah komentar multiline
penulis : titik nur latifah
tanggal : 6 oktober 2023
mata kuliah : pemrograman dasar
''
'
#review materi input, output, proses

#output (print)
print("selamat datang di MK pemrograman dasar")

#input
input ("masukkan nama anda:")

#proses - variabel dan tipe data
print ("selamat datang di sistem informasi akademik")

#variabel - tipe data string - nama : username

#database user sistem infrasi akademik
username = "1152700001"
password = "123456"

print(username)
print(password)

user = input("masukkan username anda:")
keyt = input("masukkan password anda:")

#percabangan - if elif else
#Operator perbandingan == (sama dengan)

#cara 1 (if bersarang)
if user == username :
    print(user == username) #true
    print("username anda sudah sesuai dengan database")
    if key == password:
        print(key == password) #true
        print("password anda sudah sesuai dengan database")
    else:
        print(key == password) #false
        print("password anda salah")
else:
    print(user == username) #false
    print("username anda salah")

#cara 2 : operator logika and, or, not, xor
if (user == username) and (key == password):
    print("username dan passwoed sudah sesuai")
else:
    print("username atau password anda tidak sesuai")
'''

#list - tempat untuk menyimpan lebih dari satu data
#struktur data / tipe data list
#cara membuat list menggunakan []

# menyimpan 4 data mahasiswa
listUsername = [1152700001,1152700003,1152700005,1152700007]
print(listUsername)

#ada beberapa fungsi dalam list
#len = fungsi untuk mengetahui jumlah data pada list
print(len(listUsername)) #4

#ada beberapa metode dalam list: extend, append
#extend, append = metode u/ menambah anggota dalam list
listUsername.append(1152700009) # + 1 anggota
listUsername.extend([11527000011,1152700013]) # > 1 anggota
print(listUsername)

#extend = sequence (list,string)

#akses element / anggota dalam list menggunakan indeks mulai 0
print(listUsername[0]) #mencetak 01

#akses juga bisa menggunakan indeks negatif nilai -1
print(listUsername[-2]) #mencetak 11

#list juga bisa bersarang
listUsername.append([1152700015,1152700017])
print(listUsername)
print(listUsername[-1][1])

#list juga bisa menampilkan > 1 data - sliccing / memotong
print(listUsername[0:5])

#list juga bisa menampilkam > 1 data - start,stop,step
print(listUsername[0:5:2])

#kombinasi list dan fungsi input

tambahAnggota = int(input("masukan nim anggota baru:"))
listUsername.append(tambahAnggota)
print(listUsername)
