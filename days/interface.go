package days

type AdventDay interface {
	Name() string
	Run() error
	Init()
}