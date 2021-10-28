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

func decodeCreateSnapshotRequest(r *http.Request, span trace.Span) (req SnapshotCreateParams, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SnapshotCreateParams
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreateSyncActionRequest(r *http.Request, span trace.Span) (req InstanceActionInfo, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request InstanceActionInfo
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeLoadSnapshotRequest(r *http.Request, span trace.Span) (req SnapshotLoadParams, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SnapshotLoadParams
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeMmdsConfigPutRequest(r *http.Request, span trace.Span) (req MmdsConfig, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MmdsConfig
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeMmdsPatchRequest(r *http.Request, span trace.Span) (req MmdsPatchReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MmdsPatchReq
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeMmdsPutRequest(r *http.Request, span trace.Span) (req MmdsPutReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MmdsPutReq
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchBalloonRequest(r *http.Request, span trace.Span) (req BalloonUpdate, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BalloonUpdate
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchBalloonStatsIntervalRequest(r *http.Request, span trace.Span) (req BalloonStatsUpdate, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BalloonStatsUpdate
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchGuestDriveByIDRequest(r *http.Request, span trace.Span) (req PartialDrive, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PartialDrive
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchGuestNetworkInterfaceByIDRequest(r *http.Request, span trace.Span) (req PartialNetworkInterface, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PartialNetworkInterface
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchMachineConfigurationRequest(r *http.Request, span trace.Span) (req MachineConfiguration, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MachineConfiguration
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchVmRequest(r *http.Request, span trace.Span) (req VM, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request VM
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutBalloonRequest(r *http.Request, span trace.Span) (req Balloon, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Balloon
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutGuestBootSourceRequest(r *http.Request, span trace.Span) (req BootSource, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BootSource
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutGuestDriveByIDRequest(r *http.Request, span trace.Span) (req Drive, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Drive
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutGuestNetworkInterfaceByIDRequest(r *http.Request, span trace.Span) (req NetworkInterface, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request NetworkInterface
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutGuestVsockRequest(r *http.Request, span trace.Span) (req Vsock, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Vsock
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutLoggerRequest(r *http.Request, span trace.Span) (req Logger, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Logger
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutMachineConfigurationRequest(r *http.Request, span trace.Span) (req MachineConfiguration, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MachineConfiguration
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutMetricsRequest(r *http.Request, span trace.Span) (req Metrics, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Metrics
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}