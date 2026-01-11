import pandas as pd
from sklearn.preprocessing import MinMaxScaler

# Memuat kumpulan data
df = pd.read_csv('C:/Users/Administrator/Documents/TITIK NUR LATIFAH/DATA TITIK NUR LATIFAH//PYTHON/diabetes.csv')


# Mengidentifikasi nilai yang tidak valid (nol) dalam kolom di mana angka nol tidak diharapkan
invalid_columns = ['Glucose', 'BloodPressure', 'SkinThickness', 'Insulin', 'BMI']

# Mengisi nilai yang hilang (nol) menggunakan median dari setiap kolom
for column in invalid_columns:
    median_value = df[column].median()
    df[column] = df[column].replace(0, median_value)

# Menormalkan kolom numerik menggunakan penskalaan Min-Maks
numeric_columns = ['Pregnancies', 'Glucose', 'BloodPressure', 'SkinThickness', 'Insulin', 'BMI', 'DiabetesPedigreeFunction', 'Age']
scaler = MinMaxScaler()
df[numeric_columns] = scaler.fit_transform(df[numeric_columns])

# Menampilkan beberapa baris pertama dari kumpulan data yang telah dibersihkan dan dinormalisasi
print(df.head())

# Menyimpan kumpulan data yang telah dibersihkan dan dinormalisasi ke file CSV baru
df.to_csv('cleaned_diabetes.csv', index=False)
