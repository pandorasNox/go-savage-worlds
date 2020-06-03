package savage

import (
	"fmt"
	"regexp"
	"strconv"
)

type Dice struct {
	value        int
	accumulation int
}

var diceValueToPointsUsedMap = map[string]int{
	"4":  0,
	"6":  1,
	"8":  2,
	"10": 3,
	"12": 4,
}

func ParseDice(dice string) (Dice, error) {
	var re = regexp.MustCompile(`^d(4|6|8|10|12)(\+([1-9][0-9]?))?$`)

	found := re.FindAllStringSubmatch(dice, -1)

	if found == nil || (len(found[0]) != 2 && len(found[0]) != 4) {
		return Dice{}, fmt.Errorf(
			"validation error: invalid dice value \"%s\"",
			dice,
		)
	}

	accumulation, err := strconv.Atoi(found[0][1])
	if err != nil {
		return Dice{}, err
	}

	return Dice{value: diceValueToPointsUsedMap[found[0][1]], accumulation: accumulation}, nil
}
