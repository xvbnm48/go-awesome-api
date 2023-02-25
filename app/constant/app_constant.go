package constant

type ResponseStatus int
type Header int
type General int

// constant API
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UknownError
	InvalidRequest
	unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data not found", "Unknown error", "Invalid request", "Unauthorized"}[r-1]
}
