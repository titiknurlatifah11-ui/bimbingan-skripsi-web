#dictionary
# thislist = ["saya", "anda", "mereka"] #saya = 0, anda = 1, mereka = 2

# thisDict = {
#     "merek" : "ford",
#     "jenis" : "mustang",
#     "tahun" : 1984
# } #key : value

# print(len(thisDict))




#constructor
# dictBaru = dict()
# dictBaru["aku"] = "kamu"
# dictBaru["dia"] = "mereka"
# print(dictBaru)


# nilaiDict = {
#      "Danang" : [20,90,50,40,70,80],
#      "Darto" : [30,90,40,90,80,80]
#  }

# # print(nilaiDict["Danang"])

# # print(type(nilaiDict))


# dictProdi = {
#     "Informatika" : {
#         "Nanda" : {
#             "progdas1" : 100,
#             "ilmuData" : 90
#         },
#         "Dendi" : {
#             "progdas1" : 100,
#             "ilmuData" : 90
#         }
#     },
#     "manajemen" : {
#         "fajar" : {
#             "ips" : 90,
#             "ekonomi" : 100
#        }
#    }
# }

# print(dictProdi.value())
# print()



# nilaiDict = {
#      "Danang" : [20,90,50,40,70,80],
#      "Darto" : [30,90,40,90,80,80]
#  }

# total = 0
# for item in nilaiDict["Danang"]:
#     total += item

# # print(total/len(nilaiDict["Danang"]))
# print(f'nilai total adalah : {total/len(nilaiDict["Danang"])}')

# luasWilayah = {
#     "Bandung" : 1390,
#     "Padang" : 2100,
#     "tangerang" : 3109,
#     "Serpong" : 1100
# }

# simpanKeys = luasWilayah.Keys()
# simpanValues = luasWilayah.values()

# print(f"""ini method get Keys( : {simpanKeys}, \n
#      ini method get Values() : {simpanValues}""")

# print(type(simpanKeys))
# print(type(simpanValues))

# luasWilayah["Serang"] = 4102

# print(luasWilayah)

# if "tangerang" in luasWilayah:
#     print("Ya, tangerang ada di luasWilayah")
# else:
#     print("GA ADA")

# luasWilayah2 = {
#     "Jambi" : 11000,
#     "Bengkulu" : 12000
# }

# luasWilayah.update(luasWilayah)
# print(luasWilayah)

# luasWilayah.pop("tangerang")

# print(luasWilayah)



mahasiswa = {
    "mhs1" : "raihan",
    "mhs2" : "randi",
    "mhs3" : "roza",
    "mhs4" : "raze",
    "mhs5" : "roma",
}

for item in mahasiswa:
    print(item, " : ",mahasiswa[item])