package tomlx

type GenericDocument struct {
	Value GenericValue
}

func (doc *GenericDocument) Parse(toml string) error {
	return nil
}
