package rulebook

import (
	"fmt"
	"log"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

func addRequiredHindranceModBuilder(hindranceName HindranceName, wantedDegree Degree, hindrances Hindrances, ca CharacterAggregation) CharacterAggregation {
	hIndex, hFound := hindrances.FindHindrance(string(hindranceName))
	if hFound == false {
		log.Fatalf("couldn't find %s hindrance in application data for modifierAddIgnoredPacifistHindrance function", hindranceName)
	}

	foundHindrance := hindrances[hIndex]

	_, dFound := foundHindrance.FindDegree(wantedDegree.String())
	if dFound == false {
		log.Fatalf("couldn't find degree %s for %s hindrance in application data for modifierAddIgnoredPacifistHindrance function", wantedDegree.String(), hindranceName)
	}

	ca.HindrancesRequired[hindranceName] = wantedDegree
	ca.HindrancesIgnored[hindranceName] = wantedDegree

	return ca
}

func plusShakenRecoveryAdjustmentMod(ca CharacterAggregation) CharacterAggregation {
	ca.ShakenRecoveryAdjusment++

	return ca
}

func plusBaseToughnessMod(ca CharacterAggregation) CharacterAggregation {
	ca.BaseToughness++

	return ca
}

func minusBaseToughnessMod(ca CharacterAggregation) CharacterAggregation {
	ca.BaseToughness--

	return ca
}

func minusSkillPointsUsedMod(ca CharacterAggregation) CharacterAggregation {
	ca.SkillPointsUsed--

	return ca
}

func skillStartsAtModBuilder(skillName SkillName, dice dice.Dice, ca CharacterAggregation) CharacterAggregation {
	ca = minusSkillPointsUsedMod(ca)

	if pointsRequired, ok := ca.MinimumSkillPointsRequiredFor[skillName]; ok {
		if pointsRequired >= dice.Points() {
			return ca
		}
	}

	ca.MinimumSkillPointsRequiredFor[skillName] = dice.Points()

	return ca
}

func minusAttributePointsUsedMod(ca CharacterAggregation) CharacterAggregation {
	ca.AttributePointsUsed--

	return ca
}

func attributeStartsAtModBuilder(attributeName AttributeName, dice dice.Dice, ca CharacterAggregation) CharacterAggregation {
	ca = minusAttributePointsUsedMod(ca)

	if pointsRequired, ok := ca.MinimumAttributePointsRequiredFor[attributeName]; ok {
		if pointsRequired >= dice.Points() {
			return ca
		}
	}

	ca.MinimumAttributePointsRequiredFor[attributeName] = dice.Points()

	return ca
}

func skillAdjusmentModBuilder(skillName SkillName, adjustment int, skills Skills, ca CharacterAggregation) CharacterAggregation {
	_, found := skills.FindSkill(string(skillName))
	if found == false {
		log.Fatalf("skillAdjusmentModBuilder: skill \"%s\" not found.", skillName)
	}

	ca.SkillsAdjustments[skillName] += adjustment

	skillAdjustmentValidator := func(ca CharacterAggregation, s Sheet, rb Rulebook) error {
		sheetSkill, err := s.SheetSkill(skillName)
		if err != nil {
			return err
		}

		sDice, err := dice.Parse(sheetSkill.Dice)
		if err != nil {
			return fmt.Errorf("skill \"%s\" does not contain valid dice \"%s\"", skillName, sheetSkill.Dice)
		}

		if _, ok := ca.SkillsAdjustments[skillName]; ok == false {
			return fmt.Errorf("internal error during skilladjustment validation for skill \"%s\"", skillName)
		}

		if ca.SkillsAdjustments[skillName] != sDice.Adjustment() {
			return fmt.Errorf("skill \"%s\" is expected to have \"%d\" adjustment, but got \"%d\"", skillName, ca.SkillsAdjustments[skillName], sDice.Adjustment())
		}

		return nil
	}

	ca.additionalValidators = append(ca.additionalValidators, skillAdjustmentValidator)

	return ca
}

func attributeAdjusmentModBuilder(attributeName AttributeName, adjustment int, attributes Attributes, ca CharacterAggregation) CharacterAggregation {
	_, found := attributes.FindAttribute(string(attributeName))
	if found == false {
		log.Fatalf("attributeAdjusmentModBuilder: attribute \"%s\" not found.", attributeName)
	}

	ca.AttributeAdjustments[attributeName] += adjustment

	attributeAdjustmentValidator := func(ca CharacterAggregation, s Sheet, rb Rulebook) error {
		sheetAttribute, err := s.SheetAttribute(attributeName)
		if err != nil {
			return err
		}

		aDice, err := dice.Parse(sheetAttribute.Dice)
		if err != nil {
			return fmt.Errorf("attribute \"%s\" does not contain valid dice \"%s\"", attributeName, sheetAttribute.Dice)
		}

		if _, ok := ca.AttributeAdjustments[attributeName]; ok == false {
			return fmt.Errorf("internal error during attribute adjustment validation for attribute \"%s\"", attributeName)
		}

		if ca.AttributeAdjustments[attributeName] != aDice.Adjustment() {
			return fmt.Errorf("attribute \"%s\" is expected to have \"%d\" adjustment, but got \"%d\"", attributeName, ca.AttributeAdjustments[attributeName], aDice.Adjustment())
		}

		return nil
	}

	ca.additionalValidators = append(ca.additionalValidators, attributeAdjustmentValidator)

	return ca
}

func freeEdgeMod(ca CharacterAggregation) CharacterAggregation {
	ca.HindrancePointsUsed += -2
	ca.MinimumChosenEdges++

	return ca
}

func freeNoviceEdgeMod(ca CharacterAggregation) CharacterAggregation {
	ca.HindrancePointsUsed += -2
	ca.MinimumChosenEdges++

	hasNoviceEdgeValidator := func(ca CharacterAggregation, s Sheet, rb Rulebook) error {
		for _, sheetEdge := range s.Character.Edges {
			eIndex, eFound := rb.Edges().FindEdge(edgeName(sheetEdge))
			if eFound == false {
				return fmt.Errorf("unknown edge \"%s\" in sheet", sheetEdge)
			}

			edge := rb.Edges()[eIndex]
			if edge.requirement.rank == Novice {
				return nil
			}
		}

		return fmt.Errorf("hasNoviceEdgeValidator: no novice edge found")
	}
	ca.additionalValidators = append(ca.additionalValidators, hasNoviceEdgeValidator)

	return ca
}

func minusSizeMod(ca CharacterAggregation) CharacterAggregation {
	ca.Size--

	return ca
}

func plusSizeMod(ca CharacterAggregation) CharacterAggregation {
	ca.Size++

	return ca
}

func plusArmorMod(ca CharacterAggregation) CharacterAggregation {
	ca.Armor++

	return ca
}

func addRequiredEdgeModBuilder(edgeName edgeName, edges Edges, ca CharacterAggregation) CharacterAggregation {
	index, found := edges.FindEdge(edgeName)
	if found == false {
		log.Fatalf("addRequiredEdgeModBuilder: edge \"%s\" not found.", edgeName)
	}

	ca.EdgesRequired = append(ca.EdgesRequired, edges[index])

	return ca
}

func plusSkillPointsAvailableMod(ca CharacterAggregation) CharacterAggregation {
	ca.SkillPointsAvailable++

	return ca
}

func minusArmorRequiredStrenghtPointsCorrectionMod(ca CharacterAggregation) CharacterAggregation {
	ca.ArmorRequiredStrenghtPointsCorrection--

	return ca
}
