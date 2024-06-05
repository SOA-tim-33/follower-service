package handler

import (
	"context"
	"database-example/proto/follower"
	"database-example/service"
	"fmt"
)

type UserHandler struct {
	follower.UnimplementedFollowerServiceServer
	FollowerService *service.FollowService
}

func (handler *UserHandler) toInt64Array(intArray []int) []int64 {
	result := make([]int64, len(intArray))
	for i, e := range intArray {
		result[i] = int64(e)
	}
	return result
}
func (handler *UserHandler) Follow(ctx context.Context, request *follower.MultiIdRequest) (*follower.EmptyResponse, error) {
	id1 := int(request.Id1)
	id2 := int(request.Id2)

	var err = handler.FollowerService.Follow(id1, id2)
	if err == nil {
		return &follower.EmptyResponse{}, nil
	}
	fmt.Println(err)
	return &follower.EmptyResponse{}, err

}

func (handler *UserHandler) CheckIfFollowing(ctx context.Context, request *follower.MultiIdRequest) (*follower.IsFollowingResponse, error) {

	id1 := int(request.Id1)
	id2 := int(request.Id2)

	result, err := handler.FollowerService.CheckFollowing(id1, id2)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &follower.IsFollowingResponse{IsFollowing: result}, nil
}

func (handler *UserHandler) GetRecommendation(ctx context.Context, request *follower.Request) (*follower.MultiIdResponse, error) {

	id := int(request.Id)
	result, err := handler.FollowerService.GetRecommendations(id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &follower.MultiIdResponse{Ids: handler.toInt64Array(result)}, nil
}
