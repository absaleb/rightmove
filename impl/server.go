package impl

import (
	"context"
	"gitlab.okta-solutions.com/mashroom/backend/common/errs"
	"net"

	"google.golang.org/grpc"

	"gitlab.okta-solutions.com/mashroom/backend/common/health"
	"gitlab.okta-solutions.com/mashroom/backend/common/log"
	"gitlab.okta-solutions.com/mashroom/backend/rightmove"
	"gitlab.okta-solutions.com/mashroom/backend/rightmove/version"
)

type Server interface {
	rightmove.RightmoveServiceServer
	Serve(addr string)
	Background()
}

type serverImpl struct {
}

func ToRightmoveSendPropertyRequest(request *rightmove.ListingUpdateRequest) (*RightmoveSendCallRequest, error) {
	var media []Media
	for _, v := range request.Medias {
		media = append(media, Media{
			MediaType: v.MediaType,
			MediaURL:  v.MediaURL,
		})
	}

	result := &RightmoveSendCallRequest{
		Branch:  Branch{BranchID: request.BranchID},
		Network: Network{NetworkID: request.NetworkID},
		Property: Property{
			Address: Address{
				HouseNameNumber: request.Location.Address,
				Town:            request.Location.TownOrCity,
				Postcode1:       request.Location.PostalCode,
				Postcode2:       request.Location.PostalCode,
				DisplayAddress:  request.Location.Address,
				Longitude:       request.Location.Coordinates.Longitude,
				Latitude:        request.Location.Coordinates.Latitude,
			},
			AgentRef: request.AgentRef,
			Details: Details{
				Summary:     request.Details.Summary,
				Description: request.Details.Description,
				Bedrooms:    request.Details.Bedrooms,
			},
			Media: media,
			PriceInformation: PriceInformation{
				Price:             request.Price,
				Deposit:           &request.Deposit,
				RentFrequency:     &request.RentFrequency,
				AdministrationFee: &request.AdministrationFee,
			},
			PropertyType   : request.PropertyType,
			Published : true,
			Status    : 1,
		},
	}

	return result, nil
}

func (server *serverImpl) SendProperty(ctx context.Context, request *rightmove.ListingUpdateRequest) (*rightmove.ListingResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}
	req, err := ToRightmoveSendPropertyRequest(request)

	resp, err := SendPropertyImpl(req)
	if err != nil {
		return nil, err
	}

	result := &rightmove.ListingResponse{
		RequestId: resp.RequestID,
		Success:   resp.Success,
	}
	return result, nil
}

func (server *serverImpl) DeleteProperty(ctx context.Context, request *rightmove.ListingDeleteRequest) (*rightmove.ListingResponse, error) {
	panic("implement me")
}

func (server *serverImpl) Listing(ctx context.Context, request *rightmove.ListingListRequest) (*rightmove.ListingListResponse, error) {
	panic("implement me")
}

func (server *serverImpl) Background() {
	// background processes
}

func (server *serverImpl) Serve(addr string) {
	if listener, err := net.Listen("tcp", addr); err != nil {
		panic(err)
	} else {
		log.SetHost("rightmove")
		grpcServer := grpc.NewServer()

		rightmove.RegisterRightmoveServiceServer(grpcServer, server)

		healthServer := version.NewHealthServer()
		health.RegisterHealthServiceServer(grpcServer, healthServer)

		log.Infoln("rightmove started")
		if err := grpcServer.Serve(listener); err != nil {
			log.Errorln("gRPC error", err)
		}
	}
}

func NewServer() Server {
	server := &serverImpl{}
	go server.Background()
	return server
}
