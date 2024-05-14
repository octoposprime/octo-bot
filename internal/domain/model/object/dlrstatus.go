package domain

// OctoBotStatus is a status that represents the status of a octoBot.
type OctoBotStatus int8

const (
	OctoBotStatusNONE OctoBotStatus = iota
	OctoBotStatusACTIVE
	OctoBotStatusINACTIVE
)
