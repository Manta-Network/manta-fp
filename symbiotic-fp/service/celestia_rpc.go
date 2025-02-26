package service

import (
	"context"

	"github.com/Manta-Network/manta-fp/symbiotic-fp/protobuf/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CelestiaService interface {
	StateRootSignIDs(ids string, l1BlockNumber uint64) (*pb.StateRootSignIDsResponse, error)
}

type symbioticRpcService struct {
	sRpcService pb.CelestiaServiceClient
}

func (s symbioticRpcService) StateRootSignIDs(ids string, l1BlockNumber uint64) (*pb.StateRootSignIDsResponse, error) {
	ctx := context.Background()
	request := &pb.StateRootSignIDsRequest{
		Ids:           ids,
		L1BlockNumber: l1BlockNumber,
	}
	status, err := s.sRpcService.StateRootSignIDs(ctx, request)
	return status, err
}

func NewSymbioticRpcService(rpcUrl string) (CelestiaService, error) {
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	celestiaServiceClient := pb.NewCelestiaServiceClient(conn)
	brService := &symbioticRpcService{celestiaServiceClient}
	return brService, nil
}
