package gosi

import (
	"log"
	"reflect"
	//	"os"

	"testing"
)

const (
	testIsmHome string = "testdata"
	testIsmLog  string = "var/log"
)

// setLogger ログ出力の調整
func setLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func TestMain(m *testing.M) {
	setLogger()

	m.Run()
}

func TestInitializeMonitoring(t *testing.T) {
	testDir := "testdata"
	//currentdir, _ := os.Getwd()
	//td = filepath.Clean(filepath.Join(currentdir, testDir))

	// keyFile指定
	except := []string{"testdir1"}
	dirs, err := CrawlDirs(testDir, "keyfile")
	if err != nil {
		t.Errorf("CrawlDirs Failed, %s", err)
		return
	}
	if !reflect.DeepEqual(except, dirs) {
		t.Errorf("CrawlDirs unmatch, %v : %v", except, dirs)
	}

	// keyFile省略
	except = []string{"testdir1", "testdir2"}
	dirs, err = CrawlDirs(testDir)
	if err != nil {
		t.Errorf("CrawlDirs Failed, %s", err)
		return
	}
	if !reflect.DeepEqual(except, dirs) {
		t.Errorf("CrawlDirs unmatch, %v : %v", except, dirs)
	}

	// 2つめのkeyFile指定は無視
	except = []string{"testdir1"}
	dirs, err = CrawlDirs(testDir, "keyfile", "aaa")
	if err != nil {
		t.Errorf("CrawlDirs Failed, %s", err)
		return
	}
	if !reflect.DeepEqual(except, dirs) {
		t.Errorf("CrawlDirs unmatch, %v : %v", except, dirs)
	}
}

/*
 */
