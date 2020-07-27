package rulebook

import (
	"fmt"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

func aggregateAndUpdate(cas *CharacterAggregationState, s Sheet, rb Rulebook) error {
	var err error
	var fn CharacterAggregationModifier

	attributePointsUsed, err := aggregateAttributePointsUsed(s, rb.Traits().Attributes)
	if err != nil {
		return err
	}

	skillPointsUsed, err := aggregateSkillPointsUsed(s, rb.Traits().Skills)
	if err != nil {
		return err
	}

	hindrancePointsEarned, err := aggregateHindrancePointsEarned(s, rb.Hindrances())
	if err != nil {
		return err
	}

	fn = func(currentState CharacterAggregation) CharacterAggregation {
		currentState.AttributePointsUsed = attributePointsUsed
		currentState.SkillPointsUsed = skillPointsUsed
		currentState.HindrancePointsEarned = hindrancePointsEarned
		return currentState
	}
	cas.Update(fn)

	cModifiers, err := collectModifier(s, rb)
	if err != nil {
		return err
	}
	cas.Updates(cModifiers)

	hindrancePointsUsedModifier, err := aggregateHindrancePointsUsed(cas.CharacterAggregation(), s, rb)
	if err != nil {
		return err
	}
	cas.Update(hindrancePointsUsedModifier)

	additionalValidators, err := collectAdditionalValidator(s, rb)
	if err != nil {
		return err
	}
	cas.Update(func(currentState CharacterAggregation) CharacterAggregation {
		currentState.additionalValidators = append(
			currentState.additionalValidators,
			additionalValidators...,
		)
		return currentState
	})

	return nil
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
		aDice, err := dice.Parse(sheetAttr.Dice)
		if err != nil {
			return 0, fmt.Errorf(
				"parsing dice for attribute \"%s\" failed: %s",
				sheetAttr.Name, err,
			)
		}

		for _, sheetSkill := range sheetAttr.Skills {
			index, found := skills.FindSkill(sheetSkill.Name)
			if found == false {
				return 0, fmt.Errorf(
					"Skill \"%s\" does not exist",
					sheetSkill.Name,
				)
			}
			skill := skills[index]

			sDice, err := dice.Parse(sheetSkill.Dice)
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

			higherLevelCostModifier := 0
			if aDice.Points() < sDice.Points() {
				higherLevelCostModifier = sDice.Points() - aDice.Points()
			}

			skillPointsUsed += sDice.Points() + pointCostModifier + higherLevelCostModifier
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

func collectModifier(s Sheet, rb Rulebook) (CharacterAggregationModifiers, error) {
	var modifier CharacterAggregationModifiers
	var err error

	raceMods, err := collectRaceModifier(s, rb.Races())
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}
	modifier = append(modifier, raceMods...)

	hindranceMods, err := collectHindranceModifier(s, rb.Hindrances())
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}
	modifier = append(modifier, hindranceMods...)

	edgeMods, err := collectEdgeModifier(s, rb.Edges())
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}
	modifier = append(modifier, edgeMods...)

	return modifier, nil
}

func collectRaceModifier(s Sheet, rbRaces Races) (CharacterAggregationModifiers, error) {
	sheetRace := s.Character.Info.Race

	index, found := rbRaces.FindRace(sheetRace)
	if found == false {
		return CharacterAggregationModifiers{}, fmt.Errorf("Unknown race \"%s\" in sheet", sheetRace)
	}
	race := rbRaces[index]

	return race.Modifiers(), nil
}

func collectHindranceModifier(s Sheet, rbHinds Hindrances) (CharacterAggregationModifiers, error) {
	var modifier CharacterAggregationModifiers

	for _, sheetHindrance := range s.Character.Hindrances {
		index, foundHin := rbHinds.FindHindrance(sheetHindrance.Name)
		if foundHin == false {
			return CharacterAggregationModifiers{}, fmt.Errorf("hindrance \"%s\" doesn't exist", sheetHindrance.Name)
		}
		matchedHindrance := SwadeHindrances[index]

		index, foundDeg := matchedHindrance.FindDegree(sheetHindrance.Degree)
		if foundDeg == false {
			return CharacterAggregationModifiers{}, fmt.Errorf("hindrance \"%s\" doesn't have a \"%s\" degree", sheetHindrance.Name, sheetHindrance.Degree)
		}
		matchedDegree := matchedHindrance.AvailableDegrees[index]

		modifier = append(modifier, matchedDegree.Modifiers...)
	}

	return modifier, nil
}

func collectEdgeModifier(s Sheet, es Edges) (CharacterAggregationModifiers, error) {
	var modifier CharacterAggregationModifiers

	for _, sheetEdge := range s.Character.Edges {
		index, found := es.FindEdge(sheetEdge)
		if found == false {
			return CharacterAggregationModifiers{}, fmt.Errorf("edge \"%s\" doesn't exist", sheetEdge)
		}

		edge := es[index]

		modifier = append(modifier, edge.modifiers...)
	}

	return modifier, nil
}

func aggregateHindrancePointsUsed(ca CharacterAggregation, s Sheet, rb Rulebook) (CharacterAggregationModifier, error) {
	hindrancePointsUsed := 0
	emptyFn := func(ca CharacterAggregation) CharacterAggregation {
		return ca
	}

	for _, sheetEdge := range s.Character.Edges {
		_, eFound := rb.Edges().FindEdge(sheetEdge)
		if eFound == false {
			return emptyFn, fmt.Errorf("unknown edge \"%s\" in sheet", sheetEdge)
		}

		hindrancePointsUsed += 2
	}

	extraUsedAttributePoints := max(0, (ca.AttributePointsUsed - ca.AttributePointsAvailable))

	hindrancePointsUsed += extraUsedAttributePoints * 2

	hindrancePointsUsed += max(0, (ca.SkillPointsUsed - ca.SkillPointsAvailable))

	//todo: calc hindrancePointsUsed for start wealth

	//todo: accumulate for what hindrancePointsUsed were added and put this into a map
	//into character_aggregation

	fn := func(ca CharacterAggregation) CharacterAggregation {
		ca.AttributePointsAvailable += extraUsedAttributePoints
		ca.HindrancePointsUsed = hindrancePointsUsed
		return ca
	}

	return fn, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func collectAdditionalValidator(s Sheet, rb Rulebook) (validators, error) {
	var v validators
	var err error

	edgeValidators, err := collectEdgeValidator(s, rb.Edges())
	if err != nil {
		return validators{}, err
	}
	v = append(v, edgeValidators...)

	return v, nil
}

func collectEdgeValidator(s Sheet, es Edges) (validators, error) {
	var v validators

	for _, sheetEdge := range s.Character.Edges {
		i, found := es.FindEdge(sheetEdge)
		if found == false {
			return validators{}, fmt.Errorf("invalid edge in sheet \"%s\"", sheetEdge)
		}

		edge := es[i]

		v = append(v, edge.requirement.validators...)
	}

	return v, nil
}
