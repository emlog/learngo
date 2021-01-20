
type CreateTagRequest struct {
	Name     string `form:name`
	CreateBy string `form:created_by`
	State    uint8  `form:state`
}