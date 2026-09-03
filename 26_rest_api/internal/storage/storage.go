package storage

import "github.com/itz-prashant/student-api/internal/types"

type Storage interface {
	CreatStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
