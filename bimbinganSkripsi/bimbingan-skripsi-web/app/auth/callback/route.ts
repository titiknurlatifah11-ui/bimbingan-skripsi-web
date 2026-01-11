import { NextResponse } from "next/server";
import { createClient } from "@supabase/supabase-js";

export async function GET(request: Request) {
  const url = new URL(request.url);
  const code = url.searchParams.get("code");

  if (!code) return NextResponse.redirect("/login");

  const supabase = createClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!
  );

  const { data, error } = await supabase.auth.exchangeCodeForSession(code);

  if (error) return NextResponse.redirect("/login");

  // Kalau login berhasil → redirect sesuai role
  const user = data.user;

  const role = user?.user_metadata?.role || "mahasiswa";

  if (role === "admin") return NextResponse.redirect("/admin");
  if (role === "dosen") return NextResponse.redirect("/dosen");

  return NextResponse.redirect("/mahasiswa");
}
