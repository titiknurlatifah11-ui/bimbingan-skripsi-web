"use client";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { supabase } from "@/lib/supabase";
import "./sidebar-mhs.css"; // Pastikan nama file CSS sudah benar

interface SidebarProps {
  nama: string;
  role?: string;
}

export default function SidebarMahasiswa({ nama, role }: SidebarProps) {
  const pathname = usePathname();
  const router = useRouter();

  const handleLogout = async () => {
    await supabase.auth.signOut();
    router.push("/login");
  };

  const menuItems = [
    { name: "Home", path: "/mahasiswa/dashboard", icon: "🏠" },
    { name: "Data Dosen", path: "/mahasiswa/dosen", icon: "📋" },
    { name: "Data Mahasiswa", path: "/mahasiswa/datamahasiswa", icon: "🧑‍🎓" },
    { name: "Data Bimbingan", path: "/mahasiswa/bimbingan", icon: "📝" },
    { name: "Forum Diskusi", path: "/mahasiswa/diskusi", icon: "💬" },
  ];

  return (
    <aside className="sidebar-mhs-container">
      <div className="sidebar-mhs-logo-section">
        <img src="/logo.png" alt="Logo" className="mhs-logo-img" />
        <h2 className="mhs-brand-name">CoreTrack</h2>
        <p className="mhs-brand-tagline">Thesis Monitoring</p>
      </div>

      <nav className="sidebar-mhs-nav">
        {menuItems.map((item) => (
          <Link 
            key={item.name} 
            href={item.path}
            className={`nav-mhs-item ${pathname === item.path ? "active" : ""}`}
          >
            <span className="nav-mhs-icon">{item.icon}</span>
            <span className="nav-mhs-text">{item.name}</span>
          </Link>
        ))}
      </nav>

      <div className="sidebar-mhs-footer">
        <div className="mhs-profile-card">
          <div className="mhs-avatar">
            {nama ? nama.substring(0, 2).toUpperCase() : "RA"}
          </div>
          <div className="mhs-info">
            <span className="mhs-name">{nama}</span>
            <span className="mhs-role">{role || "Student"}</span>
          </div>
        </div>
        <button onClick={handleLogout} className="mhs-logout-btn">
          <span>↪</span> Logout
        </button>
      </div>
    </aside>
  );
}