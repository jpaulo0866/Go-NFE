// internal/services/soap.go
package services

import "encoding/xml"

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`
	Body    SOAPBody
}

type SOAPBody struct {
	XMLName xml.Name    `xml:"soap:Body"`
	Content interface{} `xml:",omitempty"`
}

func createSOAPEnvelope(content interface{}) *SOAPEnvelope {
	return &SOAPEnvelope{
		XMLNS: "http://www.w3.org/2003/05/soap-envelope",
		Body: SOAPBody{
			Content: content,
		},
	}
}
