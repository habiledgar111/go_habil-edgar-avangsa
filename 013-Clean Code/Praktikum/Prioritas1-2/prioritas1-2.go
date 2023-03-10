package main

type kendaraan struct {
	totalroda       int
	kecepatanperjam int
	//nama dari object, variable, atau function sebaiknya menggunakan camelCase,snake_case atau PascalCase
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	//penamaan variable harus mudah dipahami
	//dapat diganti menjadi mobil *mobil
	m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	//penamaan variable harus mudah dipahami
	//dapat diganti menjadi mobil *mobil
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
	//untuk penambahan dan asign ke dalam variable yang sama dapat di singkat
	//menjado mobil.kecepatanPerjam += kecepatanBaru
}

func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()
}
