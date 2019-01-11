package impl

import (
	"context"
	"github.com/stripe/stripe-go/charge"
	"net"

	"github.com/stripe/stripe-go"
	"gitlab.okta-solutions.com/mashroom/backend/cardpayment"
	"gitlab.okta-solutions.com/mashroom/backend/cardpayment/version"
	"gitlab.okta-solutions.com/mashroom/backend/common/errs"
	"gitlab.okta-solutions.com/mashroom/backend/common/health"
	"gitlab.okta-solutions.com/mashroom/backend/common/log"
	"google.golang.org/grpc"
)

type Server interface {
	cardpayment.CardpaymentServiceServer
	Serve(addr string)
	Background()
}

type serverImpl struct {
}

func (server *serverImpl) PaymentByCard(ctx context.Context, request *cardpayment.PaymentByCardRequest) (*cardpayment.PaymentByCardResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}

	stripe.Key = "sk_test_fwK7y4sYjzWGwuzNEg3aDv2y"

	//customerParams := &stripe.CustomerParams{}
	//customerParams.SetSource("src_18eYalAHEMiOZZp1l9ZTjSU0")
	//
	//c, err := customer.New(customerParams)
	//if err != nil{
	//	return nil, err
	//}

	token := request.Token
	amount := request.Amount
	//customerID := c.ID

	//stripe.Key = "sk_test_fwK7y4sYjzWGwuzNEg3aDv2y"

	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		//Customer: stripe.String(customerID),
	}
	chargeParams.SetSource(token)
	ch, err := charge.New(chargeParams)

	//chargeParams.SetSource("src_18eYalAHEMiOZZp1l9ZTjSU0")
	//ch, err := charge.New(chargeParams)
	if err != nil {
		return nil, err
	}

	return &cardpayment.PaymentByCardResponse{
		Success: ch.Paid,
		Message: ch.FailureMessage,
	}, nil
}

func (server *serverImpl) Background() {
	// background processes
}

func (server *serverImpl) Serve(addr string) {
	if listener, err := net.Listen("tcp", addr); err != nil {
		panic(err)
	} else {
		log.SetHost("cardpayment")
		grpcServer := grpc.NewServer()

		cardpayment.RegisterCardpaymentServiceServer(grpcServer, server)

		healthServer := version.NewHealthServer()
		health.RegisterHealthServiceServer(grpcServer, healthServer)

		log.Infoln("cardpayment started")
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
