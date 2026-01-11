class Benda:
    def _init_(self, nama):
        self.nama = nama

    def deskripsi(self):
        print(f"Ini adalah sebuah benda bernama {self.nama}")


class Elektronik(Benda):
    def _init_(self, nama, merek):
        super()._init_(nama)
        self.merek = merek

    def tampilkan__info(self):
        super().deskripsi()
        print(f"Benda ini adalah produk dari merek {self.merek}")


class Gadget(Elektronik):
    def _init_(self, nama, merek, tipe):
        super()._init_(nama, merek)
        self.tipe = tipe

    def periksa__kondisi(self):
        super().tampilkan__info()
        print(f"Tipe dari gadget ini adalah {self.tipe}")

class Smartphone(Gadget):
    def _init_(self, nama, merek, tipe, layar):
        super()._init_(nama, merek, tipe)
        self.layar = layar

    def cek__layar(self):
        super().periksa__kondisi()
        print(f"Ukuran layar dari smartphone ini adalah {self.layar} inch")


class Laptop(Gadget):
    def _init_(self, nama, merek, tipe, prosesor):
        super()._init_(nama, merek, tipe)
        self.prosesor = prosesor

    def cek__prosesor(self):
        super().periksa__kondisi()
        print(f"Prosesor yang digunakan pada laptop ini adalah {self.prosesor}")
