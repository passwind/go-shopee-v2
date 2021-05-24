package goshopee

// Error {"msg": "package_width should bigger than 1", "request_id": "2894fe4fc158a114ea4bfbbd391820c4", "error": "error_param"}
type Error struct {
	RequestID string `json:"request_id"`
	Error     string `json:"error"`
	Message   string `json:"message"`
}
