#looping contitional

# import random

# print("selamat datang dipermainan tebak angka")
# userinput = int(input("masukkan angka : "))
# angkarandom = random.randint(1, 10)

# while userinput != angkarandom:
#     if userinput == angkarandom:
#         break
#     else:
#         print('salah, masukkan ulang')
#         userinput = int(input("masukkan angka : "))

# print('yey, benar')

import random

print("selamat datang dipermainan tebak angka")
angkarandom = random.randint(1, 10)

while True:
   userinput = int(input("masukkan angka : "))
   if userinput != angkarandom:
         print("salah")
else:
         print('benar')

angka = 0

while angka < 10:
     print(angka)
     angka += 1

