package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	"apigw/util"
	gokong "github.com/hbagdi/go-kong/kong"
)

// hmacAuthCRUD implements crud.Actions interface.
type hmacAuthCRUD struct {
	client *gokong.Client
}

func hmacAuthFromStuct(arg crud.Event) *kong.HMACAuth {
	hmacAuth, ok := arg.Obj.(*kong.HMACAuth)
	if !ok {
		panic("unexpected type, expected *kong.HMACAuth")
	}

	return hmacAuth
}

// Create creates a Route in Kong.
// The arg should be of type crud.Event, containing the hmacAuth to be created,
// else the function will panic.
// It returns a the created *kong.HMACAuth.
func (s *hmacAuthCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	hmacAuth := hmacAuthFromStuct(event)
	cid := ""
	if !util.Empty(hmacAuth.Consumer.Username) {
		cid = *hmacAuth.Consumer.Username
	}
	if !util.Empty(hmacAuth.Consumer.ID) {
		cid = *hmacAuth.Consumer.ID
	}
	createdHMACAuth, err := s.client.HMACAuths.Create(nil, &cid,
		&hmacAuth.HMACAuth)
	if err != nil {
		return nil, err
	}
	return &kong.HMACAuth{HMACAuth: *createdHMACAuth}, nil
}

// Delete deletes a Route in Kong.
// The arg should be of type crud.Event, containing the hmacAuth to be deleted,
// else the function will panic.
// It returns a the deleted *kong.HMACAuth.
func (s *hmacAuthCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	hmacAuth := hmacAuthFromStuct(event)
	cid := ""
	if !util.Empty(hmacAuth.Consumer.Username) {
		cid = *hmacAuth.Consumer.Username
	}
	if !util.Empty(hmacAuth.Consumer.ID) {
		cid = *hmacAuth.Consumer.ID
	}
	err := s.client.HMACAuths.Delete(nil, &cid, hmacAuth.ID)
	if err != nil {
		return nil, err
	}
	return hmacAuth, nil
}

// Update updates a Route in Kong.
// The arg should be of type crud.Event, containing the hmacAuth to be updated,
// else the function will panic.
// It returns a the updated *kong.HMACAuth.
func (s *hmacAuthCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	hmacAuth := hmacAuthFromStuct(event)

	cid := ""
	if !util.Empty(hmacAuth.Consumer.Username) {
		cid = *hmacAuth.Consumer.Username
	}
	if !util.Empty(hmacAuth.Consumer.ID) {
		cid = *hmacAuth.Consumer.ID
	}
	updatedHMACAuth, err := s.client.HMACAuths.Create(nil, &cid, &hmacAuth.HMACAuth)
	if err != nil {
		return nil, err
	}
	return &kong.HMACAuth{HMACAuth: *updatedHMACAuth}, nil
}
