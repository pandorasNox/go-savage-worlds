# go-savage-worlds
go implementation of savage world rules

## first milestone
- cli tool
- pass savage-world yaml sheet via cli pipe into it
- returns valid for a valid sheet / character
- returns invald and all the related "errors/mistakes" if the sheet is invalid

## todo's
- add testing
  - add integration test for data.go modifiers (no fatal shall happen)
- write integration test for sheet validation (cat ....)
- (maybe, bec how about testing???) make public fields from rulebook (e.g. Hindrance.Name etc) private again and provide getters
- cleanup deprecated modifier (was replaced by charachterAggregationModifier)
- races + racial effects
- add dice raise validation to d12 + 1. (e.g. racial ability keen senses, start notice on d6)
- hindrance effects
- edges + edges effects + edged requirements
- DerivedStatistics
- gear + gear effects
- ??? gold validation + inventory + gear IF init validation ???
- build a frontend app with go and the wasm compiler or go and gopherjs/vecty
- use opaque types (example creating hindrances)
- check to have hindrance twice?
- how to validate "minimum two novice edges" (eg. race + house rule)?

## out of scope v1.0
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
