package service

import (
	"crypto/x509"
	"encoding/pem"
	"main/internal/database"
	"main/internal/token"
)

type Supplier interface {
	GetAuthService() AuthService
	GetStationService() StationService
}

type serviceSupplier struct {
	authService    AuthService
	stationService StationService
}

func SetUp(issuerHost, audienceHost, privateKey, dsn string) (Supplier, error) {
	return createSupplier(issuerHost, audienceHost, privateKey, dsn)
}

// TODO: add release method to close database connection if needed
func TearDown(services Supplier) error {
	// if ss, ok := services.(*serviceSupplier); ok {
	// 	return ss.release()
	// }
	return nil
}

func createSupplier(issuerHost, audienceHost, privateKey, dsn string) (*serviceSupplier, error) {
	// connect to database
	db, err := database.MakePostgresDB(dsn)
	if err != nil {
		return nil, err
	}

	// parse private key
	var authService AuthService
	if len(privateKey) > 0 {
		block, _ := pem.Decode([]byte(privateKey))
		key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		authService = NewAuthService(db, &token.Options{
			IssuerHost:   issuerHost,
			AudienceHost: audienceHost,
			PrivateKey:   key,
		})
	}

	ss := &serviceSupplier{
		authService:    authService,
		stationService: NewStationService(db),
	}

	return ss, nil
}

func (s *serviceSupplier) GetAuthService() AuthService {
	return s.authService
}

func (s *serviceSupplier) GetStationService() StationService {
	return s.stationService
}
