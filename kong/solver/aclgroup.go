package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	"apigw/util"
	gokong "github.com/hbagdi/go-kong/kong"
)

// aclGroupCRUD implements crud.Actions interface.
type aclGroupCRUD struct {
	client *gokong.Client
}

func aclGroupFromStuct(arg crud.Event) *kong.ACLGroup {
	aclGroup, ok := arg.Obj.(*kong.ACLGroup)
	if !ok {
		panic("unexpected type, expected *ACLGroup")
	}

	return aclGroup
}

// Create creates a Route in Kong.
// The arg should be of type diff.Event, containing the aclGroup to be created,
// else the function will panic.
// It returns a the created *kong.ACLGroup.
func (s *aclGroupCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	aclGroup := aclGroupFromStuct(event)
	cid := ""
	if !util.Empty(aclGroup.Consumer.Username) {
		cid = *aclGroup.Consumer.Username
	}
	if !util.Empty(aclGroup.Consumer.ID) {
		cid = *aclGroup.Consumer.ID
	}
	createdACLGroup, err := s.client.ACLs.Create(nil, &cid,
		&aclGroup.ACLGroup)
	if err != nil {
		return nil, err
	}
	return &kong.ACLGroup{ACLGroup: *createdACLGroup}, nil
}

// Delete deletes a Route in Kong.
// The arg should be of type diff.Event, containing the aclGroup to be deleted,
// else the function will panic.
// It returns a the deleted *kong.ACLGroup.
func (s *aclGroupCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	aclGroup := aclGroupFromStuct(event)
	cid := ""
	if !util.Empty(aclGroup.Consumer.Username) {
		cid = *aclGroup.Consumer.Username
	}
	if !util.Empty(aclGroup.Consumer.ID) {
		cid = *aclGroup.Consumer.ID
	}
	err := s.client.ACLs.Delete(nil, &cid, aclGroup.ID)
	if err != nil {
		return nil, err
	}
	return aclGroup, nil
}

// Update updates a Route in Kong.
// The arg should be of type diff.Event, containing the aclGroup to be updated,
// else the function will panic.
// It returns a the updated *kong.ACLGroup
func (s *aclGroupCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	aclGroup := aclGroupFromStuct(event)

	cid := ""
	if !util.Empty(aclGroup.Consumer.Username) {
		cid = *aclGroup.Consumer.Username
	}
	if !util.Empty(aclGroup.Consumer.ID) {
		cid = *aclGroup.Consumer.ID
	}
	updatedACLGroup, err := s.client.ACLs.Create(nil, &cid, &aclGroup.ACLGroup)
	if err != nil {
		return nil, err
	}
	return &kong.ACLGroup{ACLGroup: *updatedACLGroup}, nil
}
