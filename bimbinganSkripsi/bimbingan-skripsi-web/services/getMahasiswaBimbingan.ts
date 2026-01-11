import { supabase } from "@/lib/supabase";

export async function getMahasiswaBimbingan(idDosen: number) {
  const { data, error } = await supabase
    .from("mahasiswa")
    .select("id_mahasiswa, nama_mahasiswa, nim, judul")
    .eq("id_dosen_pembimbing", idDosen); // Filter berdasarkan pembimbing

  if (error) throw error;
  return data;
}