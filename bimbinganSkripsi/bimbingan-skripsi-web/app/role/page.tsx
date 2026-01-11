"use client";

import Image from "next/image";
import { useRouter } from "next/navigation";
import "./role.css";

export default function RolePage() {
  const router = useRouter();

  const selectRole = (role: string) => {
    router.push(`/login?role=${role}`);
  };

  return (
    <div className="role-container">

      <div className="role-box">

        {/* HEADER */}
        <div className="role-header">
          <Image src="/logo1.png" width={65} height={65} alt="CoreTrack" />
          <h2 className="title">Thesis Tracking Informatics Engineering</h2>
          <p className="subtitle">Daftar Role</p>
        </div>

        {/* ROLE CARDS */}
        <div className="role-options">

          <div className="role-card" onClick={() => selectRole("mahasiswa")}>
            <Image src="/user-student.png" width={70} height={70} alt="Mahasiswa" />
            <p className="role-label">Mahasiswa</p>
          </div>

          <div className="role-card" onClick={() => selectRole("dosen")}>
            <Image src="/user-teacher.png" width={70} height={70} alt="Dosen" />
            <p className="role-label">Dosen</p>
          </div>

          <div className="role-card" onClick={() => selectRole("admin")}>
            <Image src="/user-admin.png" width={70} height={70} alt="Admin" />
            <p className="role-label">Admin</p>
          </div>

        </div>
      </div>
    </div>
  );
}
