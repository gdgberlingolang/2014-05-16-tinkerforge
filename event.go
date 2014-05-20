// Event for the bricker.
//
// Save the packet, error (if occured) and a timestamp.
type Event struct {
	Err         error
	TimeStamp   time.Time
	Packet      *ipconn.Packet
	BrickerName string
}
