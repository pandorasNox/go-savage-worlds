package savage

import (
	"fmt"
	"regexp"
	"strconv"
)

type Dice struct {
	value      int
	adjustment int
}

var diceValueToPointsUsedMap = map[string]int{
	"4":  0,
	"6":  1,
	"8":  2,
	"10": 3,
	"12": 4,
}

// ParseDice parse dice strig to struct using points map
func ParseDice(dice string) (Dice, error) {
	var re = regexp.MustCompile(`^d(4|6|8|10|12)((\+|-)([1-9][0-9]?))?$`)

	found := re.FindAllStringSubmatch(dice, -1)

	if found == nil || len(found) != 1 || (len(found[0]) != 2 && len(found[0]) != 5) {
		return Dice{}, fmt.Errorf(
			"validation error: invalid dice value \"%s\"",
			dice,
		)
	}

	foundDice := found[0][1]
	foundAdjustment := ""

	if len(found[0]) == 5 {
		foundAdjustment = found[0][2]
	}

	adjustment := 0
	var err error
	if foundAdjustment != "" {
		adjustment, err = strconv.Atoi(foundAdjustment)

		if err != nil {
			return Dice{}, err
		}
	}

	return Dice{value: diceValueToPointsUsedMap[foundDice], adjustment: adjustment}, nil
}
