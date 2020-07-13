package solver

import (
	"apigw/kong"
	"apigw/kong/crud"
	gokong "github.com/hbagdi/go-kong/kong"
)

// pluginCRUD implements crud.Actions interface.
type pluginCRUD struct {
	client *gokong.Client
}

func pluginFromStuct(arg crud.Event) *kong.Plugin {
	plugin, ok := arg.Obj.(*kong.Plugin)
	if !ok {
		panic("unexpected type, expected *kong.Plugin")
	}

	return plugin
}

// Create creates a Plugin in Kong.
// The arg should be of type crud.Event, containing the plugin to be created,
// else the function will panic.
// It returns a the created *kong.Plugin.
func (s *pluginCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	plugin := pluginFromStuct(event)

	createdPlugin, err := s.client.Plugins.Create(nil, &plugin.Plugin)
	if err != nil {
		return nil, err
	}
	return &kong.Plugin{Plugin: *createdPlugin}, nil
}

// Delete deletes a Plugin in Kong.
// The arg should be of type crud.Event, containing the plugin to be deleted,
// else the function will panic.
// It returns a the deleted *kong.Plugin.
func (s *pluginCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	plugin := pluginFromStuct(event)
	err := s.client.Plugins.Delete(nil, plugin.ID)
	if err != nil {
		return nil, err
	}
	return plugin, nil
}

// Update updates a Plugin in Kong.
// The arg should be of type crud.Event, containing the plugin to be updated,
// else the function will panic.
// It returns a the updated *kong.Plugin.
func (s *pluginCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	plugin := pluginFromStuct(event)

	updatedPlugin, err := s.client.Plugins.Create(nil, &plugin.Plugin)
	if err != nil {
		return nil, err
	}
	return &kong.Plugin{Plugin: *updatedPlugin}, nil
}
