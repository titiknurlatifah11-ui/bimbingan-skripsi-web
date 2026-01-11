#Studi Kasus Nilai

def nilaiMK ():
    uts = int(input("silahkan masukkan nilai uts:"))
    uas = int(input("silahkan masukkan nilai uas:"))
    partisipasi = int(input("silahkan masukkan nilai partisipasi:"))
    proyek = int(input("silahkan masukkan nilai proyek:"))
    #menghitung bobot/presentase
    b_uts = uts * 25 / 100
    b_uas = uas * 25 / 100
    b_partisipasi = partisipasi * 20 / 100
    b_proyek = proyek * 30 / 100

    #nilai total
    nilaiTotal = b_uts + b_uas +  b_partisipasi +  b_proyek
    print(nilaiTotal)
    return nilaiTotal

hasilFungsi = nilaiMK()

if hasilFungsi > 80:
    nilaiHuruf = "A"
elif hasilFungsi > 70:
    nilaiHuruf = "B"
else:
    nilaiHuruf = "C"

print(nilaiHuruf) #A