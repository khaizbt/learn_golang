package main

type Luas interface { //membuat kontrak dengan nama interface
	HitungLuas() int //untuk memenuhi kontrak interface tersebut harus mempunyai aturan mainnya harus mempunyai sebuah method dengan nama hitungluas
	//Interface hanya mendefinisikan nama gfungsi dan keluarannya

	HitungKeliling() int
}

type Persegi struct {
	Sisi int
}

func main(){

}

func (persegi Persegi) HitungLuas() int  { //Nama Fungsi dan Parameter dan juga outputnya harus sama perseis seperti di Interface
	return persegi.Sisi * persegi.Sisi
	//Artinya si Persegi sudah memenuhi kontraknya Interface Luas karena dalam fungsi ini sudah memenuhi smua yang diminta oleh fungsi HitungLuas dalam interface Luas
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func (persegipanjang PersegiPanjang) HitungKeliling() int { //Jadi kesimpulannya, Interface itu semacam aturan yang harus dipenuhi
	return 2*(persegipanjang.Panjang + persegipanjang.Lebar)
}