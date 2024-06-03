# Inisialisasi variabel untuk menghitung IPS dan IPK
total_mutu_semester = 0
total_sks_semester = 0
total_ips = 0
total_semester = int(input("Masukkan jumlah semester yang telah dijalani: "))

# Loop untuk setiap semester
for semester in range(1, total_semester + 1):
    print(f"\n=== Semester {semester} ===")
    total_mutu = 0
    total_sks = 0

    # Loop untuk menginput data mata kuliah
    jumlah_mata_kuliah = int(input(f"Masukkan jumlah mata kuliah di Semester {semester}: "))
    for i in range(jumlah_mata_kuliah):
        # ...

        # Menghitung nilai/mutu mata kuliah
        mutu = jumlah_mata_kuliah['Total Nilai'] * jumlah_mata_kuliah
        total_mutu += mutu
        total_sks += jumlah_mata_kuliah

        # ...

    # Menghitung IPS per semester
    ips_semester = total_mutu / total_sks
    print(f"IPS Semester {semester}: {ips_semester:.2f}")

    # Menyimpan IPS per semester untuk perhitungan IPK
    total_mutu_semester += total_mutu
    total_sks_semester += total_sks
    total_ips += ips_semester

# Mencetak hasil IPK
ipk = total_ips / total_semester
print(f"\nIPK: {ipk:.2f}")


total_mutu_semester = 0
total_sks_semester = 0
total_ips = 0
total_semester = int(input("Masukkan jumlah semester yang telah dijalani: "))

for semester in range(1, total_semester + 1):
    print(f"\n=== Semester {semester} ===")
    total_mutu = 0
    total_sks = 0

    jumlah_mata_kuliah = int(input(f"Masukkan jumlah mata kuliah di Semester {semester}: "))
    for i in range(jumlah_mata_kuliah):
        sks_mata_kuliah = float(input(f"Masukkan jumlah SKS mata kuliah ke-{i + 1}: "))
    

        mutu = matkul['Total Nilai'] * sks_mata_kuliah
        total_mutu += mutu
        total_sks += sks_mata_kuliah

    ips_semester = total_mutu / total_sks
    print(f"IPS Semester {semester}: {ips_semester:.2f}")

    total_mutu_semester += total_mutu
    total_sks_semester += total_sks
    total_ips += ips_semester

ipk = total_ips / total_semester
print(f"\nIPK: {ipk:.2f}")






