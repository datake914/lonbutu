package common

type WordStatus int
type LanguageType int

const (
	WordStatusUnapproved WordStatus = iota
	WordStatusApproved
)

const (
	LanguageTypeCommon LanguageType = iota
	LanguageTypeJava
	LanguageTypeRuby
	LanguageTypeGolang
	LanguageTypePython
	LanguageTypeSQL
)
