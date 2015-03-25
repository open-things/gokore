package gokore

// Dummy endpoint to be used in tests
type TestEndpoint struct {
	title  string
	stitle string
}

// TestEndpoint constructor
func NewTestEndpoint(title, stitle string) *TestEndpoint {
	this := TestEndpoint{
		title:  title,
		stitle: stitle,
	}
	return &this
}

// EndpointInterface methods
func (this *TestEndpoint) Title() string         { return this.title }
func (this *TestEndpoint) SingularTitle() string { return this.stitle }

// This is just to ensure TestEndpoint is a valid EndpointInterface.
func init() {
	// this line will not compile, if TestEndpoint does not fully
	// implement EndpointInterface.
	func(EndpointInterface) {}(new(TestEndpoint))
}
