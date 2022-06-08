package nats

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

func Subscribe(topic string, h broker.Handler) ServerOption {
	return func(s *Server) {
		if s.baseCtx == nil {
			s.baseCtx = context.Background()
		}

		_ = s.RegisterSubscriber(topic, h,
			broker.SubscribeContext(s.baseCtx),
			broker.Queue(topic),
		)
	}
}
