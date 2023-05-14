package badge

const (
	Behavior = "BEHAVIOR"
	Marvel   = "MARVEL"
	Romance  = "ROMANCE"
)

func GetCategory() []string {
	return []string{Behavior, Marvel, Romance}
}
