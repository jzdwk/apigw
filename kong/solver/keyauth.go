package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	"apigw/util"
	gokong "github.com/hbagdi/go-kong/kong"
)

// keyAuthCRUD implements crud.Actions interface.
type keyAuthCRUD struct {
	client *gokong.Client
}

func keyAuthFromStuct(arg crud.Event) *kong.KeyAuth {
	keyAuth, ok := arg.Obj.(*kong.KeyAuth)
	if !ok {
		panic("unexpected type, expected *kong.KeyAuth")
	}

	return keyAuth
}

// Create creates a Route in Kong.
// The arg should be of type crud.Event, containing the keyAuth to be created,
// else the function will panic.
// It returns a the created *kong.KeyAuth.
func (s *keyAuthCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	keyAuth := keyAuthFromStuct(event)
	createdKeyAuth, err := s.client.KeyAuths.Create(nil, keyAuth.Consumer.ID,
		&keyAuth.KeyAuth)
	if err != nil {
		return nil, err
	}
	return &kong.KeyAuth{KeyAuth: *createdKeyAuth}, nil
}

// Delete deletes a Route in Kong.
// The arg should be of type crud.Event, containing the keyAuth to be deleted,
// else the function will panic.
// It returns a the deleted *kong.KeyAuth.
func (s *keyAuthCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	keyAuth := keyAuthFromStuct(event)
	cid := ""
	if !util.Empty(keyAuth.Consumer.Username) {
		cid = *keyAuth.Consumer.Username
	}
	if !util.Empty(keyAuth.Consumer.ID) {
		cid = *keyAuth.Consumer.ID
	}
	err := s.client.KeyAuths.Delete(nil, &cid, keyAuth.ID)
	if err != nil {
		return nil, err
	}
	return keyAuth, nil
}

// Update updates a Route in Kong.
// The arg should be of type crud.Event, containing the keyAuth to be updated,
// else the function will panic.
// It returns a the updated *kong.KeyAuth.
func (s *keyAuthCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	keyAuth := keyAuthFromStuct(event)

	updatedKeyAuth, err := s.client.KeyAuths.Create(nil, keyAuth.Consumer.ID,
		&keyAuth.KeyAuth)
	if err != nil {
		return nil, err
	}
	return &kong.KeyAuth{KeyAuth: *updatedKeyAuth}, nil
}
