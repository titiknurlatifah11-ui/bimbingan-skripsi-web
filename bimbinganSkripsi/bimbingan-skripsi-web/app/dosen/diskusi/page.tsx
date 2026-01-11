"use client";

import { useEffect, useState, useRef } from "react";
import { supabase } from "@/lib/supabase";
import SidebarDosen from "@/app/components/SidebarDosen";
import { getDiskusi, sendPesan } from "@/services/getDiskusi";
import "../dashboard.css";
import "./diskusi.css";

export default function DiskusiPage() {
  const [messages, setMessages] = useState<any[]>([]);
  const [inputPesan, setInputPesan] = useState("");
  const [namaDosen, setNamaDosen] = useState("Dosen");
  const [listMahasiswa, setListMahasiswa] = useState<any[]>([]);
  const [selectedMhs, setSelectedMhs] = useState<any>(null);
  const chatEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    async function initDosen() {
      const { data: { user } } = await supabase.auth.getUser();
      let emailToSearch = user?.email || localStorage.getItem("userEmail");
      
      if (!emailToSearch) return;
  
      const { data: dosen } = await supabase
        .from("dosen")
        .select("id_dosen, nama_dosen")
        .eq("email_dosen", emailToSearch)
        .single();
  
      if (dosen) {
        setNamaDosen(dosen.nama_dosen);
        const { data: mhsList } = await supabase
          .from("mahasiswa")
          .select("id_mahasiswa, nama_mahasiswa, nim")
          .eq("id_dosen_pembimbing", dosen.id_dosen);
  
        if (mhsList && mhsList.length > 0) {
          setListMahasiswa(mhsList);
          setSelectedMhs(mhsList[0]);
        }
      }
    }
    initDosen();
  }, []);

  useEffect(() => {
    if (!selectedMhs) return;

    const loadChat = async () => {
      const { data: { user } } = await supabase.auth.getUser();
      const { data: dosen } = await supabase.from("dosen").select("id_dosen").eq("email_dosen", user?.email).single();

      if (dosen) {
        const chatData = await getDiskusi(dosen.id_dosen, selectedMhs.id_mahasiswa);
        setMessages(chatData || []);
      }
    };

    loadChat();
    const interval = setInterval(loadChat, 5000);
    return () => clearInterval(interval);
  }, [selectedMhs]);

  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const handleKirim = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!inputPesan.trim() || !selectedMhs) return;

    const { data: { user } } = await supabase.auth.getUser();
    const { data: dosen } = await supabase.from("dosen").select("id_dosen").eq("email_dosen", user?.email).single();

    if (dosen) {
      await sendPesan(selectedMhs.id_mahasiswa, dosen.id_dosen, inputPesan, 'dosen');
      setInputPesan("");
      const chatData = await getDiskusi(dosen.id_dosen, selectedMhs.id_mahasiswa);
      setMessages(chatData || []);
    }
  };

  return (
    <div className="dashboard-container">
      <SidebarDosen nama={namaDosen} />

      <div className="diskusi-content-wrapper">
        <div className="diskusi-header-section">
          <div className="selector-container">
            <label className="bold-text">Pilih Mahasiswa:</label>
            <select 
              className="mhs-dropdown" 
              onChange={(e) => setSelectedMhs(listMahasiswa.find(m => m.id_mahasiswa == e.target.value))}
              value={selectedMhs?.id_mahasiswa || ""}
            >
              {listMahasiswa.map((m) => (
                <option key={m.id_mahasiswa} value={m.id_mahasiswa}>
                  {m.nama_mahasiswa} ({m.nim})
                </option>
              ))}
            </select>
          </div>

          <div className="diskusi-header-card">
            <div className="info-mhs">
              <p className="mhs-label">Mahasiswa Sedang Aktif:</p>
              <h3>{selectedMhs?.nama_mahasiswa || "Memuat..."}</h3>
              <p className="mhs-nim">NIM: {selectedMhs?.nim || "-"}</p>
            </div>
          </div>
        </div>

        <div className="chat-main-container">
          <div className="chat-box">
            {messages.length > 0 ? messages.map((msg) => (
              <div key={msg.id_diskusi} className={`chat-wrapper ${msg.pengirim === 'dosen' ? 'pembimbing' : 'mahasiswa'}`}>
                <div className="chat-avatar">
                  {msg.pengirim === 'dosen' ? 'D' : 'M'}
                </div>
                <div className="chat-content">
                  <span className="sender-name">
                    {msg.pengirim === 'dosen' ? 'Anda' : selectedMhs?.nama_mahasiswa}
                  </span>
                  <div className="bubble">
                    {msg.isi_pesan}
                    <span className="chat-time">
                       {new Date(msg.waktu).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                    </span>
                  </div>
                </div>
              </div>
            )) : (
              <div className="empty-chat">Belum ada percakapan dengan {selectedMhs?.nama_mahasiswa}.</div>
            )}
            <div ref={chatEndRef} />
          </div>

          <form className="chat-input-area" onSubmit={handleKirim}>
            <input 
              type="text" 
              placeholder={`Balas pesan ${selectedMhs?.nama_mahasiswa}...`} 
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