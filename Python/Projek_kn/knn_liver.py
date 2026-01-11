import pandas as pd
from sklearn.model_selection import KFold
from sklearn.neighbors import KNeighborsClassifier
from sklearn.metrics import accuracy_score
from sklearn.preprocessing import LabelEncoder

# 1. Baca data
df = pd.read_csv('csv_KNN_Liver_CSU_1.csv')  # pastikan nama file cocok

# 2. Shuffle data
df = df.sample(frac=1, random_state=42).reset_index(drop=True)

# 3. Pisahkan fitur dan label
X = df.iloc[:, :-1]  # semua kolom kecuali kolom terakhir
y = df.iloc[:, -1]   # kolom terakhir sebagai label

# Jika labelnya bukan angka, encode dulu
if y.dtype == 'object':
    le = LabelEncoder()
    y = le.fit_transform(y)

# 4. 4-fold cross-validation
kf = KFold(n_splits=4, shuffle=False)
splits = list(kf.split(X))

# Ambil fold pertama
train_idx, test_idx = splits[0]
X_train, X_test = X.iloc[train_idx], X.iloc[test_idx]
y_train, y_test = y[train_idx], y[test_idx]

# 5. KNN (k=3)
knn = KNeighborsClassifier(n_neighbors=3)
knn.fit(X_train, y_train)
y_pred = knn.predict(X_test)

# 6. Evaluasi
acc = accuracy_score(y_test, y_pred)
print(f"Akurasi pada split pertama: {acc:.4f}")