package conf

type Config struct {
	DisableCpu   bool
	DisableMem   bool
	CpuThreshold float64
	MemThreshold float64

	ChatId int64
	App    string
	Token  string
}
