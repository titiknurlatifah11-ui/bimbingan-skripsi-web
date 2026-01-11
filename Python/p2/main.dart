// int tambah(int a, int b) {
//     return a + b;
// }

// void add (int a, int b) {
//     int c = a + b;
//     print(c);
// }

// void main() {
//     // tambah(2, 3);
//     print(tambah(2, 3));
// }

void pelatihan([string? nama, int? umur, bool? is pelatihan]) {
    print('nama: $nama');
    print('umur: $umur');
    print('pelatihan: $ispelatihan');
}

void main() {
    pelatihan('dimas', 20, true);
    pelatihan('agus', 21);
    pelatihan('budi');

}