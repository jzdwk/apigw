package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	"apigw/util"
	gokong "github.com/hbagdi/go-kong/kong"
)

// basicAuthCRUD implements crud.Actions interface.
type basicAuthCRUD struct {
	client *gokong.Client
}

func basicAuthFromStuct(arg crud.Event) *kong.BasicAuth {
	basicAuth, ok := arg.Obj.(*kong.BasicAuth)
	if !ok {
		panic("unexpected type, expected *kong.BasicAuth")
	}

	return basicAuth
}

// Create creates a Route in Kong.
// The arg should be of type crud.Event, containing the basicAuth to be created,
// else the function will panic.
// It returns a the created *kong.BasicAuth.
func (s *basicAuthCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	basicAuth := basicAuthFromStuct(event)
	cid := ""
	if !util.Empty(basicAuth.Consumer.Username) {
		cid = *basicAuth.Consumer.Username
	}
	if !util.Empty(basicAuth.Consumer.ID) {
		cid = *basicAuth.Consumer.ID
	}
	createdBasicAuth, err := s.client.BasicAuths.Create(nil, &cid,
		&basicAuth.BasicAuth)
	if err != nil {
		return nil, err
	}
	return &kong.BasicAuth{BasicAuth: *createdBasicAuth}, nil
}

// Delete deletes a Route in Kong.
// The arg should be of type crud.Event, containing the basicAuth to be deleted,
// else the function will panic.
// It returns a the deleted *kong.BasicAuth.
func (s *basicAuthCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	basicAuth := basicAuthFromStuct(event)
	cid := ""
	if !util.Empty(basicAuth.Consumer.Username) {
		cid = *basicAuth.Consumer.Username
	}
	if !util.Empty(basicAuth.Consumer.ID) {
		cid = *basicAuth.Consumer.ID
	}
	err := s.client.BasicAuths.Delete(nil, &cid, basicAuth.ID)
	if err != nil {
		return nil, err
	}
	return basicAuth, nil
}

// Update updates a Route in Kong.
// The arg should be of type crud.Event, containing the basicAuth to be updated,
// else the function will panic.
// It returns a the updated *kong.BasicAuth.
func (s *basicAuthCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	basicAuth := basicAuthFromStuct(event)

	cid := ""
	if !util.Empty(basicAuth.Consumer.Username) {
		cid = *basicAuth.Consumer.Username
	}
	if !util.Empty(basicAuth.Consumer.ID) {
		cid = *basicAuth.Consumer.ID
	}
	updatedBasicAuth, err := s.client.BasicAuths.Create(nil, &cid, &basicAuth.BasicAuth)
	if err != nil {
		return nil, err
	}
	return &kong.BasicAuth{BasicAuth: *updatedBasicAuth}, nil
}
