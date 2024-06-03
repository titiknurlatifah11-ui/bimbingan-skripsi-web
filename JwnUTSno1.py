#JAWABAN SOAL NO 1
mata_kuliah = []

def hitung_total_nilai(uts, uas, praktikum, quiz, partisipasi, project):
    total_nilai = 0.15 * uts + 0.20 * uas + 0.10 * praktikum + 0.05 * quiz + 0.15 * partisipasi + 0.35 * project
    return total_nilai

def nilai_huruf(total_nilai):
    if total_nilai >= 90:
        return 'A'
    elif total_nilai >= 80:
        return 'B'
    elif total_nilai >= 70:
        return 'C'
    elif total_nilai >= 60:
        return 'D'
    else:
        return 'E'

jumlah_mata_kuliah = int(input("Masukkan jumlah mata kuliah: "))

for i in range(jumlah_mata_kuliah):
    print(f"\nMata Kuliah ke-{i+1}:")
    uts = float(input("Nilai UTS: "))
    uas = float(input("Nilai UAS: "))
    praktikum = float(input("Nilai Praktikum: "))
    quiz = float(input("Nilai Quiz: "))
    partisipasi = float(input("Nilai Aktivitas Partisipasi: "))
    project = float(input("Nilai Hasil Project: "))

    total_nilai = hitung_total_nilai(uts, uas, praktikum, quiz, partisipasi, project)
    grade = nilai_huruf(total_nilai)

    mata_kuliah.append({
        'UTS': uts,
        'UAS': uas,
        'Praktikum': praktikum,
        'Quiz': quiz,
        'Partisipasi': partisipasi,
        'Project': project,
        'Total Nilai': total_nilai,
        'Grade': grade
    })

print("\nHasil per Mata Kuliah:")
for i, matkul in enumerate(mata_kuliah, start=1):
    print(f"\nMata Kuliah ke-{i}:")
    print(f"Total Nilai: {matkul['Total Nilai']}")
    print(f"Grade: {matkul['Grade']}")

mata_kuliah_lulus = [matkul for matkul in mata_kuliah if matkul['Grade'] == 'C']
ipk = len(mata_kuliah_lulus) / jumlah_mata_kuliah * 4.0

print(f"\nIPK: {ipk}")

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





