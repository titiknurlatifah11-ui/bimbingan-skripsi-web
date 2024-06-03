#studi kasus - Diskon 50%
baju = 100000
diskon = baju * 50 / 100
bajuDiskon = baju - diskon
print(bajuDiskon)

#studi kasus - Diskon 70% + 20%
baju = 100000
diskon = baju * 70 / 100 #70.000

bajuDiskon = baju - diskon #30.000

diskonGab = (baju * 70 / 100) + (baju * 20 / 100)
bajudiskonGab = baju - diskonGab

bajuDiskon = baju - diskon
diskon20 = bajuDiskon * 20 / 100
bajudiskon20 = bajuDiskon - diskon20
print("harga baju asli:", baju)
print("harga baju setelah diskon (70 + 20):", bajudiskon20)

diskonGab = (baju * 70 / 100) + (baju * 20 / 100)
bajudiskonGab = baju - diskonGab
print("harga baju setelah dsikon (70 + 20) - Gabungan:", bajudiskonGab)