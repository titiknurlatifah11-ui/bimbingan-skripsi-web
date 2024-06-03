def hitung_nilai(nilaipersen):
    if nilaipersen >= 90:
        return 4.0
    elif nilaipersen >= 80:
        return 3.7
    elif nilaipersen >= 70:
        return 3.3
    elif nilaipersen >= 60:
        return 3.0
    elif nilaipersen >= 50:
        return 2.7
   
def nilai_ke_huruf(nilai):
    if nilai == 4.0:
        return 'A'
    elif nilai == 3.7:
        return 'B'
    elif nilai == 3.3:
        return 'C'
    elif nilai == 3.0:
        return 'D'
    elif nilai == 2.7:
        return 'E-'
    
def buat_matakuliah():
    jumlah_matakuliah = int(input("Masukkan jumlah matakuliah: "))
    
    matakuliah_list = []
    
    for i in range(jumlah_matakuliah):
        nama_mk = input("Masukkan nama matakuliah ke-{}: ".format(i + 1))
        sks = int(input("Masukkan SKS matakuliah {}: ".format(nama_mk)))
        nilai = float(input("Masukkan nilai matakuliah {}: ".format(nama_mk)))
        huruf = input("Masukkan huruf nilai matakuliah {}: ".format(nama_mk))
        
        matakuliah = {
            'namaMK': nama_mk,
            'SKS': sks,
            'Nilai': nilai,
            'Huruf': huruf
        }
        
        matakuliah_list.append(matakuliah)
    
    return matakuliah_list

list_matakuliah = buat_matakuliah()

print("List Matakuliah:")
for matakuliah in list_matakuliah:
    print(matakuliah)



