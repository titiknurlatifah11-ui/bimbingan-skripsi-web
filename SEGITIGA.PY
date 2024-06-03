import math
class segitiga:
    a = 2
    b = 3
    c = 4
    
    def set_nilai(self, x, y, z):
        self.a = x
        self.b = y
        self.c = z

    def keliling(self):
        kel = self.a + self.b + self.c
        return kel
    
    def luas(self):
        s = (self.a + self.b + self.c)/2
        L = math.sqrt(s*(s - self.a)*(s - self.b)*(s - self.c))
        return L

st1 = segitiga()
st1.set_nilai(4,5,6)
print(st1.keliling())
print(st1.luas())

st2 = segitiga()
st2.set_nilai(8,9,10)
print(st2.keliling())
print(st2.luas())