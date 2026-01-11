// /services/getBimbinganByDosen.ts
import { supabase } from "@/lib/supabase";

// services/getBimbinganByDosen.ts

export const getBimbinganByDosen = async (idDosen: number) => {
  const { data, error } = await supabase
    .from("bimbingan")
    .select(`
      *,
      mahasiswa (
        nama_mahasiswa,
        nim,
        judul
      )
    `) // Jangan masukkan komentar -- di dalam string ini
    .eq("id_dosen", idDosen)
    .order("id_bimbingan", { ascending: false });

  if (error) {
    console.error("Error fetching bimbingan:", error.message);
    return [];
  }
  return data;
};