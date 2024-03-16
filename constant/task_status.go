package constant

type TaskStatus string

const (
	TaskStatusDone       TaskStatus = "Done!"
	TaskStatusNotDone    TaskStatus = "Not done!"
	TaskStatusOnGoing    TaskStatus = "On going!"
	TaskStatusOnPlanning TaskStatus = "On planning!"
	TaskStatusCancel     TaskStatus = "Cancel!"
)

func (t TaskStatus) Value() string { return string(t) }
