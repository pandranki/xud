# Exchange Union Golang Version **⚠️ WIP** [![Build Status](https://travis-ci.org/indxcrypto/xud.svg?branch=dev)](https://travis-ci.org/indxcrypto/xud)

[Exchange Union](https://www.exchangeunion.com/) (XU) is a decentralized exchange layer built on the [Lightning](http://lightning.network/) and [Raiden](https://raiden.network/) networks to enable trustless and instant cryptocurrency swaps and order fulfillment between exchanges.

This repo contains the early stages of the Exchange Union Daemon ("xud") port in golang which encompasses the following components:

* Integration with [lnd](https://github.com/lightningnetwork/lnd) and [raiden](https://github.com/raiden-network/raiden) nodes.
* Orderbook data stored in a local sqlite database.
* Peer-to-peer networking with other XU nodes via TCP.
* A JSON-RPC API to serve other applications and a command-line interface.

Contributions, feedback, and questions are welcome.

## Install

First, clone the repository from GitHub and install dependencies.

```bash
#Let go automatically compile and install
$ go install github.com/indxcrypto/xud

#OR do the hard way
$ git clone https://github.com/indxcrypto/xud
$ cd xud
$ govendor sync #This Project uses govendor to manage dependencies.
$ go build && go install ./..
```

XUD uses Sqlite for now but not limited to it and can be extended to other sql databases.

## Starting the Daemon

Assuming you have copied or placed the `xud` and `xucli` programs in `~/xud/bin` folder / path on your system.

```bash
~/xud $ cd bin
~/xud/bin $ ./xud
2018-3-2 11:36:06 - info: config loaded
2018-3-2 11:36:06 - info: connected to database
2018-3-2 11:36:06 - info: P2P server listening on port 8885
2018-3-2 11:36:06 - info: RPC server listening on port 8886
```

## Command-Line Interface

```bash
~/xud/bin $ ./xucli placeorder --price 0.12 --quantity 5
{"fieldCount":0,"affectedRows":1,"insertId":1,"serverStatus":2,"warningCount":0,"message":"","protocol41":true,"changedRows":0}
~/xud/bin $ ./xucli getorders
{"bids":[{"price":0.12,"quantity":5}],"asks":[]}
```

## Configuration

The configuration file uses [TOML](https://github.com/toml-lang/toml) and by default is located at  `~/.xud/xud.conf` on Linux or `AppData\Local\Xud\xud.conf` on Windows. Default settings which can be overridden are shown below.

```toml
lndDir = "~/.lnd"
rpcPort = 8886

[db]
username = "xud"
password = ""
database = "xud"
port = 3306
host = "localhost"

[p2p]
listen = true
port = 8885
```
