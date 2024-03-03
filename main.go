package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	_ "github.com/Safmica/CRUD-GOLANG/database_connection"
        "github.com/Safmica/CRUD-GOLANG/crud"
)

func main () {
	status := true
	for status {
		MenuAwal()

		fmt.Print("Pilihan anda = ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			input := scanner.Text()

			fmt.Println()
			switch input {
			case "1":
				fmt.Println("=====Membuat Data Mahasiswa=====")
				crud.CreateMahasiswa(InputMahasiswa())
			case "2":
		    	        fmt.Println("=====Melihat Data Mahasiswa=====")
 				crud.ReadMahasiswa()
			case "3":
				fmt.Println("=====Memperbarui Data Mahasiswa=====")
				crud.UpdateMahasiswa(InputNIM(),InputMahasiswa())
			case "4":
				fmt.Println("=====Menghapus Data Mahasiswa=====")
				crud.DeleteMahasiswa(InputNIM())
			case "5":
				fmt.Println("Anda Berhasil Keluar!")
				status = false
				return
			default:
				fmt.Println("Pilihan tidak valid, Mohon hanya pilih 1-5")
			}
			fmt.Println()
			break
		}
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

func InputMahasiswa() crud.Mahasiswa {
	var mahasiswa crud.Mahasiswa

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("NIM : ")
	scanner.Scan()
	mahasiswa.NIM = scanner.Text()

	fmt.Print("Nama : ")
        scanner.Scan()
        mahasiswa.Nama = scanner.Text()

	fmt.Print("Jurusan : ")
        scanner.Scan()
        mahasiswa.Jurusan = scanner.Text()

	fmt.Print("Semester : ")
        scanner.Scan()
	input := scanner.Text()
	number, _ := strconv.Atoi(input)
	mahasiswa.Semester = number

	fmt.Print("Nomor HP : ")
        scanner.Scan()
        mahasiswa.Nomor_hp = scanner.Text()

	return mahasiswa
}

func InputNIM() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan NIM : ")
	scanner.Scan()
	NIM := scanner.Text()

	return NIM
}
