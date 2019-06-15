package resourcemapper

// Mapper is an interface which maps given string to certain name
type Mapper interface {
	Map(interface{}) (string, error)
}
