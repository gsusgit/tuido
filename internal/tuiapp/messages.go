package tuiapp

type toastTickMsg struct{}
type deleteExpireMsg struct{}
type saveDoneMsg struct{}
type saveErrMsg struct{ err error }
