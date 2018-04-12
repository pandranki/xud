package core

import (
	"github.com/indxcrypto/xud/lib/rpc"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type XUD struct {
	config   *viper.Viper
	db       *gorm.DB
	shutdown chan bool
	xuddir   string
	lnddir   string
}

func NewXUD(db *gorm.DB, config *viper.Viper) *XUD {
	xud := &XUD{config: config, db: db, xuddir: viper.GetString("xuddir"), lnddir: viper.GetString("lndDir")}
	return xud
}

func (xud *XUD) Start() error {
	if xud.config.GetBool("p2p.listen") {
		//Start P2P
		log.Infoln("Started P2P Server on Port", xud.config.GetString("p2p.port"))
		go xud.startP2PServer(xud.config.GetString("p2p.port"))
	}
	//Start RPC
	go rpc.StartRPCServer(xud.config.GetString("rpcport"))
	log.Infoln("Started RPC Server on Port", xud.config.GetString("rpcport"))

	select {}
	return nil
}

func (xud *XUD) Shutdown() error {
	//Stop P2P
	log.Warningln("Stopping P2P server running on port", xud.config.GetString("p2p.port"))
	//Stop RPC
	log.Warningln("Stopping RPC server running on port", xud.config.GetString("rpcport"))
	return nil
}

func (xud *XUD) startP2PServer(listenPort string) p2p.Server {
	nodekey, _ := crypto.GenerateKey()
	server := p2p.Server{
		Config: p2p.Config{
			MaxPeers:   10,
			PrivateKey: nodekey,
			Name:       "my node name",
			ListenAddr: ":30300",
			Protocols:  []p2p.Protocol{},
		},
	}
	if err := server.Start(); err != nil {
		log.Fatalf("Could not start P2P server: %v", err)
	}
	return server
}
