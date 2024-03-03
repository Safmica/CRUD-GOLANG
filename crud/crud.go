package crud

import (
        "context"
        "fmt"
	"database/sql"
        "github.com/Safmica/CRUD-GOLANG/database_connection"
)

type Mahasiswa struct {
        NIM string
        Nama string
        Jurusan string
        Semester int
        Nomor_hp string
}

func CreateMahasiswa (mahasiswa Mahasiswa){
        db := database_connection.GetConnection()

        ctx := context.Background()
        _, err := db.ExecContext(ctx, "INSERT INTO mahasiswa (nim, nama, jurusan,semester, nomor_hp) VALUE(?, ?, ?, ?, ?)", mahasiswa.NIM, mahasiswa.Nama, mahasiswa.Jurusan, mahasiswa.Semester, mahasiswa.Nomor_hp)
        if err != nil {
                fmt.Println("Error : ", err)
                return
        }
        fmt.Println("Data  Mahasiswa Berhasil Dibuat")
}

func ReadMahasiswa () *sql.Rows {
	db := database_connection.GetConnection()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT * FROM mahasiswa")
	if err != nil {
		fmt.Println("Error : ", err)
	}

	return rows
}

func UpdateMahasiwa (NIM string, mahasiswa Mahasiswa) {
	db := database_connection.GetConnection()
	defer db.Close()

	ctx := context.Background()
	result, err := db.ExecContext(ctx, "UPDATE mahasiswa SET nim=?, nama=?, jurusan=?, semester=?, nomor_hp=? WHERE nim=?", mahasiswa.NIM, mahasiswa.Nama, mahasiswa.Jurusan, mahasiswa.Semester, mahasiswa.Nomor_hp, NIM)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Printf("Tidak ada Mahasiswa dengan NIM : %s. Tidak ada data yang diupdate.\n", NIM)
		return
	}

	fmt.Println("Data Mahasiswa Berhasil diUpdate!")
}

func DeleteMahasiwa (NIM string) {
        db := database_connection.GetConnection()
        defer db.Close()

        ctx := context.Background()
        result, err := db.ExecContext(ctx, "DELETE mahasiswa WHERE nim=?", NIM)

        if err != nil {
                fmt.Println("Error : ", err)
                return
        }

        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 0 {
                fmt.Printf("Tidak ada Mahasiswa dengan NIM : %s. Tidak ada data yang dihapus.\n", NIM)
                return
        }

        fmt.Println("Data Mahasiswa Berhasil diHapus!")
}
