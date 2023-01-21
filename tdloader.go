package tdloader

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var ErrGetProjectRoot = errors.New("failed to get project root directory")
var ErrTestDataNotExists = errors.New("test data not exists")

func GetPath(relative string) (string, error) {
	root, err := getProjectRoot()
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrGetProjectRoot, err)
	}

	fullpath := filepath.Join(root, relative)
	if !fileExists(fullpath) {
		return "", fmt.Errorf("%w: (%s)", ErrTestDataNotExists, fullpath)
	}

	return fullpath, nil
}

func MustGetPath(relative string) string {
	fullpath, err := GetPath(relative)
	if err != nil {
		panic(err)
	}

	return fullpath
}

func GetPathGlob(relative string) ([]string, error) {
	root, err := getProjectRoot()
	if err != nil {
		return []string{}, fmt.Errorf("%w: %s", ErrGetProjectRoot, err)
	}

	fullpath := filepath.Join(root, relative)
	matches, err := filepath.Glob(fullpath)
	if err != nil {
		return []string{}, err
	}

	return matches, nil
}

func MustGetPathGlob(relative string) []string {
	matches, err := GetPathGlob(relative)
	if err != nil {
		panic(err)
	}

	return matches
}

func GetFile(relative string) (*os.File, error) {
	fullpath, err := GetPath(relative)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func MustGetFile(relative string) *os.File {
	file, err := GetFile(relative)
	if err != nil {
		panic(err)
	}

	return file
}

func GetBytes(relative string) (*[]byte, error) {
	file, err := GetFile(relative)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func MustGetBytes(relative string) *[]byte {
	content, err := GetBytes(relative)
	if err != nil {
		panic(err)
	}

	return content
}

func GetText(relative string) (string, error) {
	bytes, err := GetBytes(relative)
	if err != nil {
		return "", err
	}

	return string(*bytes), nil
}

func MustGetText(relative string) string {
	text, err := GetText(relative)
	if err != nil {
		panic(err)
	}

	return text
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

func getProjectRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
