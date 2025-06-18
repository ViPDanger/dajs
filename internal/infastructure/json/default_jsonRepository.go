package json

import (
	"DAJ/internal/domain/entity"
	"os"
	"path/filepath"
)

var defaultFileType = ".json"

type defaultJSONRepository[T entity.Identifiable, Tdto any] struct {
	toDTO         func(T) Tdto
	ToEntity      func(Tdto) T
	filePaths     map[string]string
	pathFunc      func(*T) string
	fileDirectory string
	fileType      string
	file          jsonDB[Tdto]
}

func NewJSONRepository[T entity.Identifiable, Tdto any](filepath string, toDTO func(T) Tdto, ToEntity func(Tdto) T, pathFunc func(*T) string) (f *defaultJSONRepository[T, Tdto], err error) {
	var db jsonDB[Tdto]
	repository := defaultJSONRepository[T, Tdto]{
		fileDirectory: filepath,
		fileType:      defaultFileType,
		file:          db,
		toDTO:         toDTO,
		ToEntity:      ToEntity,
	}
	mass, err := repository.getExistingFilePaths()
	if err != nil {
		return nil, err
	}
	repository.filePaths = make(map[string]string, len(mass))
	for i := range mass {
		dto, err := repository.file.Compile(mass[i])
		if err != nil {
			return nil, err
		}
		object := ToEntity(*dto)
		repository.filePaths[object.GetID()] = mass[i]
	}
	repository.pathFunc = pathFunc
	return &repository, nil
}

func (r *defaultJSONRepository[T, Tdto]) GetArray(ids []string) (ret []T, err error) {
	for i := range ids {
		ids[i] = r.setPath(ids[i])
	}
	objects, err := r.file.CompileArray(ids)
	if err != nil {
		return
	}
	ret = make([]T, len(objects))
	for i := range objects {
		ret[i] = r.ToEntity(objects[i])
	}
	return
}

func (r *defaultJSONRepository[T, Tdto]) Insert(object *T) error {
	path := r.pathFunc(object)
	r.filePaths[(*object).GetID()] = path
	dto := r.toDTO(*object)
	return r.file.Add(r.pathFunc(object), &dto)
}
func (r *defaultJSONRepository[T, Tdto]) GetByID(id string) (ret *T, err error) {

	dto, err := r.file.Compile(r.setPath(id))
	object := r.ToEntity(*dto)
	return &object, err

}
func (r *defaultJSONRepository[T, Tdto]) GetAll() (ret []T, err error) {
	dtos, err := r.file.CompileDir(r.fileDirectory)
	if err != nil {
		return
	}
	ret = make([]T, len(dtos))
	for i := range dtos {
		ret[i] = r.ToEntity(dtos[i])
	}
	return
}
func (r *defaultJSONRepository[T, Tdto]) Update(object *T) error {
	path := r.pathFunc(object)
	dto := r.toDTO(*object)
	return r.file.Patch(path, &dto)

}
func (r *defaultJSONRepository[T, Tdto]) Delete(id string) (err error) {
	return r.file.Delete(r.setPath(id))
}

// Вспомогательные функции
func (r *defaultJSONRepository[T, Tdto]) getExistingFilePaths() ([]string, error) {
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

func (r *defaultJSONRepository[T, Tdto]) setPath(id string) string {
	if itemPath, exists := r.filePaths[id]; exists {
		return itemPath
	}
	return r.fileDirectory + id + r.fileType
}
