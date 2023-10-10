package model

type Progression struct {
	Level   int      `toml:"level"`
	Tags    []string `toml:"tags,omitempty"`
	Choices []string `toml:"choices,omitempty"`
}

type Progressions []*Progression

func (p Progressions) Level(i int) *Progression {
	for _, progression := range p {
		if progression.Level == i {
			return progression
		}
	}
	return nil
}
