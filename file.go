package goss

import (
	"fmt"
	"os"
	"path/filepath"
)

// CrawlDirs 渡されたcriteriaの直下にあるディレクトリの中に、keyFileが存在している場合のみディレクトリ名を配列で返す
// keyFileが空なら直下のディレクトリを全部返す
func CrawlDirs(criteria string, keyFile string) ([]string, error) {
	if criteria == "" {
		return nil, fmt.Errorf("criteria is nothing.")
	}

	// criteria存在チェック
	p, err := filepath.Abs(criteria)
	if err != nil {
		return nil, err
	}
	p = filepath.Clean(p)
	if !isDirExist(p) {
		return nil, fmt.Errorf("criteria dir is not found.")
	}

	// criteriaの下にあるディレクトリ取得
	dirs, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, d := range dirs {
		if d.IsDir() {
			if keyFile == "" {
				// keyFileが空なら直下のディレクトリを全部返す
				var dir string = d.Name()
				result = append(result, dir)
			} else {
				k, err := filepath.Abs(filepath.Join(criteria, d.Name(), keyFile))
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				// keyFileがあるディレクトリのみ返す
				if isFileExist(k) {
					var dir string = d.Name()
					result = append(result, dir)
				}
			}
		}
	}

	return result, nil
}

// isDirExist 渡されたpathがディレクトリとして存在するか
func isDirExist(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

// isFileExist 渡されたpathがファイルとして存在するか
func isFileExist(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
}
