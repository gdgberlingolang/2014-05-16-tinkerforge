// Interface to the connector. It should send and receive events to or from the hardware.
// The bricker works with the packets inside. The bricker must be thread safe.
type Bricker interface {
	Send(*event.Event)
	Receive() *event.Event
	Done()
}

// BrickerBuffered is the bricker with to bufferd channels for reading (In) and Writing(Out).
// Every channel will controled by a go routine. The bricker puts all of his readed packets into
// events in the In channel. He waits for packets to write out to the hardware on the Out channel.
// A close on the Quit channel let the bricker stops all go routines and disconnect to the hardware.
type BrickerBuffered struct {
	conn *ipconn.IPConn    // internal, the real connection
	seq  uint8             // internal, actual sequence number
	In   chan *event.Event // input channel, here the bricker put in the readed packets as events
	Out  chan *event.Event // output channel, here the bricker read out the events, which should be send
	Quit chan struct{}     // quit channel, if closed, the bricker stop working and release resources
}
