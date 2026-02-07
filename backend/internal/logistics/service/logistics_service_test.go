package service

import (
	"context"
	"testing"
	"time"

	"github.com/emergency-material-system/backend/internal/logistics/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock of TrackingRepository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, tracking *model.Tracking) error {
	args := m.Called(ctx, tracking)
	return args.Error(0)
}

func (m *MockRepository) GetByID(ctx context.Context, id uint) (*model.Tracking, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Tracking), args.Error(1)
}

func (m *MockRepository) GetByRequestID(ctx context.Context, requestID uint) ([]*model.Tracking, error) {
	args := m.Called(ctx, requestID)
	return args.Get(0).([]*model.Tracking), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, tracking *model.Tracking) error {
	args := m.Called(ctx, tracking)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) AddNode(ctx context.Context, node *model.TrackingNode) error {
	args := m.Called(ctx, node)
	return args.Error(0)
}

func (m *MockRepository) GetWithNodes(ctx context.Context, id uint) (*model.Tracking, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Tracking), args.Error(1)
}

func TestRecordTrajectoryNode_WithDataCompletion(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewLogisticsService(mockRepo).(*logisticsService) // Access internal struct
	ctx := context.Background()

	// Note: We are using the internal gaodeClient initialized in NewLogisticsService.
	// In a real test we'd mock the API calls, but here we verify the service structure.

	trackingID := uint(1)
	location := "Beijing"
	lat, lng := 0.0, 0.0 // Missing coordinates
	status := "in_transit"
	description := "Passing through Beijing"

	// Mock GetByID for update
	mockRepo.On("GetByID", ctx, trackingID).Return(&model.Tracking{ID: trackingID}, nil)

	// Mock AddNode
	mockRepo.On("AddNode", ctx, mock.AnythingOfType("*model.TrackingNode")).Return(nil)

	// Mock Update
	mockRepo.On("Update", ctx, mock.AnythingOfType("*model.Tracking")).Return(nil)

	// Note: Since BaiduMapsClient is a struct not an interface, testing it requires
	// either a real AK (not recommended for CI) or wrapping it in an interface.
	// For now, we verify that the service handles the flow.
	err := svc.RecordTrajectoryNode(ctx, trackingID, location, lat, lng, status, description)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTrajectory(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewLogisticsService(mockRepo)
	ctx := context.Background()

	trackingID := uint(1)
	expectedTracking := &model.Tracking{
		ID: trackingID,
		Nodes: []model.TrackingNode{
			{ID: 1, Location: "Start", TrackedAt: time.Now().Add(-1 * time.Hour)},
			{ID: 2, Location: "End", TrackedAt: time.Now()},
		},
	}

	mockRepo.On("GetWithNodes", ctx, trackingID).Return(expectedTracking, nil)

	result, err := svc.GetTrajectory(ctx, trackingID)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Nodes))
	assert.Equal(t, "Start", result.Nodes[0].Location)
	mockRepo.AssertExpectations(t)
}
