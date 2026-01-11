"use client";
import { useEffect, useState } from "react";
import { supabase } from "@/lib/supabase";
import SidebarAdmin from "@/app/components/SidebarAdmin";
import "./pembimbing.css";

export default function PenetapanPembimbingPage() {
  const [namaAdmin, setNamaAdmin] = useState("");
  const [listDosen, setListDosen] = useState<any[]>([]);
  const [listMahasiswa, setListMahasiswa] = useState<any[]>([]);
  
  // State untuk Input
  const [inputNIP, setInputNIP] = useState("");
  const [inputNIM, setInputNIM] = useState("");

  // State untuk Tampilan UI Hasil (Setelah di-input)
  const [displayData, setDisplayData] = useState<any>(null);
  const [mhsBimbingan, setMhsBimbingan] = useState<any[]>([]);

  useEffect(() => {
    fetchInitialData();
  }, []);

  async function fetchInitialData() {
    const { data: { user } } = await supabase.auth.getUser();
    if (user) {
      const { data } = await supabase.from("User").select("nama").eq("email", user.email).single();
      if (data) setNamaAdmin(data.nama);
    }
    // Ambil data dosen & mhs untuk keperluan dropdown/validasi
    const { data: dsn } = await supabase.from("dosen").select("*");
    const { data: mhs } = await supabase.from("mahasiswa").select("*");
    setListDosen(dsn || []);
    setListMahasiswa(mhs || []);
  }

  const handleProsesInput = async (e: React.FormEvent) => {
    e.preventDefault();
    
    // 1. Cari Dosen berdasarkan NIP
    const dsn = listDosen.find(d => d.nip === inputNIP);
    // 2. Cari Mahasiswa berdasarkan NIM
    const mhs = listMahasiswa.find(m => m.nim === inputNIM);

    if (!dsn || !mhs) return alert("NIP Dosen atau NIM Mahasiswa tidak ditemukan!");

    // 3. Update di Supabase (Hubungkan mhs ke dosen)
    const { error } = await supabase
      .from("mahasiswa")
      .update({ id_dosen_pembimbing: dsn.id_dosen })
      .eq("nim", inputNIM);

    if (error) return alert("Gagal update data: " + error.message);

    // 4. Set data ke tampilan UI seperti gambar
    setDisplayData(dsn);
    
    // Ambil semua mahasiswa yang dibimbing dosen tersebut untuk daftar di kanan
    const { data: allMhs } = await supabase
      .from("mahasiswa")
      .select("*")
      .eq("id_dosen_pembimbing", dsn.id_dosen);
    
    setMhsBimbingan(allMhs || []);
    alert("Berhasil menetapkan pembimbing!");
  };

  return (
    <div className="dashboard-container">
    <SidebarAdmin nama={namaAdmin} />

    <div className="admin-main-content">
        {/* Judul di Box Biru */}
        <div className="header-banner-container">
        <h2 className="title-penetapan">Penetapan Pembimbing</h2>
        </div>

        {/* Form Input */}
        <div className="input-pembimbing-card">
        <form onSubmit={handleProsesInput} className="input-form-inline">
            <div className="input-group">
            <label>NIP Dosen</label>
            <input type="text" placeholder="NIP Dosen" value={inputNIP} onChange={(e) => setInputNIP(e.target.value)} />
            </div>
            <div className="input-group">
            <label>NIM Mahasiswa</label>
            <input type="text" placeholder="NIM Mahasiswa" value={inputNIM} onChange={(e) => setInputNIM(e.target.value)} />
            </div>
            <button type="submit" className="proses-btn">Proses Data</button>
        </form>
        </div>

        {/* Hasil Data (Muncul di Bawah) */}
        {displayData && (
        <div className="pembimbing-grid animate-fade-in">
            <div className="card-pembimbing">
            <h3 className="card-title-text">Data Pembimbing</h3>
            <div className="info-box-container">
                <div className="detail-row">
                <span className="label">NAMA DOSEN</span>
                <span className="value">{displayData.nama_dosen}</span>
                </div>
                <div className="detail-row">
                <span className="label">NIDN</span>
                <span className="value">{displayData.nip}</span>
                </div>
                <div className="detail-row">
                <span className="label">STATUS</span>
                <span className="value status-active">Dosen Pembimbing Aktif</span>
                </div>
            </div>
            </div>

            <div className="card-pembimbing">
            <h3 className="card-title-text">Daftar Mahasiswa Bimbingan</h3>
            <div className="list-mahasiswa-scroll">
                {mhsBimbingan.map((m, idx) => (
                <div key={idx} className="mhs-item-row">
                    <div className="mhs-info-pill">{m.nama_mahasiswa}</div>
                    <button className="remove-btn">✖</button>
                </div>
                ))}
            </div>
            <button className="detail-selengkapnya-btn">Detail Selengkapnya</button>
            </div>
        </div>
        )}
    </div>
    </div>
  );
}