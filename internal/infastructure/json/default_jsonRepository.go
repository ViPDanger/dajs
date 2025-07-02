package json

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/ViPDanger/dajs/internal/domain/entity"
)

var defaultFileType = ".json"

type defaultJSONRepository[T entity.Identifiable, Tdto any] struct {
	toDTO         func(T) Tdto
	ToEntity      func(Tdto) T
	filePaths     *sync.Map
	pathFunc      func(*T) string
	fileDirectory string
	fileType      string
	file          jsonDB[Tdto]
}

func NewJSONRepository[T entity.Identifiable, Tdto any](filepath string, toDTO func(T) Tdto, ToEntity func(Tdto) T, pathFunc func(*T) string) (f *defaultJSONRepository[T, Tdto], err error) {
	if filepath == "" || toDTO == nil || ToEntity == nil || pathFunc == nil {
		var o T
		TName := reflect.TypeOf(o).Name()
		return nil, fmt.Errorf("NewJSONRepository[%s](): Some parameters are empty", TName)
	}
	var db jsonDB[Tdto]
	repository := defaultJSONRepository[T, Tdto]{
		fileDirectory: filepath,
		fileType:      defaultFileType,
		filePaths:     &sync.Map{},
		file:          db,
		toDTO:         toDTO,
		ToEntity:      ToEntity,
	}
	mass, err := repository.getExistingFilePaths()
	if err != nil {
		return nil, err
	}
	for i := range mass {
		dto, err := repository.file.Compile(mass[i])
		if err != nil {
			return nil, err
		}
		object := ToEntity(*dto)
		repository.filePaths.Store(object.GetID(), mass[i])
	}
	repository.pathFunc = pathFunc
	return &repository, nil
}

func (r *defaultJSONRepository[T, Tdto]) GetArray(ids []string) (ret []T, err error) {

	for i := range ids {
		if ids[i] == "" {
			var o T
			TName := reflect.TypeOf(o).Name()
			return nil, fmt.Errorf("JSONRepositrory[%s].GetArray()/Some Id in id array is empty", TName)
		}
		ids[i] = r.setPath(ids[i])
	}
	objects, err := r.file.CompileArray(ids)
	if err != nil {
		var o T
		TName := reflect.TypeOf(o).Name()
		return nil, fmt.Errorf("JSONRepositrory[%s].GetArray()/%w", TName, err)
	}
	ret = make([]T, len(objects))
	for i := range objects {
		ret[i] = r.ToEntity(objects[i])
	}
	return
}

func (r *defaultJSONRepository[T, Tdto]) Insert(object *T) error {
	if object == nil {
		var o T
		TName := reflect.TypeOf(o).Name()
		return fmt.Errorf("JSONRepositrory[%s].Insert(): Nill pointer", TName)
	}
	path := r.pathFunc(object)
	r.filePaths.Store((*object).GetID(), path)
	dto := r.toDTO(*object)
	return r.file.Add(r.pathFunc(object), &dto)
}
func (r *defaultJSONRepository[T, Tdto]) GetByID(id string) (ret *T, err error) {
	dto, err := r.file.Compile(r.setPath(id))
	if err != nil {
		var o T
		TName := reflect.TypeOf(o).Name()
		return nil, fmt.Errorf("JSONRepositrory[%s].GetByID()/%w", TName, err)
	}
	object := r.ToEntity(*dto)
	return &object, err

}
func (r *defaultJSONRepository[T, Tdto]) GetAll() (ret []T, err error) {
	dtos, err := r.file.СompileDir(r.fileDirectory)
	if err != nil {
		var o T
		TName := reflect.TypeOf(o).Name()
		return nil, fmt.Errorf("JSONRepositrory[%s].GetAll()/%w", TName, err)
	}
	ret = make([]T, len(dtos))
	for i := range dtos {
		ret[i] = r.ToEntity(dtos[i])
	}
	return
}
func (r *defaultJSONRepository[T, Tdto]) Update(object *T) error {
	if object == nil {
		var o T
		TName := reflect.TypeOf(o).Name()
		return fmt.Errorf("JSONRepositrory[%s].Update(): Nill pointer", TName)
	}
	path := r.pathFunc(object)
	dto := r.toDTO(*object)
	return r.file.Patch(path, &dto)

}
func (r *defaultJSONRepository[T, Tdto]) Delete(id string) (err error) {
	if id == "" {
		var o T
		TName := reflect.TypeOf(o).Name()
		return fmt.Errorf("JSONRepositrory[%s].Delete(): Empty id", TName)
	}
	return r.file.Delete(r.setPath(id))
}

// Вспомогательные функции
func (r *defaultJSONRepository[T, Tdto]) getExistingFilePaths() ([]string, error) {
	var files []string
	err := filepath.Walk(r.fileDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			var o T
			TName := reflect.TypeOf(o).Name()
			return fmt.Errorf("JSONRepositrory[%s].getExistingFilePath()/%w", TName, err)

		}
		if !info.IsDir() && strings.Contains(info.Name(), defaultFileType) {
			files = append(files, path) // добавляем только файлы
		}
		return nil
	})

	return files, err
}

func (r *defaultJSONRepository[T, Tdto]) setPath(id string) string {
	if itemPath, exists := r.filePaths.Load(id); exists {
		return itemPath.(string)
	}
	return r.fileDirectory + "/" + id + r.fileType
}
