"use client";
import { useEffect, useState } from "react";
import { supabase } from "@/lib/supabase";
import SidebarMahasiswa from "@/app/components/SidebarMahasiswa";
import "./dosen.css";

export default function DataDosenMahasiswa() {
  const [dosen, setDosen] = useState<any>(null);
  const [namaMhs, setNamaMhs] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchDosenPembimbing() {
      // 1. Dapatkan identitas dari Auth atau LocalStorage
      const { data: { user } } = await supabase.auth.getUser();
      const emailToSearch = user?.email || localStorage.getItem("userEmail");

      if (!emailToSearch) {
        setLoading(false);
        return;
      }

      // 2. Ambil data mahasiswa berdasarkan email
      const { data: mhs, error: mhsError } = await supabase
        .from("mahasiswa")
        .select("nama_mahasiswa, id_dosen_pembimbing")
        .eq("email_mahasiswa", emailToSearch)
        .single();

      if (mhsError) {
        console.error("Gagal mengambil data mahasiswa:", mhsError.message);
        setLoading(false);
        return;
      }

      if (mhs) {
        setNamaMhs(mhs.nama_mahasiswa);
        
        // 3. Gunakan id_dosen_pembimbing untuk mencari data di tabel dosen
        if (mhs.id_dosen_pembimbing) {
          const { data: dsn, error: dsnError } = await supabase
            .from("dosen")
            .select("*")
            .eq("id_dosen", mhs.id_dosen_pembimbing)
            .single();
          
          if (dsnError) {
            console.error("Gagal mengambil data dosen:", dsnError.message);
          } else if (dsn) {
            setDosen(dsn);
          }
        }
      }
      setLoading(false);
    }
    fetchDosenPembimbing();
  }, []);

  if (loading) return <div className="loading-screen">Memuat Profil Dosen...</div>;

  return (
    <div className="dashboard-container">
      {/* Sidebar menampilkan nama mahasiswa */}
      <SidebarMahasiswa nama={namaMhs || "Mahasiswa"} role="Student" />
      
      <div className="dosen-content">
        <div className="profile-header-banner">
          <h3>Profil Dosen Pembimbing</h3>
          <p>Informasi detail mengenai dosen pembimbing skripsi Anda</p>
        </div>

        <div className="profile-card-main">
          <div className="profile-grid">
            <div className="profile-info-column">
              <div className="info-item">
                <label>Nama Lengkap</label>
                <p>{dosen?.nama_dosen || "Belum ditentukan"}</p>
              </div>
              <div className="info-item">
                <label>NIP / NIDN</label>
                <p>{dosen?.nip || "-"}</p>
              </div>
              <div className="info-item">
                <label>E-mail</label>
                <p>{dosen?.email_dosen || "-"}</p>
              </div>
              <div className="info-item">
                <label>Status</label>
                {/* Tampilkan badge status dinamis jika ada di DB */}
                <span className={dosen?.status === "Aktif" ? "status-badge-active" : "status-badge-inactive"}>
                  {dosen?.status || "Aktif"}
                </span>
              </div>
            </div>

            <div className="profile-photo-column">
              <div className="photo-wrapper">
                {dosen?.avatar_dosen ? (
                  <img 
                    src={dosen.avatar_dosen} 
                    alt="Foto Dosen" 
                    style={{ width: '100%', height: '100%', borderRadius: '15px', objectFit: 'cover' }} 
                  />
                ) : (
                  <div style={{ fontSize: "100px", textAlign: "center" }}>👤</div>
                )}
              </div>
              <p className="photo-label">Dosen Pembimbing Utama</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}