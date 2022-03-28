package services

import (
	"net/url"
)

type HttpService interface {
	SetClient(withCookie bool) error
	Get(url string, data interface{}) error
	PostForm(url string, form url.Values, data interface{}) error
}
