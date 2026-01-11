import { supabase } from "@/lib/supabase";

export async function getDosenProfile(idDosen: number) {
  const { data, error } = await supabase
    .from("dosen") // ✅ SESUAIKAN: Gunakan "Dosen" (D kapital) sesuai SQL Anda
    .select("*")
    .eq("id_dosen", idDosen)
    .maybeSingle(); // Menggunakan maybeSingle lebih aman daripada .single() jika data belum ada

  if (error) {
    // 🔥 PERBAIKAN: Agar error tidak muncul {}, panggil properti .message
    console.error("Error fetching profile:", error.message);
    return null;
  }
  return data;

}
