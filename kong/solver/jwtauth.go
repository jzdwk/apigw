package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	"apigw/util"
	gokong "github.com/hbagdi/go-kong/kong"
)

// jwtAuthCRUD implements crud.Actions interface.
type jwtAuthCRUD struct {
	client *gokong.Client
}

func jwtAuthFromStuct(arg crud.Event) *kong.JWTAuth {
	jwtAuth, ok := arg.Obj.(*kong.JWTAuth)
	if !ok {
		panic("unexpected type, expected *kong.JWTAuth")
	}

	return jwtAuth
}

// Create creates a Route in Kong.
// The arg should be of type crud.Event, containing the jwtAuth to be created,
// else the function will panic.
// It returns a the created *kong.JWTAuth.
func (s *jwtAuthCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	jwtAuth := jwtAuthFromStuct(event)
	cid := ""
	if !util.Empty(jwtAuth.Consumer.Username) {
		cid = *jwtAuth.Consumer.Username
	}
	if !util.Empty(jwtAuth.Consumer.ID) {
		cid = *jwtAuth.Consumer.ID
	}
	createdJWTAuth, err := s.client.JWTAuths.Create(nil, &cid,
		&jwtAuth.JWTAuth)
	if err != nil {
		return nil, err
	}
	return &kong.JWTAuth{JWTAuth: *createdJWTAuth}, nil
}

// Delete deletes a Route in Kong.
// The arg should be of type crud.Event, containing the jwtAuth to be deleted,
// else the function will panic.
// It returns a the deleted *kong.JWTAuth.
func (s *jwtAuthCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	jwtAuth := jwtAuthFromStuct(event)
	cid := ""
	if !util.Empty(jwtAuth.Consumer.Username) {
		cid = *jwtAuth.Consumer.Username
	}
	if !util.Empty(jwtAuth.Consumer.ID) {
		cid = *jwtAuth.Consumer.ID
	}
	err := s.client.JWTAuths.Delete(nil, &cid, jwtAuth.ID)
	if err != nil {
		return nil, err
	}
	return jwtAuth, nil
}

// Update updates a Route in Kong.
// The arg should be of type crud.Event, containing the jwtAuth to be updated,
// else the function will panic.
// It returns a the updated *kong.JWTAuth.
func (s *jwtAuthCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	jwtAuth := jwtAuthFromStuct(event)

	cid := ""
	if !util.Empty(jwtAuth.Consumer.Username) {
		cid = *jwtAuth.Consumer.Username
	}
	if !util.Empty(jwtAuth.Consumer.ID) {
		cid = *jwtAuth.Consumer.ID
	}
	updatedJWTAuth, err := s.client.JWTAuths.Create(nil, &cid, &jwtAuth.JWTAuth)
	if err != nil {
		return nil, err
	}
	return &kong.JWTAuth{JWTAuth: *updatedJWTAuth}, nil
}
