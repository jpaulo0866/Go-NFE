// internal/services/webservice.go
package services

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jpaulo0866/GO-NFE/internal/config"
	"github.com/jpaulo0866/GO-NFE/internal/models"
	"github.com/jpaulo0866/GO-NFE/internal/utils"
)

type WebService struct {
	config     *config.Config
	cert       *Certificate
	httpClient *http.Client
	endpoint   string
}

func NewWebService(cfg *config.Config, cert *Certificate) (*WebService, error) {
	endpoint, err := getEndpoint(cfg.UF, cfg.Environment)
	if err != nil {
		return nil, err
	}

	// Create TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{cert.Cert.Raw},
				PrivateKey:  cert.PrivKey,
			},
		},
	}

	// Create HTTP client
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(cfg.Timeout) * time.Second,
	}

	return &WebService{
		config:     cfg,
		cert:       cert,
		httpClient: client,
		endpoint:   endpoint,
	}, nil
}

func (ws *WebService) callSOAP(ctx context.Context, action string, payload interface{}) ([]byte, error) {
	envelope := createSOAPEnvelope(payload)

	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", ws.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Set("SOAPAction", action)

	resp, err := ws.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (ws *WebService) StatusServico(ctx context.Context) (*models.StatusServicoResponse, error) {
	params := map[string]interface{}{
		"versao": "4.00",
		"tpAmb":  ws.config.Environment,
		"cUF":    ws.config.UF,
	}

	response, err := ws.callSOAP(ctx, "nfeStatusServicoNF", params)
	if err != nil {
		return nil, err
	}

	var result models.StatusServicoResponse
	if err := xml.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (ws *WebService) AutorizaNFe(ctx context.Context, nfe *models.NFe) (*models.AutorizacaoResponse, error) {
	xmlData, err := utils.Marshal(nfe)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"nfeDadosMsg": string(xmlData),
	}

	response, err := ws.callSOAP(ctx, "nfeAutorizacaoLote", params)
	if err != nil {
		return nil, err
	}

	var result models.AutorizacaoResponse
	if err := xml.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (ws *WebService) ConsultaRecibo(ctx context.Context, recibo string) (*models.ConsultaReciboResponse, error) {
	params := map[string]interface{}{
		"versao": "4.00",
		"tpAmb":  ws.config.Environment,
		"nRec":   recibo,
	}

	response, err := ws.callSOAP(ctx, "nfeRetAutorizacaoLote", params)
	if err != nil {
		return nil, err
	}

	var result models.ConsultaReciboResponse
	if err := xml.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
