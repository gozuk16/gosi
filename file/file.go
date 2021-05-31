// file ファイルやディレクトリを扱うライブラリ
package file

import (
	"fmt"
	"os"
	"path/filepath"
)

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

	for _, d := range dirs {
		if d.IsDir() {
		}
	}

	// keyFileが無ければ除外

	return nil, nil
}

// isDirExist 渡されたpathがディレクトリとして存在するか
func isDirExist(path string) bool {
	f, err := os.Stat(path)
	return os.IsNotExist(err) || !f.IsDir()
}
