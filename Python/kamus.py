#Kamus Python{} unsur1: kunci/key unsur2:nilai/value, kunci hrs identik dlm kamus

pemrograman = {"namaMK":"pemrograman" ,"sks":4, "nilai":90}
#merubah nilai
pemrograman["nilai"] = 80

#menambah data baru
pemrograman["huruf"] = "A"

#cara mengakses nilai/value menggunakan kunci/key
print(pemrograman["nilai"])
print(pemrograman["huruf"])
print(pemrograman)

#menghapus 1 data (hnya menerima 1 input)
del pemrograman["huruf"]
print(pemrograman)

#menghapus semua data 
#pemrograman.clear()
#print(pemrograman)

#mengintegrasikan kamus dgn list
matematika = {"namaMK":"Matematika" ,"sks":2,"nilai":70}

mataKuliah = [pemrograman,matematika]
print(mataKuliah[0]["nilai"]) #80