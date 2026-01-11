#Dictionary

# thisDict = (
#     "Nama" : "kian santang",
#     "umur" : 17,
#     "kelamin" : "LK"
# )   #key : val

# for i in thisDict:
#     print(thisDict, " : ", thisDict(1))

# for key, value in thisDict.items():
#     print(key, " : ", value)

# thisDict = {
#     "Nama" : "kian santang",
#     "umur" : 17,
#     "kelamin" : "LK"
# }

# # print("ini awal (thisDict)")
# # key1 = input("masukkan key : ")
# # val1 = input("masukkan value : ")

# # thisDict[key1] = val1
# # print(thisDict)

# thisDictKosong = ()
# while True:
#     key1 = input("masukkan key : ")
#     val1 = input("masukkan value : ")
#     userInput = input("apakah data ingin diinput? v lanjut/N program stop)")
#     if userInput == "N":
#        break
#     else :
#        thisDict[key1] = val1
#        print(thisDictKosong)

# print(thisDictKosong)

kamus_terjemahan = {
    "i" : "saya",
    "hate" : "benci",
    "you" : "kamu"
}

def terjemahkan_kalimat(kalimat):
    return''.join(kamus_terjemahan.get(kata, kata)
for kata in kalimat.upper().split())

def main():
    input_user = input("masukkan kalimat(dalam bahasa inggris): ")
    hasil_terjemahan = terjemahkan_kalimat(input_user)

    print("terjemahan:", hasil_terjemahan.capitalize())
    if hasil_terjemahan != input_user.upper()
    else:
    ("kata tidak ditemukan dalam kamus.")



english = input("masukkan kalimat(inggris): ")



