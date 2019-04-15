package task

type ShowTasksInput struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type TaskOutput struct {
	Name     string `json:"name"`
	Uid      string `json:"uid"`
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	Status   string `json:"status"`
	TaskName string `json:"task-name"`
}

type ShowTasksOutput struct {
	From    int           `json:"from"`
	Objects *[]TaskOutput `json:"objects"`
	To      int           `json:"to"`
	Total   int           `json:"total"`
}

func (c *TaskService) ShowTasks() (*ShowTasksOutput, error) {
	input := &ShowTasksInput{
		Limit:  500,
		Offset: 0,
	}
	out := &ShowTasksOutput{}
	req := c.NewPostRequest("show-tasks", input)

	return out, c.Client.Send(req, out)
}
