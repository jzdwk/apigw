package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	gokong "github.com/hbagdi/go-kong/kong"
)

// certificateCRUD implements crud.Actions interface.
type certificateCRUD struct {
	client *gokong.Client
}

func certificateFromStuct(arg crud.Event) *kong.Certificate {
	certificate, ok := arg.Obj.(*kong.Certificate)
	if !ok {
		panic("unexpected type, expected *kong.certificate")
	}
	return certificate
}

// Create creates a Certificate in Kong.
// The arg should be of type crud.Event, containing the certificate to be created,
// else the function will panic.
// It returns a the created *kong.Certificate.
func (s *certificateCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	certificate := certificateFromStuct(event)
	createdCertificate, err := s.client.Certificates.Create(nil, &certificate.Certificate)
	if err != nil {
		return nil, err
	}
	return &kong.Certificate{Certificate: *createdCertificate}, nil
}

// Delete deletes a Certificate in Kong.
// The arg should be of type crud.Event, containing the certificate to be deleted,
// else the function will panic.
// It returns a the deleted *kong.Certificate.
func (s *certificateCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	certificate := certificateFromStuct(event)
	err := s.client.Certificates.Delete(nil, certificate.ID)
	if err != nil {
		return nil, err
	}
	return certificate, nil
}

// Update updates a Certificate in Kong.
// The arg should be of type crud.Event, containing the certificate to be updated,
// else the function will panic.
// It returns a the updated *kong.Certificate.
func (s *certificateCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	certificate := certificateFromStuct(event)

	updatedCertificate, err := s.client.Certificates.Create(nil, &certificate.Certificate)
	if err != nil {
		return nil, err
	}
	return &kong.Certificate{Certificate: *updatedCertificate}, nil
}
