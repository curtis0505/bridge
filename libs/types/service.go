package types

type Service string

const (
	NeopinService  = "NEOPIN" // Legacy
	StablezService = "STABLEZ"
)

func (s Service) String() string {
	return string(s)
}
