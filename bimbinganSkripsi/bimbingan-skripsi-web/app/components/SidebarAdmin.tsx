"use client";

import Image from "next/image";
import { useRouter, usePathname } from "next/navigation";
import { supabase } from "@/lib/supabase";
import "./sidebar-admin.css"; // Pastikan nama file ini benar di folder components

export default function SidebarAdmin({ nama }: { nama: string }) {
  const router = useRouter();
  const pathname = usePathname();

  const isActive = (path: string) => pathname === path;

  const handleLogout = async () => {
    const confirmLogout = confirm("Apakah Anda yakin ingin keluar?");
    if (!confirmLogout) return;

    try {
      await supabase.auth.signOut();
      router.push("/login");
      router.refresh();
    } catch (error: any) {
      alert("Gagal logout: " + error.message);
    }
  };

  return (
    <div className="sidebar">
      <div className="sidebar-top">
        <div style={{ textAlign: "center", marginBottom: "20px" }}>
          <Image src="/logo1.png" width={70} height={70} alt="CoreTrack" style={{ margin: "0 auto" }} />
          <h3>CoreTrack</h3>
          <p style={{ fontSize: "10px", color: "#666", marginTop: "-10px" }}>Thesis Monitoring</p>
        </div>

        <ul className="menu-list">
          <li className={isActive("/admin/dosen") ? "active" : ""} onClick={() => router.push("/admin/dosen")}>
            📋 Data Dosen
          </li>
          <li className={isActive("/admin/mahasiswa") ? "active" : ""} onClick={() => router.push("/admin/mahasiswa")}>
            🧑‍🎓 Data Mahasiswa
          </li>
          <li className={isActive("/admin/pembimbing") ? "active" : ""} onClick={() => router.push("/admin/pembimbing")}>
            👥 Data Pembimbing
          </li>
        </ul>
      </div>

      <div className="sidebar-bottom">
        <div className="user-box">
          <div className="avatar">{nama ? nama[0].toUpperCase() : "A"}</div>
          <div>
            <p className="name">{nama || "ADMIN"}</p>
            <p className="role">ADMINISTRATOR</p>
          </div>
        </div>
        <button className="logout" onClick={handleLogout}>
          Logout ↪
        </button>
      </div>
    </div>
  );
}