"use client";
import { useEffect, useState, useRef } from "react";
import { supabase } from "@/lib/supabase";
import SidebarAdmin from "@/app/components/SidebarAdmin";
import "./dosen.css"; 
import "../admin.css";

export default function KelolaDosenPage() {
  const [listDosen, setListDosen] = useState<any[]>([]);
  const [selectedDosenId, setSelectedDosenId] = useState("");
  const [namaAdmin, setNamaAdmin] = useState("");
  const [loading, setLoading] = useState(false);
  const [uploading, setUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [formData, setFormData] = useState({
    nama_dosen: "",
    nip: "",
    email_dosen: "",
    status: "Aktif",
    avatar_dosen: "" 
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
    const { data: dosen } = await supabase.from("dosen").select("*");
    if (dosen) setListDosen(dosen);
  }

  const handlePrepareNew = () => {
    setSelectedDosenId("");
    setFormData({
      nama_dosen: "",
      nip: "",
      email_dosen: "",
      status: "Aktif",
      avatar_dosen: ""
    });
  };

  const handleSelectDosen = (id: string) => {
    const dosen = listDosen.find(d => d.id_dosen == id);
    if (dosen) {
      setSelectedDosenId(id);
      setFormData({
        nama_dosen: dosen.nama_dosen || "",
        nip: dosen.nip || "",
        email_dosen: dosen.email_dosen || "",
        status: dosen.status || "Aktif",
        avatar_dosen: dosen.avatar_dosen || "",
      });
    }
  };

  const handleUploadAvatar = async (event: React.ChangeEvent<HTMLInputElement>) => {
    try {
      const file = event.target.files?.[0];
      if (!file) return;
      setUploading(true);
      const fileExt = file.name.split('.').pop();
      const fileName = `${Date.now()}.${fileExt}`;
      const filePath = `avatars/${fileName}`;
      const { error: uploadError } = await supabase.storage
        .from('bimbingan_files')
        .upload(filePath, file, { upsert: true });
      if (uploadError) throw uploadError;
      const { data } = supabase.storage.from('bimbingan_files').getPublicUrl(filePath);
      setFormData(prev => ({ ...prev, avatar_dosen: data.publicUrl }));
      alert("Foto berhasil dipilih! Jangan lupa klik 'Save' untuk menyimpan permanen.");
    } catch (error: any) {
      alert("Gagal memproses foto: " + error.message);
    } finally {
      setUploading(false);
    }
  };

  // --- FUNGSI HANDLE SAVE YANG SUDAH DIPERBAIKI (UPDATE DUA TABEL) ---
  const handleSave = async () => {
    if (!formData.nama_dosen || !formData.nip || !formData.email_dosen) {
      return alert("Mohon lengkapi Nama, NIP, dan Email!");
    }
    
    setLoading(true);

    try {
      if (selectedDosenId) {
        // --- 1. Ambil email lama dosen sebelum diupdate ---
        // Ini penting sebagai kunci untuk mencari data yang benar di tabel User
        const { data: oldDosenData } = await supabase
          .from("dosen")
          .select("email_dosen")
          .eq("id_dosen", selectedDosenId)
          .single();

        // --- 2. Update tabel 'dosen' ---
        const { error: errorDosen } = await supabase
          .from("dosen")
          .update({
            nama_dosen: formData.nama_dosen,
            nip: formData.nip,
            email_dosen: formData.email_dosen,
            status: formData.status,
            avatar_dosen: formData.avatar_dosen
          })
          .eq("id_dosen", selectedDosenId);
        
        if (errorDosen) throw errorDosen;

        // --- 3. Update juga tabel 'User' agar email login sinkron ---
        if (oldDosenData) {
          const { error: errorUser } = await supabase
            .from("User")
            .update({
              nama: formData.nama_dosen,
              email: formData.email_dosen, // Mengupdate ke email baru
              password: formData.nip // Mengupdate password jika NIP berubah
            })
            .ilike("email", oldDosenData.email_dosen); // Cari berdasarkan email lama

          if (errorUser) throw errorUser;
        }

        alert("Data Dosen dan Akun Login berhasil diperbarui!");

      } else {
        // --- LOGIKA INPUT DOSEN BARU (INSERT) ---
        const { data: existingUser } = await supabase
          .from("User")
          .select("email")
          .eq("email", formData.email_dosen)
          .single();

        if (existingUser) throw new Error("Email sudah terdaftar!");

        const { data: newUser, error: userError } = await supabase
          .from("User")
          .insert([{ 
            nama: formData.nama_dosen, 
            email: formData.email_dosen, 
            password: formData.nip, 
            role: 'dosen' 
          }])
          .select('id_user')
          .single();

        if (userError) throw userError;

        const { error: dosenError } = await supabase
          .from("dosen")
          .insert([{
            id_dosen: newUser.id_user, 
            nama_dosen: formData.nama_dosen,
            nip: formData.nip,
            email_dosen: formData.email_dosen,
            status: formData.status,
            avatar_dosen: formData.avatar_dosen
          }]);

        if (dosenError) {
          await supabase.from("User").delete().eq("id_user", newUser.id_user);
          throw dosenError;
        }
        alert("Dosen baru berhasil ditambahkan!");
      }

      fetchInitialData();
      handlePrepareNew();
    } catch (err: any) {
      alert("Terjadi kesalahan: " + err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!selectedDosenId) return alert("Pilih dosen yang ingin dihapus!");
    if (!confirm("Apakah Anda yakin ingin menghapus dosen ini? Data akun di tabel User juga akan ikut terhapus.")) return;

    // Ambil info dosen untuk hapus user terkait
    const dosenToDelete = listDosen.find(d => d.id_dosen == selectedDosenId);

    const { error } = await supabase.from("dosen").delete().eq("id_dosen", selectedDosenId);
    if (!error) {
      // Hapus juga di tabel User berdasarkan email
      await supabase.from("User").delete().eq("email", dosenToDelete.email_dosen);
      alert("Dosen dan akun berhasil dihapus");
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
            <span>Profil Dosen</span>
            <div className="header-icons">
              <button title="Tambah Baru" onClick={handlePrepareNew}>➕</button>
              <button title="Hapus Dosen" onClick={handleDelete}>🗑️</button>
            </div>
          </div>
          <div className="profile-content-grid">
            <div className="profile-form-section">
              <div className="form-group-custom">
                <label>Pilih Dosen (Untuk Edit)</label>
                <select onChange={(e) => handleSelectDosen(e.target.value)} value={selectedDosenId}>
                  <option value="">-- Mode Input Baru --</option>
                  {listDosen.map(d => (
                    <option key={d.id_dosen} value={d.id_dosen}>{d.nama_dosen}</option>
                  ))}
                </select>
              </div>
              <div className="form-group-custom">
                <label>Nama Dosen</label>
                <input type="text" value={formData.nama_dosen} onChange={(e) => setFormData({...formData, nama_dosen: e.target.value})} />
              </div>
              <div className="form-group-custom">
                <label>NIP/NIDN</label>
                <input type="text" value={formData.nip} onChange={(e) => setFormData({...formData, nip: e.target.value})} />
              </div>
              <div className="form-group-custom">
                <label>e-mail</label>
                <input type="email" value={formData.email_dosen} onChange={(e) => setFormData({...formData, email_dosen: e.target.value})} />
              </div>
              <div className="form-group-custom">
                <label>Status</label>
                <select value={formData.status} onChange={(e) => setFormData({...formData, status: e.target.value})}>
                    <option value="Aktif">Aktif</option>
                    <option value="Cuti">Cuti</option>
                    <option value="Tidak Aktif">Tidak Aktif</option>
                </select>
              </div>
            </div>
            <div className="profile-right-section">
              <div className="profile-avatar-placeholder">
                {formData.avatar_dosen ? (
                  <img src={formData.avatar_dosen} alt="Profile" className="avatar-img" />
                ) : (
                  <div className="avatar-icon">👤</div>
                )}
              </div>
              <input type="file" ref={fileInputRef} onChange={handleUploadAvatar} accept="image/*" style={{ display: 'none' }} />
              <button className="edit-btn-text" onClick={() => fileInputRef.current?.click()} disabled={uploading}>
                {uploading ? "Mengunggah..." : "Ganti Foto"}
              </button>
              <button className="save-btn-large" onClick={handleSave} disabled={loading}>
                {loading ? "Memproses..." : selectedDosenId ? "Update Data" : "Simpan Dosen Baru"}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}