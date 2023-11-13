package validation

type ValidationError struct {
	Path    interface{}
	Message string
}
type ValidationErrors []ValidationError
