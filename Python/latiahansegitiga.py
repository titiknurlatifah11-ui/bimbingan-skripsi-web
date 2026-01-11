import math

class SegitigaSikuSiku:
    def __init__(self, alas, tinggi):
        """
        Metode konstruktor untuk menginisialisasi objek SegitigaSikuSiku dengan alas dan tinggi.
        
        Args:
        alas: Panjang alas segitiga.
        tinggi: Tinggi segitiga.
        """
        self.alas = alas
        self.tinggi = tinggi
    
    def luas(self):
        """
        Menghitung luas segitiga siku-siku.
        
        Returns:
        Luas segitiga.
        """
        return 0.5 * self.alas * self.tinggi
    
    def keliling(self):
        """
        Menghitung keliling segitiga siku-siku.
        
        Returns:
        Keliling segitiga.
        """
        sisi_miring = math.sqrt(self.alas**2 + self.tinggi**2)
        return self.alas + self.tinggi + sisi_miring

# Contoh penggunaan
segitiga = SegitigaSikuSiku(3, 4)
print("Luas segitiga:", segitiga.luas())
print("Keliling segitiga:", segitiga.keliling())
