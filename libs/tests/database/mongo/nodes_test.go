package mongo

import (
	"context"
	models "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	mocks "github.com/curtis0505/bridge/libs/mocks/mongo_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestFindNodes(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 생성된 목 객체
	mockClient := mocks.NewMockNodesRepository(ctrl)

	// 테스트할 코드 실행
	mockClient.EXPECT().FindNodes(ctx, gomock.Any()).Return(models.Nodes{
		ID:            primitive.NewObjectID(),
		Chain:         "BASE",
		Symbol:        "",
		PublicNodeURL: "",
		ServerNodeURL: []string{},
	}, nil)

	// 실제 테스트 실행
	result, err := mockClient.FindNodes(ctx, bson.M{"chain": "BASE"})
	// 결과 검증
	assert.NoError(t, err)

	// 예상 결과 확인
	assert.NotNil(t, result)
}
