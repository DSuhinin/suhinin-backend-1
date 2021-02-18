package config

// ParameterType configuration parameter type.
type ParameterType string

const (
	// ParameterTypeString config parameter with "string" type.
	ParameterTypeString ParameterType = "string"
	// ParameterTypeLogger config parameter with "logger" type.
	ParameterTypeLogger ParameterType = "logger"
)
