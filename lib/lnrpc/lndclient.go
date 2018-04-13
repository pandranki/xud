package lnrpc

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

//LndClient instance of a Lightning Client
type LndClient struct {
	lightning LightningClient
	grpcmeta  metadata.MD
}

//NewLNDClient creates new lightning network client
func NewLNDClient(host, port string) *LndClient {
	lndc := &LndClient{}
	var opts []grpc.DialOption
	certFile := viper.GetString("lndDir") + "tls.cert"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(host + ":" + port)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	adminMacaroon := viper.GetString("lndDir") + "admin.macaroon"
	//TODO Add adminMacaroon to GRPC Metadata
	log.Debugln("AdminMacaroon:", adminMacaroon)
	lightning := NewLightningClient(conn)
	lndc.lightning = lightning
	return lndc
}

//GetInfo Return general information concerning the lightning node including it’s identity pubkey,
//alias, the chains it is connected to, and information concerning the number of
//open+pending channels.
func (lc *LndClient) GetInfo() (*GetInfoResponse, error) {
	resp, err := lc.lightning.GetInfo(context.Background(), &GetInfoRequest{})
	if err != nil {
		log.Errorln(err)
		return resp, err
	}
	return resp, nil
}

//AddInvoice Attempt to add a new invoice to the lnd invoice database.
//value - The value of this invoice in satoshis
func (lc *LndClient) AddInvoice(value int64) (*AddInvoiceResponse, error) {
	invoice := &Invoice{Memo: "XU", Value: value}
	resp, err := lc.lightning.AddInvoice(context.Background(), invoice)
	if err != nil {
		log.Errorln(err)
		return resp, err
	}
	return resp, nil

}

//PayInvoice Pay an invoice through the Lightning Network.
//paymentRequest - An invoice for a payment within the Lightning Network.
func (lc *LndClient) PayInvoice(paymentRequest string) (*SendResponse, error) {
	resp, err := lc.lightning.SendPaymentSync(context.Background(), &SendRequest{PaymentRequest: paymentRequest})
	if err != nil {
		log.Errorln(err)
		return resp, err
	}
	return resp, nil
}
