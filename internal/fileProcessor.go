package internal

import (
	"bufio"
	"io"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strings"

	"github.com/BreCkver/st-codeChallenge/models"
)

/*Request */
type Request struct {
	Email    string
	File     string
	FileName string
	Errors   map[string]string
}

var rxEmail = regexp.MustCompile(".+@.+\\..+")

func (req *Request) Validate() bool {
	req.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(req.Email))
	if match == false {
		req.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(req.File) == "" {
		req.Errors["File"] = "Please select a valid file"
	}

	return len(req.Errors) == 0
}

func (req *Request) LoadFile(file multipart.File) error {

	f, err := os.OpenFile(req.FileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return err
	} else {

		io.Copy(f, file)
		return nil
	}

}

func (req *Request) ReadFile() ([]*models.Transaction, error) {
	var results []*models.Transaction
	if fileN, err := os.Open(req.FileName); err != nil {
		return nil, err
	} else {
		scanner := bufio.NewScanner(fileN)
		for scanner.Scan() {

			if t, err := ConvertTransaction(scanner.Text()); err != nil {
				log.Printf("Error to convert Transaction: \n %v", err.Error())
			} else {
				results = append(results, t)
			}
		}

		return results, nil
	}
}
