import { supabase } from './bimbingan-skripsi-web/lib/supabaseClient.js'

async function testConnection() {
  const { data, error } = await supabase
    .from('User')   // ← BENERINI
    .select('*')

  if (error) {
    console.error('❌ Error:', error)
  } else {
    console.log('✅ Data User:', data)
  }
}

testConnection()
