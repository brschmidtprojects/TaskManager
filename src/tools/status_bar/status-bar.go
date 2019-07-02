package status_bar

// statusBar struct to track and create a status bar for monitoring a task's status
type StatusBar struct {
	id      string   // uuid of specific status bar instance
	taskId  string   // uuid of the specific task being processed
	size    int64    // total size of task to be completed (to be normalized to a %)
	current int64    // current progress x out of total bytes
	bar     []string // visual representation of progress to be updated as the task is progressed
	state   string   // current state of the statusBar ('ready', 'processing', 'done')
	done    bool     // flag to determine whether or not a task has finished
}
