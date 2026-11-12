package utils

import (
	"path/filepath"
	"runtime"
)

// ProjectFilePath возвращает абсолютный путь к файлу относительно корня проекта,
// чтобы тесты находили файлы независимо от того, из какого пакета они запущены
func ProjectFilePath(relativePath string) string {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("Не удалось определить путь к текущему файлу")
	}
	projectRoot := filepath.Join(filepath.Dir(currentFile), "..", "..")
	return filepath.Join(projectRoot, relativePath)
}
