package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	gokong "github.com/hbagdi/go-kong/kong"
)

// routeCRUD implements crud.Actions interface.
type routeCRUD struct {
	client *gokong.Client
}

func routeFromStuct(arg crud.Event) *kong.Route {
	route, ok := arg.Obj.(*kong.Route)
	if !ok {
		panic("unexpected type, expected *kong.Route")
	}

	return route
}

// Create creates a Route in Kong.
// The arg should be of type crud.Event, containing the route to be created,
// else the function will panic.
// It returns a the created *kong.Route.
func (s *routeCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	route := routeFromStuct(event)
	createdRoute, err := s.client.Routes.Create(nil, &route.Route)
	if err != nil {
		return nil, err
	}
	return &kong.Route{Route: *createdRoute}, nil
}

// Delete deletes a Route in Kong.
// The arg should be of type crud.Event, containing the route to be deleted,
// else the function will panic.
// It returns a the deleted *kong.Route.
func (s *routeCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	route := routeFromStuct(event)
	err := s.client.Routes.Delete(nil, route.ID)
	if err != nil {
		return nil, err
	}
	return route, nil
}

// Update updates a Route in Kong.
// The arg should be of type crud.Event, containing the route to be updated,
// else the function will panic.
// It returns a the updated *kong.Route.
func (s *routeCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	route := routeFromStuct(event)

	updatedRoute, err := s.client.Routes.Create(nil, &route.Route)
	if err != nil {
		return nil, err
	}
	return &kong.Route{Route: *updatedRoute}, nil
}
