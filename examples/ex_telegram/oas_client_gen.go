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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type config struct {
	TracerProvider trace.TracerProvider
	Tracer         trace.Tracer
	Client         HTTPClient
}

const defaultTracerName = "github.com/ogen-go/ogen/otelogen"

func newConfig(opts ...Option) config {
	cfg := config{
		TracerProvider: otel.GetTracerProvider(),
		Client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	cfg.Tracer = cfg.TracerProvider.Tracer(
		defaultTracerName,
		trace.WithInstrumentationVersion(otelogen.SemVersion()),
	)
	return cfg
}

type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.TracerProvider = provider
		}
	})
}

func WithHTTPClient(client HTTPClient) Option {
	return optionFunc(func(cfg *config) {
		if client != nil {
			cfg.Client = client
		}
	})
}

type Client struct {
	serverURL *url.URL
	cfg       config
}

func NewClient(serverURL string, opts ...Option) *Client {
	u, err := url.Parse(serverURL)
	if err != nil {
		panic(err) // TODO: fix
	}
	return &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
}

func (c *Client) AnswerCallbackQueryPost(ctx context.Context, req AnswerCallbackQueryPostReq) (res AnswerCallbackQueryPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `AnswerCallbackQueryPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeAnswerCallbackQueryPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/answerCallbackQuery"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeAnswerCallbackQueryPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) AnswerPreCheckoutQueryPost(ctx context.Context, req AnswerPreCheckoutQueryPostReq) (res AnswerPreCheckoutQueryPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `AnswerPreCheckoutQueryPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeAnswerPreCheckoutQueryPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/answerPreCheckoutQuery"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeAnswerPreCheckoutQueryPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) AnswerShippingQueryPost(ctx context.Context, req AnswerShippingQueryPostReq) (res AnswerShippingQueryPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `AnswerShippingQueryPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeAnswerShippingQueryPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/answerShippingQuery"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeAnswerShippingQueryPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) ClosePost(ctx context.Context) (res ClosePostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `ClosePost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	u := uri.Clone(c.serverURL)
	u.Path += "/close"

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeClosePostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) DeleteStickerFromSetPost(ctx context.Context, req DeleteStickerFromSetPostReq) (res DeleteStickerFromSetPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `DeleteStickerFromSetPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeDeleteStickerFromSetPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/deleteStickerFromSet"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeDeleteStickerFromSetPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) DeleteWebhookPost(ctx context.Context, req DeleteWebhookPostReq) (res DeleteWebhookPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `DeleteWebhookPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeDeleteWebhookPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/deleteWebhook"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeDeleteWebhookPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetFilePost(ctx context.Context, req GetFilePostReq) (res GetFilePostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetFilePost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeGetFilePostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/getFile"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetFilePostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetGameHighScoresPost(ctx context.Context, req GetGameHighScoresPostReq) (res GetGameHighScoresPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetGameHighScoresPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeGetGameHighScoresPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/getGameHighScores"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetGameHighScoresPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetMePost(ctx context.Context) (res GetMePostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetMePost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	u := uri.Clone(c.serverURL)
	u.Path += "/getMe"

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetMePostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetMyCommandsPost(ctx context.Context) (res GetMyCommandsPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetMyCommandsPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	u := uri.Clone(c.serverURL)
	u.Path += "/getMyCommands"

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetMyCommandsPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetStickerSetPost(ctx context.Context, req GetStickerSetPostReq) (res GetStickerSetPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetStickerSetPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeGetStickerSetPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/getStickerSet"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetStickerSetPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetUpdatesPost(ctx context.Context, req GetUpdatesPostReq) (res GetUpdatesPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetUpdatesPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeGetUpdatesPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/getUpdates"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetUpdatesPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetUserProfilePhotosPost(ctx context.Context, req GetUserProfilePhotosPostReq) (res GetUserProfilePhotosPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetUserProfilePhotosPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeGetUserProfilePhotosPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/getUserProfilePhotos"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetUserProfilePhotosPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) GetWebhookInfoPost(ctx context.Context) (res GetWebhookInfoPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `GetWebhookInfoPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	u := uri.Clone(c.serverURL)
	u.Path += "/getWebhookInfo"

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeGetWebhookInfoPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) LogOutPost(ctx context.Context) (res LogOutPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `LogOutPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	u := uri.Clone(c.serverURL)
	u.Path += "/logOut"

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeLogOutPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) SendGamePost(ctx context.Context, req SendGamePostReq) (res SendGamePostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `SendGamePost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeSendGamePostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/sendGame"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeSendGamePostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) SendInvoicePost(ctx context.Context, req SendInvoicePostReq) (res SendInvoicePostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `SendInvoicePost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeSendInvoicePostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/sendInvoice"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeSendInvoicePostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) SetMyCommandsPost(ctx context.Context, req SetMyCommandsPostReq) (res SetMyCommandsPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `SetMyCommandsPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeSetMyCommandsPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/setMyCommands"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeSetMyCommandsPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) SetStickerPositionInSetPost(ctx context.Context, req SetStickerPositionInSetPostReq) (res SetStickerPositionInSetPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `SetStickerPositionInSetPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeSetStickerPositionInSetPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/setStickerPositionInSet"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeSetStickerPositionInSetPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) SetWebhookPost(ctx context.Context, req SetWebhookPostReq) (res SetWebhookPostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `SetWebhookPost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeSetWebhookPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/setWebhook"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeSetWebhookPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) UploadStickerFilePost(ctx context.Context, req UploadStickerFilePostReq) (res UploadStickerFilePostRes, err error) {
	ctx, span := c.cfg.Tracer.Start(ctx, `UploadStickerFilePost`,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()
	buf, contentType, err := encodeUploadStickerFilePostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/uploadStickerFile"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeUploadStickerFilePostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}