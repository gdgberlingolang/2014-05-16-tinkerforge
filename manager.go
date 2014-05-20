// The Manager controls the created bricker (connections) and dispatch the events to the
// subscriber. The subscriber consums the events once or more times.
type Manager interface {
	AddBricker(bricker.Bricker, string) error
	ReleaseBricker(string) error
	Subscribe(bricker.Subscriber, interface{}) error
	Unsubscribe(bricker.Subscriber) error
	Done()
}
