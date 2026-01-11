import { supabase } from "@/lib/supabase";

export async function getStatusBimbingan(idDosen: number) {
  const { data, error } = await supabase
    .from("Bimbingan")
    .select("status_bimbingan")
    .eq("id_dosen", idDosen);

  if (error) return { menunggu: 0, revisi: 0, disetujui: 0 };

  return {
    menunggu: data.filter((d) => d.status_bimbingan === "Menunggu").length,
    revisi: data.filter((d) => d.status_bimbingan === "Revisi").length,
    disetujui: data.filter((d) => d.status_bimbingan === "Disetujui").length,
  };
}
