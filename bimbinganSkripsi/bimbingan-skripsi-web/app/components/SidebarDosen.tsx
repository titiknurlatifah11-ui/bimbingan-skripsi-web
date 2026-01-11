"use client";

import Image from "next/image";
import { useRouter, usePathname } from "next/navigation";
import { supabase } from "@/lib/supabase"; // Import client supabase
import "./sidebar.css";

export default function SidebarDosen({ nama }: { nama: string }) {
  const router = useRouter();
  const pathname = usePathname();

  const isActive = (path: string) => pathname === path;

  // 🔥 FUNGSI LOGOUT AKTIF
  const handleLogout = async () => {
    const confirmLogout = confirm("Apakah Anda yakin ingin keluar?");
    if (!confirmLogout) return;

    try {
      const { error } = await supabase.auth.signOut();
      
      if (error) throw error;

      // Jika berhasil, arahkan ke halaman login
      router.push("/login");
      router.refresh(); // Memastikan state aplikasi segar kembali
    } catch (error: any) {
      console.error("Error logging out:", error.message);
      alert("Gagal logout: " + error.message);
    }
  };

  return (
    <div className="sidebar">
      <div className="sidebar-top">
        <Image src="/logo1.png" width={70} height={70} alt="CoreTrack" />
        <h3>CoreTrack</h3>

        <ul className="menu-list">
          <li 
            className={isActive("/dosen") ? "active" : ""} 
            onClick={() => router.push("/dosen/dashboard")}
          >
            🏠 Home
          </li>
          <li 
            className={isActive("/dosen/data") ? "active" : ""} 
            onClick={() => router.push("/dosen/data")}
          >
            📋 Data Dosen
          </li>
          <li 
            className={isActive("/dosen/mahasiswa") ? "active" : ""} 
            onClick={() => router.push("/dosen/mahasiswa")}
          >
            🧑‍🎓 Data Mahasiswa
          </li>
          <li 
            className={isActive("/dosen/bimbingan") ? "active" : ""} 
            onClick={() => router.push("/dosen/bimbingan")}
          >
            📝 Data Bimbingan
          </li>
          <li 
            className={isActive("/dosen/diskusi") ? "active" : ""} 
            onClick={() => router.push("/dosen/diskusi")}
          >
            💬 Forum Diskusi</li>
        </ul>
      </div>

      <div className="sidebar-bottom">
        <div className="user-box">
          <div className="avatar">{nama ? nama[0] : "D"}</div>
          <div>
            <p className="name">{nama || "Nama Dosen"}</p>
            <p className="role">Teacher</p>
          </div>
        </div>
        {/* Tombol Logout Sekarang Memanggil handleLogout */}
        <button className="logout" onClick={handleLogout}>
          Logout ↪
        </button>
      </div>
    </div>
  );
}