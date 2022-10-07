package enum

type Entity int

const (
	Admin Entity = iota
	Student
	Lecture
	Parent
)

func (e Entity) getString() string {
	return []string{"Admin", "Mahasiswa", "Dosen", "Admin"}[e]
}

func (e Entity) EnumIndex() int {
	return int(e)
}
