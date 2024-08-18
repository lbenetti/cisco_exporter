package environment

type EnvironmentItem struct {
	Name        string
	Status      string
	OK          bool
	IsTemp      bool `default:"false"`
	IsFan       bool `default:"false"`
	IsPower     bool `default:"false"`
	Temperature float64
	Fan         float64
	Power       float64
}
