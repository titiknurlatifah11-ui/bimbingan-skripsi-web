// services/getJadwalKalender.ts
import { supabase } from "@/lib/supabase";

export async function getJadwalKalender() {
  try {
    const { data, error } = await supabase
      .from("jadwal") // atau nama tabel yang sesuai
      .select("*")
      .order("hari", { ascending: true });
    
    if (error) throw error;
    return data || [];
  } catch (error) {
    console.error("Error fetching jadwal kalender:", error);
    return [];
  }
}