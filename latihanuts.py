# def tarik_tunai(nominal):
#     saldo = 1000000 #saldo awalsebesar 1.000.000

#     if  nominal > 1000000:
#         print("maaf dana mengendap direkening harus 100.000")
#     elif nominal < 900000:
#         saldo = nominal
#         print("nominal tarik tunai:", nominal)
#         print("saldo akhir: ", saldo)
#     else: 
#         print("nominal tidak valid. pastikan nominal tarik tunai kurang dari 900000")

#         return saldo

# # contoh pemanggilan fungsi:
# nominal_uang = int(input("masukkan nominal uang yang akan ditarik tunai:"))
# saldo_akhir =  tarik_tunai(nominal_uang)

# def tarik_tunai(nominal):
#     saldo = 1000000 #saldo awalsebesar 1.000.000

#     if  nominal > 1000000:
#         print("saldo anda tidak cukup")
#     elif nominal ("900000 smp dgn 1000000") :
#          print("maaf dana mengendap direkening harus 100000")
#     elif nominal < 900000:
#         saldo = nominal
#         print("nominal tarik tunai:", nominal)
#         print("saldo akhir: ", saldo)
#     else: 
#         print("nominal tidak valid. pastikan nominal tarik tunai kurang dari 900000")

#         return saldo

# # contoh pemanggilan fungsi:
# nominal_uang = int(input("masukkan nominal uang yang akan ditarik tunai:"))
# saldo_akhir = tarik_tunai(nominal_uang)

def tarik_tunai(nominal):
    # Saldo awal
    saldo = 1000000
    
    # Pengecekan nominal
    if nominal > 1000000:
        print("Saldo anda tidak cukup")
    elif 900000 <= nominal <= 1000000:
        saldo -= 100000
        print("Maaf, dana mengendap di rekening harus 100.000")
        print("Nominal tarik tunai:", nominal)
        print("Saldo akhir:", saldo)
    elif nominal < 900000:
        saldo -= nominal
        print("Nominal tarik tunai:", nominal)
        print("Saldo akhir:", saldo)
    
    return saldo

# Contoh pemanggilan fungsi
nominal_tarik = int(input("Masukkan nominal tarik tunai: "))
saldo_akhir = tarik_tunai(nominal_tarik)






