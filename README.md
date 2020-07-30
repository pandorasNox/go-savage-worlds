# go-savage-worlds
go implementation of savage world rules

## nice to know
- run single test: `go test -v -run TestHindrances_FindHindrance ./...`
  - where `TestHindrances_FindHindrance` is the function name

## first milestone
- cli tool
- pass savage-world yaml sheet via cli pipe into it
- returns valid for a valid sheet / character
- returns invald and all the related "errors/mistakes" if the sheet is invalid

## todo's
- complete data.go / swadeHindrances swadeEdges
- split data.go
- add testing
  - add integration test for data.go modifiers (no fatal shall happen)
- hindrancePointsUsed
	- calc hindrancePointsUsed for start wealth
	- accumulate for what hindrancePointsUsed were added and put this into a map into character_aggregation
  - enhance errors OR stop counting hindrancePointsUsed after we reach earned/limit
- write integration test for sheet validation (cat ....)
- (maybe, bec how about testing???) make public fields from rulebook (e.g. Hindrance.Name etc) private again and provide getters
- add dice raise validation to d12 + 1. (e.g. racial ability keen senses, start notice on d6)
- DerivedStatistics
- gear + gear effects
  - add ArmorRequiredStrenghtPointsCorrection validator
- ??? gold validation + inventory + gear IF init validation ???
- use opaque types (example creating hindrances)
- how to validate "minimum two novice edges" (eg. race + house rule)?

## questions
- maybe move minimumAttributePointsRequiredForValidator &minimumSkillPointsRequiredForValidator into modifier (out of sheet_validator)

## out of scope v1.0
- build a frontend app with go and the wasm compiler or go and gopherjs/vecty
- racial abilities
  - android
    - hindrance: OUTSIDER (Major): Androids subtract 2 from Persuasion rolls when interacting with anyone besides other androids, and have no legal rights in most areas (they’re generally considered property).
  - Rakashans
    - BITE/CLAWS
    - CAN’T SWIM
    - RACIAL ENEMY
  - Saurians
    - ENVIRONMENTAL WEAKNESS
    - OUTSIDER (Minor)
- subskills (roadmap v1.1)
  - in racial ability
    - pace modifiers
      - e.g. agurian, avion, half folk ...
    - CAN’T SWIM (-2 to swim)
  - edge adjustments
  - hindrance adjustment
- pace
  - character enrichment with pace info about diffrent skills (roadmap v2)
