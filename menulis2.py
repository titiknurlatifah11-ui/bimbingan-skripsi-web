#fungsi sebuah file

#cara1
with open("dokumen1","a") as file:
    file.write("teknik informatika")

#cara2
file2 = open("dokumen2.docx","w")
file2.write("isi dari file dokumen2 berubah")
file2.close()