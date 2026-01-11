import sqlite3

def buatTabelMhs():
    #menguhungkan ke database sia.db
    #membuat koneksi
    koneksi = sqlite3.connect("sia.db")
    #membuat kursor
    kursor = koneksi.cursor()
    #membuat tabel mahasiswa dengan kursor
    #Query SQL
    querySQL = (""" CREATE TABLE mahasiswa
    (nama TEXT,jurusan TEXT)
    )""")
    #eksekusi query tsb
    kursor.execute(querySQL)

buatTabelMhs()