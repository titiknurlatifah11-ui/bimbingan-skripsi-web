

def penjumlahan3bil():
      a = int(input("masukkan bil pertama : "))
      b = int(input("masukkan bil kedua :"))
      c = int(input("masukkan bil ketiga:"))
      print(a+b+c)

def pengurangan3bil():
      a = int(input("masukkan bil pertama : "))
      b = int(input("masukkan bil kedua :"))
      c = int(input("masukkan bil ketiga:"))
      print(a-b-c)

def perkalian3bil():
      a = int(input("masukkan bil pertama : "))
      b = int(input("masukkan bil kedua :"))
      c = int(input("masukkan bil ketiga:"))
      print(a*b*c)

      print("menu")
      print("1. penjumlahan")
      print("2. pengurangan")
      print("3. perkalian")
userinput = input("pilih menu : ")
if userinput =='1':
      penjumlahan3bil()
elif userinput =='2':
      pengurangan3bil()
elif userinput == '3':
      perkalian3bil()
