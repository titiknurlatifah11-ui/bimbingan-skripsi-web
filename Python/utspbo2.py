class Buku:
    def __init__(self, judul, penulis, harga, stok):
        self.judul = judul
        self.penulis = penulis
        self.harga = harga
        self.stok = stok

class Pelanggan:
    def __init__(self, nama, alamat, saldo):
        self.nama = nama
        self.alamat = alamat
        self.saldo = saldo

def beli_buku(pelanggan, buku):
    if pelanggan.saldo >= buku.harga:
        pelanggan.saldo -= buku.harga
        buku.stok -= 1
        print("Pembelian berhasil!")
        print("Pelanggan:", pelanggan.nama)
        print("Buku yang dibeli:", buku.judul)
        print("Harga:", buku.harga)
        print("Sisa saldo:", pelanggan.saldo)
        print("Stok buku", buku.judul, "tersisa:", buku.stok)
    else:
        print("Saldo tidak mencukupi untuk membeli buku", buku.judul)

if __name__ == "__main__":
    buku1 = Buku("Laut Bercerita", "Leila S. Chudori", 150000, 10)
    pelanggan1 = Pelanggan("Nurul Rahmadani", "Jl. Tigaraksa", 200000)
    
    beli_buku(pelanggan1, buku1)
