"use client";
import { useEffect, useState } from "react";
import { supabase } from "../lib/supabaseClient";

// Props yang DITERIMA dari DashboardMahasiswa:
// skripsiId, currentProgress, onProgressUpdated
export default function UploadBimbingan({ mahasiswaId, skripsiId, currentProgress, onProgressUpdated }) {
  const [judul, setJudul] = useState(''); // Menggunakan 'judul'
  const [bab, setBab] = useState('BAB 1'); // Menggunakan 'bab'
  const [file, setFile] = useState(null);
  const [isUploading, setIsUploading] = useState(false); // Status upload
  const [message, setMessage] = useState(''); // Pesan notifikasi
  const [isError, setIsError] = useState(false); // Status error
  const [riwayat, setRiwayat] = useState([]); // Riwayat bimbingan

  // Nilai progres untuk setiap Bab (Target progres, bukan penambahan)
  const BAB_WEIGHTS = {
      'BAB 1': 20,
      'BAB 2': 40,
      'BAB 3': 60,
      'BAB 4': 80,
      'BAB 5': 100,
      // Progres tidak akan dihitung lebih dari 100
  };

  /**
   * Fungsi untuk mengambil riwayat bimbingan dari Supabase
   */
  const fetchRiwayat = async (id) => {
      const { data, error } = await supabase
          .from("bimbingan")
          .select("*") 
          .eq("skripsi_id", id) // Filter berdasarkan ID skripsi
          .order("created_at", { ascending: false }); 

      if (error) console.error("Error ambil riwayat:", error);
      else setRiwayat(data);
  };

  // Panggil fetchRiwayat saat komponen dimuat atau skripsiId berubah
  useEffect(() => {
      if (skripsiId) fetchRiwayat(skripsiId);
  }, [skripsiId]);


  const handleFileUpload = (e) => {
      setFile(e.target.files[0]);
  };

  const handleSubmit = async (e) => {
      e.preventDefault();
      setMessage('');
      setIsError(false);

      if (!skripsiId) {
          setMessage('Data skripsi tidak ditemukan. Pastikan judul sudah disetujui.');
          setIsError(true);
          return;
      }

      if (!judul || !file) {
          setMessage('Semua field (Judul dan File) harus diisi.');
          setIsError(true);
          return;
      }

      setIsUploading(true);
      // Menggunakan mahasiswaId sebagai folder unik untuk upload
      const fileName = `${mahasiswaId}/${Date.now()}_${file.name}`;
      
      try {
          // 1. Upload File ke Storage ('bimbingan_files')
          const { error: uploadError } = await supabase.storage
              .from('bimbingan_files') 
              .upload(fileName, file);

          if (uploadError) throw uploadError;

          // 2. Dapatkan URL Publik
          const { data: { publicUrl: fileUrl } } = supabase.storage.from("bimbingan_files").getPublicUrl(fileName);

          // 3. Hitung Progres Baru (Menggunakan logic BAB_WEIGHTS)
          const newProgressTarget = BAB_WEIGHTS[bab];
          // Progres baru adalah yang tertinggi antara progres saat ini dan target bab
          const newProgress = Math.min(100, Math.max(currentProgress, newProgressTarget)); 

          // 4. Catat Bimbingan ke tabel 'bimbingan'
          const { error: insertError } = await supabase
              .from('bimbingan')
              .insert({
                  skripsi_id: skripsiId,
                  mahasiswa_id: mahasiswaId,
                  judul: judul,
                  bab: bab,
                  file_url: fileUrl, 
                  // status dan komentar akan diisi oleh dosen di komponen lain
              });

          if (insertError) throw insertError;

          // 5. Update Progres dan Status Bab di tabel 'skripsi'
          const { error: updateError } = await supabase
              .from('skripsi')
              .update({ 
                  persentase_progres: newProgress,
                  status_progres: bab // Update status ke bab terbaru (misal: BAB 1)
              })
              .eq('id', skripsiId);

          if (updateError) throw updateError;
          
          // 6. Sukses dan Refresh Data Dashboard
          setMessage(`Bimbingan ${judul} berhasil diunggah! Progres diperbarui menjadi ${newProgress}%.`);
          setJudul('');
          setFile(null);
          
          // Panggil fungsi refresh di parent component
          onProgressUpdated(); 
          // Refresh Riwayat di komponen ini
          fetchRiwayat(skripsiId);

      } catch (error) {
          console.error("Upload Error:", error);
          // Menampilkan pesan error Supabase jika tersedia
          const errorMsg = error.message || error.error_description || 'Unknown error';
          setMessage(`Gagal mengunggah bimbingan: ${errorMsg}`);
          setIsError(true);
      } finally {
          setIsUploading(false);
      }
  };

  return (
      <div className="bg-gray-50 p-6 rounded-xl border border-gray-200 shadow-lg">
          <h2 className="text-xl font-bold text-purple-700 mb-4 border-b pb-2">Upload Bimbingan Baru</h2>
          
          {/* FORM UPLOAD */}
          <form onSubmit={handleSubmit} className="space-y-4">
              {/* Input Judul Bimbingan */}
              <div>
                  <label htmlFor="judul" className="block text-sm font-medium text-gray-700 mb-1">
                      Judul Bimbingan
                  </label>
                  <input
                      id="judul"
                      type="text"
                      value={judul}
                      onChange={(e) => setJudul(e.target.value)}
                      placeholder="Misalnya: Revisi bab 1 pendahuluan"
                      required
                      className="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-purple-500 focus:border-purple-500"
                      disabled={isUploading}
                  />
              </div>

              {/* Dropdown Bab Skripsi */}
              <div>
                  <label htmlFor="bab" className="block text-sm font-medium text-gray-700 mb-1">
                      Bab Skripsi
                  </label>
                  <select
                      id="bab"
                      value={bab}
                      onChange={(e) => setBab(e.target.value)}
                      className="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-purple-500 focus:border-purple-500 bg-white"
                      disabled={isUploading}
                  >
                      {Object.keys(BAB_WEIGHTS).map(b => (
                          <option key={b} value={b}>{b}</option>
                      ))}
                  </select>
              </div>

              {/* Input File */}
              <div>
                  <label htmlFor="file" className="block text-sm font-medium text-gray-700 mb-1">
                      Pilih File (PDF/DOCX)
                  </label>
                  <input
                      id="file"
                      type="file"
                      accept=".pdf,.docx,.doc"
                      onChange={handleFileUpload}
                      required
                      className="w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-purple-50 file:text-purple-700 hover:file:bg-purple-100"
                      disabled={isUploading}
                  />
                  {file && <p className="mt-1 text-xs text-gray-500">File terpilih: {file.name}</p>}
              </div>

              {/* Tombol Kirim */}
              <button
                  type="submit"
                  className={`w-full py-2 px-4 border border-transparent rounded-lg text-sm font-medium text-white 
                      ${isUploading ? 'bg-purple-300 cursor-not-allowed' : 'bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 transition duration-150'}`}
                  disabled={isUploading}
              >
                  {isUploading ? 'Mengunggah...' : 'Kirim Bimbingan'}
              </button>

              {/* Message Box */}
              {message && (
                  <div 
                      className={`p-3 rounded-lg text-sm ${isError ? 'bg-red-100 text-red-800' : 'bg-green-100 text-green-800'}`}
                      role="alert"
                  >
                      {message}
                  </div>
              )}
          </form>

          {/* Riwayat Bimbingan */}
          <h3 className="text-lg font-bold text-gray-700 mt-8 mb-3 pt-4 border-t border-gray-300">Riwayat Bimbingan</h3>
          {riwayat.length === 0 ? (
              <p className="text-gray-500 italic">Belum ada riwayat bimbingan yang diunggah.</p>
          ) : (
              <ul className="space-y-3">
                  {riwayat.map((item) => (
                      <li key={item.id} className="p-3 bg-white border border-gray-100 rounded-lg shadow-sm">
                          <div className="flex justify-between items-start">
                              <div>
                                  <strong className="text-base text-gray-800">{item.judul}</strong>
                                  <span className={`ml-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium 
                                      ${item.status === 'ACC' ? 'bg-green-100 text-green-800' : item.status === 'Direvisi' ? 'bg-red-100 text-red-800' : 'bg-yellow-100 text-yellow-800'}`}>
                                      {item.status || 'Menunggu'}
                                  </span>
                              </div>
                              <small className="text-xs text-gray-500">
                                  {new Date(item.created_at || item.tanggal_upload).toLocaleString("id-ID", {dateStyle: 'medium', timeStyle: 'short'})}
                              </small>
                          </div>
                          <p className="text-sm text-purple-600 font-semibold mt-1">{item.bab}</p>
                          {item.file_url && (
                              <a 
                                  href={item.file_url} 
                                  target="_blank" 
                                  rel="noopener noreferrer"
                                  className="text-sm text-blue-500 hover:text-blue-700 underline"
                              >
                                  📥 Lihat File
                              </a>
                          )}
                          {item.komentar && (
                              <p className="text-sm text-gray-600 mt-2 border-l-2 border-red-400 pl-2 italic">
                                  Komentar: {item.komentar}
                              </p>
                          )}
                      </li>
                  ))}
              </ul>
          )}
      </div>
  );
}
