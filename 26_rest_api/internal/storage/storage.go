package storage

type Storage interface {
	CreatStudent(name string, email string, age int) (int64, error)
}