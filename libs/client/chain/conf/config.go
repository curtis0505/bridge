package conf

type Config map[string]ClientConfig

type Chains []string

func (c Chains) GetClientChains() []string { return c }

type ClientConfig struct {
	Url                 string `toml:"url"`
	Chain               string `toml:"chain"`
	FinalizedBlockCount int    `toml:"finalizedBlockCount"`
}
