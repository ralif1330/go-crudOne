Website Crud Golang

Golang
HTML
Bootstrap
mySQL

MVC sebagai pattern<br>
Model View Controllers
Jalankan Program go run . (titik)
Matikan Program CTRL + C

Folder  : Models, Views, Controllers, 
Folder tambahan : config, entities

Kegunaan Folder = 
   - Views untuk menampung file file layout tampilan websitenya
   - Models untuk file yang komunikasi dengan database
   - entities file file yang didalamnya merupakan struct gambaran dari table di database
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