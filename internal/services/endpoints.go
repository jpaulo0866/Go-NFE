// internal/services/endpoints.go
package services

import (
	"fmt"

	"github.com/jpaulo0866/GO-NFE/internal/config"
)

var endpoints = map[string]map[config.Environment]string{
	"AM": {
		config.Production:   "https://nfe.sefaz.am.gov.br/services2/services/NfeAutorizacao4",
		config.Homologation: "https://homnfe.sefaz.am.gov.br/services2/services/NfeAutorizacao4",
	},
	"BA": {
		config.Production:   "https://nfe.sefaz.ba.gov.br/webservices/NFeAutorizacao4/NFeAutorizacao4.asmx",
		config.Homologation: "https://hnfe.sefaz.ba.gov.br/webservices/NFeAutorizacao4/NFeAutorizacao4.asmx",
	},
	"GO": {
		config.Production:   "https://nfe.sefaz.go.gov.br/nfe/services/NFeAutorizacao4",
		config.Homologation: "https://homolog.sefaz.go.gov.br/nfe/services/NFeAutorizacao4",
	},
	"MG": {
		config.Production:   "https://nfe.fazenda.mg.gov.br/nfe2/services/NFeAutorizacao4",
		config.Homologation: "https://hnfe.fazenda.mg.gov.br/nfe2/services/NFeAutorizacao4",
	},
	"MS": {
		config.Production:   "https://nfe.sefaz.ms.gov.br/ws/NFeAutorizacao4",
		config.Homologation: "https://hom.nfe.sefaz.ms.gov.br/ws/NFeAutorizacao4",
	},
	"MT": {
		config.Production:   "https://nfe.sefaz.mt.gov.br/nfews/v2/services/NfeAutorizacao4",
		config.Homologation: "https://homologacao.sefaz.mt.gov.br/nfews/v2/services/NfeAutorizacao4",
	},
	"PE": {
		config.Production:   "https://nfe.sefaz.pe.gov.br/nfe-service/services/NFeAutorizacao4",
		config.Homologation: "https://nfehomolog.sefaz.pe.gov.br/nfe-service/services/NFeAutorizacao4",
	},
	"PR": {
		config.Production:   "https://nfe.sefa.pr.gov.br/nfe/NFeAutorizacao4",
		config.Homologation: "https://homologacao.nfe.sefa.pr.gov.br/nfe/NFeAutorizacao4",
	},
	"RS": {
		config.Production:   "https://nfe.sefazrs.rs.gov.br/ws/NfeAutorizacao/NfeAutorizacao4.asmx",
		config.Homologation: "https://nfe-homologacao.sefazrs.rs.gov.br/ws/NfeAutorizacao/NfeAutorizacao4.asmx",
	},
	"SP": {
		config.Production:   "https://nfe.fazenda.sp.gov.br/ws/nfeautorizacao4.asmx",
		config.Homologation: "https://homologacao.nfe.fazenda.sp.gov.br/ws/nfeautorizacao4.asmx",
	},
	"SVAN": {
		config.Production:   "https://www.sefazvirtual.fazenda.gov.br/NFeAutorizacao4/NFeAutorizacao4.asmx",
		config.Homologation: "https://hom.sefazvirtual.fazenda.gov.br/NFeAutorizacao4/NFeAutorizacao4.asmx",
	},
	"SVRS": {
		config.Production:   "https://nfe.svrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx",
		config.Homologation: "https://nfe-homologacao.svrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx",
	},
}

func getEndpoint(uf string, env config.Environment) (string, error) {
	if stateEndpoints, ok := endpoints[uf]; ok {
		if endpoint, ok := stateEndpoints[env]; ok {
			return endpoint, nil
		}
	}
	return "", fmt.Errorf("endpoint not found for UF %s and environment %v", uf, env)
}
