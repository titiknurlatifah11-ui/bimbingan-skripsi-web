// "use client";
// import { useEffect, useState, useCallback } from "react";
// import { supabase } from "@/lib/supabaseClient";
// import Sidebar from "@/components/Sidebar";
// import UploadBimbingan from "@/components/UploadBimbingan";
// import DaftarJudul from "@/components/DaftarJudul";

// export default function DashboardMahasiswa() {
//   const [user, setUser] = useState(null); // Data user dari tabel 'users'
//   const [skripsi, setSkripsi] = useState(null); // Data skripsi

//   // Fungsi refresh data diletakkan di luar karena akan sering dipanggil
//   const fetchData = useCallback(async () => {
//     // 1. Ambil sesi otentikasi
//     const { data: { user: authUser } } = await supabase.auth.getUser();
//     if (!authUser) {
//       return; 
//     } 

//     // 2. Ambil data dari tabel users (data profil mahasiswa)
//     const { data: userData } = await supabase
//       .from("users")
//       .select("*")
//       .eq("auth_id", authUser.id) // Gunakan ID dari sesi auth
//       .single();
    
//     if (!userData) {
//       console.error("User data (mahasiswa) not found in 'users' table.");
//       return; 
//     }
    
//     setUser(userData); 

//     // 3. Ambil data skripsi berdasarkan ID user yang ditemukan
//     const { data: skripsiData } = await supabase
//       .from("skripsi")
//       .select("*")
//       .eq("mahasiswa_id", userData.id) 
//       .maybeSingle();
    
//     setSkripsi(skripsiData); 
//   }, []); // Dependency array kosong karena supabase stabil

//   useEffect(() => {
//     fetchData();
//   }, [fetchData]); 

//   // Fungsi yang dipanggil child component (DaftarJudul atau UploadBimbingan) untuk me-refresh data 
//   const handleDataRefresh = () => {
//     fetchData(); 
//   };
  
//   // Ambil progres saat ini. Jika null, anggap 0.
//   const currentProgress = skripsi?.persentase_progres ?? 0;

//   // Tampilkan loading state
//   if (!user) {
//     return (
//       <div className="flex justify-center items-center min-h-screen bg-purple-50">
//         <p className="text-xl text-purple-600">Memuat data dashboard...</p>
//       </div>
//     );
//   }

//   return (
//     <div className="flex bg-purple-50 min-h-screen">
//       {/* Sidebar */}
//       <Sidebar />

//       {/* Konten utama */}
//       <main className="flex-1 p-10">
//         <header className="flex justify-between items-center mb-6">
//           <h1 className="text-2xl font-bold text-gray-700">Dashboard Mahasiswa</h1>
//           <input
//             type="text"
//             placeholder="Search..."
//             className="border rounded-lg px-3 py-2 text-sm w-64 focus:ring-2 focus:ring-purple-400 outline-none"
//           />
//         </header>

//         {/* User dipastikan ada (user state sudah di-set) */}
//         <div className="grid grid-cols-3 gap-6">
//           {/* Kiri: info skripsi & upload bimbingan / daftar judul */}
//           <div className="col-span-2 bg-white p-5 rounded-2xl shadow-lg">
//             <h2 className="text-lg font-semibold text-purple-700 mb-3">Skripsi Saya</h2>
            
//             {/* ✅ LOGIC TAMPILAN: Conditional Rendering */}
//             {skripsi ? (
//               // 1. Skripsi sudah terdaftar
//               <>
//                 <div className="mb-4 space-y-1 text-gray-700 border-b pb-4">
//                   <p>
//                     <strong>Judul:</strong> <span className="font-medium">{skripsi.judul}</span>
//                   </p>
//                   <p>
//                     {/* STATUS YANG DIBULETIN: Menampilkan Status Bab dan Persentase Progres */}
//                     <strong>Status:</strong> 
//                     <span 
//                         className={`font-semibold ml-2 inline-block px-3 py-0.5 rounded-full text-xs 
//                         ${currentProgress >= 90 ? 'bg-green-100 text-green-800' : 
//                           currentProgress > 0 ? 'bg-blue-100 text-blue-800' : 
//                           'bg-yellow-100 text-yellow-800'}`
//                         }
//                     >
//                         {skripsi.status_progres} ({currentProgress}%)
//                     </span>
//                   </p>
//                   {/* Progres di baris terpisah (opsional) */}
//                   <p><strong>Progres:</strong> {currentProgress}%</p>
//                 </div>
                
//                 {/* Tampilkan upload bimbingan dan Meneruskan semua props yang diperlukan */}
//                 <UploadBimbingan 
//                     mahasiswaId={user.id} 
//                     skripsiId={skripsi.id} // ID Skripsi dibutuhkan untuk UPDATE
//                     currentProgress={currentProgress} // Progres saat ini dibutuhkan untuk perhitungan progres BARU
//                     onProgressUpdated={handleDataRefresh} // Pemicu refresh setelah upload
//                 />
//               </>
//             ) : (
//               // 2. Skripsi belum terdaftar: Tampilkan Form Pendaftaran
//               <DaftarJudul 
//                   mahasiswaId={user.id} 
//                   onJudulRegistered={handleDataRefresh} // Gunakan handleDataRefresh
//               />
//             )}
//           </div>

//           {/* Kanan: profil + progres */}
//           <div className="bg-white p-5 rounded-2xl shadow-lg h-fit">
//             <div className="text-center mb-5 border-b pb-4">
//               <img
//                 src="https://i.pravatar.cc/100"
//                 alt="Avatar"
//                 className="w-20 h-20 mx-auto rounded-full mb-2 border-2 border-purple-400"
//               />
//               <p className="font-semibold text-gray-700 text-lg">{user.nama}</p>
//               <p className="text-sm text-gray-500">{user.nim}</p>
//             </div>
//             <div className="mt-4">
//               <h3 className="text-sm font-semibold mb-2 text-purple-700">Task Progress</h3>
//               {/* Progres BAR akan berubah sesuai currentProgress */}
//               <div className="w-full bg-gray-200 rounded-full h-2 mb-2">
//                 <div
//                   className="bg-purple-600 h-2 rounded-full transition-all duration-500"
//                   style={{ width: `${currentProgress}%` }}
//                 />
//               </div>
//               <p className="text-xs text-right text-gray-500">
//                 {currentProgress}% Selesai
//               </p>
//             </div>
//           </div>
//         </div>
//       </main>
//     </div>
//   );
// }
