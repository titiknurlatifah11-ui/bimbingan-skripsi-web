#membaca sebuah file

#cara 1
with open("dokumen.txt","a") as file:
    file.write(", teknik informatik")

#cara2
file2 = open("dokumen2.docx","w")
file2.write("isi dari file dokumen2 berubah")
file2.close()

#cara1
with open("dokumen.txt","r") as file:
    hasil = file.read()
    print(hasil)

#cara2
file = open("dokumen2.docx")
hasil = file.read()
print(hasil)