import { supabase } from "@/lib/supabase";

// Di dalam services/getMahasiswaByDosen.js
export const getMahasiswaByDosen = async (idDosen) => {
  const { data } = await supabase
    .from("mahasiswa")
    .select("*")
    .eq("id_dosen_pembimbing", idDosen);
  return data || [];
};