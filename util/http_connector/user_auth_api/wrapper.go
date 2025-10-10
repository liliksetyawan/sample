package user_auth_api

import (
	"net/http"
)


// HTTPClientDTO and Variants
type HTTPClientDTO struct {
	Header        http.Header
	Unsuccessfull ErrorResponse
}

func (h *HTTPClientDTO) SetHeader(header http.Header) {
	h.Header = header
}

func (h *HTTPClientDTO) GetHeader() http.Header {
	return h.Header
}

func (h *HTTPClientDTO) Unsuccessfully() interface{} {
	return &h.Unsuccessfull
}
