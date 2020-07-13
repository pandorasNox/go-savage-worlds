package rulebook

import (
	"fmt"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)


func aggregate(cas CharacterAggregationState, s Sheet, rb Rulebook) (CharacterAggregationModifiers, error) {
	var err error
	var fn CharacterAggregationModifier
	var modifiers CharacterAggregationModifiers

	attributePointsUsed, err := aggregateAttributePointsUsed(s, rb.Traits().Attributes)
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}

	skillPointsUsed, err := aggregateSkillPointsUsed(s, rb.Traits().Skills)
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}

	hindrancePointsEarned, err := aggregateHindrancePointsEarned(s, rb.Hindrances())
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}

	fn = func(currentState CharacterAggregation) CharacterAggregation {
		currentState.AttributePointsUsed = attributePointsUsed
		currentState.SkillPointsUsed = skillPointsUsed
		currentState.HindrancePointsEarned = hindrancePointsEarned
		return currentState
	}

	modifiers, err = s.collectModifier(rb)
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}

	modifiers = append(modifiers, fn)

	return modifiers, nil
}

func aggregateAttributePointsUsed(s Sheet, attributes Attributes) (pointsUsed int, err error) {
	attributePointsUsed := 0
	for _, attribute := range s.Character.Traits.Attributes {
		_, ok := attributes.FindAttribute(attribute.Name)
		if ok == false {
			return 0, fmt.Errorf("\"%s\" is no valid attribute", attribute.Name)
		}

		dice, err := dice.Parse(attribute.Dice)
		if err != nil {
			return 0, fmt.Errorf(
				"parsing dice for attribute \"%s\" failed: %s",
				attribute.Name, err,
			)
		}

		attributePointsUsed += dice.Points()
	}

	return attributePointsUsed, nil
}

func aggregateSkillPointsUsed(s Sheet, skills Skills) (pointsUsed int, err error) {
	skillPointsUsed := 0

	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, found := skills.FindSkill(sheetSkill.Name)
			if found == false {
				return 0, fmt.Errorf(
					"Skill \"%s\" does not exist",
					sheetSkill.Name,
				)
			}
			skill := skills[index]

			dice, err := dice.Parse(sheetSkill.Dice)
			if err != nil {
				return 0, fmt.Errorf(
					"parsing dice for skill \"%s\" failed: %s",
					sheetSkill.Name, err,
				)
			}

			pointCostModifier := 1
			if skill.IsCore {
				pointCostModifier = 0
			}

			skillPointsUsed += dice.Points() + pointCostModifier
		}
	}

	return skillPointsUsed, nil
}

func aggregateHindrancePointsEarned(s Sheet, hs Hindrances) (pointsEarned int, err error) {
	hindrancePoints := 0

	for _, sheetHindrance := range s.Character.Hindrances {
		index, found := hs.FindHindrance(sheetHindrance.Name)

		if !found {
			return 0, fmt.Errorf("\"%s\" is no valid hindrance", sheetHindrance.Name)
		}

		if _, found := hs[index].FindDegree(sheetHindrance.Degree); !found {
			return 0, fmt.Errorf(
				"\"%s\" is no valid degree of \"%s\"",
				sheetHindrance.Degree,
				SwadeHindrances[index].Name,
			)
		}

		if sheetHindrance.Degree == Minor.String() {
			hindrancePoints++
		}

		if sheetHindrance.Degree == Major.String() {
			hindrancePoints += 2
		}
	}

	return hindrancePoints, nil
}
