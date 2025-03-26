package enum

//go:generate enumer -type=JournalManagerStatus -trimprefix=JournalManagerStatus -yaml -json -text -transform=snake --output=zzz_enumer_journalManagerStatus.go
type JournalManagerStatus int32

const (
	JournalManagerStatusFailed JournalManagerStatus = iota
	JournalManagerStatusSuccess
)

//go:generate enumer -type=JournalManagerAction -trimprefix=JournalManagerAction -yaml -json -text -sql -transform=snake --output=zzz_enumer_journalManagerAction.go
type JournalManagerAction int32

const (
	JournalManagerActionUnknown JournalManagerAction = iota
)

//go:generate enumer -type=JournalManagerProto -trimprefix=JournalManagerProto -yaml -json -text -transform=snake --output=zzz_enumer_journalManagerProto.go
type JournalManagerProto int32

const (
	JournalManagerProtoUnknown JournalManagerProto = iota
	JournalManagerProtoHttp                        // http
	JournalManagerProtoSftp                        // sftp
)
