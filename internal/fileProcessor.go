package internal

import (
	"bufio"
	"io"
	"mime/multipart"
	"os"
	"regexp"
	"strings"

	"github.com/BreCkver/st-codeChallenge/models"
)

/*Request */
type Request struct {
	Account  models.Account
	FileName string
	Errors   map[string]string
}

var rxEmail = regexp.MustCompile(".+@.+\\..+")
var rxUserName = regexp.MustCompile("[^a-zA-Z0-9 ]*$")

func (req *Request) Validate() bool {
	req.Errors = make(map[string]string)

	emailMatch := rxEmail.Match([]byte(req.Account.Email))
	if emailMatch == false {
		req.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(req.FileName) == "" {
		req.Errors["File"] = "Please select a valid file"
	}

	userMatch := rxUserName.Match([]byte(req.Account.UserName))
	if strings.TrimSpace(req.Account.UserName) == "" || userMatch == false {
		req.Errors["AccountName"] = "Please enter a beneficiary name, alphanumeric only"
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

func (req *Request) ReadFile() ([]*models.Transaction, []string) {

	var errString []string
	var results []*models.Transaction

	if fileN, err := os.Open(req.FileName); err != nil {
		errString = append(errString, err.Error())
		return nil, errString
	} else {

		fi, err := fileN.Stat()
		if err != nil {
			errString = append(errString, err.Error())
			return nil, errString
		}

		if fi.Size() == 0 {
			errString = append(errString, "Can't process file empty")
			return nil, errString
		}

		scanner := bufio.NewScanner(fileN)
		if scanner.Scan() {
			for scanner.Scan() {

				if t, err := ConvertTransaction(scanner.Text()); err != nil {
					errString = append(errString, err.Error())
				} else {
					results = append(results, t)
				}
			}
		}

		return results, errString
	}
}
