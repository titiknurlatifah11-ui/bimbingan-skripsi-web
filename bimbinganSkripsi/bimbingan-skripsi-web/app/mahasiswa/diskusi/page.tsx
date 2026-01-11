"use client";

import { useEffect, useState, useRef } from "react";
import { supabase } from "@/lib/supabase";
import SidebarMahasiswa from "@/app/components/SidebarMahasiswa";
import { getDiskusi, sendPesan } from "@/services/getDiskusi";
import "./diskusi.css";

export default function ForumDiskusiMahasiswaPage() {
  const [messages, setMessages] = useState<any[]>([]);
  const [inputPesan, setInputPesan] = useState("");
  const [namaMhs, setNamaMhs] = useState(""); 
  const [profileMhs, setProfileMhs] = useState<any>(null);
  const [dosenInfo, setDosenInfo] = useState<any>(null);
  const chatEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    async function initData() {
      // 1. Ambil identitas (Cek Auth dulu, kalau kosong ambil dari LocalStorage)
      const { data: { user } } = await supabase.auth.getUser();
      const emailToSearch = user?.email || localStorage.getItem("userEmail");

      if (!emailToSearch) return;

      // 2. Ambil data profil mahasiswa & join ke tabel dosen pembimbing
      const { data: mhs, error } = await supabase
        .from("mahasiswa")
        .select("*, dosen:id_dosen_pembimbing(id_dosen, nama_dosen, nip)")
        .eq("email_mahasiswa", emailToSearch)
        .single();

      if (error) {
        console.error("Gagal memuat data:", error.message);
        setNamaMhs("Error Load");
        return;
      }

      if (mhs) {
        setNamaMhs(mhs.nama_mahasiswa); 
        setProfileMhs(mhs);
        setDosenInfo(mhs.dosen); 
        
        // 3. Load chat pertama kali
        if (mhs.dosen?.id_dosen) {
          const chatData = await getDiskusi(mhs.dosen.id_dosen, mhs.id_mahasiswa);
          setMessages(chatData || []);
        }
      }
    }
    initData();
  }, []);

  // Polling Chat (Setiap 5 detik)
  useEffect(() => {
    if (!profileMhs?.id_mahasiswa || !dosenInfo?.id_dosen) return;
    
    const interval = setInterval(async () => {
      const chatData = await getDiskusi(dosenInfo.id_dosen, profileMhs.id_mahasiswa);
      setMessages(chatData || []);
    }, 5000);
    
    return () => clearInterval(interval);
  }, [profileMhs, dosenInfo]);

  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const handleKirim = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!inputPesan.trim() || !profileMhs || !dosenInfo) return;

    try {
      await sendPesan(profileMhs.id_mahasiswa, dosenInfo.id_dosen, inputPesan, 'mahasiswa');
      setInputPesan("");
      const chatData = await getDiskusi(dosenInfo.id_dosen, profileMhs.id_mahasiswa);
      setMessages(chatData || []);
    } catch (err) {
      console.error("Gagal mengirim pesan:", err);
    }
  };

  return (
    <div className="dashboard-container">
      <SidebarMahasiswa nama={namaMhs || "Mahasiswa"} role="Student" />

      <div className="diskusi-main-content">
        <div className="diskusi-header-section">
          <div className="diskusi-header-card">
            <div className="info-mhs">
              <p className="mhs-label">Dosen Pembimbing Anda:</p>
              <h3 className="nama-dosen-aktif">{dosenInfo?.nama_dosen || "Memuat Nama Dosen..."}</h3>
              <p className="mhs-nim">NIP: {dosenInfo?.nip || "-"}</p>
            </div>
          </div>
        </div>

        <div className="chat-main-container">
          <div className="chat-box">
            {messages.length > 0 ? messages.map((msg, idx) => (
              <div key={idx} className={`chat-wrapper ${msg.pengirim === 'mahasiswa' ? 'pembimbing' : 'mahasiswa'}`}>
                <div className="chat-avatar">
                  {msg.pengirim === 'mahasiswa' ? 'M' : 'D'}
                </div>
                <div className="chat-content">
                  <span className="sender-name">
                    {msg.pengirim === 'mahasiswa' ? 'Anda' : dosenInfo?.nama_dosen}
                  </span>
                  <div className="bubble">
                    {msg.isi_pesan}
                    <span className="chat-time">
                       {msg.waktu ? new Date(msg.waktu).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) : ""}
                    </span>
                  </div>
                </div>
              </div>
            )) : <div className="empty-chat">Belum ada percakapan bimbingan.</div>}
            <div ref={chatEndRef} />
          </div>

          <form className="chat-input-area" onSubmit={handleKirim}>
            <input 
              type="text" 
              placeholder="Tulis pesan bimbingan..." 
              value={inputPesan}
              onChange={(e) => setInputPesan(e.target.value)}
            />
            <button type="submit" className="send-btn">➤ Kirim</button>
          </form>
        </div>
      </div>
    </div>
  );
}