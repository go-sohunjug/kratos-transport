package kafka

import (
	"context"
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sohunjug/kratos-transport/broker"
)

type ServerOption func(o *Server)

func Address(addrs []string) ServerOption {
	return func(s *Server) {
		s.bOpts = append(s.bOpts, broker.Addrs(addrs...))
	}
}

func Logger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(logger)
	}
}

func TLSConfig(c *tls.Config) ServerOption {
	return func(s *Server) {
		if c != nil {
			s.bOpts = append(s.bOpts, broker.Secure(true))
		}
		s.bOpts = append(s.bOpts, broker.TLSConfig(c))
	}
}

func Subscribe(ctx context.Context, topic, queue string, disableAutoAck bool, h broker.Handler, opts ...broker.SubscribeOption) ServerOption {
	return func(s *Server) {
		_ = s.RegisterSubscriber(ctx, topic, queue, disableAutoAck, h, opts...)
	}
}
