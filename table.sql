USE gdsc;

CREATE TABLE mahasiswa (
	nim INT PRIMARY KEY NOT NULL,
	nama VARCHAR(50) NOT NULL,
	jurusan VARCHAR(50) NOT NULL,
	semester INT NOT NULL,
	nomor_hp VARCHAR(50) NOT NULL
)

