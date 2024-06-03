# """
# cetak: gunakan fungsi input
# kode MataKuliah: IF1111
# Nama MataKuliah: pemrograman
# Nilai UTS: 90
# Nilau UAS: 90
# Nilai Praktikum: 90
# Nilai Aktivitas: 90
# Nilai Proyek: 90
# Total Nilai:
# Nilai Huruf:
# """
#keterangan total nilai
# UTS - 20%, UAS - 20%, Praktikum - 10%, aktivitas - 20%, proyek - 20%

#total nilai = %UTS + %UAS + %PRK + %AKT + %PRO
#total nilai = 18 + 18 + 9 + 18 + 27 = 90

#nilai huruf:
#90 - 100 :
#80 - 90 :
#70 - 80 :
#60 - 70 :
# 50 - 60 :
#<50 : F

KodeMataKuliah = input("masukkan kode MataKuliah:")
NamaMataKuliah = input("Masukkan nama MataKuliah:")
NilaiUTS = int(input("masukkan nilai UTS:"))
NilaiUAS = int(input("masukkan nilai UAS:"))
NilaiPRAK = int(input("masukkan nilai PRAK:"))
NilaiAKT = int(input("masukkan nilai AKT:"))
NilaiPRO = int(input("masukkan nilai PRO:"))

print("KodeMataKuliah,NamaMataKuliah,NilaiUTS,NilaiUAS,NilaiPRAK,NialiAKT,NialiPRO")

totalNilai = (NilaiUTS * 20 /100) + (NilaiUAS * 20 /100) + (NilaiPRAK * 10 /100) + (NilaiAKT * 30 /100) + (NilaiPRO * 20 /100)

print(totalNilai)

if totalNilai >=90:
    print("grade A")
elif totalNilai <90 and totalNilai >=80:
    print("grade B")
elif totalNilai <80 and totalNilai >=70:
    print("grade C")
elif totalNilai <70 and totalNilai >=60:
    print("grade D")
elif totalNilai <60 and totalNilai >=50:
    print("grade E")
elif totalNilai <50:
    print("grade F")