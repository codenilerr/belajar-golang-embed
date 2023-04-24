package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

/** untuk melakukan golang embed, kita mesti membuat variabel di luar function, sehingga untuk membuat variabel,
kita mesti pakai yang var, bukan yang := (karena := hanya berlaku untuk pembuatan variabel di dalam function.
sementara untuk golang embed kita tidak boleh membuat variabel di dalam function*/

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed alpukat.png
var logo []byte

/** kita akan load data gambar dari file alpukat.png ke dalam variabel logo. lalu kita akan save lagi ke
file baru yakni "logo_new.png". nanti kita akan cek apakah datanya beneran sama atau tidak
*/

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	// kita akan pindahkan/save lagi menggunakan ioutil.writefile ke logo_new.png
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var document embed.FS

func TestMultipleFiles(t *testing.T) {
	// readfile adalah method untuk struct FS
	/** return values nya ada 2, yang pertama adalah []byte dan yang kedua adlah error.
	jadi kalau ternyata file a.txt, b.txt, c.txt adalah bertipe data text yang notabene nya adalah string,
	lalu ketika dia dikenai Readfile() yang akan menghasilkan return value dalam bentuk []byte,
	maka di akhirnya nanti kita akan konversi lagi dia ke string
	*/
	a, _ := document.ReadFile("files/a.txt")
	// disini error nya kita ignore aja
	fmt.Println(string(a))

	b, _ := document.ReadFile("files/b.txt")
	// disini error nya kita ignore aja
	fmt.Println(string(b))

	c, _ := document.ReadFile("files/c.txt")
	// disini error nya kita ignore aja
	fmt.Println(string(c))
}

//go:embed files/*.txt
/** kalau dia bentuknya hanya //go:embed files/* itu artinya seluruh file yang ada di dalam folder
files, tidak perduli apakah ekstensi nya .txt ataupun yang lain semuanya akan di-embed.
tetapi kalau //go:embed files/*.txt maka yang di-embed seluruh file yang hanya ber-ekstensi .txt di
folder files saja.
*/
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	// error nya kita ignore aja dulu
	for _, entry := range dirEntries {
		// disini index nya kita ignore saja dengan menggunakan _
		/** balikan nilai dari dirEntries sebenernya adalah []Fs.DirEntry. entry disini adalah DirEntry,
		tetapi kita harus cek dulu DirEntry nay ini adalah directory atau bukan. karena kan ga perduli
		apakah dia directory.
		*/
		if !entry.IsDir() {
			// ini bacanay jika bukan directory
			// jika hanya ingin tau namanya
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			// disini kita akan baca masing-masing file nya. return value error nya kita ignore aja
			fmt.Println(string(file))
		}
	}
}

func SayHello(name string) string {
	return "Hello " + name
}