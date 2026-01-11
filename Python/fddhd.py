# Mendefinisikan matriks m1, m2, dan m3
m1 = [[1, 2], [3, 4]]
m2 = [[5, 6], [7, 8]]
m3 = [[1*5 + 2*7, 1*6 + 2*8], [3*5 + 4*7, 3*6 + 4*8]]

# Mendefinisikan kelas matriks
class Matriks:
    def __init__(self, a, b):
        self.a = a
        self.b = b

    def __mul__(self, other):
        # Operasi perkalian matriks
        baris_a = len(self.a)
        kolom_a = len(self.a[0])
        kolom_b = len(other.b[0])

        result = [[0 for j in range(kolom_b)] for i in range(baris_a)]

        for i in range(baris_a):
            for j in range(kolom_b):
                for k in range(kolom_a):
                    result[i][j] += self.a[i][k] * other.b[k][j]

        return result

# Menggunakan kelas Matriks
matriks1 = Matriks(m1, m2)
hasil = matriks1.__mul__(Matriks(m1, m2))
print(hasil)