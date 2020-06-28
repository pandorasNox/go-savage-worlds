package rulebook

import (
	"log"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

type Modifier struct {
	kind     ModifierKind
	value    int
	selector Selector
}

type Modifiers []Modifier

type ModifierKind int

const (
	ModifierKindDiceValue ModifierKind = iota
	ModifierKindDiceAdjustment
)

func (mk ModifierKind) String() string {
	return [...]string{"diceValue", "diceAdjustment"}[mk]
}

type Selector struct {
	kind   SelectorKind
	target string
}

type SelectorKind int

const (
	SelectorKindAttribute SelectorKind = iota
	SelectorKindSkill
)

func (sk SelectorKind) String() string {
	return [...]string{"attribute", "skill"}[sk]
}

func addHindranceModBuilder(hindranceName HindranceName, wantedDegree Degree, ca CharacterAggregation) CharacterAggregation {
	hIndex, hFound := SwadeHindrances.FindHindrance(string(hindranceName))
	if hFound == false {
		log.Fatalf("couldn't find %s hindrance in application data for modifierAddIgnoredPacifistHindrance function", hindranceName)
	}

	foundHindrance := SwadeHindrances[hIndex]

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

func plusToughnessAdjustmentMod(ca CharacterAggregation) CharacterAggregation {
	ca.ToughnessAdjustment++

	return ca
}

func minusToughnessAdjustmentMod(ca CharacterAggregation) CharacterAggregation {
	ca.ToughnessAdjustment--

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
