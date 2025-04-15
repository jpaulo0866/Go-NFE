// internal/models/nfe.go
package models

import "time"

// NFe represents the main NFe document structure
type NFe struct {
	ID        string    `xml:"Id,attr"`
	Version   string    `xml:"versao,attr"`
	InfNFe    InfNFe    `xml:"infNFe"`
	Signature Signature `xml:"Signature,omitempty"`
}

// InfNFe contains the main NFe information
type InfNFe struct {
	ID      string       `xml:"Id,attr"`
	Version string       `xml:"versao,attr"`
	Ide     Ide          `xml:"ide"`
	Emit    Emitente     `xml:"emit"`
	Dest    Destinatario `xml:"dest"`
	Det     []Det        `xml:"det"`
	Total   Total        `xml:"total"`
	Transp  Transporte   `xml:"transp"`
	Cobr    Cobranca     `xml:"cobr,omitempty"`
	Pag     Pagamento    `xml:"pag"`
	InfAdic InfAdic      `xml:"infAdic,omitempty"`
}

// Ide contains identification information
type Ide struct {
	CUF      string    `xml:"cUF"`
	CNF      string    `xml:"cNF"`
	NatOp    string    `xml:"natOp"`
	Mod      string    `xml:"mod"`
	Serie    string    `xml:"serie"`
	NNF      string    `xml:"nNF"`
	DhEmi    time.Time `xml:"dhEmi"`
	TpNF     string    `xml:"tpNF"`
	IdDest   string    `xml:"idDest"`
	CMunFG   string    `xml:"cMunFG"`
	TpImp    string    `xml:"tpImp"`
	TpEmis   string    `xml:"tpEmis"`
	TpAmb    string    `xml:"tpAmb"`
	FinNFe   string    `xml:"finNFe"`
	IndFinal string    `xml:"indFinal"`
	IndPres  string    `xml:"indPres"`
	ProcEmi  string    `xml:"procEmi"`
	VerProc  string    `xml:"verProc"`
}

// Emitente represents the issuer
type Emitente struct {
	CNPJ      string   `xml:"CNPJ"`
	Nome      string   `xml:"xNome"`
	Fantasia  string   `xml:"xFant,omitempty"`
	EnderEmit Endereco `xml:"enderEmit"`
	IE        string   `xml:"IE"`
	CRT       string   `xml:"CRT"`
}

// Destinatario represents the recipient
type Destinatario struct {
	CNPJ      string   `xml:"CNPJ,omitempty"`
	CPF       string   `xml:"CPF,omitempty"`
	Nome      string   `xml:"xNome"`
	EnderDest Endereco `xml:"enderDest"`
	IndIEDest string   `xml:"indIEDest"`
	IE        string   `xml:"IE,omitempty"`
	Email     string   `xml:"email,omitempty"`
}

// Endereco represents an address
type Endereco struct {
	Logradouro  string `xml:"xLgr"`
	Numero      string `xml:"nro"`
	Complemento string `xml:"xCpl,omitempty"`
	Bairro      string `xml:"xBairro"`
	CodigoMun   string `xml:"cMun"`
	Municipio   string `xml:"xMun"`
	UF          string `xml:"UF"`
	CEP         string `xml:"CEP"`
	CodigoPais  string `xml:"cPais"`
	Pais        string `xml:"xPais"`
	Fone        string `xml:"fone,omitempty"`
}

// Det represents product details
type Det struct {
	NItem     string  `xml:"nItem,attr"`
	Prod      Produto `xml:"prod"`
	Imposto   Imposto `xml:"imposto"`
	InfAdProd string  `xml:"infAdProd,omitempty"`
}

// Produto represents product information
type Produto struct {
	CProd    string  `xml:"cProd"`
	CEAN     string  `xml:"cEAN"`
	XProd    string  `xml:"xProd"`
	NCM      string  `xml:"NCM"`
	CFOP     string  `xml:"CFOP"`
	UCom     string  `xml:"uCom"`
	QCom     float64 `xml:"qCom"`
	VUnCom   float64 `xml:"vUnCom"`
	VProd    float64 `xml:"vProd"`
	CEANTrib string  `xml:"cEANTrib"`
	UTrib    string  `xml:"uTrib"`
	QTrib    float64 `xml:"qTrib"`
	VUnTrib  float64 `xml:"vUnTrib"`
	IndTot   string  `xml:"indTot"`
}

// Imposto represents tax information
type Imposto struct {
	ICMS   ICMS   `xml:"ICMS"`
	PIS    PIS    `xml:"PIS"`
	COFINS COFINS `xml:"COFINS"`
}

// ICMS tax information
type ICMS struct {
	ICMS00 *ICMS00 `xml:"ICMS00,omitempty"`
	ICMS10 *ICMS10 `xml:"ICMS10,omitempty"`
	// Add other ICMS types as needed
}

// ICMS00 represents ICMS tax type 00
type ICMS00 struct {
	Orig  string  `xml:"orig"`
	CST   string  `xml:"CST"`
	ModBC string  `xml:"modBC"`
	VBC   float64 `xml:"vBC"`
	PICMS float64 `xml:"pICMS"`
	VICMS float64 `xml:"vICMS"`
}

// ICMS00 represents ICMS tax type 00
type ICMS10 struct {
	Orig  string  `xml:"orig"`
	CST   string  `xml:"CST"`
	ModBC string  `xml:"modBC"`
	VBC   float64 `xml:"vBC"`
	PICMS float64 `xml:"pICMS"`
	VICMS float64 `xml:"vICMS"`
}

// PIS tax information
type PIS struct {
	PISAliq *PISAliq `xml:"PISAliq,omitempty"`
	// Add other PIS types as needed
}

// PISAliq represents PIS tax calculation
type PISAliq struct {
	CST  string  `xml:"CST"`
	VBC  float64 `xml:"vBC"`
	PPIS float64 `xml:"pPIS"`
	VPIS float64 `xml:"vPIS"`
}

// COFINS tax information
type COFINS struct {
	COFINSAliq *COFINSAliq `xml:"COFINSAliq,omitempty"`
	// Add other COFINS types as needed
}

// COFINSAliq represents COFINS tax calculation
type COFINSAliq struct {
	CST     string  `xml:"CST"`
	VBC     float64 `xml:"vBC"`
	PCOFINS float64 `xml:"pCOFINS"`
	VCOFINS float64 `xml:"vCOFINS"`
}

// Total represents total values
type Total struct {
	ICMSTot ICMSTot `xml:"ICMSTot"`
}

// ICMSTot represents ICMS totals
type ICMSTot struct {
	VBC     float64 `xml:"vBC"`
	VICMS   float64 `xml:"vICMS"`
	VProd   float64 `xml:"vProd"`
	VFrete  float64 `xml:"vFrete"`
	VSeg    float64 `xml:"vSeg"`
	VDesc   float64 `xml:"vDesc"`
	VPIS    float64 `xml:"vPIS"`
	VCOFINS float64 `xml:"vCOFINS"`
	VNF     float64 `xml:"vNF"`
}

// Transporte represents transportation information
type Transporte struct {
	ModFrete   string          `xml:"modFrete"`
	Transporta *Transportadora `xml:"transporta,omitempty"`
	Vol        []Volume        `xml:"vol,omitempty"`
}

// Transportadora represents carrier information
type Transportadora struct {
	CNPJ     string `xml:"CNPJ,omitempty"`
	CPF      string `xml:"CPF,omitempty"`
	Nome     string `xml:"xNome"`
	IE       string `xml:"IE,omitempty"`
	Endereco string `xml:"xEnder,omitempty"`
	UF       string `xml:"UF,omitempty"`
}

// Volume represents volume information
type Volume struct {
	QVol  string  `xml:"qVol"`
	Esp   string  `xml:"esp,omitempty"`
	Marca string  `xml:"marca,omitempty"`
	PesoL float64 `xml:"pesoL,omitempty"`
	PesoB float64 `xml:"pesoB,omitempty"`
}

// Cobranca represents billing information
type Cobranca struct {
	Fat Fatura      `xml:"fat,omitempty"`
	Dup []Duplicata `xml:"dup,omitempty"`
}

// Fatura represents invoice information
type Fatura struct {
	NFat  string  `xml:"nFat,omitempty"`
	VOrig float64 `xml:"vOrig,omitempty"`
	VDesc float64 `xml:"vDesc,omitempty"`
	VLiq  float64 `xml:"vLiq,omitempty"`
}

// Duplicata represents payment installment
type Duplicata struct {
	NDup  string    `xml:"nDup"`
	DVenc time.Time `xml:"dVenc"`
	VDup  float64   `xml:"vDup"`
}

// Pagamento represents payment information
type Pagamento struct {
	DetPag []DetalhePagamento `xml:"detPag"`
}

// DetalhePagamento represents payment details
type DetalhePagamento struct {
	TPag string  `xml:"tPag"`
	VPag float64 `xml:"vPag"`
}

// InfAdic represents additional information
type InfAdic struct {
	InfCpl     string `xml:"infCpl,omitempty"`
	InfAdFisco string `xml:"infAdFisco,omitempty"`
}

// Signature represents the digital signature
type Signature struct {
	SignedInfo     SignedInfo `xml:"SignedInfo"`
	SignatureValue string     `xml:"SignatureValue"`
	KeyInfo        KeyInfo    `xml:"KeyInfo"`
}

// SignedInfo represents signature information
type SignedInfo struct {
	CanonicalizationMethod Algorithm `xml:"CanonicalizationMethod"`
	SignatureMethod        Algorithm `xml:"SignatureMethod"`
	Reference              Reference `xml:"Reference"`
}

// Algorithm represents signature algorithm
type Algorithm struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// Reference represents signature reference
type Reference struct {
	URI          string     `xml:"URI,attr"`
	Transforms   Transforms `xml:"Transforms"`
	DigestMethod Algorithm  `xml:"DigestMethod"`
	DigestValue  string     `xml:"DigestValue"`
}

// Transforms represents signature transforms
type Transforms struct {
	Transform []Algorithm `xml:"Transform"`
}

// KeyInfo represents key information
type KeyInfo struct {
	X509Data X509Data `xml:"X509Data"`
}

// X509Data represents certificate data
type X509Data struct {
	X509Certificate string `xml:"X509Certificate"`
}
