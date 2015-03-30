package gokore

// Defines a common Endpoint interface for gokore.
type EndpointInterface interface {
	// Returns title of the endpoint (used to construct url etc.)
	// Examples: "articles", "users", "categories"
	Title() string

	// Returns Singular form of the Title (used in some json output)
	// Examples: "article", "user", "category"
	SingularTitle() string

	// Returns a list of objects.
	// Example "[{..}, {..}]"
	List(args *ListArguments) ([]interface{}, GokoreError)

	// Returns HTTP 201 status code and
	// created resource.
	Create(args *CreateArguments) (interface{}, GokoreError)

	// Returns a resource.
	Read(args *ReadArguments) (interface{}, GokoreError)

	// Updates a resource and returns it.
	Update(args *UpdateArguments) (interface{}, GokoreError)

	// Deletes a resource.
	Delete(args *DeleteArguments) GokoreError
}

// Endpoints List
type Endpoints []EndpointInterface

// Helper used when sorting Endpoints
type EndpointsByTitle Endpoints

func (this EndpointsByTitle) Len() int           { return len(this) }
func (this EndpointsByTitle) Swap(a, b int)      { this[a], this[b] = this[b], this[a] }
func (this EndpointsByTitle) Less(a, b int) bool { return this[a].Title() < this[b].Title() }
