package bybit

import (
	"github.com/gorilla/websocket"
	"net"
)

// V5WebsocketServiceI :
type V5WebsocketServiceI interface {
	Public(CategoryV5) (V5WebsocketPublicService, error)
	PublicWithSourceIP(CategoryV5, sourceIP string) (V5WebsocketPublicService, error)
	Private() (V5WebsocketPrivateService, error)
	PrivateWithSourceIP(sourceIP string) (V5WebsocketPrivateService, error)
	Trade() (V5WebsocketTradeService, error)
	TradeWithSourceIP(sourceIP string) (V5WebsocketPrivateService, error)
}

// V5WebsocketService :
type V5WebsocketService struct {
	client *WebSocketClient
}

// Public :
func (s *V5WebsocketService) Public(category CategoryV5) (V5WebsocketPublicServiceI, error) {
	url := s.client.baseURL + V5WebsocketPublicPathFor(category)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPublicService{
		client:              s.client,
		connection:          c,
		category:            category,
		paramOrderBookMap:   make(map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error),
		paramKlineMap:       make(map[V5WebsocketPublicKlineParamKey]func(V5WebsocketPublicKlineResponse) error),
		paramTickerMap:      make(map[V5WebsocketPublicTickerParamKey]func(V5WebsocketPublicTickerResponse) error),
		paramTradeMap:       make(map[V5WebsocketPublicTradeParamKey]func(V5WebsocketPublicTradeResponse) error),
		paramLiquidationMap: make(map[V5WebsocketPublicLiquidationParamKey]func(V5WebsocketPublicLiquidationResponse) error),
	}, nil
}

// PublicWithSourceIP :
func (s *V5WebsocketService) PublicWithSourceIP(category CategoryV5, sourceIP string) (V5WebsocketPublicServiceI, error) {
	url := s.client.baseURL + V5WebsocketPublicPathFor(category)

	// 创建一个自定义的 指定了SourceIP的Dialer
	dialer := genDialerWithSourceIP(sourceIP)
	c, _, err := dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPublicService{
		client:              s.client,
		connection:          c,
		category:            category,
		paramOrderBookMap:   make(map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error),
		paramKlineMap:       make(map[V5WebsocketPublicKlineParamKey]func(V5WebsocketPublicKlineResponse) error),
		paramTickerMap:      make(map[V5WebsocketPublicTickerParamKey]func(V5WebsocketPublicTickerResponse) error),
		paramTradeMap:       make(map[V5WebsocketPublicTradeParamKey]func(V5WebsocketPublicTradeResponse) error),
		paramLiquidationMap: make(map[V5WebsocketPublicLiquidationParamKey]func(V5WebsocketPublicLiquidationResponse) error),
	}, nil
}

// Private :
func (s *V5WebsocketService) Private() (V5WebsocketPrivateServiceI, error) {
	url := s.client.baseURL + V5WebsocketPrivatePath
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPrivateService{
		client:            s.client,
		connection:        c,
		paramOrderMap:     make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateOrderResponse) error),
		paramPositionMap:  make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponse) error),
		paramExecutionMap: make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateExecutionResponse) error),
		paramWalletMap:    make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateWalletResponse) error),
	}, nil
}

// PrivateWithSourceIP :
func (s *V5WebsocketService) PrivateWithSourceIP(sourceIP string) (V5WebsocketPrivateServiceI, error) {
	url := s.client.baseURL + V5WebsocketPrivatePath
	// 创建一个自定义的 指定了SourceIP的Dialer
	dialer := genDialerWithSourceIP(sourceIP)
	c, _, err := dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPrivateService{
		client:            s.client,
		connection:        c,
		paramOrderMap:     make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateOrderResponse) error),
		paramPositionMap:  make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponse) error),
		paramExecutionMap: make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateExecutionResponse) error),
		paramWalletMap:    make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateWalletResponse) error),
	}, nil
}

// Trade :
func (s *V5WebsocketService) Trade() (V5WebsocketTradeServiceI, error) {
	url := s.client.baseURL + V5WebsocketTradePath
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketTradeService{
		client:     s.client,
		connection: c,
	}, nil
}

// TradeWithSourceIP :
func (s *V5WebsocketService) TradeWithSourceIP(sourceIP string) (V5WebsocketTradeServiceI, error) {
	url := s.client.baseURL + V5WebsocketTradePath
	// 创建一个自定义的 指定了SourceIP的Dialer
	dialer := genDialerWithSourceIP(sourceIP)
	c, _, err := dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketTradeService{
		client:     s.client,
		connection: c,
	}, nil
}

// V5 :
func (c *WebSocketClient) V5() *V5WebsocketService {
	return &V5WebsocketService{c}
}

func genDialerWithSourceIP(sourceIP string) *websocket.Dialer {
	dialer := websocket.DefaultDialer
	dialer.NetDial = func(network, addr string) (net.Conn, error) {
		// 创建本地 TCP 地址
		localAddr, err := net.ResolveTCPAddr(network, sourceIP+":0") // 端口可以是 0，表示随机端口
		if err != nil {
			return nil, err
		}

		// 解析远程地址
		remoteAddr, err := net.ResolveTCPAddr(network, addr)
		if err != nil {
			return nil, err
		}

		// 使用指定的本地地址进行拨号
		return net.DialTCP(network, localAddr, remoteAddr)
	}
	return dialer
}
