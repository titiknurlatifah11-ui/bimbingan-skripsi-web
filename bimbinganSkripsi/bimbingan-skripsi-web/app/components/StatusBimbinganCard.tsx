"use client";

export default function StatusBimbinganCard({
  menunggu,
  disetujui,
  revisi,
}: {
  menunggu: number;
  disetujui: number;
  revisi: number;
}) {
  return (
    <div className="status-card">
      <h3>Status Bimbingan</h3>

      <div className="item">
        ⏳ Menunggu review <span>{menunggu} file</span>
      </div>

      <div className="item">
        ✔ File disetujui <span>{disetujui} file</span>
      </div>

      <div className="item">
        🔄 File revisi <span>{revisi} file</span>
      </div>
    </div>
  );
}
