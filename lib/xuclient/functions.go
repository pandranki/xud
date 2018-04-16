package xuclient

import (
	"gopkg.in/urfave/cli.v2"

	"bytes"
	"net/http"

	"io/ioutil"
	"math/rand"

	"encoding/json"
)

func GetInfo(c *cli.Context) error {
	resp, err := callrpc("", "getInfo")
	if err != nil {
		return err
	}
	println(resp)
	return nil
}

func GetOrders(c *cli.Context) error {
	resp, err := callrpc("", "getOrders")
	if err != nil {
		return err
	}
	println(resp)
	return nil
}

func PlaceOrder(c *cli.Context) error {
	resp, err := callrpc("", "placeOrder")
	if err != nil {
		return err
	}
	println(resp)
	return nil
}

func Connect(c *cli.Context) error {
	resp, err := callrpc("", "connect")
	if err != nil {
		return err
	}
	println(resp)
	return nil
}

func TokenSwap(c *cli.Context) error {
	resp, err := callrpc("", "tokenSwap")
	if err != nil {
		return err
	}
	println(resp)
	return nil
}

func Shutdown(c *cli.Context) error {
	resp, err := callrpc("", "shutdown")
	if err != nil {
		return err
	}
	println(resp)
	return nil
}

func callrpc(address string, method string, params ...interface{}) (string, error) {
	newdata, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     rand.Uint32(),
		"params": params,
	})

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", address, bytes.NewBuffer(newdata))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodystr := string(body)
		return bodystr, nil
	}

	return "", nil
}
