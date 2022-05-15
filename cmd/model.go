package root

type User struct {
	UUID string `json:UUID`
	Name string `json:"name"`
	Age int `json:"age"`
	Version int `json:"version"`
}
