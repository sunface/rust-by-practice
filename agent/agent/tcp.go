package agent

// TcpClient vgo通信tcp客户端
type TcpClient struct {
}

// NewTcpClient ...
func NewTcpClient() *TcpClient {
	return &TcpClient{}
}

// Start ...
func (t *TcpClient) Start() error {
	return nil
}

// Close ....
func (t *TcpClient) Close() error {
	return nil
}
