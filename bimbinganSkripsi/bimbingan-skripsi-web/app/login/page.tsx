"use client";

import { useState } from "react";
import Image from "next/image";
import { useSearchParams, useRouter } from "next/navigation";
import { supabase } from "@/lib/supabase";
import "./login.css";

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMsg, setErrorMsg] = useState("");

  const params = useSearchParams();
  const role = params.get("role") ?? "mahasiswa";

  const handleGoogleLogin = async () => {
    await supabase.auth.signInWithOAuth({
      provider: 'google',
      options: {
        redirectTo: `${window.location.origin}/auth/callback`,
      },
    });
  };

// ... (kode atas tetap sama)

  async function handleLogin(e: any) {
    e.preventDefault();
    setErrorMsg("");

    // --- LANGKAH 1: COBA LOGIN VIA SUPABASE AUTH ---
    const { data: authData, error: authError } = await supabase.auth.signInWithPassword({
      email: email,
      password: password,
    });

    if (!authError && authData.user) {
      // Jika login via Auth berhasil, kita juga simpan emailnya agar sinkron
      localStorage.setItem("userEmail", email); // <-- TAMBAHKAN INI

      const { data: adminData } = await supabase
        .from("User")
        .select("role")
        .eq("email", email)
        .single();

      if (adminData) {
        redirectUser(adminData.role);
        return;
      }
    }

    // --- LANGKAH 2: JIKA AUTH GAGAL, CEK TABEL USER MANUAL ---
    const { data: userData, error: userError } = await supabase
      .from("User")
      .select("*")
      .eq("email", email)
      .eq("password", password)
      .single();

    if (userData) {
      if (userData.role !== role) {
        setErrorMsg(`Email ini terdaftar sebagai ${userData.role}, bukan ${role}.`);
        return;
      }

      // KUNCI PERBAIKAN: Simpan email ke localStorage sebelum pindah halaman
      localStorage.setItem("userEmail", email); // <-- TAMBAHKAN INI

      redirectUser(userData.role);
    } else {
      setErrorMsg("Email atau password salah.");
    }
  }

// ... (kode sisanya tetap sama)
  // Fungsi pembantu untuk arahkan halaman sesuai role
  function redirectUser(userRole: string) {
    if (userRole === 'admin') {
      router.push("/admin/dosen");
    } else if (userRole === 'mahasiswa') {
      router.push("/mahasiswa/dashboard");
    } else if (userRole === 'dosen') {
      router.push("/dosen/dashboard");
    }
  }

  return (
    <div className="login-container">
      {/* SISI KIRI: Visual Brand */}
      <div className="login-left">
        <Image 
          src="/logo.png" 
          alt="CoreTrack Logo" 
          width={300} 
          height={300} 
          className="logo-img"
          priority 
        />
        <h1>CoreTrack</h1>
        <p>Thesis Monitoring</p>
      </div>

      {/* SISI KANAN: Form Input */}
      <div className="login-right">
        <h2>Login sebagai {role}</h2>

        <form onSubmit={handleLogin}>
          <div className="input-group">
            <label>Email</label>
            <input
              type="email"
              value={email}
              placeholder="Masukkan email..."
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>

          <div className="input-group">
            <label>Kata sandi</label>
            <input
              type="password"
              value={password}
              placeholder="Masukkan password..."
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>

          {errorMsg && <p className="error">{errorMsg}</p>}

          <button type="submit" className="submit-btn">
            MASUK
          </button>

          <div className="divider">
            <span>atau</span>
          </div>

          <button 
            type="button" 
            className="google-btn" 
            onClick={handleGoogleLogin}
          >
            <img 
              src="https://www.gstatic.com/firebasejs/ui/2.0.0/images/auth/google.svg" 
              alt="Google" 
              width="18" 
              style={{ marginRight: '10px' }} 
            />
            Login melalui Google
          </button>
        </form>
      </div>
    </div>
  );
}