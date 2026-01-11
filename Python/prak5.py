#hubungan antar ibjek

# 6.1 asosiasi
# class penjual():
#     def __init___(self, nama, barang, harga):
#         self.nama = nama
#         self.barang = barang
#         self.harga = harga
#     def jual(self, pembeli):
#         #prinsip(self.nama, "Menjual", self.barang, "Kepada", pembeli, "Dengan harga", self.harga, "")
#         print(f"{self.nama} menjual {self.barang} kepada {pembeli.nama} dengan harga {self.harga}")
# class pembeli():
#     def __init__(self, nama):
#         self.nama = nama

# def mainProgram():
#     print("Hubungan Antar Penjual dan Pembeli")
#     Guntur = penjual("Guntur", "Narkoboy", "Rp 1.0000.000")
#     Ahmid = pembeli("Ahmid")
#     Guntur.jual(Ahmid)

# mainProgram()

#6.2 Agregasi
# class karyawan():
#     def __init__(self, nama, umur):
#         self.nama = nama
#         self.umur = umur
#     def info(self, honor):
#         totalGaji = honor.gajiPokok + honor.lembur + honor.tunjangan
#         totalGaji -= honor.pajak
#         print(f"{self.nama} mendapatkan honor sebanyak Rp. {totalGaji} per bulan")
# class honor():
#     def __init__(self, gajiPokok, lembur, tunjangan, pajak):
#         self.gajiPokok = gajiPokok
#         self.lembur = lembur
#         self.tunjangan = tunjangan
#         self.pajak = pajak
# def main():
#     print("Agregasi (saling memiliki)")
#     faris = karyawan("faris", 24)
#     honor = honor(5000000, 3000000, 2000000, 1000)
#     faris.info(honor)
# main()  

#6.3 komposisi
class warga():
    def __init__(self, nama):
        self.nama = nama
        self.ktp = KTP(None) 
class KTP():
    def __init__(self, NIK):
        self.NIK = NIK
    def cetak(self, warga):
        print(f"kepemilikan NIK KTP ini {self.NIK} adalah {warga.nama}")
def main():
    rafi = warga("Rafi")
    rafi.ktp.NIK = input("silahkan masukkan NIK KTP: ")
    rafi.ktp.cetak(rafi)
main()
