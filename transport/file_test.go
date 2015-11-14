package transport

import "testing"

func TestZip(t *testing.T) {
	dest := "/tmp/TestZip.zip"
	defer FileRemove(dest)

	err := Zip(".", dest)
	if err != nil {
		t.Errorf("zip failed:%v", err)
	}
}

func TestUnzip(t *testing.T) {
	dest := "/tmp/TestZip.zip"
	defer FileRemove(dest)

	err := Zip(".", dest)
	if err != nil {
		t.Errorf("zip failed:%v", err)
	}
	err = Unzip(dest, "/tmp")
	if err != nil {
		t.Errorf("unzip failed:%v", err)
	}
	FileRemove("/tmp/file.go")
	FileRemove("/tmp/file_test.go")
	FileRemove("/tmp/http.go")
}
