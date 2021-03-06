package rtmp

import (
	"context"

	"github.com/pkg/errors"
)

func (conn *defaultConn) OnStatus(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, infoObject map[string]interface{}) error {
	p := NewOnStatus(infoObject, conn.encodingAMFType)
	b, err := p.MarshalBinary()
	if err != nil {
		return errors.Wrap(err, "failed to MarshalBinary")
	}
	var msgTypeID MessageTypeID
	if conn.encodingAMFType == EncodingAMFTypeAMF0 {
		msgTypeID = MessageTypeIDCommandAMF0
	} else {
		msgTypeID = MessageTypeIDCommandAMF3
	}

	m := NewMessage(
		chunkStreamID,
		msgTypeID,
		conn.Timestamp(),
		messageStreamID,
		b,
	)
	_, err = conn.Writer().WriteMessage(m)
	if err != nil {
		return errors.Wrap(err, "failed to WriteMessage")
	}
	if err := conn.Writer().Flush(); err != nil {
		return errors.Wrap(err, "failed to Flush Writer")
	}
	return nil
}

func (conn *defaultConn) Play(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, streamName string, start uint32, duration uint32, reset bool) error {
	panic("not implemented")
}

func (conn *defaultConn) Play2(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, parameters map[string]interface{}) error {
	panic("not implemented")
}

func (conn *defaultConn) DeleteStream(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, streamID uint32) error {
	panic("not implemented")
}

func (conn *defaultConn) CloseStream(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, streamID uint32) error {
	panic("not implemented")
}

func (conn *defaultConn) ReceiveAudio(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, boolFlag bool) error {
	panic("not implemented")
}

func (conn *defaultConn) ReceiveVideo(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, boolFlag bool) error {
	panic("not implemented")
}

func (conn *defaultConn) Publish(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, publishingName string, publishingType PublishingType) error {
	p := NewPublish(publishingName, publishingType, conn.encodingAMFType)
	b, err := p.MarshalBinary()
	if err != nil {
		return errors.Wrap(err, "failed to MarshalBinary")
	}
	var msgTypeID MessageTypeID
	if conn.encodingAMFType == EncodingAMFTypeAMF0 {
		msgTypeID = MessageTypeIDCommandAMF0
	} else {
		msgTypeID = MessageTypeIDCommandAMF3
	}

	m := NewMessage(
		chunkStreamID,
		msgTypeID,
		conn.Timestamp(),
		messageStreamID,
		b,
	)
	_, err = conn.Writer().WriteMessage(m)
	if err != nil {
		return errors.Wrap(err, "failed to WriteMessage")
	}
	if err := conn.Writer().Flush(); err != nil {
		return errors.Wrap(err, "failed to Flush Writer")
	}
	return nil
}

func (conn *defaultConn) Seek(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, milliSeconds uint32) error {
	panic("not implemented")
}

func (conn *defaultConn) Pause(ctx context.Context, chunkStreamID uint32, messageStreamID uint32, pauseUnpauseFlag bool, milliSeconds uint32) error {
	panic("not implemented")
}
