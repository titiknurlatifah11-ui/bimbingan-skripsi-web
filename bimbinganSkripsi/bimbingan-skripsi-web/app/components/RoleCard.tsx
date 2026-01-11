import Link from "next/link";

interface RoleCardProps {
  title: string;
  href: string;
}

export default function RoleCard({ title, href }: RoleCardProps) {
  return (
    <Link href={href}>
      <div className="bg-[#b9c4ff] hover:bg-[#9daaff] transition p-6 rounded-xl shadow-md cursor-pointer flex flex-col items-center">
        <div className="w-20 h-20 rounded-full bg-white mb-4"></div>
        <span className="bg-white px-4 py-2 rounded-full font-semibold">
          {title}
        </span>
      </div>
    </Link>
  );
}
