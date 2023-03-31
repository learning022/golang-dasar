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
