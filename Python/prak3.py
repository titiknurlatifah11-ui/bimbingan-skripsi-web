#overiding
#overloading ada parameter kyk (tipe)
# overloading lebih efisien 

class hero:

    def __init__(self, name, health):
        self.name = name
        self.health = health

    def info(self,tipe):
        print(f"name = {self.name}\nhealth = {self.health} \ntipe = {tipe}")

class stat_str(hero):
    def __init__(self, name):
            super().__init__(name, 100)
            self.tipe = "strenght"

    def info(self):
            super().info("strenght")
            #print(f"name = {self.name} \nhealth = {self.health}") 

ares = stat_str("ares")
ares.info()

class Buah:
    def __repr__(self):
        return "ini kelas buah"

    def __str__(self):
        return"ini kelas buah1"

objek = Buah

