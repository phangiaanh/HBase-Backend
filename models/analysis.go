package models

type CategoryResponse struct {
	Labels []string `json:"label"`
	Data   []int64  `json:"data"`
}

func GetAllCategory() *CategoryResponse {
	var res *CategoryResponse

	res = &CategoryResponse{
		Labels: []string{
			"A", "N", "H",
		},
		Data: []int64{
			1, 2, 3,
		},
	}

	return res
}
