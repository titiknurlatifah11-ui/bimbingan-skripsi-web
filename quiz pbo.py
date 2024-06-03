#buatlah sebuah program, minimal terdiri dari 3 kelas dengan ketentuan kandidat kelas:
# pasien, dokter, apotek, resep, obat

class Obat:
    def __init__(self, nama, harga):
        self.nama = nama
        self.harga = harga

class Resep:
    def __init__(self, obat, jumlah):
        self.obat = obat
        self.jumlah = jumlah

    def hitung_total_harga(self):
        return self.obat.harga * self.jumlah

class Pasien:
    def __init__(self, nama, umur):
        self.nama = nama
        self.umur = umur
        self.resep = []

    def tambah_resep(self, resep):
        self.resep.append(resep)

class Dokter:
    def __init__(self, nama):
        self.nama = nama

    def tulis_resep(self, obat, jumlah):
        return Resep(obat, jumlah)

class Apotek:
    def __init__(self):
        self.daftar_obat = {}

    def tambah_obat(self, obat):
        self.daftar_obat[obat.nama] = obat

    def beli_obat(self, pasien):
        total_harga = 0
        for resep in pasien.resep:
            total_harga += resep.hitung_total_harga()
        return total_harga

# Membuat beberapa obat
obat1 = Obat("Paracetamol", 5000)
obat2 = Obat("Amoxicillin", 10000)
obat3 = Obat("Omeprazole", 8000)

# Membuat apotek dan menambahkan obat ke dalam daftar obat di apotek
apotek = Apotek()
apotek.tambah_obat(obat1)
apotek.tambah_obat(obat2)
apotek.tambah_obat(obat3)

# Membuat dokter
dokter = Dokter("Dr. Smith")

# Membuat pasien dan dokter menulis resep
pasien = Pasien("John", 30)
resep1 = dokter.tulis_resep(obat1, 2)
resep2 = dokter.tulis_resep(obat2, 1)

# Pasien menambahkan resep ke daftar resepnya
pasien.tambah_resep(resep1)
pasien.tambah_resep(resep2)

# Pasien membeli obat di apotek
total_harga = apotek.beli_obat(pasien)


print(f"Total harga obat yang harus dibayar oleh {pasien.nama}: Rp {total_harga}")
