package osslinked

import "time"

//go:generate protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:. osslinked.proto

// Event .
type Event struct {
	Message string `json:"string"`
	Error   error  `json:"error"`
}

// Events .
type Events <-chan *Event

// Registry .
type Registry interface {
	Create(project *Project) (Events, error)
	Update(project *Project) (Events, error)
}

// Passport .
type Passport struct {
	ID          string    `json:"id" xorm:"pk"`
	Description string    `json:"description"`
	Created     time.Time `json:"created" xorm:"created"`
}

// Passports .
type Passports interface {
	Create(description string, password string) (id string, mnemonic string, err error)
	PrivateKey(id string, password string) ([]byte, error)
	List() ([]*Passport, error)
	Delete(id string, password string) error
}
