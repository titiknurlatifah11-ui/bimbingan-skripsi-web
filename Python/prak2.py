# #konsep PBO

# #pewarisan

# class uang:
#     def __init__(self, nominal):
#         self.nominal = nominal

#     def belanja(self, totalbelanja):
#         self.totalbelanja = totalbelanja
#         print(self.nominal - totalbelanja)

# class koin(uang):
#     def __init__(self, bahan):
#         self.bahan = bahan
#     def beli(self, totalbeli):
#         print(f"terimakasih sisa uang anda adalah {self.nominal - totalbeli}")

# uang1 = uang(10000)
# uang1.belanja = 5000

# koin1 = koin("logam")
# koin1.nominal = 2000
# koin1.beli(500)
# koin1.belanja(500)



class Ortu:
    def __init__(self, sifatParam, tinggiParam, rasParam):
        self.sifat = sifatParam
        self.kategoriTinggi = tinggiParam
        self.ras = rasParam

class Anak(Ortu):
    def __init__(self, sifatParam, tinggiParam, rasParam):
        super().__init__(sifatParam, tinggiParam, rasParam)

    def printStatus(self):
        print(f"sifat bawaan: {self.sifat}\ntinggi bawaan: {self.kategoriTinggi}\nras bawaan: {self.ras}\n")

anak = Anak("Pemarah", "Tinggi", "Eropa")
anak.printStatus()














