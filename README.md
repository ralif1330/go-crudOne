Website Crud Golang

Golang<br>
HTML<br>
Bootstrap<br>
mySQL<br>

Jalankan mySQL di Xampp / Adminer <br>

MVC sebagai pattern<br>
Model View Controllers<br>

![mvc](/mvc.png)

Jalankan Program go run . (titik)<br>
Matikan Program CTRL + C

Folder  : Models, Views, Controllers, <br>
Folder tambahan : config, entities

Kegunaan Folder = 
   - Views untuk menampung file file layout tampilan websitenya
   - Models untuk file yang komunikasi dengan database
   - entities adalah file file yang didalamnya merupakan struct gambaran dari table di database
   - controllers logic
   - config koneksi ke database 

package bawaan golang 
   -HTTP 


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