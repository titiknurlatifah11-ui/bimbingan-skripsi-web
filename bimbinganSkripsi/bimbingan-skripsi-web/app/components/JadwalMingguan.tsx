"use client";

export default function JadwalMingguan({ jadwal }: { jadwal: any }) {
  return (
    <div className="jadwal-card">
      <h3>Jadwal Minggu Ini</h3>

      <div className="hari">
        {Object.entries(jadwal).map(([hari, isi]) => (
          <div key={hari} className="hari-item">
            <span>{hari}</span>
            <p>{String(isi)}</p> {/* FIX ERROR */}
          </div>
        ))}
      </div>
    </div>
  );
}
