"use client";
import { Home, Calendar, Book, BarChart2, MessageSquare, Settings, LogOut } from "lucide-react";

export default function Sidebar() {
  const menu = [
    { name: "Home", icon: <Home size={18} /> },
    { name: "Schedule", icon: <Calendar size={18} /> },
    { name: "Course", icon: <Book size={18} /> },
    { name: "Grades", icon: <BarChart2 size={18} /> },
    { name: "Chats", icon: <MessageSquare size={18} /> },
  ];

  return (
    <aside className="w-64 bg-purple-700 text-white min-h-screen flex flex-col justify-between p-5">
      <div>
        <h2 className="text-2xl font-bold mb-8 text-center">🎓 Dashboard</h2>
        <nav className="flex flex-col gap-4">
          {menu.map((item) => (
            <button
              key={item.name}
              className="flex items-center gap-3 px-4 py-2 rounded-lg hover:bg-purple-600 transition"
            >
              {item.icon}
              <span>{item.name}</span>
            </button>
          ))}
        </nav>
      </div>

      <div className="border-t border-purple-500 pt-4">
        <button className="flex items-center gap-3 w-full px-4 py-2 rounded-lg hover:bg-purple-600 transition">
          <Settings size={18} />
          <span>Settings</span>
        </button>
        <button className="flex items-center gap-3 w-full px-4 py-2 rounded-lg hover:bg-purple-600 transition mt-2">
          <LogOut size={18} />
          <span>Log Out</span>
        </button>
      </div>
    </aside>
  );
}
