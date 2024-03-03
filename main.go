package main

import(
	"fmt"
	"bufio"
	"os"
	"github.com/Safmica/CRUD-GOLANG/database_connection"
        "github.com/Safmica/CRUD-GOLANG/crud"
)

func main () {
	status := true
	while status {
		MenuAwal()
	}
}

func MenuAwal () {
	fmt.Println(`=====DATABASE MAHASISWA=====
	1. Membuat Data Mahasiswa
	2. Melihat Data Mahasiswa
	3. Memperbarui Data Mahasiswa
	4. Menghapus Data Mahasiswa
	5. Keluar`)

}

