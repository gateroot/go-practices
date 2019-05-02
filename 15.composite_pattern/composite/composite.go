package composite

import "fmt"

type DirectoryEntry interface {
	Remove()
}

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{name}
}

func (f *File) Remove() {
	fmt.Println(f.name, "を削除しました")
}

type Directory struct {
	name string
	list []DirectoryEntry
}

func NewDirectory(name string) *Directory {
	return &Directory{name, make([]DirectoryEntry, 0)}
}

func (d *Directory) Add(entry DirectoryEntry) {
	d.list = append(d.list, entry)
}

func (d *Directory) Remove() {
	for _, e := range d.list {
		e.Remove()
	}
	fmt.Println(d.name, "を削除しました")
}
