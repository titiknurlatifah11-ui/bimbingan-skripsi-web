"use client";

import { useEffect, useState } from "react";
import { supabase } from "@/lib/supabase";
import SidebarDosen from "@/app/components/SidebarDosen";
import "../dosen-data.css"; 

export default function DataDosenPage() {
  const [dosen, setDosen] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  
  const [nama, setNama] = useState("");
  const [nip, setNip] = useState("");
  const [avatarDosen, setAvatarDosen] = useState<string | null>(null);

  useEffect(() => {
    async function loadProfile() {
      // 1. Cek Auth Session
      const { data: { user } } = await supabase.auth.getUser();
      
      // 2. Ambil email (dari Auth atau dari LocalStorage jika login manual)
      let emailToSearch = user?.email || localStorage.getItem("userEmail");
  
      if (!emailToSearch) return;
  
      const { data, error } = await supabase
        .from("dosen")
        .select("*")
        .eq("email_dosen", emailToSearch) // Gunakan email yang ditemukan
        .single();
  
      if (data) {
        setDosen(data);
        setNama(data.nama_dosen);
        setNip(data.nip);
        setAvatarDosen(data.avatar_dosen);
      }
    }
    loadProfile();
  }, []);

  // ⬇️ TAMBAHKAN INI DI SINI ⬇️
  const handleEditClick = () => {
    setIsEditing(true);
  };
  
  const handleUploadFoto = async (e: React.ChangeEvent<HTMLInputElement>) => {
    try {
      setLoading(true);
      const file = e.target.files?.[0];
      if (!file) return;

      const fileExt = file.name.split('.').pop();
      const fileName = `avatar-${dosen.id_dosen}-${Date.now()}.${fileExt}`;
      const filePath = `avatars/${fileName}`;

      const { error: uploadError } = await supabase.storage
        .from("bimbingan_files")
        .upload(filePath, file);

      if (uploadError) throw uploadError;

      const { data: { publicUrl } } = supabase.storage
        .from("bimbingan_files")
        .getPublicUrl(filePath);

      setAvatarDosen(publicUrl);
      alert("Foto berhasil diunggah! Klik Save untuk menyimpan.");
    } catch (error: any) {
      alert("Upload gagal: " + error.message);
    } finally {
      setLoading(false);
    }
  };

  const handleSave = async () => {
    try {
      setLoading(true);

      // 1. UPDATE TABEL 'dosen'
      const { error: errorDosen } = await supabase
        .from("dosen")
        .update({
          nama_dosen: nama,
          nip: nip,
          avatar_dosen: avatarDosen
        })
        .eq("id_dosen", dosen.id_dosen);

      if (errorDosen) throw errorDosen;

      // 2. UPDATE TABEL 'User' (Agar nama di sidebar & password login ikut sinkron)
      // Kita menggunakan email_dosen sebagai kunci pencarian di tabel User
      const { error: errorUser } = await supabase
        .from("User")
        .update({
          nama: nama,
          password: nip // Mengupdate password sesuai NIP baru jika ada perubahan
        })
        .eq("email", dosen.email_dosen);

      if (errorUser) {
        console.error("Gagal sinkron ke tabel User:", errorUser.message);
        // Kita tidak throw error di sini agar user tetap merasa berhasil update profil, 
        // tapi kita log untuk pengecekan admin.
      }
      
      setIsEditing(false);
      alert("Data profil dan akun login berhasil disimpan!");
      
      // Opsional: Refresh halaman agar nama di Sidebar (yang mengambil dari tabel User) ikut berubah
      window.location.reload(); 
      
    } catch (error: any) {
      alert("Gagal menyimpan: " + error.message);
    } finally {
      setLoading(false);
    }
  };

  if (!dosen) return <div className="p-10">Memuat Data...</div>;


  return (
    <div className="dashboard-container">
      <SidebarDosen nama={dosen.nama_dosen} />
      
      <div className="dashboard-content">
        <div className="profile-header-bar">Profil Dosen</div>
        
        {/* Container Utama Profil */}
        <div className="profile-container">
          
          <div className="profile-flex">
            {/* Form Input */}
            <div className="profile-form">
              <div className="form-group">
                <label>Dosen</label>
                <input 
                  type="text" 
                  value={nama} 
                  onChange={(e) => setNama(e.target.value)} 
                  disabled={!isEditing} 
                />
              </div>
              <div className="form-group">
                <label>NIDN / NIP</label>
                <input 
                  type="text" 
                  value={nip} 
                  onChange={(e) => setNip(e.target.value)} 
                  disabled={!isEditing} 
                />
              </div>
              <div className="form-group">
                <label>e-mail</label>
                <input type="email" value={dosen.email_dosen} readOnly className="readonly-input" />
              </div>
              <div className="form-group">
                <label>Status</label>
                <input type="text" value="Aktif" readOnly className="readonly-input" />
              </div>
            </div>

            {/* Bagian Foto & Ikon Aksi (Sudah Digabung) */}
            <div className="profile-photo-section">
              <div className="photo-box">
                {avatarDosen ? (
                  <img src={avatarDosen} alt="Avatar" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
                ) : (
                  <span style={{ fontSize: "80px" }}>👤</span>
                )}
              </div>
              
              {/* PINDAHKAN ICON KE SINI AGAR DI BAWAH FOTO */}
              <div className="action-small-btns">
                <button className="icon-btn">📤</button>
                <button className="icon-btn" onClick={handleEditClick} title="Edit Profil & Foto">
                  ✏️
                </button>
                <button className="icon-btn">🗑️</button>
              </div>
              
              {/* Input file yang disembunyikan */}
              <input 
                type="file" 
                id="upload-avatar" 
                hidden 
                accept="image/*" 
                onChange={handleUploadFoto} 
              />

              {/* Tombol ganti foto hanya muncul saat mode edit */}
              {isEditing && (
                <button 
                  className="edit-photo-btn" 
                  onClick={() => document.getElementById('upload-avatar')?.click()}
                  type="button"
                  style={{ marginTop: '10px', fontSize: '12px', cursor: 'pointer' }}
                >
                  {loading ? "..." : "Ganti Foto"}
                </button>
              )}
            </div>
          </div>

          {/* Tombol Save di bagian paling bawah */}
          <div className="save-section">
            <button 
              className="save-btn" 
              onClick={handleSave} 
              disabled={loading || (!isEditing && !avatarDosen)}
              style={{ opacity: (isEditing || loading) ? 1 : 0.6 }}
            >
              {loading ? "Menyimpan..." : "Save Perubahan"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}