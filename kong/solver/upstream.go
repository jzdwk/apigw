package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	gokong "github.com/hbagdi/go-kong/kong"
)

// upstreamCRUD implements crud.Actions interface.
type upstreamCRUD struct {
	client *gokong.Client
}

func upstreamFromStuct(arg crud.Event) *kong.Upstream {
	upstream, ok := arg.Obj.(*kong.Upstream)
	if !ok {
		panic("unexpected type, expected *kong.upstream")
	}
	return upstream
}

// Create creates a Upstream in Kong.
// The arg should be of type crud.Event, containing the upstream to be created,
// else the function will panic.
// It returns a the created *kong.Upstream.
func (s *upstreamCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	upstream := upstreamFromStuct(event)
	createdUpstream, err := s.client.Upstreams.Create(nil, &upstream.Upstream)
	if err != nil {
		return nil, err
	}
	return &kong.Upstream{Upstream: *createdUpstream}, nil
}

// Delete deletes a Upstream in Kong.
// The arg should be of type crud.Event, containing the upstream to be deleted,
// else the function will panic.
// It returns a the deleted *kong.Upstream.
func (s *upstreamCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	upstream := upstreamFromStuct(event)
	err := s.client.Upstreams.Delete(nil, upstream.ID)
	if err != nil {
		return nil, err
	}
	return upstream, nil
}

// Update updates a Upstream in Kong.
// The arg should be of type crud.Event, containing the upstream to be updated,
// else the function will panic.
// It returns a the updated *kong.Upstream.
func (s *upstreamCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	upstream := upstreamFromStuct(event)

	updatedUpstream, err := s.client.Upstreams.Create(nil, &upstream.Upstream)
	if err != nil {
		return nil, err
	}
	return &kong.Upstream{Upstream: *updatedUpstream}, nil
}
