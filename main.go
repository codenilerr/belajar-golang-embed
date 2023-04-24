package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"

	_ "embed"
)

//go:embed version.txt
var version string

//go:embed alpukat.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	// untuk yang version
	fmt.Println(version)

	// untuk yang logo
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	// kita akan pindahkan/save lagi menggunakan ioutil.writefile ke logo_new.png
	if err != nil {
		panic(err)
	}

	// untuk yang path matcher
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
