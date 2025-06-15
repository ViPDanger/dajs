package fileRepository

import (
	"DAJ/internal/db"
	"DAJ/internal/domain/entity"
	"errors"
	"os"
	"path/filepath"
)

var defaultFileType = ".json"

type fileRepository[T entity.Identifiable] struct {
	ItemIds       map[string]string
	fileDirectory string
	fileType      string
	file          db.FileDB[T]
}

func NewFileRepository[T entity.Identifiable](filepath string) (f *fileRepository[T], err error) {
	var db db.FileDB[T]

	repository := fileRepository[T]{fileDirectory: filepath, fileType: defaultFileType, file: db}
	mass, err := repository.getAllFilePaths()
	if err != nil {
		return nil, err
	}

	repository.ItemIds = make(map[string]string, len(mass))
	for i := range mass {
		object, err := repository.file.Compile(mass[i])
		if err != nil {
			return nil, err
		}
		repository.ItemIds[(*object).GetID()] = filepath
	}

	return &repository, nil
}
func (r *fileRepository[T]) GetArray(ids []string) (ret []T, err error) {
	for i := range ids {
		ids[i] = r.fileDirectory + ids[i] + r.fileType
	}
	return r.file.CompileArray(ids)
}

func (r *fileRepository[T]) Insert(item *T) error {
	id := (*item).GetID()
	r.ItemIds[id] = id

	return r.file.Add(r.fileDirectory+r.ItemIds[id]+r.fileType, item)
}
func (r *fileRepository[T]) GetByID(id string) (ret *T, err error) {
	if itemId, exists := r.ItemIds[id]; exists {
		return r.file.Compile(r.fileDirectory + itemId + r.fileType)
	}
	return r.file.Compile(r.fileDirectory + id + r.fileType)
}
func (r *fileRepository[T]) GetAll() (ret []T, err error) {
	return r.file.CompileDir(r.fileDirectory)
}
func (r *fileRepository[T]) Update(item *T) error {
	id := (*item).GetID()
	if itemId, exists := r.ItemIds[id]; exists {
		return r.file.Patch(r.fileDirectory+itemId+r.fileType, item)
	}
	return errors.New("no object found")

}
func (r *fileRepository[T]) Delete(id string) (err error) {
	if itemId, exists := r.ItemIds[id]; exists {
		return r.file.Delete(r.fileDirectory + itemId + r.fileType)
	}
	return r.file.Delete(r.fileDirectory + id + r.fileType)
}

func (r *fileRepository[T]) getAllFilePaths() ([]string, error) {
	var files []string
	err := filepath.Walk(r.fileDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // если ошибка доступа к файлу/папке — пробрасываем её
		}
		if !info.IsDir() {
			files = append(files, path) // добавляем только файлы
		}
		return nil
	})

	return files, err
}
