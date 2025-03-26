package errortool

var (
	Codes = DefaultGroup()
)

func DefaultGroup() *group {
	return &group{
		define: &define{
			separated: "-",
			codes:     newCodeRepository(),
		},
		code:     "",
		groupSeq: newSequence(0, maxGroupCodeSequence),
		codeSeq:  newSequence(0, maxErrorCodeSequence),
	}
}
