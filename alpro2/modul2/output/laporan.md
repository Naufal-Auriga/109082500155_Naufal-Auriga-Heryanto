laporan praktikum modul 2

Nama  : Naufal Auriga Heryanto
Kelas : S1 if 03
Nim   : 109082500155

SOAL 1 

Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

package main 

import “fmt” func main() { 
    var ( satu, dua, tiga string temp string ) 
    fmt.Print("Masukan input string: ")
     fmt.Scanln(&satu) 
     fmt.Print("Masukan input string: ") 
     fmt.Scanln(&dua) fmt.Print("Masukan input string: ")
      fmt.Scanln(&tiga)
       fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
        temp = satu
         satu = dua 
         dua = tiga
          tiga = temp 
          fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
 }
 

DESKRIPSI jawabanya

 
kodigan di atas eror karna ya import fmt nya bukan menggunakan petik dua "" 
beserta pembuaatan variable tidak pakai tanda kurung

ini output erornya
curly quotation mark '“' (use neutral '"')

link fotonya
https://github.com/Naufal-Auriga/109082500155_Naufal-Auriga-Heryanto/blob/main/alpro2/modul2/output/soal1.png


SOAL 2

Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.

link poto
https://github.com/Naufal-Auriga/109082500155_Naufal-Auriga-Heryanto/blob/main/alpro2/modul2/output/soal2.png

SOAL 3

PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

link poto
https://github.com/Naufal-Auriga/109082500155_Naufal-Auriga-Heryanto/blob/main/alpro2/modul2/output/soal3.png


