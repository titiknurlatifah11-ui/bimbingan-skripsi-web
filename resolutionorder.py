#resolutionOrder
#diliat dri perintah mana yang jalan duluan

class A:

    def info(self):
        print("ini kelas A")

class B:

    def info(self):
        print("ini kelas B")

class C(A,B):
    # def info(self):
    #     print("ini kelas c")
    pass

objek = C()
objek.info()
# help(objek)