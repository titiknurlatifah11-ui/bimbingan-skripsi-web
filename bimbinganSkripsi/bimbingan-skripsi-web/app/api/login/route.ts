import { NextResponse } from "next/server";
import { supabase } from "@/lib/supabase";

export async function POST(req: Request) {
  const { email, password } = await req.json();

  const { data: user, error } = await supabase
    .from("User")
    .select("*")
    .eq("email", email)
    .eq("password", password)
    .single();

  if (error || !user) {
    return NextResponse.json(
      { message: "Email atau password salah" },
      { status: 401 }
    );
  }

  return NextResponse.json({ role: user.role });
}
