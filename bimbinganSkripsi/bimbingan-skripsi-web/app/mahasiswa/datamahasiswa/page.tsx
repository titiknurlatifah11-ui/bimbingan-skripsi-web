"use client";
import { useEffect, useState, useRef } from "react";
import { supabase } from "@/lib/supabase";
import SidebarMahasiswa from "@/app/components/SidebarMahasiswa";
import "./profile-mhs.css";

export default function ProfileMahasiswaPage() {
  const [mhs, setMhs] = useState<any>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [loading, setLoading] = useState(true);
  
  const [formData, setFormData] = useState({
    nama_mahasiswa: "",
    nim: "",
    judul: "",
    angkatan: ""
  });
  
  const [previewPhoto, setPreviewPhoto] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    async function getProfile() {
      // 1. Ambil identitas dari Auth atau LocalStorage
      const { data: { user } } = await supabase.auth.getUser();
      const emailToSearch = user?.email || localStorage.getItem("userEmail");

      if (!emailToSearch) {
        setLoading(false);
        return;
      }

      // 2. Query ke tabel mahasiswa
      const { data, error } = await supabase
        .from("mahasiswa")
        .select("*, dosen:id_dosen_pembimbing(nama_dosen)") 
        .eq("email_mahasiswa", emailToSearch)
        .single();

      if (data) {
        setMhs(data);
        setPreviewPhoto(data.foto_url || null);
        setFormData({
          nama_mahasiswa: data.nama_mahasiswa || "",
          nim: data.nim || "",
          judul: data.judul || "",
          angkatan: data.angkatan || ""
        });
      }
      setLoading(false);
    }
    getProfile();
  }, []);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleAddPhoto = () => {
    if (isEditing) fileInputRef.current?.click();
    else alert("Klik ikon pensil terlebih dahulu untuk mengedit!");
  };

  const onFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => setPreviewPhoto(reader.result as string);
      reader.readAsDataURL(file);
    }
  };

  const handleDeletePhoto = () => {
    if (isEditing) setPreviewPhoto(null);
    else alert("Klik ikon pensil terlebih dahulu untuk mengedit!");
  };

  const handleSave = async () => {
    try {
      // 1. Update tabel 'mahasiswa'
      const { error: errorMhs } = await supabase
        .from("mahasiswa")
        .update({ 
          nama_mahasiswa: formData.nama_mahasiswa,
          nim: formData.nim,
          judul: formData.judul,
          angkatan: formData.angkatan
        })
        .eq("id_mahasiswa", mhs.id_mahasiswa);

      if (errorMhs) throw errorMhs;

      // 2. Update tabel 'User' agar nama di akun login sinkron
      const { error: errorUser } = await supabase
        .from("User")
        .update({
          nama: formData.nama_mahasiswa,
          password: formData.nim // Opsional: Update password jika NIM berubah
        })
        .eq("email", mhs.email_mahasiswa);

      if (errorUser) throw errorUser;
      
      setMhs({ ...mhs, ...formData });
      setIsEditing(false);
      alert("Data berhasil diperbarui!");
    } catch (err: any) {
      alert("Gagal menyimpan: " + err.message);
    }
  };

  if (loading) return <div className="p-10">Memuat Profil...</div>;

  return (
    <div className="dashboard-container">
      <SidebarMahasiswa nama={mhs?.nama_mahasiswa || "Loading..."} />
      
      <div className="profile-content">
        <div className="header-wrapper">
          <div className="title-banner">Profil Mahasiswa</div>
        </div>

        <div className="profile-main-card">
          <div className="profile-form">
            
            {/* INPUT NAMA */}
            <div className="input-group">
              <label>Mahasiswa</label>
              {isEditing ? (
                <input 
                  name="nama_mahasiswa"
                  className="edit-input"
                  value={formData.nama_mahasiswa}
                  onChange={handleInputChange}
                />
              ) : (
                <div className="display-value">{mhs?.nama_mahasiswa || "-"}</div>
              )}
            </div>

            {/* INPUT NIM */}
            <div className="input-group">
              <label>NIM</label>
              {isEditing ? (
                <input 
                  name="nim"
                  className="edit-input"
                  value={formData.nim}
                  onChange={handleInputChange}
                />
              ) : (
                <div className="display-value">{mhs?.nim || "-"}</div>
              )}
            </div>

            {/* PEMBIMBING (Tetap Read-Only karena ditentukan sistem) */}
            <div className="input-group">
              <label>Pembimbing</label>
              <div className="display-value disabled">{mhs?.dosen?.nama_dosen || "-"}</div>
            </div>

            {/* INPUT JUDUL */}
            <div className="input-group">
              <label>Judul</label>
              {isEditing ? (
                <input 
                  name="judul"
                  className="edit-input"
                  value={formData.judul}
                  onChange={handleInputChange}
                  placeholder="Masukkan judul skripsi..."
                />
              ) : (
                <div className="display-value">{mhs?.judul || "Belum ada judul"}</div>
              )}
            </div>

            {/* INPUT ANGKATAN / TAHUN AKADEMIK */}
            <div className="input-group">
              <label>Tahun Akademik</label>
              {isEditing ? (
                <input 
                  name="angkatan"
                  className="edit-input"
                  value={formData.angkatan}
                  onChange={handleInputChange}
                />
              ) : (
                <div className="display-value">{mhs?.angkatan || "-"}</div>
              )}
            </div>
          </div>

          <div className="profile-right-section">
            <div className="photo-box">
               {previewPhoto ? (
                 <img src={previewPhoto} alt="Profile" className="profile-img-preview" />
               ) : (
                 <span className="user-icon-large">👤</span>
               )}
            </div>
            
            <input 
              type="file" 
              hidden 
              ref={fileInputRef} 
              accept="image/*" 
              onChange={onFileChange} 
            />

            <div className="action-icons-vertical">
              <button className="icon-btn" onClick={handleAddPhoto} title="Tambah Foto">➕</button>
              <button 
                className={`icon-btn ${isEditing ? "active-edit" : ""}`} 
                onClick={() => setIsEditing(!isEditing)}
                title="Edit Data"
              >
                ✏️
              </button>
              <button className="icon-btn" onClick={handleDeletePhoto} title="Hapus Foto">🗑️</button>
            </div>
            
            <div className="save-container">
                <button className="save-btn" onClick={handleSave} disabled={!isEditing}>
                  Save
                </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}