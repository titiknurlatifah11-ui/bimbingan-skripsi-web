class PBO:
    pass

IF_A = PBO() #object
print (IF_A)

class PBO:
    def __init__(self, nama, nilai):
        self.NamaMahasiswa = nama
        self.NilaiMahasiswa = nilai

IF_A = PBO("Wahyudi", 80) #object
IF_B = PBO("Kekez", 80) #object
#mencetak atribut
print("Nama Mahasiswa: ", IF_A.NamaMahasiswa, "\nNilai Mahasiswa: ", IF_A.NilaiMahasiswa)
print("Nama Mahasiswa: ", IF_B.NamaMahasiswa, "\nNilai Mahasiswa: ", IF_B.NilaiMahasiswa)


#class/blueprint/template
class PBO:
    def __init__(self, nama, umur):
        self.namaSaya = nama
        self.umurSaya = umur
        
    #method
    def cetakNamaUmur(self):
        return f"nama saya: {self.namaSaya}, umur saya: {self.umurSaya}"

# test = PBO("jonathan", 100) 
# print(test.cetakNamaUmur())
# print(test.umurSaya)

objectOrang = []
pilihan = int(input("Masukkan Berapa Object: "))
for i in range(pilihan):
    InputNama = input("Masukkan Nama saya: ")
    InputUmur = int(input("Masukkan Umur Saya: "))
    object = PBO(InputNama, InputUmur)
    objectOrang.append(object)

for object in objectOrang:
    print("Nama",object.namaSaya, "Umur", object.umurSaya)