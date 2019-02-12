package param

/*------------------------------------ 一层 -------------------------------------*/

type IDParam struct {
	ID string `json:"id" query:"id" validate:"required"`
}

type URLParam struct {
	URL string `json:"url" query:"url" validate:"required"`
}

type GroupIDParam struct {
	GroupID string `json:"group_id" query:"group_id" validate:"required"`
}
