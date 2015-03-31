package gokore

import "appengine/user"

type CommonArguments struct  {
	User *user.User
}

type ListArguments struct {
	CommonArguments
	Offset    int
	Limit     int
	Sort      string
	Direction int
}

type CreateArguments struct  {
	CommonArguments
	Object interface{}
}

type ReadArguments struct  {
	CommonArguments
	Id string
}

type UpdateArguments struct  {
	CommonArguments
	Id string
	Object interface{}
}

type DeleteArguments struct  {
	CommonArguments
	Id string
}

func NewListArguments(user *user.User, offset, limit int, sort string, direction int) *ListArguments {
	return &ListArguments{
		CommonArguments: CommonArguments{User: user},
		Offset:    offset,
		Limit:     limit,
		Sort:      sort,
		Direction: direction,
	}
}

func NewCreateArguments(user *user.User, object interface{}) *CreateArguments {
	return &CreateArguments{
		CommonArguments: CommonArguments{
			User: user,
		},
		Object: object,
	}
}

func NewReadArguments(user *user.User, id string) *ReadArguments {
	return &ReadArguments{
		CommonArguments: CommonArguments{
			User: user,
		},
		Id: id,
	}
}

func NewUpdateArguments(user *user.User, id string, object interface{}) *UpdateArguments {
	return &UpdateArguments{
		CommonArguments: CommonArguments{
			User: user,
		},
		Id: id,
		Object: object,
	}
}

func NewDeleteArguments(user *user.User, id string) *DeleteArguments {
	return &DeleteArguments{
		CommonArguments: CommonArguments{
			User: user,
		},
		Id: id,
	}
}
