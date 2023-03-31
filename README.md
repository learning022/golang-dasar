# golang-dasar

Ini adalah repository untuk yang ingin belajar bahasa pemrograman Go

# langkah yang dilakukan

- Membuat file baru dengan nama main.go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World! Yuk Belajar Golang Dasar")
}
```

### Kenapa disini pakai package main, saya jelaskan dulu ya

1. Package Executabel (main) : Program yang akan di eksekusi langsung pada cmd
2. Library (selain main): sebelum eksekusi harus ada sebuah file yang dijadikan sebagai executable

sebelum running silahkan buat modul requirement dan sums nya terlebih dahulu agar golangnya dapat di jalankan, caranya seperti ini, silahkan copas

```go
go mod init masukanurlpackagenya
```

ikuti langkahnya
![image](https://user-images.githubusercontent.com/15622730/229246571-cfa8d2d3-d7fa-4fe7-a44b-d309fe056e98.png)

Langkah Selanjutnya belajat Looping
![image](https://user-images.githubusercontent.com/15622730/229247863-03300bdb-9ed6-4a5c-819b-10591f10ffa6.png)<br>
Disini kita akan melakukan looping sebanyak 15 kali <br>
berikut script codenya
```go
	for i := 1; i <= 15; i++ {
		fmt.Println("Hello, World! Yuk Belajar Golang Dasar Looping : ", i)

	}
```

![image](https://user-images.githubusercontent.com/15622730/229247913-0a9f927b-5929-4cac-bb78-983133a58af9.png)<br>

### Membuat Looping ke tiga
```go
package main

import "fmt"

func Looping2() {
	// berikut loopingnya ketiga
	title := "Ngoding Golang bersama RPL Class"

	for index, letter := range title {
		fmt.Println("index :", index, "leter :", string(letter))
	}
}
```

Berikut hasilnya
![image](https://user-images.githubusercontent.com/15622730/229248977-500d12f8-b62a-4fb3-9d94-e2522794b1aa.png)








