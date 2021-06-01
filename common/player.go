package common

//player.go
// common player type structure used by updater service and updater tool

type Application struct {
	ApplicationID string `json:"applicationId"`
	Version       string `json:"version"`
}
type Profile struct {
	Applications []Application `json:"applications"`
}

type Player struct {
	Profile Profile `json:"profile"`
}

type Players []Player
