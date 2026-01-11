"use client";

import { useEffect, useState } from "react";
import SidebarDosen from "@/app/components/SidebarDosen";
import { uploadReviewFile } from "@/services/uploadReview";
import { supabase } from "@/lib/supabase";
import "../dosen-data.css";
import "../dashboard.css";

export default function BimbinganPage() {
  const [dataBimbingan, setDataBimbingan] = useState<any[]>([]);
  const [listMahasiswa, setListMahasiswa] = useState<any[]>([]);
  const [filteredMahasiswa, setFilteredMahasiswa] = useState<any[]>([]);
  const [selected, setSelected] = useState<any>(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [loading, setLoading] = useState(false);
  const [namaDosen, setNamaDosen] = useState("Loading...");
  const [idDosenAktif, setIdDosenAktif] = useState<number | null>(null); // Simpan ID Dosen

  useEffect(() => {
    async function loadData() {
      setLoading(true);
      
      // 1. Ambil email dari Auth Session atau LocalStorage
      const { data: { user } } = await supabase.auth.getUser();
      let emailToSearch = user?.email || localStorage.getItem("userEmail");

      if (emailToSearch) {
        // 2. Ambil Profil Dosen
        const { data: profile } = await supabase
          .from("dosen")
          .select("nama_dosen, id_dosen")
          .eq("email_dosen", emailToSearch)
          .single();
        
        if (profile) {
          setNamaDosen(profile.nama_dosen);
          setIdDosenAktif(profile.id_dosen); // Simpan ID asli dosen

          // 3. Ambil Semua Riwayat Bimbingan Dosen Ini
          const { data: bimbinganAll } = await supabase
            .from("bimbingan")
            .select("*, mahasiswa(*)")
            .eq("id_dosen", profile.id_dosen);
          
          if (bimbinganAll) setDataBimbingan(bimbinganAll);

          // 4. Ambil Daftar Mahasiswa Bimbingan Dosen Ini
          const { data: mhs } = await supabase
            .from("mahasiswa")
            .select("*")
            .eq("id_dosen_pembimbing", profile.id_dosen);

          if (mhs) {
            setListMahasiswa(mhs);
            setFilteredMahasiswa(mhs);
          }
        }
      }
      setLoading(false);
    }
    loadData();
  }, []);

  // Filter Search Tetap Sama
  useEffect(() => {
    const query = searchTerm.toLowerCase().trim();
    if (!query) {
      setFilteredMahasiswa(listMahasiswa);
      return;
    }
    const hasil = listMahasiswa.filter((m) => {
      const nama = m.nama_mahasiswa?.toLowerCase() || "";
      const nim = m.nim?.toString() || "";
      return nama.includes(query) || nim.includes(query);
    });
    setFilteredMahasiswa(hasil);
  }, [searchTerm, listMahasiswa]);

  // --- FUNGSI SIMPAN BIASA (Tombol +) ---
  const handleSaveBimbingan = async () => {
    const mhsId = selected?.mahasiswa?.id_mahasiswa || selected?.id_mahasiswa;
    const topik = selected?.topik_bimbingan;
    
    // Ambil judul dari tabel mahasiswa
    const judulSimpan = selected?.mahasiswa?.judul || selected?.judul || "";

    if (!mhsId || !topik || !idDosenAktif) {
      return alert("Pilih mahasiswa dan isi topik!");
    }

    setLoading(true);
    try {
      const formatJam = new Date().toTimeString().split(' ')[0];
      const { error } = await supabase.from("bimbingan").insert({
        id_mahasiswa: mhsId,
        id_dosen: idDosenAktif,
        topik_bimbingan: topik,
        judul_saat_bimbingan: judulSimpan, // SUDAH DISESUAIKAN DENGAN DB
        status_bimbingan: "Menunggu",
        jam_pertemuan: formatJam 
      });

      if (error) throw error;
      alert("Data bimbingan berhasil disimpan!");
      window.location.reload();
    } catch (err: any) {
      alert("Gagal Simpan: " + err.message);
    } finally {
      setLoading(false);
    }
  };

  // --- FUNGSI UPLOAD REVIEW (Tombol Upload) ---
  const handleUploadReview = async () => {
    const mhsId = selected?.mahasiswa?.id_mahasiswa || selected?.id_mahasiswa;
    const topik = selected?.topik_bimbingan;
  
    if (!mhsId || !topik) return alert("Pilih mahasiswa dan isi topik dahulu!");
  
    const fileInput = document.createElement("input");
    fileInput.type = "file";
    fileInput.accept = ".pdf,.doc,.docx";
    
    fileInput.onchange = async (e: any) => {
      const file = e.target.files[0];
      if (!file) return;
      
      setLoading(true);
      try {
        let currentBimbinganId = selected?.id_bimbingan;

        // 1. JIKA BELUM ADA DATA, INSERT DULU
        if (!currentBimbinganId) {
          const formatJam = new Date().toTimeString().split(' ')[0];
          const judulSimpan = selected?.mahasiswa?.judul || selected?.judul || "";

          // GUNAKAN ARRAY [] SAAT INSERT AGAR LEBIH STABIL
          const { data: newBim, error: errBim } = await supabase
            .from("bimbingan")
            .insert([{
              id_mahasiswa: mhsId,
              id_dosen: idDosenAktif,
              topik_bimbingan: topik,
              judul_saat_bimbingan: judulSimpan,
              status_bimbingan: "Disetujui",
              jam_pertemuan: formatJam
            }])
            .select()
            .single();

          if (errBim) throw new Error("Gagal Simpan Data: " + errBim.message);
          currentBimbinganId = newBim.id_bimbingan;
        }

        // 2. UPLOAD FILE KE STORAGE
        const fileExt = file.name.split(".").pop();
        const fileName = `rev_${currentBimbinganId}_${Date.now()}.${fileExt}`;
        
        const { error: uploadError } = await supabase.storage
          .from("bimbingan_files")
          .upload(fileName, file, { cacheControl: '3600', upsert: true });
        
        if (uploadError) throw new Error("Gagal Storage: " + uploadError.message);

        const { data: { publicUrl } } = supabase.storage
          .from("bimbingan_files")
          .getPublicUrl(fileName);

        // 3. UPDATE URL REVIEW KE DATABASE
        const { error: updateError } = await supabase
          .from("bimbingan")
          .update({ 
            review_bimbingan: publicUrl, 
            status_bimbingan: "Disetujui" 
          })
          .eq("id_bimbingan", currentBimbinganId);

        if (updateError) throw new Error("Gagal Update Review: " + updateError.message);
        
        alert("Berhasil! Review bimbingan tersimpan.");
        window.location.reload();
      } catch (err: any) {
        alert(err.message);
      } finally {
        setLoading(false);
      }
    };
    fileInput.click();
  };
  
  const handleUpdateBimbingan = async () => {
    if (!selected?.id_bimbingan) return alert("Pilih data bimbingan!");
    try {
      const { error } = await supabase.from("bimbingan").update({
        topik_bimbingan: selected.topik_bimbingan,
        status_bimbingan: selected.status_bimbingan
      }).eq("id_bimbingan", selected.id_bimbingan);
      if (error) throw error;
      alert("Data berhasil diperbarui");
      window.location.reload();
    } catch (err: any) { alert(err.message); }
  };

  const handleDeleteBimbingan = async () => {
    if (!selected?.id_bimbingan) return alert("Pilih data!");
    if (!confirm("Yakin hapus?")) return;
    try {
      const { error } = await supabase.from("bimbingan").delete().eq("id_bimbingan", selected.id_bimbingan);
      if (error) throw error;
      alert("Data berhasil dihapus");
      window.location.reload();
    } catch (err: any) { alert(err.message); }
  };

  return (
    <div className="dashboard-container">
      {/* 1. Sidebar Harus di Sini */}
      <SidebarDosen nama={namaDosen} />

      <div className="dashboard-content">
        {/* 2. Banner Judul */}
        <div className="header-banner-dosen">
          Detail Bimbingan
        </div>

        {/* 3. Search Bar */}
        <div className="search-section-wrapper">
        <div className="search-bar-container">
          <div className="search-wrapper">
            <input 
              type="text" 
              className="search-input" 
              placeholder="Ketik Nama Mahasiswa atau NIM..." 
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
            <button className="search-icon-btn">🔍</button>
          </div>
        </div>
      </div>

        {/* 4. Form Utama */}
        <div className="profile-container">
          <div className="bimbingan-form">
            <div className="form-group">
              <label>Mahasiswa</label>
              <select 
                className="form-control"
                onChange={(e) => {
                  const idx = parseInt(e.target.value);
                  if (isNaN(idx)) { setSelected(null); return; }
                  const mhs = filteredMahasiswa[idx];
                  const riwayat = dataBimbingan
                    .filter(b => b.id_mahasiswa === mhs.id_mahasiswa)
                    .sort((a, b) => b.id_bimbingan - a.id_bimbingan)[0];
                    setSelected(riwayat ? { ...riwayat, mahasiswa: mhs } : { 
                      mahasiswa: mhs, 
                      topik_bimbingan: "", 
                      id_mahasiswa: mhs.id_mahasiswa,
                      judul: mhs.judul // Menyimpan judul dari tabel mahasiswa ke state selected
                    });
                  }}
                >
          <option value="">-- Pilih Mahasiswa ({filteredMahasiswa.length}) --</option>
              {filteredMahasiswa.map((m, idx) => (
                <option key={m.id_mahasiswa} value={idx}>
                  {m.nama_mahasiswa} - {m.nim}
                </option>
                ))}
              </select>
            </div>

            <div className="form-group">
              <label>Judul</label>
              <input type="text" value={selected?.mahasiswa?.judul || "Belum ada judul"} readOnly className="readonly-input" />
            </div>

            <div className="form-group">
              <label>Topik</label>
              <div className="topic-group">
                <input 
                  type="text" 
                  className="topic-input"
                  value={selected?.topik_bimbingan || ""} 
                  placeholder="Contoh: Revisi Bab 1"
                  onChange={(e) => setSelected((prev: any) => ({ ...prev, topik_bimbingan: e.target.value }))}
                />
                <button className="upload-review-btn" onClick={handleUploadReview}>
                  📤 Upload Review
                </button>
              </div>
            </div>

            <div className="action-small-btns">
              <button onClick={handleSaveBimbingan} disabled={loading}>{loading ? "..." : "➕"}</button>
              <button onClick={handleUpdateBimbingan} disabled={loading}>✏️</button>
              <button onClick={handleDeleteBimbingan} disabled={loading}>🗑️</button>
            </div>

            <div className="bottom-info-grid" style={{ display: 'flex', gap: '20px', marginTop: '30px' }}>
              <div className="form-group" style={{ flex: 2 }}>
                <label>Pembimbing</label>
                <input type="text" value={namaDosen} readOnly />
              </div>
              <div className="form-group" style={{ flex: 1 }}>
                <label>Status</label>
                <input type="text" value={selected?.status_bimbingan || "Belum ada riwayat"} readOnly />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}