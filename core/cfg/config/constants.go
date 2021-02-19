package config

// ParameterType configuration parameter type.
type ParameterType string

const (
	// ParameterTypeString config parameter with "string" type.
	ParameterTypeString ParameterType = "string"
	// ParameterTypeBool config parameter with "boolean" type.
	ParameterTypeBool ParameterType = "bool"
	// ParameterTypeLogger config parameter with "logger" type.
	ParameterTypeLogger ParameterType = "logger"
)
