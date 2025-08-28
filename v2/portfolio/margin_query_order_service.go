package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginQueryOrderService service to query margin orders
type MarginQueryOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *MarginQueryOrderService) Symbol(symbol string) *MarginQueryOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *MarginQueryOrderService) OrderID(orderID int64) *MarginQueryOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *MarginQueryOrderService) OrigClientOrderID(origClientOrderID string) *MarginQueryOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *MarginQueryOrderService) RecvWindow(recvWindow int64) *MarginQueryOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginQueryOrderService) Do(ctx context.Context) (*MarginQueryOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginQueryOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginQueryOrderResponse define query order response
type MarginQueryOrderResponse struct {
	ClientOrderID           string                  `json:"clientOrderId"`
	CummulativeQuoteQty     string                  `json:"cummulativeQuoteQty"`
	ExecutedQty             string                  `json:"executedQty"`
	IcebergQty              string                  `json:"icebergQty"`
	IsWorking               bool                    `json:"isWorking"`
	OrderID                 int64                   `json:"orderId"`
	OrigQty                 string                  `json:"origQty"`
	Price                   string                  `json:"price"`
	Side                    SideType                `json:"side"`
	Status                  OrderStatusType         `json:"status"`
	StopPrice               string                  `json:"stopPrice"`
	Symbol                  string                  `json:"symbol"`
	Time                    int64                   `json:"time"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	UpdateTime              int64                   `json:"updateTime"`
	AccountID               int64                   `json:"accountId"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	PreventedMatchID        string                  `json:"preventedMatchId"`
	PreventedQuantity       string                  `json:"preventedQuantity"`
}
