class Dokter:
    def __init__(self, nama):
        self.nama = nama
    @property
    def deskripsi(self):
        print(f"Ini adalah seorang dokter bernama {self.nama}")


class Pasien:
    def __init__(self, nama, penyakit):
        self.nama = nama
        self.penyakit = penyakit

    def tampilkan__info(self):
        print(f"pasien ini memiliki penyakit {self.penyakit}")


class Resep(Pasien):
    def __init__(self, nama, penyakit, dosis):
        super().__init__(nama, penyakit)
        self.dosis = dosis

    def periksa__kondisi(self):
        super().tampilkan__info()
        print(f"resep ini memiliki dosis {self.dosis}")


class Obat(Resep):
    def __init__(self, nama, penyakit, dosis, umur):
        super().__init__(nama, penyakit, dosis)
        self.umur = umur

    def cek__layar(self):
        super().periksa__kondisi()
        print(f"obat ini untuk {self.nama} yang memiliki nama paracetamol")
        if self.umur < 12:
            dosis_asli = self.dosis.split()[0]  # Ambil bagian angka dari dosis
            satuan = self.dosis.split()[-1]     # Ambil bagian satuan dari dosis
            dosis_baru = str(int(dosis_asli) * 2) + " " + satuan  # Hitung dosis baru
            print(f"Pasien di bawah 12 tahun, dosis ditingkatkan menjadi {dosis_baru}")

class Apotek(Obat):
    def __init__(self, nama, penyakit, dosis, instruksi):
        super().__init__(nama, penyakit, dosis)
        self.instruksi = instruksi

    def cek__prosesor(self):
        super().periksa__kondisi()
        print(f"intsruksi digunakan untuk {self.pasien} dengan {self.dosis} dengan instruksi dari {self.dokter}")

dokter = Dokter("kekez")
dokter.deskripsi

pasien = Obat("eha", "demam", "tiga kali sehari", 18 )
pasien.cek__layar()


