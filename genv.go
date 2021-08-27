package genv

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

func Load(filenames ...string) error {
	files := getFileNames(filenames)
	for _, name := range files {

		file, err := os.Open(name)
		if err != nil {
			if name == ".env" {
				return errors.New(".env file not found")
			}
			return errors.New("file not found")
		}
		defer file.Close()
		
		scanner := bufio.NewScanner(file)

		rg, err := regexp.Compile(`\w+=\w*`)
		if err != nil {
			return err
		}

		for scanner.Scan() {
			if !rg.MatchString(scanner.Text()) {
				continue
			}
			envArr := strings.Split(scanner.Text(), "=")
			err := os.Setenv(envArr[0], envArr[1])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func getFileNames(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}

