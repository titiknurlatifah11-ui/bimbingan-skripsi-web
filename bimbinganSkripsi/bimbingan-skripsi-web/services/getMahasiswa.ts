import { supabase } from "@/lib/supabase";

export async function getMahasiswa(idDosen: number) {
  return supabase
    .from("Mahasiswa")
    .select("*")
    .eq("id_dosen_pembimbing", idDosen);
}

