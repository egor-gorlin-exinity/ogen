// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = fasthttp.Client{}
)

func decodeGetBookResponse(resp *http.Response, span trace.Span) (res GetBookRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response Book
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 403:
		return &GetBookResForbidden{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeGetPageCoverImageResponse(resp *http.Response, span trace.Span) (res GetPageCoverImageRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "image/*":
			return res, fmt.Errorf("image/* decoder not implemented")
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 403:
		return &GetPageCoverImageResForbidden{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeGetPageImageResponse(resp *http.Response, span trace.Span) (res GetPageImageRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "image/*":
			return res, fmt.Errorf("image/* decoder not implemented")
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 403:
		return &GetPageImageResForbidden{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeGetPageThumbnailImageResponse(resp *http.Response, span trace.Span) (res GetPageThumbnailImageRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "image/*":
			return res, fmt.Errorf("image/* decoder not implemented")
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 403:
		return &GetPageThumbnailImageResForbidden{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeSearchResponse(resp *http.Response, span trace.Span) (res SearchRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response SearchOKApplicationJSON
			if err := func() error {
				if err := fmt.Errorf(`decoding of "SearchOKApplicationJSON" (alias) is not implemented`); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 403:
		return &SearchResForbidden{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeSearchByTagIDResponse(resp *http.Response, span trace.Span) (res SearchByTagIDRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response SearchByTagIDOKApplicationJSON
			if err := func() error {
				if err := fmt.Errorf(`decoding of "SearchByTagIDOKApplicationJSON" (alias) is not implemented`); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 403:
		return &SearchByTagIDResForbidden{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}