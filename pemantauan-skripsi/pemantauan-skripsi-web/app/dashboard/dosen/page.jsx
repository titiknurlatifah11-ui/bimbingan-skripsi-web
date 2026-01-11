// "use client";
// import { useState, useEffect } from "react";
// import { supabase } from "@/lib/supabaseClient";

// export default function DashboardDosen() {
//   const [dosen, setDosen] = useState(null);
//   const [mahasiswaList, setMahasiswaList] = useState([]);
//   const [selectedSkripsi, setSelectedSkripsi] = useState(null);
//   const [riwayat, setRiwayat] = useState([]);

//   // 🔹 Ambil data dosen dan mahasiswa bimbingannya
//   useEffect(() => {
//     const fetchDosenData = async () => {
//       const { data: { user } } = await supabase.auth.getUser();
//       if (!user) return;

//       const { data: dosenData } = await supabase
//         .from("users")
//         .select("*")
//         .eq("auth_id", user.id)
//         .single();
//       setDosen(dosenData);

//       const { data: skripsiList } = await supabase
//         .from("skripsi")
//         .select("*, users(nama, nim)")
//         .eq("dosen_id", dosenData.id);
//       setMahasiswaList(skripsiList);
//     };

//     fetchDosenData();
//   }, []);

//   // 🔹 Ambil riwayat bimbingan berdasarkan skripsi_id
//   const handleSelectSkripsi = async (skripsiId) => {
//     setSelectedSkripsi(skripsiId);
//     const { data: bimbinganList } = await supabase
//       .from("bimbingan")
//       .select("*")
//       .eq("skripsi_id", skripsiId)
//       .order("tanggal_upload", { ascending: false });

//     setRiwayat(bimbinganList);
//   };

//   // 🔹 Dosen memberi feedback
//   const handleFeedback = async (bimbinganId, feedback, mahasiswaId) => {
//     await supabase
//       .from("bimbingan")
//       .update({ feedback_dosen: feedback })
//       .eq("id", bimbinganId);

//     await supabase.from("notifikasi").insert({
//       user_id: mahasiswaId,
//       pesan: "Dosen memberikan feedback pada bimbingan kamu.",
//       dibuat_pada: new Date().toISOString(),
//     });
//     alert("Feedback berhasil dikirim!");
//   };

//   return (
//     <div className="p-10 bg-purple-50 min-h-screen">
//       <h1 className="text-2xl font-bold text-gray-700 mb-6">Dashboard Dosen</h1>

//       {!dosen ? (
//         <p className="text-gray-500">Memuat data dosen...</p>
//       ) : (
//         <>
//           <h2 className="text-lg text-purple-700 font-semibold mb-3">
//             Mahasiswa Bimbingan
//           </h2>

//           {mahasiswaList.length === 0 ? (
//             <p className="text-gray-500 italic">
//               Belum ada mahasiswa bimbingan.
//             </p>
//           ) : (
//             mahasiswaList.map((mhs) => (
//               <div
//                 key={mhs.id}
//                 className="bg-white p-4 rounded-lg shadow mb-3 cursor-pointer"
//                 onClick={() => handleSelectSkripsi(mhs.id)}
//               >
//                 <h3 className="font-semibold text-purple-700">{mhs.users.nama}</h3>
//                 <p className="text-sm text-gray-600">Judul: {mhs.judul}</p>
//                 <p className="text-xs text-gray-500">
//                   Progres: {mhs.persentase_progres}%
//                 </p>
//               </div>
//             ))
//           )}

//           {selectedSkripsi && (
//             <div className="mt-6 bg-white p-5 rounded-xl shadow">
//               <h3 className="font-bold text-lg text-purple-700 mb-3">
//                 Riwayat Bimbingan
//               </h3>
//               {riwayat.map((b) => (
//                 <div key={b.id} className="border-b py-3">
//                   <p className="text-sm text-gray-700 mb-1">
//                     Catatan Mahasiswa: {b.catatan_mahasiswa}
//                   </p>
//                   <p className="text-xs text-gray-500 mb-1">
//                     File:{" "}
//                     <a
//                       href={b.file_bimbingan}
//                       className="text-blue-500 underline"
//                       target="_blank"
//                     >
//                       Lihat
//                     </a>
//                   </p>
//                   <textarea
//                     placeholder="Tulis feedback..."
//                     className="w-full border rounded-lg p-2 text-sm"
//                     onBlur={(e) =>
//                       handleFeedback(b.id, e.target.value, mhs.mahasiswa_id)
//                     }
//                   />
//                 </div>
//               ))}
//             </div>
//           )}
//         </>
//       )}
//     </div>
//   );
// }
