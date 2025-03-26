package interfaces

type CronTask interface {
	Schedule() string
	Run()
	Name() string
}
