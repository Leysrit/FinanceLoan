package payload

type PaginateInfo struct {
	NextPage   string `json:"nextPage"`
	PrevPage   string `json:"prevPage"`
	TotalPages int    `json:"totalPages"`
}
