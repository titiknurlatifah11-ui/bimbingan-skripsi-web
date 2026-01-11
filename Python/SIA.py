#koordinasi dengan modul yang lain(import)
import datamahasiswa#import 1 modul
import dosen , matakuliah #import lebih dari 1 modul

#untuk memanggil entitas(.)
print (datamahasiswa.nama)
print (datamahasiswa.NIM)
print (datamahasiswa.alamat)

print("Data Dosen\n")
print(dosen.nama)

print("Data matakuliah\n")
print(matakuliah.NamaMK)