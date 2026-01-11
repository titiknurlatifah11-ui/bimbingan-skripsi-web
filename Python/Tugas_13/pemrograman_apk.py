def longest_common_subsequence(str1, str2):
    m, n = len(str1), len(str2)

    # Membuat matriks kosong untuk menyimpan panjang LCS
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    # Mengisi matriks dp
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if str1[i - 1] == str2[j - 1]:  # Jika karakter cocok
                dp[i][j] = dp[i - 1][j - 1] + 1
            else:  # Jika karakter tidak cocok
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

    # Menelusuri kembali untuk mendapatkan LCS
    lcs = []
    i, j = m, n
    while i > 0 and j > 0:
        if str1[i - 1] == str2[j - 1]:
            lcs.append(str1[i - 1])
            i -= 1
            j -= 1
        elif dp[i - 1][j] > dp[i][j - 1]:
            i -= 1
        else:
            j -= 1

    # Karena ditambahkan dari belakang, urutan harus dibalik
    lcs.reverse()

    return dp[m][n], ''.join(lcs)


# Input dari pengguna
string1 = input("Masukkan string pertama: ")
string2 = input("Masukkan string kedua: ")

# Memanggil fungsi
length, lcs_string = longest_common_subsequence(string1, string2)

# Output hasil
print(f"Panjang LCS: {length}")
print(f"LCS: {lcs_string}")
