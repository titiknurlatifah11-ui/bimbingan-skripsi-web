#fungsi penarikan

def tarikTunai(akun,nominal):
    print("transaksi anda berhasil")
    saldoAkhir = akun.saldo - nominal
    print("saldo anda menjadi:" , saldoAkhir)
    #apdet saldo yang ada diakun
    akun.saldo = saldoAkhir
    return akun