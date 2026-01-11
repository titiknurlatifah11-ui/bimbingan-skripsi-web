"use client";
import { useEffect, useState } from "react";
import { supabase } from "@/lib/supabase";
import SidebarMahasiswa from "@/app/components/SidebarMahasiswa";
import "./laporan-mhs.css";

export default function LaporanMahasiswa() {
  const [namaMhs, setNamaMhs] = useState("Loading...");
  const [riwayatBimbingan, setRiwayatBimbingan] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function loadLaporanData() {
      const { data: { user } } = await supabase.auth.getUser();
      if (!user) return;

      const { data: profile } = await supabase
        .from("mahasiswa")
        .select("nama_mahasiswa, id_mahasiswa")
        .eq("email_mahasiswa", user.email)
        .single();

      if (profile) {
        setNamaMhs(profile.nama_mahasiswa);

        const { data: bimbingan } = await supabase
          .from("bimbingan")
          .select("*")
          .eq("id_mahasiswa", profile.id_mahasiswa)
          .order("id_bimbingan", { ascending: false });

        if (bimbingan) setRiwayatBimbingan(bimbingan);
      }
      setLoading(false);
    }
    loadLaporanData();
  }, []);

  return (
    <div className="dashboard-container">
      <SidebarMahasiswa nama={namaMhs} role="Student" />

      <div className="dashboard-content">
        <div className="header-banner-mahasiswa">
          <h3>Riwayat Laporan Bimbingan</h3>
          <p>Daftar semua progres dan review dari dosen pembimbing</p>
        </div>

        <div className="laporan-grid">
          {/* Kolom Menunggu Review */}
          <div className="laporan-card">
            <h4 className="title-waiting">⌛ Menunggu Review</h4>
            <div className="list-container">
              {riwayatBimbingan.filter(b => b.status_bimbingan === "Menunggu").length > 0 ? (
                riwayatBimbingan.filter(b => b.status_bimbingan === "Menunggu").map((item) => (
                  <div key={item.id_bimbingan} className="item-bimbingan">
                    <div className="item-header">
                        <strong>{item.topik_bimbingan}</strong>
                        <span className="date-tag">
                            {item.tanggal_pertemuan ? new Date(item.tanggal_pertemuan).toLocaleDateString('id-ID') : 'No Date'}
                        </span>
                    </div>
                    <p>Status: Sedang diproses</p>
                  </div>
                ))
              ) : <p className="empty-state">Tidak ada data</p>}
            </div>
          </div>

          {/* Kolom File Disetujui */}
          <div className="laporan-card">
            <h4 className="title-approved">✅ File Disetujui</h4>
            <div className="list-container">
              {riwayatBimbingan.filter(b => b.status_bimbingan === "Disetujui").length > 0 ? (
                riwayatBimbingan.filter(b => b.status_bimbingan === "Disetujui").map((item) => (
                  <div key={item.id_bimbingan} className="item-bimbingan approved">
                    <div className="item-header">
                        <strong>{item.topik_bimbingan}</strong>
                        <span className="date-tag">
                            {item.tanggal_pertemuan ? new Date(item.tanggal_pertemuan).toLocaleDateString('id-ID') : 'No Date'}
                        </span>
                    </div>
                    {item.review_bimbingan && (
                      <a href={item.review_bimbingan} target="_blank" className="btn-download-small">📂 Lihat File</a>
                    )}
                  </div>
                ))
              ) : <p className="empty-state">Tidak ada data</p>}
            </div>
          </div>

          {/* Kolom File Revisi */}
          <div className="laporan-card">
            <h4 className="title-revision">🔄 File Revisi</h4>
            <div className="list-container">
              {riwayatBimbingan.filter(b => b.status_bimbingan === "Revisi").length > 0 ? (
                riwayatBimbingan.filter(b => b.status_bimbingan === "Revisi").map((item) => (
                  <div key={item.id_bimbingan} className="item-bimbingan revision">
                    <div className="item-header">
                        <strong>{item.topik_bimbingan}</strong>
                        <span className="date-tag">
                            {item.tanggal_pertemuan ? new Date(item.tanggal_pertemuan).toLocaleDateString('id-ID') : 'No Date'}
                        </span>
                    </div>
                    {item.review_bimbingan && (
                      <a href={item.review_bimbingan} target="_blank" className="btn-download-small">📂 Download Review</a>
                    )}
                  </div>
                ))
              ) : <p className="empty-state">Tidak ada data</p>}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}