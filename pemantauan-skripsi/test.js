import { supabase } from './pemantauan-skripsi-web/src/lib/supabaseClient.js/index.js.js.js'

async function testConnection() {
  const { data, error } = await supabase.from('skripsi').select('*')

  if (error) console.error('Error:', error)
  else console.log('Data:', data)
}

testConnection()
