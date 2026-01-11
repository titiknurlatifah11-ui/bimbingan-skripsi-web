"use client";

import { useEffect, useState } from "react";
import { supabase } from "@/lib/supabase";
import SidebarDosen from "@/app/components/SidebarDosen";
import "./laporan.css";

export default function LaporanBimbingan() {
  const [dosen, setDosen] = useState<any>(null);
  const [dataBimbingan, setDataBimbingan] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  // Fungsi load data dipisah agar bisa dipanggil ulang setelah update status
  async function loadLaporan() {
    // 1. Ambil identitas (Cek Auth dulu, kalau kosong ambil dari LocalStorage)
    const { data: { user } } = await supabase.auth.getUser();
    const emailToSearch = user?.email || localStorage.getItem("userEmail");

    if (!emailToSearch) {
      setLoading(false);
      return;
    }

    // 2. Ambil Profil Dosen berdasarkan email
    const { data: profile } = await supabase
      .from("dosen")
      .select("*")
      .eq("email_dosen", emailToSearch)
      .single();

    if (profile) {
      setDosen(profile);
      
      // 3. Ambil data bimbingan mahasiswa milik dosen ini
      const { data: bimbingan } = await supabase
        .from("bimbingan") 
        .select(`
          id_bimbingan,
          tanggal_pertemuan,
          topik_bimbingan,
          status_bimbingan,
          mahasiswa (nama_mahasiswa, nim)
        `)
        .eq("id_dosen", profile.id_dosen)
        .order("id_bimbingan", { ascending: false });

      setDataBimbingan(bimbingan || []);
    }
    setLoading(false);
  }

  useEffect(() => {
    loadLaporan();
  }, []);

  // FUNGSI AKSI: Update status dan refresh data
  const updateStatus = async (idBimbingan: number, statusBaru: string) => {
    try {
      const { error } = await supabase
        .from("bimbingan") 
        .update({ status_bimbingan: statusBaru })
        .eq("id_bimbingan", idBimbingan);

      if (error) throw error;
      
      alert(`Berhasil: Data dipindahkan ke kolom ${statusBaru}`);
      loadLaporan(); // Panggil ulang data agar kartu pindah kolom otomatis
    } catch (err: any) {
      alert("Gagal memperbarui status: " + err.message);
    }
  };

  const filterByStatus = (status: string) => {
    return dataBimbingan.filter(item => item.status_bimbingan === status);
  };

  if (loading) return <div className="loading">Memuat Laporan...</div>;

  return (
    <div className="dashboard-container">
      {/* Sidebar mengambil nama dosen dari state */}
      <SidebarDosen nama={dosen?.nama_dosen || "Dosen"} />
      
      <div className="laporan-content">
        <div className="profile-header-banner">
          <h3>Rincian Laporan Bimbingan</h3>
          <p>Dosen: {dosen?.nama_dosen || "Memuat..."}</p>
        </div>
  
        <div className="laporan-grid">
          {/* Kolom 1: Menunggu Review */}
          <div className="laporan-column">
            <h3 className="title-waiting">⌛ Menunggu Review</h3>
            <div className="column-divider"></div>
            {filterByStatus("Menunggu").length > 0 ? filterByStatus("Menunggu").map((item, idx) => (
              <div key={idx} className="laporan-item-card">
                <p className="mhs-name">{item.mahasiswa?.nama_mahasiswa}</p>
                <p className="topic">{item.topik_bimbingan}</p>
                <span className="date">{item.tanggal_pertemuan || "Baru saja"}</span>
                
                <div className="action-buttons">
                  <button className="btn-approve" onClick={() => updateStatus(item.id_bimbingan, "Disetujui")}>
                    ✅ Setujui
                  </button>
                  <button className="btn-revise" onClick={() => updateStatus(item.id_bimbingan, "Revisi")}>
                    🔄 Revisi
                  </button>
                </div>
              </div>
            )) : <p className="empty-text">Tidak ada data</p>}
          </div>
  
          {/* Kolom 2: Disetujui */}
          <div className="laporan-column">
            <h3 className="title-approved">✅ File Disetujui</h3>
            <div className="column-divider"></div>
            {filterByStatus("Disetujui").length > 0 ? filterByStatus("Disetujui").map((item, idx) => (
              <div key={idx} className="laporan-item-card approved">
                <p className="mhs-name">{item.mahasiswa?.nama_mahasiswa}</p>
                <p className="topic">{item.topik_bimbingan}</p>
                <span className="date">{item.tanggal_pertemuan}</span>
              </div>
            )) : <p className="empty-text">Tidak ada data</p>}
          </div>
  
          {/* Kolom 3: Revisi */}
          <div className="laporan-column">
            <h3 className="title-revision">🔄 File Revisi</h3>
            <div className="column-divider"></div>
            {filterByStatus("Revisi").length > 0 ? filterByStatus("Revisi").map((item, idx) => (
              <div key={idx} className="laporan-item-card revision">
                <p className="mhs-name">{item.mahasiswa?.nama_mahasiswa}</p>
                <p className="topic">{item.topik_bimbingan}</p>
                <span className="date">{item.tanggal_pertemuan}</span>
              </div>
            )) : <p className="empty-text">Tidak ada data</p>}
          </div>
        </div>
      </div>
    </div>
  );
}