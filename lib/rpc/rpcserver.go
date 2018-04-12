package rpc

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	json "github.com/gorilla/rpc/v2/json2"
	log "github.com/sirupsen/logrus"
)

type RPCService struct {
}

type Response struct {
	Result interface{}
}

type Args struct {
}

func (rm *RPCService) getInfo(r *http.Request, args *Args, result *Response) error {
	return nil
}

func (rm *RPCService) getOrders(r *http.Request, args *Args, result *Response) error {
	return nil
}

func (rm *RPCService) placeOrder(r *http.Request, args *Args, result *Response) error {
	return nil
}

func (rm *RPCService) connect(r *http.Request, args *Args, result *Response) error {
	return nil
}

func (rm *RPCService) tokenSwap(r *http.Request, args *Args, result *Response) error {
	return nil
}

func (rm *RPCService) shutdown(r *http.Request, args *Args, result *Response) error {
	os.Exit(0)
	return nil
}

func StartRPCServer(rpcport string) {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	rpcservice := new(RPCService)
	rpcServer.RegisterService(rpcservice, "")
	router := mux.NewRouter()
	router.Handle("/rpc", rpcServer)
	if err := http.ListenAndServe(":"+rpcport, router); err != nil {
		log.Fatalln(err)
	}
}
