dataMhs = [{"nama":"nama","nim":"nim","jurusan":"jurusan"},{"nama":"titik nur latifah","nim":"1152700001","jurusan":"teknik informatika"}]

print(type(dataMhs[0]))

print(dataMhs[1]["nim"])

#fungsi menulis file
def menulisFile(linkFile,data):
    with open(linkFile,'w') as file:
        for mhs in data:
            stringMhs = f"{mhs['nama']},{mhs['nim']},{mhs['jurusan']}\n"
            print(stringMhs)
            file.write(stringMhs)

menulisFile('mhs.csv',dataMhs)
print("data berhasil ditulis kedalam sebuah file")
