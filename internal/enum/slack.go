package enum

//go:generate enumer -type=SlackNotifierType -trimprefix=SlackNotifierType -yaml -json -text -transform=snake --output=zzz_enumer_slackNotifierType.go
type SlackNotifierType int32

const (
	SlackNotifierTypeUnknown SlackNotifierType = iota
	SlackNotifierTypeRiskControl
	SlackNotifierTypeFinance
	SlackNotifierTypeSystem
)

//go:generate enumer -type=SlackRiskLevel -trimprefix=SlackRiskLevel -yaml -json -text -transform=snake --output=zzz_enumer_slackRiskLevel.go
type SlackRiskLevel int32

const (
	// SlackRiskLevelUnknown represents unknown risk level
	SlackRiskLevelUnknown SlackRiskLevel = iota
	// SlackRiskLevelLow represents low risk level which is not critical. hint: if it happens too often, it may become critical
	SlackRiskLevelLow
	// SlackRiskLevelMedium represents medium risk level which is critical but not urgent
	SlackRiskLevelMedium
	// SlackRiskLevelHigh represents high risk level which is critical and urgent
	SlackRiskLevelHigh
)
