"use client";
import { useEffect, useState, useRef } from "react";
import { supabase } from "@/lib/supabase";
import SidebarMahasiswa from "@/app/components/SidebarMahasiswa";
import "./bimbingan-mhs.css";

export default function DataBimbinganMahasiswa() {
  const [namaMhs, setNamaMhs] = useState("Loading...");
  const [mhsId, setMhsId] = useState<number | null>(null);
  const [dosen, setDosen] = useState<any>(null);
  const [jadwal, setJadwal] = useState<any[]>([]);
  
  // State Form
  const [topik, setTopik] = useState("");
  const [bab, setBab] = useState("");
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    async function loadInitialData() {
      // 1. Ambil identitas dari Auth atau LocalStorage
      const { data: { user } } = await supabase.auth.getUser();
      const emailToSearch = user?.email || localStorage.getItem("userEmail");

      if (!emailToSearch) return;

      // 2. Ambil Profile & ID Dosen Pembimbing
      const { data: mhs } = await supabase
        .from("mahasiswa")
        .select("id_mahasiswa, nama_mahasiswa, id_dosen_pembimbing")
        .eq("email_mahasiswa", emailToSearch)
        .single();

      if (mhs) {
        setNamaMhs(mhs.nama_mahasiswa);
        setMhsId(mhs.id_mahasiswa);

        // 3. Ambil Detail Dosen
        if (mhs && mhs.id_dosen_pembimbing) {
          // Pastikan data dosen ini berhasil di-set agar handleSubmitBimbingan tidak ditolak
          const { data: dsn } = await supabase
            .from("dosen")
            .select("*")
            .eq("id_dosen", mhs.id_dosen_pembimbing)
            .single();
          setDosen(dsn);
        }

        // 4. Ambil Jadwal Spesifik Minggu Ini
        const { data: listJadwal } = await supabase
          .from("bimbingan")
          .select("tanggal_pertemuan, jam_pertemuan, topik_bimbingan")
          .eq("id_mahasiswa", mhs.id_mahasiswa)
          .not("tanggal_pertemuan", "is", null);
        setJadwal(listJadwal || []);
      }
    }
    loadInitialData();
  }, []);
  
  const handleUploadClick = () => fileInputRef.current?.click();

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files?.[0]) setSelectedFile(e.target.files[0]);
  };

  const handleResetFile = () => setSelectedFile(null);

  const handleSubmitBimbingan = async () => {
    if (!topik || !bab || !selectedFile || !mhsId || !dosen) {
      return alert("Mohon lengkapi semua data dan pastikan profil dosen termuat!");
    }
  
    try {
      setIsUploading(true);

      // 1. Proses Upload File
      const fileExt = selectedFile.name.split('.').pop();
      const fileName = `${mhsId}-${Date.now()}.${fileExt}`;
      const filePath = `dokumen_skripsi/${fileName}`;
  
      const { error: uploadError } = await supabase.storage
        .from('bimbingan_files')
        .upload(filePath, selectedFile);
  
      if (uploadError) throw uploadError;
  
      const { data: { publicUrl } } = supabase.storage
        .from('bimbingan_files')
        .getPublicUrl(filePath);
  
      // 2. Format Waktu
      const sekarang = new Date();
      const formatJam = sekarang.toTimeString().split(' ')[0];

      // 3. Insert Data ke Tabel Bimbingan
      const { error: insertError } = await supabase
        .from("bimbingan")
        .insert({
          id_mahasiswa: mhsId,
          id_dosen: dosen.id_dosen,
          topik_bimbingan: `${bab}: ${topik}`, // Menggabungkan BAB dan Topik agar rapi
          file_bimbingan: publicUrl,
          status_bimbingan: "Menunggu",
          jam_pertemuan: formatJam 
        });
  
      if (insertError) throw insertError;
  
      alert("Bimbingan berhasil dikirim!");
      
      setTopik("");
      setBab("");
      setSelectedFile(null);
      window.location.reload(); 
  
    } catch (error: any) {
      console.error("Error detail:", error);
      alert("Gagal mengirim bimbingan: " + error.message);
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <div className="dashboard-container">
      <SidebarMahasiswa nama={namaMhs} role="Student" />
      
      <div className="bimbingan-content">
        <div className="welcome-banner-mhs">
          <h3 className="bold-text">Detail Bimbingan, {namaMhs}</h3>
          <p>Berikut ringkasan aktivitas bimbingan Anda</p>
        </div>

        <div className="top-row-grid">
          <div className="info-box-mhs">
            <div className="box-header bold-text">Dosen 👥</div>
            <p className="box-subtext">Informasi dosen pembimbing utama Anda</p>
            <div className="dosen-pill-wrapper">
              <div className="white-pill-long bold-text">
                {dosen ? dosen.nama_dosen : "Data Dosen..."}
              </div>
              <button 
                className="detail-btn-small bold-text" 
                onClick={() => window.location.href='/mahasiswa/dosen'}
              >
                Detail
              </button>
            </div>
          </div>

          <div className="info-box-mhs">
            <div className="box-header bold-text">Jadwal Minggu Ini</div>
            <div className="jadwal-list">
              {jadwal.length > 0 ? (
                jadwal.slice(0, 2).map((j, i) => (
                  <div key={i} className="mini-jadwal-pill">
                    📅 {j.tanggal_pertemuan} | 🕒 {j.jam_pertemuan}
                  </div>
                ))
              ) : (
                <div className="no-schedule-pill bold-text">
                  📝 <span>Tidak ada jadwal</span>
                </div>
              )}
            </div>
          </div>
        </div>

        <div className="detail-bimbingan-card">
          <h4 className="bold-text section-title">Formulir Pengajuan Bimbingan</h4>
          
          <div className="form-group-mhs">
            <label className="bold-text">Topik Bimbingan</label>
            <input 
              type="text" 
              placeholder="Contoh: Revisi Latar Belakang" 
              className="input-mhs-full"
              value={topik}
              onChange={(e) => setTopik(e.target.value)}
            />
          </div>

          <div className="form-group-mhs">
            <label className="bold-text">BAB Skripsi</label>
            <input 
              type="text" 
              placeholder="Contoh: BAB 1" 
              className="input-mhs-full"
              value={bab}
              onChange={(e) => setBab(e.target.value)}
            />
          </div>

          <div className="file-upload-section">
            <label className="upload-label bold-text">📁 Pilih File (PDF/DOC)</label>
            <div className="upload-actions">
              <input 
                type="text" 
                readOnly 
                placeholder={selectedFile ? selectedFile.name : "Belum ada file dipilih..."} 
                className="input-mhs-short" 
              />

              <input 
                type="file" 
                hidden 
                ref={fileInputRef} 
                onChange={handleFileChange}
                accept=".pdf,.doc,.docx"
              />

              <div className="action-buttons-group">
                <button 
                  className="action-mhs-btn plus" 
                  onClick={handleUploadClick} 
                  disabled={isUploading}
                  title="Pilih File"
                >➕</button>
                
                <button 
                  className="action-mhs-btn upload" 
                  onClick={handleSubmitBimbingan}
                  disabled={isUploading}
                  title="Kirim ke Dosen"
                >
                  {isUploading ? "⏳" : "📤"}
                </button>
                
                <button 
                  className="action-mhs-btn trash" 
                  onClick={handleResetFile}
                  disabled={isUploading}
                  title="Hapus"
                >🗑️</button>
              </div>
            </div>
            {isUploading && <p className="upload-status-text">Sedang mengunggah file, mohon tunggu...</p>}
          </div>
        </div>
      </div>
    </div>
  );
}