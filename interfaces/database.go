package interfaces

type CommandDatabase interface {
	Connect(connectionString string) error
	Close() error
	Create(collection string, data interface{}) error
	Update(collection string, filter interface{}, update interface{}) error
	Delete(collection string, filter interface{}) error
}

type QueryDatabase interface {
	Connect(connectionString string) error
	Close() error
	Read(collection string, filter interface{}) (interface{}, error)
}
