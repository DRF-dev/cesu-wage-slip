package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type HttpServiceImpl struct{
	Client http.Client
}

func (receiver *HttpServiceImpl) SetClient(withCookie bool) error {
	if withCookie == false {
		receiver.Client = http.Client{Timeout: 40 * time.Second}
	}

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	receiver.Client = http.Client{
		Timeout: 40 * time.Second,
		Jar: cookieJar,
	}

	return nil
}

func (receiver HttpServiceImpl) Get(url string, data interface{}) error {
	res, err := receiver.Client.Get(url)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &data)

	return nil
}

func (receiver HttpServiceImpl) PostForm(url string, form url.Values, data interface{}) error {
	res, err := receiver.Client.PostForm(url, form)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &data)

	return nil
}


