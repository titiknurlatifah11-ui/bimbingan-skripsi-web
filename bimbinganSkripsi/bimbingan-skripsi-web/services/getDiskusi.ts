// services/getDiskusi.ts
import { supabase } from "@/lib/supabase";

export const getDiskusi = async (idDosen: number, idMhs: number) => {
  const { data, error } = await supabase
    .from("forum_diskusi") // Pastikan nama tabel di Supabase adalah forum_diskusi
    .select("*")
    .eq("id_dosen", idDosen)
    .eq("id_mahasiswa", idMhs)
    .order("created_at", { ascending: true }); // Mengurutkan chat agar tidak berantakan

  if (error) {
    // Memberikan informasi error lebih detail di console
    console.error("Supabase Error:", error.message);
    throw error;
  }
  return data;
};

// Fungsi sendPesan Anda yang sudah ada di gambar
export const sendPesan = async (idMhs: number, idDosen: number, pesan: string, pengirim: 'dosen' | 'mahasiswa') => {
  const { data, error } = await supabase
    .from("forum_diskusi")
    .insert([{ id_mahasiswa: idMhs, id_dosen: idDosen, isi_pesan: pesan, pengirim: pengirim }]);
  if (error) throw error;
  return data;
};