def nilaiMK():
    UTS = int(input("Silahkan masukkan nilai UTS: "))
    UAS = int(input("Silahkan masukkan nilai UAS: "))
    Tugas = int(input("Silahkan masukkan nilai Tugas: "))
    Quiz = int(input("Silahkan masukkan nilai Quiz: "))
    PARTISIPASI = int(input("Silahkan masukkan nilai PARTISIPASI: "))
    PROYEK = int(input("Silahkan masukkan nilai PROYEK: "))

    b_UTS = UTS * 10/100
    b_UAS = UAS * 20/100
    b_Tugas = Tugas * 30/100
    b_Quiz = Quiz * 15/100
    b_PARTISIPASI = PARTISIPASI * 20/100
    b_PROYEK = PROYEK * 20/100

    nilaiTotal = b_UTS + b_UAS + b_Tugas + b_Quiz + b_PARTISIPASI + b_PROYEK
    return nilaiTotal

def nilaiHuruf(nilaiTotal):
    if nilaiTotal >= 90:
        return "A"
    elif nilaiTotal >= 80:
        return "B"
    elif nilaiTotal >= 70:
        return "C"
    elif nilaiTotal >= 60:
        return "D"
    else:
        return "E"

mata_kuliah =[]

JumlahMK = int(input("Masukkan jumlah matakuliah: "))

for i in range(JumlahMK):
    print(f"Masukkan data untuk Mata Kuliah ke-{i + 1}:")
    namaMK = input("Nama MataKuliah: ")
    SKS = int(input("SKS: "))
    Nilai = nilaiMK()
    Huruf = nilaiHuruf(Nilai)
    
    mata_kuliah.append({"Nama Mata Kuliah": namaMK,"SKS": SKS,"Nilai": Nilai,"Nilai Huruf": Huruf})
    
print(mata_kuliah)
