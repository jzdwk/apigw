package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	gokong "github.com/hbagdi/go-kong/kong"
)

// consumerCRUD implements crud.Actions interface.
type consumerCRUD struct {
	client *gokong.Client
}

func consumerFromStuct(arg crud.Event) *kong.Consumer {
	consumer, ok := arg.Obj.(*kong.Consumer)
	if !ok {
		panic("unexpected type, expected *kong.consumer")
	}
	return consumer
}

// Create creates a Consumer in Kong.
// The arg should be of type crud.Event, containing the consumer to be created,
// else the function will panic.
// It returns a the created *kong.Consumer.
func (s *consumerCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	consumer := consumerFromStuct(event)
	createdConsumer, err := s.client.Consumers.Create(nil, &consumer.Consumer)
	if err != nil {
		return nil, err
	}
	return &kong.Consumer{Consumer: *createdConsumer}, nil
}

// Delete deletes a Consumer in Kong.
// The arg should be of type crud.Event, containing the consumer to be deleted,
// else the function will panic.
// It returns a the deleted *kong.Consumer.
func (s *consumerCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	consumer := consumerFromStuct(event)
	err := s.client.Consumers.Delete(nil, consumer.ID)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// Update updates a Consumer in Kong.
// The arg should be of type crud.Event, containing the consumer to be updated,
// else the function will panic.
// It returns a the updated *kong.Consumer.
func (s *consumerCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	consumer := consumerFromStuct(event)

	updatedConsumer, err := s.client.Consumers.Create(nil, &consumer.Consumer)
	if err != nil {
		return nil, err
	}
	return &kong.Consumer{Consumer: *updatedConsumer}, nil
}
