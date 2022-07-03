package internal

type EnvMap map[string]string

type RunningContainer struct {
	Name  string
	AppID string
}
