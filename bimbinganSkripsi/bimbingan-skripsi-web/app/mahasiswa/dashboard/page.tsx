"use client";
import { useEffect, useState } from "react";
import { supabase } from "@/lib/supabase";
import SidebarMahasiswa from "@/app/components/SidebarMahasiswa";
import "./dashboard-mhs.css";

export default function DashboardMahasiswa() {
  const [namaMhs, setNamaMhs] = useState("Loading...");
  const [jadwal, setJadwal] = useState<any>(null); // State untuk jadwal
  const [stats, setStats] = useState({
    disetujui: 0,
    revisi: 0,
    persentase: 0
  });

  useEffect(() => {
    async function loadDashboardData() {
      // 1. Ambil session user dari Supabase Auth
      const { data: { user } } = await supabase.auth.getUser();
      
      // 2. Ambil email dari Auth atau dari LocalStorage (untuk login manual)
      const emailToSearch = user?.email || localStorage.getItem("userEmail");

      // Jika tidak ada identitas email, hentikan proses (atau arahkan ke login)
      if (!emailToSearch) {
        console.error("Email tidak ditemukan");
        return;
      }

      // 3. Ambil profil mahasiswa berdasarkan email yang ditemukan
      const { data: profile } = await supabase
        .from("mahasiswa")
        .select("nama_mahasiswa, id_mahasiswa")
        .eq("email_mahasiswa", emailToSearch) // Menggunakan variabel emailToSearch
        .single();

      if (profile) {
        setNamaMhs(profile.nama_mahasiswa);

        // 4. Ambil Data Bimbingan & Jadwal berdasarkan ID Mahasiswa
        const { data: bimbingan } = await supabase
          .from("bimbingan")
          .select("*")
          .eq("id_mahasiswa", profile.id_mahasiswa);

        if (bimbingan) {
          const jmlDisetujui = bimbingan.filter(b => b.status_bimbingan === "Disetujui").length;
          const jmlRevisi = bimbingan.filter(b => b.status_bimbingan === "Revisi").length;
          
          // Ambil jadwal pertemuan terbaru (yang ada tanggalnya)
          const pertemuanTerdekat = bimbingan
            .filter(b => b.tanggal_pertemuan !== null)
            .sort((a, b) => new Date(b.tanggal_pertemuan).getTime() - new Date(a.tanggal_pertemuan).getTime())[0];

          setJadwal(pertemuanTerdekat);

          const targetBimbingan = 8; 
          const kalkulasiPersen = Math.min(Math.round((jmlDisetujui / targetBimbingan) * 100), 100);

          setStats({
            disetujui: jmlDisetujui,
            revisi: jmlRevisi,
            persentase: kalkulasiPersen
          });
        }
      } else {
        console.error("Profil mahasiswa tidak ditemukan di database");
      }
    }
    loadDashboardData();
  }, []);

  return (
    <div className="dashboard-container">
      <SidebarMahasiswa nama={namaMhs} role="Student" />

      <div className="dashboard-content">
        <div className="welcome-banner">
          <h3>Welcome back, {namaMhs}</h3>
          <p>Berikut ringkasan aktivitas bimbingan Anda hari ini</p>
        </div>

        <div className="dashboard-grid-layout">
          {/* SISI KIRI: Status Bimbingan */}
          <div className="main-stats-card">
            <h4 className="stats-title">Status Bimbingan</h4>
            <div className="stats-main-row">
              <div className="chart-container">
                <div 
                  className="circle-chart" 
                  style={{ background: `conic-gradient(#4e73df ${stats.persentase * 3.6}deg, #e6e6e6 0deg)` }}
                >
                  <div className="inner-circle">
                    <span style={{fontSize: '24px'}}>{stats.persentase === 100 ? "🎓" : "📚"}</span>
                  </div>
                </div>
              </div>
              <div className="percentage-badge">
                <span>{stats.persentase}%</span>
              </div>
            </div>
            
            <div className="stats-divider"></div>
            
            <div className="stats-grid-footer">
              <div className="footer-item" onClick={() => window.location.href='/mahasiswa/laporan'}>
                <div className="item-label">☑ File disetujui</div>
                <div className="white-pill">{stats.disetujui} file</div>
              </div>
              <div className="footer-item" onClick={() => window.location.href='/mahasiswa/laporan'}>
                <div className="item-label">🔄 File revisi</div>
                <div className="white-pill">{stats.revisi} file</div>
              </div>
            </div>
          </div>

          {/* SISI KANAN: Jadwal Pertemuan (Kalender Sederhana) */}
          <div className="calendar-card">
            <h4 className="stats-title">Jadwal Pertemuan</h4>
            {jadwal ? (
              <div className="event-box">
                <div className="event-date">
                  <span className="day">{new Date(jadwal.tanggal_pertemuan).getDate()}</span>
                  <span className="month">{new Date(jadwal.tanggal_pertemuan).toLocaleString('default', { month: 'short' })}</span>
                </div>
                <div className="event-info">
                  <p className="event-topic">{jadwal.topik_bimbingan}</p>
                  <p className="event-time">🕒 {jadwal.jam_pertemuan}</p>
                </div>
              </div>
            ) : (
              <p className="no-data">Belum ada jadwal pertemuan.</p>
            )}
          </div>
        </div>

        <div className="bottom-grid-full">
          <div className="bottom-card-full">
            <h4 className="tindakan-title">Tindakan Cepat</h4>
            <div className="tindakan-buttons-row">
              <button className="btn-upload-mhs" onClick={() => window.location.href='/mahasiswa/bimbingan'}>
                📤 Unggah File Skripsi
              </button>
              <button className="btn-laporan-mhs" onClick={() => window.location.href='/mahasiswa/laporan'}>
                📄 Lihat Riwayat Laporan
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}