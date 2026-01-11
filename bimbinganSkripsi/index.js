// import modul supabase
import { createClient } from '@supabase/supabase-js'
import 'dotenv/config'  // untuk membaca file .env

// buat koneksi ke Supabase
const supabase = createClient(process.env.SUPABASE_URL, process.env.SUPABASE_KEY)

// contoh ambil data dari tabel skripsi
async function getSkripsi() {
  const { data, error } = await supabase
    .from('skripsi')
    .select('*')

  if (error) {
    console.error('❌ Error:', error.message)
  } else {
    console.log('📄 Data Skripsi:', data)
  }
}

getSkripsi()
