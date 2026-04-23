package batch

type Result struct {
	TaskID int
	Value  any
	Error  error
}

func NewResult(taskId int, value any, err error) Result {
	return Result{
		TaskID: taskId,
		Value:  value,
		Error:  err,
	}
}
