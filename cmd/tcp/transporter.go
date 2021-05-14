package tcp

import (
	"encoding/json"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
	"net"
)

type Transport interface {
	MakeRequest(data *models.DataTransportRequest) (*models.DataTransportResponse, error)
}

type tcpRequest struct {
	conn net.Conn
	log  *logrus.Logger
}

func (t tcpRequest) MakeRequest(data *models.DataTransportRequest) (*models.DataTransportResponse, error) {

	if err := json.NewEncoder(t.conn).Encode(data); err != nil {
		t.log.Errorf("TCP (MakeRequest) Encode - %s", err.Error())
		return nil, err
	}

	response := models.DataTransportResponse{}
	if err := json.NewDecoder(t.conn).Decode(&response); err != nil {
		t.log.Errorf("TCP (MakeRequest) Decode - %v", err)
		return nil, err
	}

	return &response, nil
}

func Connect(logger *logrus.Logger) (Transport, error) {
	conn, err := net.Dial("tcp", "localhost:4444")
	if err != nil {
		logger.Errorf("Connect (simpl deamon): dail conn refused - %s", err.Error())
		return nil, err
	}

	return &tcpRequest{
		conn: conn,
		log: logger,
	}, nil
}
