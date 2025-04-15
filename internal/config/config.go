package config

type Config struct {
	UF              string
	Environment     Environment
	CertificatePath string
	CertificatePass string
	Timeout         int
}

type Environment int

const (
	Production Environment = iota
	Homologation
)
