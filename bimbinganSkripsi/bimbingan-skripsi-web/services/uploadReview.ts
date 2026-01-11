import { supabase } from "@/lib/supabase";

export async function uploadReviewFile(idBimbingan: number, file: File) {
  // 1. Upload ke Storage
  const fileName = `${Date.now()}_${file.name}`;
  const { data: uploadData, error: uploadError } = await supabase.storage
    .from("bimbingan_files")
    .upload(fileName, file);

  if (uploadError) throw uploadError;

  // 2. Ambil URL Publik
  const { data: { publicUrl } } = supabase.storage
    .from("bimbingan_files")
    .getPublicUrl(fileName);

  // 3. Update Status & URL di Database
  const { error: updateError } = await supabase
    .from("bimbingan")
    .update({ 
      review_bimbingan: publicUrl, 
      status_bimbingan: 'Selesai' 
    })
    .eq("id_bimbingan", idBimbingan);

  if (updateError) throw updateError;
  return publicUrl;
}