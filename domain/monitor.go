package domain


type Status int

const (
	RUNNING Status = iota + 1
	STOPPED
)

func (s Status) String() string {
	return [...]string{"RUNNING", "STOPPED"}[s-1]
}

type Monitor struct {
	Message string `json:"message"`
}





