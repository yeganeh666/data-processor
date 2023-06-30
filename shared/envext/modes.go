package envext

type Mode string

const (
	ModeLocal       Mode = "LOCAL"
	ModeDevelopment Mode = "DEVELOPMENT"
	ModeStaging     Mode = "STAGING"
	ModeProduction  Mode = "PRODUCTION"
)
