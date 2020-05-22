package writer

import (
	"asher/internal/api"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type MigrationWriterHandler struct {
	api.WriterHandler
	migrationFileName string
}

func NewMigrationWriterHandler() *MigrationWriterHandler {
	return &MigrationWriterHandler{
		migrationFileName: "",
	}
}

func (writerHandler *MigrationWriterHandler) BeforeHandle(emitterFile api.EmitterFile) {
	writerHandler.migrationFileName = generateMigration(emitterFile.FileName())
}

func (writerHandler *MigrationWriterHandler) Handle(emitterFile api.EmitterFile) bool {
	if writerHandler.migrationFileName == "" {
		return false
	}
	// Fetching the current working directory
	basePath, err := os.Getwd()
	if err != nil {
		return false
	}

	fileContent := PhpTag + emitterFile.Generator().String()
	pathToMigrationFile := basePath + "/" + emitterFile.Path() + "/" + writerHandler.migrationFileName +
		PhpExtension
	err = ioutil.WriteFile(pathToMigrationFile, []byte(fileContent), 0644)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (writerHandler *MigrationWriterHandler) AfterHandle(emitterFile api.EmitterFile) {
}

func generateMigration(name string) string {
	migrationFileName := ""
	// Fetching the current working directory
	basePath, err := os.Getwd()
	if err != nil {
		return ""
	}

	fmt.Println("Path : " + basePath)

	// Creating a new migration
	cmd := exec.Command("php", "artisan", "make:migration", name)
	cmd.Dir = basePath
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err == nil {
		message := out.String()
		// Extracting the migration name from the command line output
		migrationFileName = parseMigrationName(message)
	}

	return migrationFileName
}

func parseMigrationName(input string) string {
	migrationName := ""
	pattern := regexp.MustCompile("Created Migration: ([\\d\\w_]*)")
	result := pattern.FindStringSubmatch(strings.Trim(input, "\n"))
	if len(result) > 1 {
		migrationName = result[1]
	}
	return migrationName
}
