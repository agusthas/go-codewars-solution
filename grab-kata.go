package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Kata struct {
	Slug string   `json:"slug"`
	Name string   `json:"name"`
	Url  string   `json:"url"`
	Rank KataRank `json:"rank"`
}

type KataRank struct {
	Name string `json:"name"`
}

// Create a struct to save file path and content
type FileDescriptor struct {
	Path    string
	Content string
}

func getKataById(id string) (*Kata, error) {
	response, err := http.Get("https://www.codewars.com/api/v1/code-challenges/" + id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("Kata not found")
	}

	var kata Kata
	json.Unmarshal(responseData, &kata)

	return &kata, nil
}

func upsertFile(file *FileDescriptor) error {
	newFile, err := os.OpenFile(file.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	if _, err := newFile.WriteString(file.Content); err != nil {
		newFile.Close()
		return fmt.Errorf(err.Error())
	}

	if err := newFile.Close(); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func upsertKyuDir(dirName string, kata *Kata) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}

	kyuReadmeFileDescriptor := FileDescriptor{
		Path:    filepath.Join(dirName, "README.md"),
		Content: fmt.Sprintf("- [%s](%s)\n", kata.Name, kata.Url),
	}

	if err := upsertFile(&kyuReadmeFileDescriptor); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func createSourceFile(kyuDir string, kata *Kata) error {
	dirName := filepath.Join(kyuDir, kata.Slug)
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}

	// Readme
	readmeFileDescriptor := FileDescriptor{
		Path:    filepath.Join(dirName, "README.md"),
		Content: fmt.Sprintf("## %s\n\n## %s\n", kata.Name, kata.Url),
	}
	if err := upsertFile(&readmeFileDescriptor); err != nil {
		return fmt.Errorf(err.Error())
	}

	// Src Files
	packageName := "src"
	packagePath := filepath.Join(dirName, packageName)
	err := os.Mkdir(packagePath, os.ModePerm)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Println("Enter function declaration: ")
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf(err.Error())
	}

	goModFileDescriptor := FileDescriptor{
		Path:    filepath.Join(packagePath, "go.mod"),
		Content: fmt.Sprintf("module %s\n\ngo 1.23.2\n\n", packageName),
	}
	mainFileDescriptor := FileDescriptor{
		Path:    filepath.Join(packagePath, kata.Slug+".go"),
		Content: fmt.Sprintf("package %s\n\n%s\n", packageName, strings.Join(lines, "\n")),
	}
	testFileDescriptor := FileDescriptor{
		Path:    filepath.Join(packagePath, kata.Slug+"_test.go"),
		Content: fmt.Sprintf("package %s\n\n// TODO: Write your test cases here\n", packageName),
	}

	if err := upsertFile(&goModFileDescriptor); err != nil {
		return fmt.Errorf(err.Error())
	}
	if err := upsertFile(&mainFileDescriptor); err != nil {
		return fmt.Errorf(err.Error())
	}
	if err := upsertFile(&testFileDescriptor); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func main() {
	cmd := flag.String("id", "", "Kata ID")
	flag.Parse()

	kata, err := getKataById(*cmd)
	if err != nil {
		panic(err)
	}
	kyuDir := strings.ReplaceAll(kata.Rank.Name, " ", "-")

	// check if solution directory already exists
	// if yes, exit
	if _, err := os.Stat(filepath.Join(kyuDir, kata.Slug)); !os.IsNotExist(err) {
		panic(fmt.Errorf("Kata solution directory already exists"))
	}

	if err := upsertKyuDir(kyuDir, kata); err != nil {
		panic(err)
	}

	if err := createSourceFile(kyuDir, kata); err != nil {
		panic(err)
	}
	fmt.Printf("Kata %s grabbed successfully!\n\n", kata.Name)

	finalPath := filepath.Join(kyuDir, kata.Slug, "src")
	fmt.Printf("To start coding, run the following command:\n")
	fmt.Printf("cd %s\n", finalPath)
}
