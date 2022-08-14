package application

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type Token struct {
	Token string `json:"token"`
}

// When it runs in a CI pipeline simulate sign-up and
// submit images to the API with an associated id.
func (a App) CiMode() error {
	log.WithFields(log.Fields{"tag": a.Tag}).Info("ci mode started")

	// mock real client
	client := &ClientMock{}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("test", "eye_1.png")
	if err != nil {
		log.WithFields(log.Fields{"tag": a.Tag,
			"err": err}).Error("writer can't be created from file")
		return err
	}
	file, err := os.Open("test/eye_1.png")
	if err != nil {
		log.WithFields(log.Fields{"tag": a.Tag,
			"err": err}).Error("can't open an eye image")
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		log.WithFields(log.Fields{"tag": a.Tag,
			"err": err}).Error("io copy error")

		return err
	}

	req, err := http.NewRequest("POST", a.Url+"authorization", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.WithFields(log.Fields{"tag": a.Tag,
			"err": err}).Error("ci mode started")

		return err
	}

	writer.WriteField("_json", `{"login":"user login", "password": "user password"}`)
	writer.Close()

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, _ := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{"tag": a.Tag,
			"err": err}).Errorf("Request failed with response code: %d", resp.StatusCode)
		return errors.New("server response with error")
	}

	// var data Token
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.WithFields(log.Fields{ "tag": a.Tag, "err": err,
	// 	}).Errorf("response read error")
	// }
	// bodyString := string(bodyBytes)
	// log.Info(bodyString)
	// log.WithFields(log.Fields{ "tag": a.Tag,
	// }).Infof("response: %s", bodyString)

	// if err = json.Unmarshal(bodyBytes, &data); err != nil {
	// 	log.WithFields(log.Fields{ "tag": a.Tag, "err": err,
	// 	}).Error("response unmarshal error")

	// 	return err
	// }

	log.WithFields(log.Fields{"tag": a.Tag}).Info("success")
	return nil

}
