package domain

// OctoBotSortField is a type that represents the sort fields of a octoBot.
type OctoBotSortField int8

const (
	OctoBotSortFieldNONE OctoBotSortField = iota
	OctoBotSortFieldId
	//OctoBotSortFieldName
	OctoBotSortFieldCreatedAt
	OctoBotSortFieldUpdatedAt
)
