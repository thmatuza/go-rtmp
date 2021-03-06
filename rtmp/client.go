package rtmp

import (
	"context"
	"io"
	"net"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Client struct {
	ctx         context.Context
	cancelFunc  context.CancelFunc
	connOptions []ConnOption

	logger *zap.Logger
}

func NewClient(
	ctx context.Context,
	logger *zap.Logger,
	connOps ...ConnOption,
) *Client {
	ctx, cancel := context.WithCancel(ctx)
	return &Client{
		ctx:         ctx,
		cancelFunc:  cancel,
		logger:      logger,
		connOptions: connOps,
	}
}

func (c *Client) Close() error {
	c.cancelFunc()
	return nil
}

func (c *Client) Connect(ctx context.Context, addr string) (Conn, error) {
	if dd, ok := c.ctx.Deadline(); ok {
		var cancel func()
		ctx, cancel = context.WithDeadline(ctx, dd)
		defer cancel()
	}
	var d net.Dialer
	nc, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial %s", addr)
	}
	if nc == nil {
		return nil, errors.New("conn is nil")
	}

	conn := NewDefaultConn(
		c.ctx,
		nc,
		false,
		c.logger,
		c.connOptions...,
	)

	go func() {
		remoteAddr := nc.RemoteAddr()
		defer func() {
			if err := conn.Close(); err != nil {
				c.logger.Error(
					"failed to close conn",
					zap.Error(err),
					zap.Stringer("remoteAddr", remoteAddr),
				)
			}
		}()
		if err := conn.Serve(); err != nil {
			if isCanceledErr(err) ||
				errors.Cause(err) == io.EOF ||
				isDone(ctx) {
				return
			}
			if e, ok := errors.Cause(err).(ConnError); ok {
				c.logger.Error(
					"failed to conn.serve",
					append(e.Fields(), zap.Error(err), zap.Stringer("remoteAddr", remoteAddr))...,
				)
			} else {
				c.logger.Error(
					"failed to conn.serve",
					zap.Error(err),
					zap.Stringer("remoteAddr", remoteAddr),
				)
			}
		}
	}()
	return conn, nil
}
