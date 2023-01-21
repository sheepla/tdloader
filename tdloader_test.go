package tdloader

import "testing"

func TestGetPath(t *testing.T) {
	fullpath, err := GetPath("_testdata/data1.json")
	if err != nil {
		t.Fatal(err, fullpath)
	}

	t.Log(fullpath)
}

func TestGetProjectRoot(t *testing.T) {
	root, err := getProjectRoot()
	if err != nil {
		t.Fatal(err, root)
	}

	if root == "" {
		t.Fatal("project root directory is empty")
	}
}

func TestGetPathGlob(t *testing.T) {
	matches, err := GetPathGlob("_testdata/*.json")
	if err != nil {
		t.Fatal(err, matches)
	}

	t.Log(matches)
}

func TestGetFile(t *testing.T) {
	f, err := GetFile("_testdata/data2.json")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(f.Name())
}

func TestGetBytes(t *testing.T) {
	content, err := GetBytes("_testdata/data3.json")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(*content)
}

func TestGetText(t *testing.T) {
	text, err := GetText("_testdata/data4.json")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(text)
}
