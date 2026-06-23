package tuiapp

type toastTickMsg struct{}
type systemThemeTickMsg struct{}
type deleteExpireMsg struct{}
type saveDoneMsg struct{}
type saveErrMsg struct{ err error }
