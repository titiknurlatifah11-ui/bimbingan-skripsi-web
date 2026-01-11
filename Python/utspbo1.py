class Karyawan:
    def __init__(self, nama, jabatan, jam_kerja, tarif_gaji):
        self.nama = nama
        self.jabatan = jabatan
        self.set_jam_kerja(jam_kerja)
        self.set_tarif_gaji(tarif_gaji)
    
    def set_jam_kerja(self, jam_kerja):
        if jam_kerja < 0:
            raise ValueError("Jam kerja tidak boleh negatif")
        self.jam_kerja = jam_kerja

    def set_tarif_gaji(self, tarif_gaji):
        if tarif_gaji < 0:
            raise ValueError("Tarif gaji tidak boleh negatif")
        self.tarif_gaji = tarif_gaji

    def tampilkan_info(self):
        print(f"Nama: {self.nama}")
        print(f"Jabatan: {self.jabatan}")
        print(f"Jam Kerja: {self.jam_kerja} jam")
        print(f"Tarif Gaji: Rp{self.tarif_gaji:.4f} per jam")

def hitung_gaji(karyawan):
    gaji = karyawan.jam_kerja * karyawan.tarif_gaji
    print(f"\nGaji karyawan {karyawan.nama} ({karyawan.jabatan}) adalah: Rp{gaji:.4f}")

karyawan1 = Karyawan("Eha", "Dosen", 8, 50000)
karyawan2 = Karyawan("Abyan", "Asisten", 8, 40000)

karyawan1.tampilkan_info()
hitung_gaji(karyawan1)

karyawan2.tampilkan_info()
hitung_gaji(karyawan2)
