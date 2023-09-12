Website Crud Golang

Golang<br>
HTML<br>
Bootstrap<br>
mySQL<br>

Jalankan mySQL di Xampp / Adminer <br>


HTML > Controller

MVC sebagai pattern<br>
Model View Controllers<br>
Model > Controller

![mvc](/mvc.png)

Jalankan Program go run . (titik)<br>
Matikan Program CTRL + C

Folder  : Models, Views, Controllers, <br>
Folder tambahan : config, entities

Folder Model akan di panggil di folder Controller = Folder controller akan manggil function dari folder model

Kegunaan Folder = 
   - Views untuk menampung file file layout tampilan websitenya
   - Models untuk file yang komunikasi dengan database
   - entities adalah file file yang didalamnya merupakan struct gambaran dari table di database
   - controllers logic
   - config koneksi ke database
   - Controller Yang Akan Mengirimkan HTML 


package bawaan golang 
   -HTTP 
<br>

Sesuaikan folder entities dengan table yang ada di database




  ###  Fungsi masing2 URL/HREF 


-Folder Root = 
   -  Http.HandleFunc  (nama file main.go)
         = Untuk dipakai di Browser

-Folder Controller = 
   - template.ParseFiles()
         = untuk menerima operan dari file main.go 

-Folder Views / HTML = 
   - HREF
         = ketika di klik akan mengarahkan ke folder controller



