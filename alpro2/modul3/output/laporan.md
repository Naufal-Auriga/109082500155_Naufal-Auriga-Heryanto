Halaman 37 | M o d u l P r a k t i k u m A l g o r i t m a d a n P e m r o g r a m a n 2
Pada contoh di atas fungsi faktorial dipanggil secara tidak langsung melalui fungsi permutasi, dan 
fungsi faktorial dan permutasi dipanggil sebagai ekspresi dari suatu statement.
3.5 Soal Latihan Modul 3

Nama : Naufal Auriga heryanto
NIM: 109082500155
kelas : if 03


1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika 
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng 
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian 
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, 
dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂
terhadap 𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.
Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan 
menggunakan persamaan berikut!
𝑃(𝑛, 𝑟) =
𝑛!
(𝑛−𝑟)!
, sedangkan 𝐶(𝑛, 𝑟) =
𝑛!
𝑟!(𝑛−𝑟)!
Contoh
No Masukan Keluaran Penjelasan
1 5 10 3 10 60 10
3628800 1
P(5,3) = 5!/2! = 120/2 = 60
C(5,3) = 5!/(3!x2!) = 120/12 = 10
P(10,10) = 10!/0! = 3628800/1 = 3628800
C(10,10) = 10!/(10!x0!) = 10!/10! = 1
2 8 0 2 0 56 28
1 1
Selesaikan program tersebut dengan memanfaatkan subprogram yang diberikan berikut ini!
function factorial(n: integer) → integer
{mengembalikan nilai faktorial dari n}
function permutation(n,r : integer) → integer
{Mengembalikan hasil n permutasi r, dan n >= r}
function combination(n,r : integer) → integer




Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥
2
, 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 +
1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔(ℎ(𝑥))). Tuliskan 𝑓(𝑥), 𝑔(𝑥) dan ℎ(𝑥)
dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat 𝑎, 𝑏 dan 𝑐 yang dipisahkan oleh spasi.
Keluaran terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), 
dan baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐)!


3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 
𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat 
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik 
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan 
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik 
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".