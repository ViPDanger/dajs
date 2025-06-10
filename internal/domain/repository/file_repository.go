package repository

import (
	"DAJ/internal/db"
)

type FileRepository[T any] struct {
	filepath string
	file     db.File[T]
}

func NewFileRepository[T any](filepath string) *FileRepository[T] {
	return &FileRepository[T]{filepath: filepath, file: db.File[T]{}}
}

func (r *FileRepository[T]) New(id string, item *T) error {
	return r.file.Add(r.filepath+id+".json", item)
}

func (r *FileRepository[T]) GetByID(id string) (ret *T, err error) {
	return r.file.Compile(r.filepath + id + ".json")
}
func (r *FileRepository[T]) GetAll() (ret *[]T, err error) {
	return r.file.CompileDir(r.filepath)
}
func (r *FileRepository[T]) Set(id string, item *T) error {
	return r.file.Patch(r.filepath+id+".json", item)

}
func (r *FileRepository[T]) Delete(id string) (err error) {
	return r.file.Delete(r.filepath + id + ".json")
}
