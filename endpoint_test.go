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

func (this *TestEndpoint) List(args *ListArguments) ([]interface{}, GokoreError) {
	return nil, nil
}

func (this *TestEndpoint) Create(object interface{}) (interface{}, GokoreError) {
	return nil, nil
}

func (this *TestEndpoint) Read(id string) (interface{}, GokoreError) {
	return nil, nil
}

func (this *TestEndpoint) Update(id string, object interface{}) (interface{}, GokoreError) {
	return nil, nil
}

func (this *TestEndpoint) Delete(id string) GokoreError {
	return nil
}

// This is just to ensure TestEndpoint is a valid EndpointInterface.
func init() {
	// this line will not compile, if TestEndpoint does not fully
	// implement EndpointInterface.
	func(EndpointInterface) {}(new(TestEndpoint))
}
