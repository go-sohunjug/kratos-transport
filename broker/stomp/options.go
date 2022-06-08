package stomp

import (
	"context"
	"github.com/go-sohunjug/kratos-transport/broker"
	"time"
)

type authKey struct{}
type connectHeaderKey struct{}
type connectTimeoutKey struct{}
type durableQueueKey struct{}
type receiptKey struct{}
type subscribeHeaderKey struct{}
type suppressContentLengthKey struct{}
type vHostKey struct{}
type subscribeContextKey struct{}
type ackSuccessKey struct{}

type authRecord struct {
	username string
	password string
}

func WithSubscribeHeaders(h map[string]string) broker.SubscribeOption {
	return broker.SubscribeContextWithValue(subscribeHeaderKey{}, h)
}

func SubscribeHeadersFromContext(ctx context.Context) (map[string]string, bool) {
	h, ok := ctx.Value(subscribeHeaderKey{}).(map[string]string)
	return h, ok
}

func WithSubscribeContext(ctx context.Context) broker.SubscribeOption {
	return broker.SubscribeContextWithValue(subscribeContextKey{}, ctx)
}

func SubscribeContextFromContext(ctx context.Context) (context.Context, bool) {
	c, ok := ctx.Value(subscribeContextKey{}).(context.Context)
	return c, ok
}

func WithAckOnSuccess() broker.SubscribeOption {
	return broker.SubscribeContextWithValue(ackSuccessKey{}, true)
}

func AckOnSuccessFromContext(ctx context.Context) (bool, bool) {
	b, ok := ctx.Value(ackSuccessKey{}).(bool)
	return b, ok
}

func WithDurable() broker.SubscribeOption {
	return broker.SubscribeContextWithValue(durableQueueKey{}, true)
}

func WithReceipt(_ time.Duration) broker.PublishOption {
	return broker.PublishContextWithValue(receiptKey{}, true)
}

func WithSuppressContentLength(_ time.Duration) broker.PublishOption {
	return broker.PublishContextWithValue(suppressContentLengthKey{}, true)
}

func WithConnectTimeout(ct time.Duration) broker.Option {
	return broker.OptionContextWithValue(connectTimeoutKey{}, ct)
}

func ConnectTimeoutFromContext(ctx context.Context) (time.Duration, bool) {
	v, ok := ctx.Value(connectTimeoutKey{}).(time.Duration)
	return v, ok
}

func WithAuth(username string, password string) broker.Option {
	return broker.OptionContextWithValue(authKey{}, &authRecord{
		username: username,
		password: password,
	})
}

func AuthFromContext(ctx context.Context) (*authRecord, bool) {
	v, ok := ctx.Value(authKey{}).(*authRecord)
	return v, ok
}

func WithConnectHeaders(h map[string]string) broker.Option {
	return broker.OptionContextWithValue(connectHeaderKey{}, h)
}

func ConnectHeadersFromContext(ctx context.Context) (map[string]string, bool) {
	v, ok := ctx.Value(connectHeaderKey{}).(map[string]string)
	return v, ok
}

func WithVirtualHost(h string) broker.Option {
	return broker.OptionContextWithValue(vHostKey{}, h)
}

func VirtualHostFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(vHostKey{}).(string)
	return v, ok
}
