package client

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_client "worldofwisdom.com/m/internal/client/mocks"
	"worldofwisdom.com/m/internal/services/proof_of_work"
	"worldofwisdom.com/m/internal/tcp/proto"
)

func TestClient_GetQuote(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockConn := mock_client.NewMockConn(ctrl)
	mockConn.EXPECT().
		Send(ctx, gomock.Eq(proto.NewInitializeRequest())).
		Return(nil)
	mockConn.EXPECT().
		Send(ctx, gomock.Eq(proto.NewSolvedChallengeRequest([]byte{123, 12, 0, 0, 0, 0, 0, 0}))).
		Return(nil)
	mockConn.EXPECT().
		Read(ctx).
		Return(&proto.Message{
			Operation:  2,
			Complexity: 2,
			Length:     8,
			Content:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
		}, nil)
	mockConn.EXPECT().
		Read(ctx).
		Return(&proto.Message{
			Operation:  4,
			Complexity: 0,
			Length:     9,
			Content:    []byte{116, 101, 115, 116, 81, 117, 111, 116, 101},
		}, nil)

	mockLog := mock_client.NewMockLogger(ctrl)
	mockLog.
		EXPECT().
		InfoContext(ctx, gomock.Any()).
		Times(5)

	mockSolver := mock_client.NewMockSolver(ctrl)
	mockSolver.EXPECT().
		SolveChallenge(gomock.Any(), byte(2), []byte{1, 2, 3, 4, 5, 6, 7, 8}).
		Return(proof_of_work.NonceFromBytes([]byte{123, 12, 0, 0, 0, 0, 0, 0}), nil)

	client := Client{
		log:     mockLog,
		solver:  mockSolver,
		conn:    mockConn,
		timeout: 150,
	}
	quote, err := client.GetQuote(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, quote)
	assert.Equal(t, "testQuote", string(quote))
}
