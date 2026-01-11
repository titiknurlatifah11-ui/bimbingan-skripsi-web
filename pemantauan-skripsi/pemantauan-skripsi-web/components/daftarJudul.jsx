"use client";
import { useState } from "react";
// FIX: Satu langkah ke atas untuk lib/supabaseClient
import { supabase } from "../lib/supabaseClient"; 

/**
 * Komponen untuk mendaftarkan judul skripsi baru.
 * Ditampilkan jika mahasiswa belum memiliki skripsi terdaftar di tabel 'skripsi'.
 */
export default function DaftarJudul({ mahasiswaId, onJudulRegistered }) {
  const [judul, setJudul] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    if (!mahasiswaId) {
      alert("ID Mahasiswa tidak ditemukan.");
      setLoading(false);
      return;
    }
    
    // Asumsi: Skripsi harus memiliki nilai default di Supabase: 
    // status_progres = 'Proposal' atau 'Diajukan'
    // persentase_progres = 0

    // 1. Simpan Judul ke Tabel Skripsi
    const { error: insertError } = await supabase.from("skripsi").insert([
      {
        mahasiswa_id: mahasiswaId,
        judul: judul,
        // Kolom lain akan menggunakan nilai default dari skema SQL
      },
    ]);

    if (insertError) {
      console.error("Gagal mendaftarkan judul:", insertError);
      alert(`Gagal mendaftarkan judul: ${insertError.message}`);
    } else {
      alert("✅ Judul skripsi berhasil didaftarkan! Menunggu persetujuan.");
      setJudul("");
      
      // Memberi tahu DashboardMahasiswa untuk me-refresh data
      if (onJudulRegistered) onJudulRegistered(); 
    }

    setLoading(false);
  };

  return (
    <div className="bg-gray-50 p-6 rounded-xl border border-purple-300">
      <h3 className="text-xl font-bold text-purple-700 mb-4">Ajukan Judul Skripsi Baru</h3>
      <p className="text-sm text-gray-600 mb-4">
          Anda belum memiliki skripsi terdaftar. Silakan ajukan usulan judul Anda.
      </p>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="judul" className="block text-sm font-medium text-gray-700">
            Usulan Judul Skripsi
          </label>
          <input
            id="judul"
            type="text"
            value={judul}
            onChange={(e) => setJudul(e.target.value)}
            placeholder="Masukkan judul skripsi yang diusulkan"
            required
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm p-2 focus:ring-purple-500 focus:border-purple-500"
          />
        </div>
        <button
          type="submit"
          disabled={loading || judul.length < 5}
          className="w-full py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 disabled:bg-gray-400"
        >
          {loading ? "Mendaftarkan..." : "Daftarkan Judul"}
        </button>
      </form>
    </div>
  );
}
