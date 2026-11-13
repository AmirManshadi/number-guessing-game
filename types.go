package main

type Difficulty uint8

const (
	DifficultyEasy   Difficulty = 1
	DifficultyMedium Difficulty = 2
	DifficultyHard   Difficulty = 3
)

func (d Difficulty) IsValid() bool {
	switch d {
	case DifficultyEasy, DifficultyMedium, DifficultyHard:
		return true
	default:
		return false
	}
}

func (d Difficulty) String() string {
	switch d {
	case DifficultyEasy:
		return "Easy"
	case DifficultyMedium:
		return "Medium"
	case DifficultyHard:
		return "Hard"
	default:
		return "Unknown Difficulty"
	}
}

func (d Difficulty) Limit() uint8 {
	switch d {
	case DifficultyEasy:
		return 10
	case DifficultyMedium:
		return 5
	case DifficultyHard:
		return 3
	default:
		return 5
	}
}

// -------------------------

type Indicator string

const (
	IndicatorLess    Indicator = "less"
	IndicatorGreater Indicator = "greater"
)
