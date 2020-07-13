package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	gokong "github.com/hbagdi/go-kong/kong"
)

// caCertificateCRUD implements crud.Actions interface.
type caCertificateCRUD struct {
	client *gokong.Client
}

func caCertFromStuct(arg crud.Event) *kong.CACertificate {
	caCert, ok := arg.Obj.(*kong.CACertificate)
	if !ok {
		panic("unexpected type, expected *kong.CACertificate")
	}
	return caCert
}

// Create creates a CACertificate in Kong.
// The arg should be of type crud.Event, containing the certificate to be created,
// else the function will panic.
// It returns a the created *kong.CACertificate.
func (s *caCertificateCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	certificate := caCertFromStuct(event)
	createdCertificate, err := s.client.CACertificates.Create(nil,
		&certificate.CACertificate)
	if err != nil {
		return nil, err
	}
	return &kong.CACertificate{CACertificate: *createdCertificate}, nil
}

// Delete deletes a CACertificate in Kong.
// The arg should be of type crud.Event, containing the certificate to be deleted,
// else the function will panic.
// It returns a the deleted *kong.CACertificate.
func (s *caCertificateCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	certificate := caCertFromStuct(event)
	err := s.client.CACertificates.Delete(nil, certificate.ID)
	if err != nil {
		return nil, err
	}
	return certificate, nil
}

// Update updates a CACertificate in Kong.
// The arg should be of type crud.Event, containing the certificate to be updated,
// else the function will panic.
// It returns a the updated *kong.CACertificate.
func (s *caCertificateCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	certificate := caCertFromStuct(event)
	updatedCertificate, err := s.client.CACertificates.Create(nil,
		&certificate.CACertificate)
	if err != nil {
		return nil, err
	}
	return &kong.CACertificate{CACertificate: *updatedCertificate}, nil
}
