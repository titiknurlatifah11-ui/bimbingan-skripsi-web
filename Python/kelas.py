#menggunakan nama alias
import matakuliah as mk
from dosen import nama 
#menggunakan from untuk memanggil 1/ lebih entitas 
pd1 = {"namaMK" : mk.nama,"kodeMK" :mk.kode,"sks" :mk.sks,"hari":"senin","jam":"08.00-12.00","ruangan":"GC", "dosen": nama}
print(pd1["dosen"])