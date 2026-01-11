#file

# Buat File
file = open("Daftar-belanja.txt","w")
file.write("Daftar belanja ibu: ")
file.write("\n1. cabai")
file.write("\n2. bawang")
file.write("\n3. garam")
file.close()


#update file
file = open("Daftar-belanja.txt","a")
file.write("\n4. tepung")
file.close()

#baca file
file = open("Daftar-belanja.txt","r")
print(file.read())
# print(file.readline())
# print(file.read(10)) #membaca 10 karakter saja
# print(file.readline())
#file.close()

#mwnghapus isi file
# file = open("Daftar-belanja.txt","w")
# file.close()

#hapus file
# import os
# os.remove("Daftar-belanja.txt")

#=================penggunaan SQLite3 pada penggunaan database===========
import sqlite3
import random
