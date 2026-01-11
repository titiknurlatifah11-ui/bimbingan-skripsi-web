"use client";

export default function MahasiswaList({ mahasiswa }: { mahasiswa: any[] }) {
  return (
    <div className="mahasiswa-card">
      <h3>Mahasiswa</h3>
      <p>Anda memiliki {mahasiswa.length} mahasiswa bimbingan</p>

      {mahasiswa.map((mhs) => (
        <div className="mhs-row" key={mhs.id_mahasiswa}>
          <div className="mhs-info">
            <p className="nama">{mhs.nama_mahasiswa}</p>
            <p className="nim">{mhs.nim} · {mhs.angkatan}</p>
          </div>
          <button className="detail">detail</button>
        </div>
      ))}
    </div>
  );
}
