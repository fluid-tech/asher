package writer

import (
	"asher/internal/api"
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const tempProjectPath = "/Users/gauravpunjabi/Coding/Projects/asher/resources/projects/laravel-demo"

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

func (writerHandler *MigrationWriterHandler) Handle(emitterFile api.EmitterFile) int {
	noOfBytes := 0
	generator := emitterFile.Generator()
	fileContent := (*generator).String()
	pathToMigrationFile := tempProjectPath + emitterFile.Path() + writerHandler.migrationFileName + ".php"
	fileWriter, err := os.Create(pathToMigrationFile)
	if err == nil {
		noOfBytes, err = fileWriter.WriteString(fileContent)
		if err != nil {
			noOfBytes = 0
		}
	}
	return noOfBytes
}

func (writerHandler *MigrationWriterHandler) AfterHandle(emitterFile api.EmitterFile) {
}

func generateMigration(name string) string {
	migrationFileName := ""

	// Creating a new migration
	cmd := exec.Command("php", "artisan", "make:migration", name)
	cmd.Dir = tempProjectPath
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
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
