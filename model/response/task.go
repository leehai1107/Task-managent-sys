package response

type TaskResponse struct {
	TaskID uint   `json:"task_id"`
	Title  string `json:"title"`
}
