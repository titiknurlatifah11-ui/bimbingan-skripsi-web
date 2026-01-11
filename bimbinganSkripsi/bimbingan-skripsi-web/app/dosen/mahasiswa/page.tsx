"use client";



import { useEffect, useState } from "react";
import { supabase } from "@/lib/supabase";
import { useRouter } from "next/navigation";
import SidebarDosen from "@/app/components/SidebarDosen";
import { getMahasiswaByDosen } from "@/services/getMahasiswaByDosen";
import "../dosen-data.css";
import "../dashboard.css";

export default function DataMahasiswaPage() {
  const router = useRouter();
  const [listMhs, setListMhs] = useState<any[]>([]);
  const [filteredMhs, setFilteredMhs] = useState<any[]>([]); // Untuk hasil search
  const [selectedMhs, setSelectedMhs] = useState<any>(null);
  const [namaDosen, setNamaDosen] = useState("Loading...");
  const [searchTerm, setSearchTerm] = useState(""); // State untuk teks search

  useEffect(() => {
    async function loadData() {
      // 1. Cek Auth Session
      const { data: { user } } = await supabase.auth.getUser();
      
      // 2. Ambil email (dari Auth atau dari LocalStorage jika login manual)
      let emailToSearch = user?.email || localStorage.getItem("userEmail");

      if (emailToSearch) {
        // 3. Cari data dosen berdasarkan email yang ditemukan
        const { data: dosenData } = await supabase
          .from("dosen")
          .select("nama_dosen, id_dosen")
          .eq("email_dosen", emailToSearch)
          .single();

        if (dosenData) {
          setNamaDosen(dosenData.nama_dosen);
          
          // 4. Ambil daftar mahasiswa bimbingan milik dosen ini
          const dataMhs = await getMahasiswaByDosen(dosenData.id_dosen);
          setListMhs(dataMhs || []);
          setFilteredMhs(dataMhs || []);

          if (dataMhs && dataMhs.length > 0) {
            setSelectedMhs(dataMhs[0]);
          }
        }
      } else {
        // Jika tidak ada identitas, arahkan kembali ke login
        router.push("/login");
      }
    }
    loadData();
  }, [router]);



  // Fungsi untuk menangani pencarian
  const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setSearchTerm(value);

    const filtered = listMhs.filter((mhs) =>
      mhs.nama_mahasiswa.toLowerCase().includes(value.toLowerCase())
    );

    setFilteredMhs(filtered);
    // Jika ada hasil, otomatis pilih yang pertama agar form berubah
    if (filtered.length > 0) {
      setSelectedMhs(filtered[0]);
    } else {
      setSelectedMhs(null);
    }
  };

  return (
    <div className="dashboard-container">
      <SidebarDosen nama={namaDosen} />
      <div className="dashboard-content">
        {/* Search Bar yang sudah berfungsi */}
        <div className="search-bar-container" style={{ marginBottom: '20px', position: 'relative' }}>
             <input
                type="text"
                className="search-input"
                placeholder="🔍 Cari Nama Mahasiswa..."
                value={searchTerm}
                onChange={handleSearch}
             />

             {/* Dropdown hasil pencarian (Opsional agar user bisa pilih) */}
             {searchTerm && filteredMhs.length > 0 && (
               <div className="search-dropdown">
                 {filteredMhs.map((mhs) => (
                   <div
                     key={mhs.id_mahasiswa}
                     className="search-item"
                     onClick={() => {
                        setSelectedMhs(mhs);
                        setSearchTerm(""); // Reset search setelah pilih
                     }}
                   >
                     {mhs.nama_mahasiswa} - {mhs.nim}
                   </div>
                 ))}
               </div>
             )}
        </div>
        <div className="profile-container">
          <div className="profile-flex">
            <div className="profile-form">
              <div className="form-group">
                <label>Mahasiswa</label>
                <input type="text" value={selectedMhs?.nama_mahasiswa || "Data tidak ditemukan"} readOnly />
              </div>

              <div className="form-group">
                <label>NIM</label>
                <input type="text" value={selectedMhs?.nim || "-"} readOnly />
              </div>

              <div className="form-group">
                <label>Pembimbing</label>
                <input type="text" value={namaDosen} readOnly />
              </div>

              <div className="form-group">
                <label>Judul</label>
                <input type="text" value={selectedMhs?.judul || "Belum ada judul"} readOnly />
              </div>

              <div className="form-group">
                <label>Tahun Akademik</label>
                <input type="text" value={selectedMhs?.angkatan || "-"} readOnly />
              </div>
            </div>

            <div className="profile-photo-section">
              <div className="photo-box">
                <span style={{ fontSize: "80px" }}>👤</span>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>
  );
}