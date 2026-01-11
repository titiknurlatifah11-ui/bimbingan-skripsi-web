class A:

    def info(self):
        print("ini kelas A")

class B(A):
    def info(self):
        print("ini kelas B")
    pass

class C(A):
    def info(self):
        print("ini kelas c")
    pass

class D(B):
    def info(self):
        print("ini kelas D")

class E(D,C):
    pass

objek = E()
objek.info()

help(objek)