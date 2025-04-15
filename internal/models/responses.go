// internal/models/responses.go
package models

type StatusServicoResponse struct {
	Version   string `xml:"versao,attr"`
	TpAmb     string `xml:"tpAmb"`
	VerAplic  string `xml:"verAplic"`
	CStat     string `xml:"cStat"`
	XMotivo   string `xml:"xMotivo"`
	CUF       string `xml:"cUF"`
	DhRecbto  string `xml:"dhRecbto"`
	TMed      string `xml:"tMed"`
	DhRetorno string `xml:"dhRetorno"`
	XObs      string `xml:"xObs,omitempty"`
}

type AutorizacaoResponse struct {
	Version  string `xml:"versao,attr"`
	TpAmb    string `xml:"tpAmb"`
	VerAplic string `xml:"verAplic"`
	CStat    string `xml:"cStat"`
	XMotivo  string `xml:"xMotivo"`
	CUF      string `xml:"cUF"`
	NRec     string `xml:"nRec,omitempty"`
}

type ConsultaReciboResponse struct {
	Version  string  `xml:"versao,attr"`
	TpAmb    string  `xml:"tpAmb"`
	VerAplic string  `xml:"verAplic"`
	CStat    string  `xml:"cStat"`
	XMotivo  string  `xml:"xMotivo"`
	CUF      string  `xml:"cUF"`
	ProtNFe  ProtNFe `xml:"protNFe,omitempty"`
}

type ProtNFe struct {
	Version string  `xml:"versao,attr"`
	InfProt InfProt `xml:"infProt"`
}

type InfProt struct {
	TpAmb    string `xml:"tpAmb"`
	VerAplic string `xml:"verAplic"`
	ChNFe    string `xml:"chNFe"`
	DhRecbto string `xml:"dhRecbto"`
	NProt    string `xml:"nProt"`
	DigVal   string `xml:"digVal"`
	CStat    string `xml:"cStat"`
	XMotivo  string `xml:"xMotivo"`
}
