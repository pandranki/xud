package raidenclient

type RaidenClient struct {
	host string
	port string
}

//NewRaidenClient Creates a New Raiden API Client
func NewRaidenClient(host, port string) *RaidenClient {
	rdnclient := &RaidenClient{host: host, port: port}
	return rdnclient
}

func (rndc *RaidenClient) TokenSwap() (bool, error) {
	return true, nil
}
