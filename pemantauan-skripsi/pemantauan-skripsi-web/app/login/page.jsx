"use client";
import { useState } from "react";
import { supabase } from "@/lib/supabaseClient";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleLogin = async (e) => {
    e.preventDefault();
    setError("");
    setIsLoading(true);

    // Cek user dari tabel User
    const { data: user, error: queryError } = await supabase
      .from("User")
      .select("*")
      .eq("email", email)
      .single();

    if (queryError || !user) {
      setError("Email tidak ditemukan.");
      setIsLoading(false);
      return;
    }

    // Cek password
    if (user.password !== password) {
      setError("Password salah.");
      setIsLoading(false);
      return;
    }

    // Login berhasil → arahkan sesuai role
    switch (user.role.toLowerCase()) {
      case "admin":
        router.push("/dashboard/admin");
        break;
      case "dosen":
        router.push("/dashboard/dosen");
        break;
      case "mahasiswa":
        router.push("/dashboard/mahasiswa");
        break;
      default:
        setError("Role tidak dikenali.");
    }

    setIsLoading(false);
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4 bg-gray-100">
      <div className="bg-white shadow-xl rounded-xl p-8 w-full max-w-md">
        
        <h1 className="text-2xl font-bold mb-6">Login CoreTrack</h1>

        {error && (
          <div className="bg-red-200 text-red-800 p-3 rounded mb-4 text-sm">
            {error}
          </div>
        )}

        <form onSubmit={handleLogin}>
          <div className="mb-4">
            <label className="text-sm font-medium">Email</label>
            <input
              type="email"
              className="w-full p-3 border rounded-lg mt-1"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>

          <div className="mb-6">
            <label className="text-sm font-medium">Password</label>
            <input
              type="password"
              className="w-full p-3 border rounded-lg mt-1"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>

          <button
            type="submit"
            disabled={isLoading}
            className={`w-full p-3 rounded-lg text-white font-semibold ${
              isLoading ? "bg-gray-400" : "bg-blue-600 hover:bg-blue-700"
            }`}
          >
            {isLoading ? "Memproses..." : "Masuk"}
          </button>
        </form>
      </div>
    </div>
  );
}
