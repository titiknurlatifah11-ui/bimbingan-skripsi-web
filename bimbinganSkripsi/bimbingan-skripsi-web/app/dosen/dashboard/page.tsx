"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { supabase } from "@/lib/supabase";
import SidebarDosen from "@/app/components/SidebarDosen";
import "../dashboard.css";

import { getStatusBimbingan } from "@/services/getStatusBimbingan";
import { getMahasiswaBimbingan } from "@/services/getMahasiswaBimbingan";
import { uploadReviewFile } from "@/services/uploadReview";

export default function DosenDashboard() {
  const router = useRouter();
  const [dosen, setDosen] = useState<any>(null);
  const [status, setStatus] = useState<any>(null);
  const [mahasiswa, setMahasiswa] = useState<any[]>([]);
  const [jadwal, setJadwal] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  // State untuk Modal Set Jadwal
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedMhs, setSelectedMhs] = useState<any>(null);
  const [tanggal, setTanggal] = useState("");
  const [jam, setJam] = useState("");

  useEffect(() => {
    async function loadData() {
      try {
        setLoading(true);
        
        // --- LANGKAH 1: IDENTIFIKASI SIAPA YANG LOGIN ---
        // Cek via Supabase Auth
        const { data: { user } } = await supabase.auth.getUser();
        let emailToSearch = user?.email;

        // Jika tidak ada di Auth, ambil dari localStorage (hasil login manual tabel User)
        if (!emailToSearch) {
          emailToSearch = localStorage.getItem("userEmail") || "";
        }

        // Jika tidak ada email sama sekali, tendang balik ke login
        if (!emailToSearch) {
          return router.replace("/login");
        }

        // --- LANGKAH 2: AMBIL PROFIL DOSEN BERDASARKAN EMAIL ---
        const { data: profile, error: profileError } = await supabase
          .from("dosen")
          .select("*")
          .eq("email_dosen", emailToSearch) // FILTER KUNCI agar data tidak tertukar
          .single();

        if (profileError || !profile) {
          console.error("Profil Dosen tidak ditemukan");
          return router.replace("/login");
        }

        setDosen(profile);

        // --- LANGKAH 3: AMBIL DATA PENDUKUNG (Status, Mahasiswa, Jadwal) ---
        // 1. Ambil data bimbingan untuk statistik
        const { data: bimbinganAll } = await supabase
          .from("bimbingan")
          .select("status_bimbingan")
          .eq("id_dosen", profile.id_dosen);

        const stats = {
          menunggu: bimbinganAll?.filter(b => b.status_bimbingan?.toLowerCase() === "menunggu").length || 0,
          disetujui: bimbinganAll?.filter(b => b.status_bimbingan?.toLowerCase() === "disetujui").length || 0,
          revisi: bimbinganAll?.filter(b => b.status_bimbingan?.toLowerCase() === "revisi").length || 0,
        };
        setStatus(stats);

        // 2. Ambil data mahasiswa bimbingan
        const m = await getMahasiswaBimbingan(profile.id_dosen);
        setMahasiswa(m);

        // 3. Ambil jadwal
        const { data: j } = await supabase
          .from("bimbingan")
          .select("tanggal_pertemuan, jam_pertemuan, id_mahasiswa")
          .eq("id_dosen", profile.id_dosen)
          .not("tanggal_pertemuan", "is", null);

        setJadwal(Array.isArray(j) ? j : []);

      } catch (error) {
        console.error("Gagal mengambil data dashboard:", error);
      } finally {
        setLoading(false);
      }
    }

    loadData();
  }, [router]);
  
  // FUNGSI 1: Set Jadwal Bimbingan (DIPERBAIKI)
  const handleSetJadwal = async () => {
    if (!tanggal || !jam || !selectedMhs) return alert("Lengkapi data!");

    try {
      setLoading(true);

      // Langsung lakukan upsert tanpa mencari data lama terlebih dahulu.
      // Jika id_mahasiswa dan id_dosen sudah ada, data akan diperbarui.
      // Jika belum ada, baris baru akan dibuat secara otomatis.
      const { error } = await supabase
        .from("bimbingan") // Pastikan nama tabel di SQL sama persis (Bimbingan atau bimbingan)
        .upsert({
          id_mahasiswa: selectedMhs.id_mahasiswa,
          id_dosen: dosen.id_dosen,
          tanggal_pertemuan: tanggal,
          jam_pertemuan: jam,
          status_bimbingan: "Menunggu",
          topik_bimbingan: "Pertemuan Terjadwal"
        });

      if (error) {
        console.error("Supabase Error:", error);
        throw error;
      }

      alert("Jadwal bimbingan berhasil dikirim!");
      setIsModalOpen(false);
      
      // Memuat ulang data dashboard agar angka status dan jadwal langsung terupdate
      window.location.reload(); 
      
    } catch (err: any) {
      console.error("Detail Error:", err);
      alert(`Gagal: ${err.message || "Gagal menyimpan jadwal ke database"}`);
    } finally {
      setLoading(false);
    }
  };

  // FUNGSI 2: Upload Review 
  const handleUploadReview = async () => {
    const fileInput = document.createElement("input");
    fileInput.type = "file";
    fileInput.onchange = async (e: any) => {
      const file = e.target.files[0];
      if (file) {
        try {
          alert("Sedang mengunggah...");
          // Gunakan ID bimbingan dari mahasiswa yang dipilih atau yang pertama
          const bimbinganId = mahasiswa[0]?.id_bimbingan || 1;
          await uploadReviewFile(bimbinganId, file);
          alert("Review Berhasil Diunggah!");
          window.location.reload();
        } catch (err) {
          alert("Gagal mengunggah file.");
        }
      }
    };
    fileInput.click();
  };

  if (loading) return <div className="loading-container"><div className="loading">⏳ Memuat Dashboard...</div></div>;

  return (
    <div className="dashboard-container">
      <SidebarDosen nama={dosen?.nama_dosen} />

      <div className="dashboard-content">
        <div className="welcome-box">
          <h2>Welcome back, {dosen?.nama_dosen}</h2>
          <p>NIP: {dosen?.nip || "Tidak Terdaftar"}</p>
        </div>

        <div className="grid-layout">
          {/* Status Card */}
          <div className="status-card">
            <h3 
            onClick={() => router.push("/dosen/laporan")} 
            style={{ cursor: 'pointer', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}
            title="Klik untuk lihat detail laporan"
          >
            Status Bimbingan
            <span style={{ fontSize: '12px', color: '#4a90e2' }}>Lihat Detail ➔</span>
          </h3>
            <div className="item">⌛ Menunggu review <span>{status?.menunggu || 0} file</span></div>
            <div className="item">✅ File disetujui <span>{status?.disetujui || 0} file</span></div>
            <div className="item">🔄 File revisi <span>{status?.revisi || 0} file</span></div>
          </div>

          {/* Mahasiswa Card */}
          <div className="mahasiswa-card">
            <h3>Mahasiswa</h3>
            <p>Klik tombol untuk atur jadwal</p>
            {mahasiswa.length > 0 ? mahasiswa.map((mhs, idx) => (
              <div className="mhs-row" key={idx}>
                <div className="mhs-info">
                  <p className="nama">{mhs.nama_mahasiswa}</p>
                  <p className="nim">{mhs.nim}</p>
                </div>
                <button className="detail" onClick={() => {
                  setSelectedMhs(mhs);
                  setIsModalOpen(true);
                }}>📅 Jadwal</button>
              </div>
            )) : <p>Belum ada mahasiswa.</p>}
          </div>

          {/* Jadwal Card */}
          <div className="jadwal-card">
            <h3>Jadwal Minggu Ini</h3>
            <div className="hari">
              {["Sen", "Sel", "Rab", "Kam", "Jum"].map((day) => {
                const daysMap: any = { "Sen": 1, "Sel": 2, "Rab": 3, "Kam": 4, "Jum": 5 };
                const bimbinganHariIni = jadwal.find(j => 
                  new Date(j.tanggal_pertemuan).getDay() === daysMap[day]
                );

                return (
                  <div className={`hari-item ${bimbinganHariIni ? "active" : ""}`} key={day}>
                    <span>{day}</span>
                    <p>{bimbinganHariIni ? bimbinganHariIni.jam_pertemuan : "Kosong"}</p>
                  </div>
                );
              })}
            </div>
          </div>

          {/* Tindakan Cepat */}
          <div className="tindakan-card">
            <h3>Tindakan Cepat</h3>
            <button className="btn" onClick={handleUploadReview}>📤 Unggah File Review</button>
            <button className="btn" onClick={() => router.push("/dosen/laporan")}>📝 Laporan Bimbingan</button>
          </div>
        </div>
      </div>

      {/* Modal Set Jadwal */}
      {isModalOpen && (
        <div className="modal-overlay">
          <div className="modal-box">
            <h3>Set Jadwal Bimbingan</h3>
            <p>Mahasiswa: <strong>{selectedMhs?.nama_mahasiswa}</strong></p>
            <div className="form-group">
              <label>Tanggal</label>
              <input type="date" onChange={(e) => setTanggal(e.target.value)} />
            </div>
            <div className="form-group">
              <label>Jam</label>
              <input type="time" onChange={(e) => setJam(e.target.value)} />
            </div>
            <div className="modal-actions">
              <button className="btn-save" onClick={handleSetJadwal}>Simpan</button>
              <button className="btn-cancel" onClick={() => setIsModalOpen(false)}>Batal</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}