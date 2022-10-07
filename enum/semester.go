package enum

type Semester int

var SemesterAll = []Semester{
	SemesterGanjil,
	SemesterGenap,
	SemesterPendek,
}

const (
	SemesterGanjil Semester = iota + 1
	SemesterGenap
	SemesterPendek
)

func (e Semester) GetString() string {
	return []string{"Semster Ganjil", "Semester Genap", "Semester Pendek"}[(e - 1)]
}

func (e Semester) EnumIndex() int {
	return int((e - 1))
}
