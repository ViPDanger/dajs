package json

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
)

type jsonDB[T any] struct {
	mutex sync.RWMutex
}

// Просмотреть файл
func (f *jsonDB[T]) Compile(filepath string) (object *T, err error) {
	object = new(T)
	data, err := f.read(filepath)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, object)
	return
}

// Просмотреть массив файлов
func (f *jsonDB[T]) CompileArray(filepaths []string) (objects []T, err error) {
	objects = make([]T, len(filepaths))
	for i, filepath := range filepaths {
		data, err := f.read(filepath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &objects[i])
		if err != nil {
			return nil, err
		}
	}
	return
}

// Изменить существующий файл
func (f *jsonDB[T]) Patch(filename string, patch *T) error {
	// Читаем оригинальный файл
	f.mutex.RLock()
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var original T
	if err := json.Unmarshal(data, &original); err != nil {
		return err
	}
	f.mutex.RUnlock()
	f.mutex.Lock()
	// Обновляем оригинальный объект
	if err := applyPatch(&original, patch); err != nil {
		f.mutex.Unlock()
		return err
	}
	f.mutex.Unlock()
	// Сохраняем обратно
	updated, err := json.MarshalIndent(original, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, updated, 0644)
}

func (f *jsonDB[T]) Add(filename string, item *T) error {
	// Проверяем отсутствие наличия файла
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if _, err := os.Stat(filename); err == nil {
		return errors.New("file already exists")
	}
	// Сохраняем обратно
	updated, err := json.MarshalIndent(*item, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, updated, 0644)
}

// read .json file
func (f *jsonDB[T]) read(filepath string) ([]byte, error) {
	f.mutex.RLock()
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	f.mutex.RUnlock()
	re := regexp.MustCompile(`(?s)/\*.*?\*/`)
	return re.ReplaceAll(data, []byte("")), nil
}

// applyPatch обновляет поля original значениями из patch, если они не default.
func applyPatch[T any](original *T, patch *T) error {

	return recursivePatch(reflect.ValueOf(original).Elem(), reflect.ValueOf(patch).Elem())
}
func recursivePatch(origVal, patchVal reflect.Value) error {
	for i := 0; i < patchVal.NumField(); i++ {
		patchField := patchVal.Field(i)
		origField := origVal.Field(i)

		if !patchField.IsValid() || !origField.CanSet() {
			continue
		}

		switch patchField.Kind() {
		case reflect.Struct:
			// Рекурсивный вызов для вложенной структуры
			err := recursivePatch(origField, patchField)
			if err != nil {
				return err
			}

		case reflect.Pointer:
			if !patchField.IsNil() {
				if origField.IsNil() {
					// Если в original nil, создаём новый экземпляр
					origField.Set(reflect.New(patchField.Type().Elem()))
				}
				// Рекурсивно патчим содержимое указателя
				err := recursivePatch(origField.Elem(), patchField.Elem())
				if err != nil {
					return err
				}
			}

		default:
			if !reflect.DeepEqual(patchField.Interface(), reflect.Zero(patchField.Type()).Interface()) {
				origField.Set(patchField)
			}
		}
	}
	return nil
}

// Просмотреть все обьекты в папке
func (f *jsonDB[T]) СompileDir(dir string) ([]T, error) {
	var result []T
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Пропустить с ошибкой
		}

		// Пропустить директории
		if d.IsDir() {
			return nil
		}

		// Пропустить не-json файлы (можно адаптировать под нужный формат)
		if filepath.Ext(path) != defaultFileType {
			return nil
		}

		obj, err := f.Compile(path)
		if err != nil {
			return err
		}
		result = append(result, *obj)
		return nil
	})
	return result, err
}

func (f *jsonDB[T]) Delete(filename string) error {
	f.mutex.Lock()
	if err := os.Remove(filename); err != nil {
		f.mutex.Unlock()
		return err
	}
	f.mutex.Unlock()
	return nil
}
