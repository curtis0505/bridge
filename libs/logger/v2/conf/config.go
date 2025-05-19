package conf

type Config struct {
	UseTerminal        bool `mapstructure:"use_terminal"`
	TerminalJSONOutput bool `mapstructure:"terminal_json_output"`
	VerbosityTerminal  int  `mapstructure:"verbosity_terminal"`
}
