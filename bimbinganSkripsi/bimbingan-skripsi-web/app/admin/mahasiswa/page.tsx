"use client";
import { useEffect, useState, useRef } from "react";
import { supabase } from "@/lib/supabase";
import SidebarAdmin from "@/app/components/SidebarAdmin";
import "./mahasiswa.css"; 
import "../admin.css";

export default function KelolaMahasiswaPage() {
  const [listMahasiswa, setListMahasiswa] = useState<any[]>([]);
  const [selectedMhsId, setSelectedMhsId] = useState("");
  const [namaAdmin, setNamaAdmin] = useState("");
  const [loading, setLoading] = useState(false);
  const [uploading, setUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [formData, setFormData] = useState({
    nama_mahasiswa: "",
    nim: "",
    email_mahasiswa: "",
    angkatan: new Date().getFullYear(),
    avatar_mahasiswa: "" 
  });

  useEffect(() => {
    fetchInitialData();
  }, []);

  async function fetchInitialData() {
    const { data: { user } } = await supabase.auth.getUser();
    if (user) {
      const { data: admin } = await supabase.from("User").select("nama").eq("email", user.email).single();
      if (admin) setNamaAdmin(admin.nama);
    }
    const { data: mhs } = await supabase.from("mahasiswa").select("*");
    if (mhs) setListMahasiswa(mhs);
  }

  const handlePrepareNew = () => {
    setSelectedMhsId("");
    setFormData({
      nama_mahasiswa: "",
      nim: "",
      email_mahasiswa: "",
      angkatan: new Date().getFullYear(),
      avatar_mahasiswa: ""
    });
  };

  const handleSelectMhs = (id: string) => {
    const mhs = listMahasiswa.find(m => m.id_mahasiswa == id);
    if (mhs) {
      setSelectedMhsId(id);
      setFormData({
        nama_mahasiswa: mhs.nama_mahasiswa || "",
        nim: mhs.nim || "",
        email_mahasiswa: mhs.email_mahasiswa || "",
        angkatan: mhs.angkatan || "",
        avatar_mahasiswa: mhs.avatar_mahasiswa || "",
      });
    }
  };

  const handleUploadAvatar = async (event: React.ChangeEvent<HTMLInputElement>) => {
    try {
      const file = event.target.files?.[0];
      if (!file) return;
      setUploading(true);
      
      const fileExt = file.name.split('.').pop();
      const fileName = `mhs-${Date.now()}.${fileExt}`;
      const filePath = `avatars/${fileName}`;

      const { error: uploadError } = await supabase.storage
        .from('bimbingan_files')
        .upload(filePath, file, { upsert: true });

      if (uploadError) throw uploadError;

      const { data } = supabase.storage.from('bimbingan_files').getPublicUrl(filePath);

      setFormData(prev => ({ ...prev, avatar_mahasiswa: data.publicUrl }));
      alert("Foto berhasil dipilih! Klik Save untuk menyimpan.");
    } catch (error: any) {
      alert("Gagal upload: " + error.message);
    } finally {
      setUploading(false);
    }
  };

  // --- FUNGSI HANDLE SAVE DENGAN SINKRONISASI DUA TABEL ---
  const handleSave = async () => {
    if (!formData.nama_mahasiswa || !formData.nim || !formData.email_mahasiswa) {
      return alert("Mohon lengkapi Nama, NIM, dan Email!");
    }
    
    setLoading(true);
    try {
      if (selectedMhsId) {
        // --- 1. Ambil email lama mahasiswa sebelum diupdate ---
        const { data: oldMhsData } = await supabase
          .from("mahasiswa")
          .select("email_mahasiswa")
          .eq("id_mahasiswa", selectedMhsId)
          .single();

        // --- 2. Update tabel 'mahasiswa' ---
        const { error: errorMhs } = await supabase
          .from("mahasiswa")
          .update({
            nama_mahasiswa: formData.nama_mahasiswa,
            nim: formData.nim,
            email_mahasiswa: formData.email_mahasiswa,
            angkatan: formData.angkatan,
            avatar_mahasiswa: formData.avatar_mahasiswa
          })
          .eq("id_mahasiswa", selectedMhsId);
        
        if (errorMhs) throw errorMhs;

        // --- 3. Update juga tabel 'User' agar akun login sinkron ---
        if (oldMhsData) {
          const { error: errorUser } = await supabase
            .from("User")
            .update({
              nama: formData.nama_mahasiswa,
              email: formData.email_mahasiswa, // Update email baru
              password: formData.nim // Update password jadi NIM baru
            })
            .ilike("email", oldMhsData.email_mahasiswa); // Cari berdasarkan email lama

          if (errorUser) throw errorUser;
        }

        alert("Data Mahasiswa dan Akun Login berhasil diperbarui!");
      } else {
        // --- LOGIKA INSERT BARU ---
        const { data: existingUser } = await supabase.from("User").select("email").eq("email", formData.email_mahasiswa).single();
        if (existingUser) throw new Error("Email sudah terdaftar!");

        const { data: newUser, error: userError } = await supabase
          .from("User")
          .insert([{ 
            nama: formData.nama_mahasiswa, 
            email: formData.email_mahasiswa, 
            password: formData.nim, 
            role: 'mahasiswa' 
          }])
          .select('id_user')
          .single();

        if (userError) throw userError;

        const { error: mhsError } = await supabase
          .from("mahasiswa")
          .insert([{
            id_mahasiswa: newUser.id_user, 
            ...formData
          }]);

        if (mhsError) {
          await supabase.from("User").delete().eq("id_user", newUser.id_user);
          throw mhsError;
        }
        alert("Mahasiswa baru berhasil ditambahkan!");
      }
      fetchInitialData();
      handlePrepareNew();
    } catch (err: any) {
      alert("Kesalahan: " + err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!selectedMhsId) return alert("Pilih mahasiswa!");
    if (!confirm("Hapus mahasiswa ini? Akun login di tabel User juga akan ikut terhapus.")) return;

    // Ambil info email untuk hapus user terkait
    const mhsToDelete = listMahasiswa.find(m => m.id_mahasiswa == selectedMhsId);

    const { error } = await supabase.from("mahasiswa").delete().eq("id_mahasiswa", selectedMhsId);
    if (!error) {
      // Hapus di tabel User berdasarkan email
      await supabase.from("User").delete().eq("email", mhsToDelete.email_mahasiswa);
      alert("Mahasiswa dan akun berhasil dihapus");
      fetchInitialData();
      handlePrepareNew();
    }
  };

  return (
    <div className="dashboard-container">
      <SidebarAdmin nama={namaAdmin || "Admin"} />

      <div className="admin-main-content">
        <div className="profile-edit-container">
          <div className="profile-header-bar">
            <span>Profil Mahasiswa</span>
            <div className="header-icons">
              <button title="Tambah Baru" onClick={handlePrepareNew}>➕</button>
              <button title="Hapus" onClick={handleDelete}>🗑️</button>
            </div>
          </div>

          <div className="profile-content-grid">
            <div className="profile-form-section">
              <div className="form-group-custom">
                <label>Pilih Mahasiswa (Edit)</label>
                <select onChange={(e) => handleSelectMhs(e.target.value)} value={selectedMhsId}>
                  <option value="">-- Mode Input Baru --</option>
                  {listMahasiswa.map(m => (
                    <option key={m.id_mahasiswa} value={m.id_mahasiswa}>{m.nama_mahasiswa}</option>
                  ))}
                </select>
              </div>

              <div className="form-group-custom">
                <label>Nama Mahasiswa</label>
                <input type="text" value={formData.nama_mahasiswa} onChange={(e) => setFormData({...formData, nama_mahasiswa: e.target.value})} />
              </div>

              <div className="form-group-custom">
                <label>NIM</label>
                <input type="text" value={formData.nim} onChange={(e) => setFormData({...formData, nim: e.target.value})} />
              </div>

              <div className="form-group-custom">
                <label>e-mail</label>
                <input type="email" value={formData.email_mahasiswa} onChange={(e) => setFormData({...formData, email_mahasiswa: e.target.value})} />
              </div>

              <div className="form-group-custom">
                <label>Angkatan</label>
                <input type="number" value={formData.angkatan} onChange={(e) => setFormData({...formData, angkatan: parseInt(e.target.value)})} />
              </div>
            </div>

            <div className="profile-right-section">
              <div className="profile-avatar-placeholder">
                {formData.avatar_mahasiswa ? (
                  <img src={formData.avatar_mahasiswa} alt="Profile" className="avatar-img" />
                ) : (
                  <div className="avatar-icon">👤</div>
                )}
              </div>
              
              <input type="file" ref={fileInputRef} onChange={handleUploadAvatar} accept="image/*" style={{ display: 'none' }} />
              
              <button className="edit-btn-text" onClick={() => fileInputRef.current?.click()} disabled={uploading}>
                {uploading ? "Mengunggah..." : "Ganti Foto"}
              </button>
              
              <button className="save-btn-large" onClick={handleSave} disabled={loading}>
                {loading ? "Memproses..." : selectedMhsId ? "Update Data" : "Simpan Mahasiswa Baru"}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}