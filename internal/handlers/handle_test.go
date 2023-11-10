package handlers

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	mock_handlers "worldofwisdom.com/m/internal/handlers/mocks"
	"worldofwisdom.com/m/internal/tcp/proto"
)

func TestQuotesHandler_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockLog := mock_handlers.NewMockLogger(ctrl)
	mockLog.
		EXPECT().
		InfoContext(ctx, gomock.Any()).
		Times(8)

	mockConn := mock_handlers.NewMockConn(ctrl)
	mockConn.
		EXPECT().
		Read(ctx).
		Return(&proto.Message{
			Operation: proto.Initiate,
		}, nil)
	mockConn.
		EXPECT().
		Read(ctx).
		Return(&proto.Message{
			Operation: proto.SendNonce,
			Content:   []byte{123, 12, 0, 0, 0, 0, 0, 0},
		}, nil)
	mockConn.
		EXPECT().
		Send(ctx, gomock.Eq(proto.NewSendChallengeRequest(2, []byte{1, 2, 3, 4, 5, 6, 7, 8}))).
		Times(1)
	mockConn.
		EXPECT().
		Send(ctx, gomock.Eq(proto.NewSendDataRequest([]byte{116, 101, 115, 116, 81, 117, 111, 116, 101}))).
		Times(1)
	mockConn.
		EXPECT().
		Close(ctx).
		Times(1)

	mockStorage := mock_handlers.NewMockQuotesStorage(ctrl)
	mockStorage.
		EXPECT().
		GetRandomQuote().
		Return("testQuote")

	mockGenerator := mock_handlers.NewMockPOWGenerator(ctrl)
	mockGenerator.
		EXPECT().
		Generate(ctx).
		Return([]byte{1, 2, 3, 4, 5, 6, 7, 8}, byte(2), nil)
	mockGenerator.
		EXPECT().
		CheckNonce([]byte{123, 12, 0, 0, 0, 0, 0, 0}, []byte{1, 2, 3, 4, 5, 6, 7, 8}).
		Return(true)

	handler := NewQuotesHandler(mockLog, mockStorage, mockGenerator)
	handler.Handle(ctx, mockConn)
}
