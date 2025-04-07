package search

import (
	"io/fs"
	"path/filepath"
	"strings"

	"backend/models"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

func LoadParquetDataFromDir(dir string) ([]models.LogEntry, error) {
	var allEntries []models.LogEntry

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".parquet") {
			entries, err := loadSingleFile(path)
			if err != nil {
				return err
			}
			allEntries = append(allEntries, entries...)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return allEntries, nil
}

func loadSingleFile(path string) ([]models.LogEntry, error) {
	fr, err := local.NewLocalFileReader(path)
	if err != nil {
		return nil, err
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(models.LogEntry), 4)
	if err != nil {
		return nil, err
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	all := make([]models.LogEntry, 0, num)

	for i := 0; i < num; i += 100 {
		count := 100
		if i+count > num {
			count = num - i
		}
		batch := make([]models.LogEntry, count)
		if err := pr.Read(&batch); err != nil {
			break
		}
		all = append(all, batch...)
	}
	return all, nil
}
