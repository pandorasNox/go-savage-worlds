# go-savage-worlds
go implementation of savage world rules

## first milestone
- cli tool
- pass savage-world yaml sheet via cli pipe into it
- returns valid for a valid sheet / character
- returns invald and all the related "errors/mistakes" if the sheet is invalid

## todo's
- add testing
- write integration test for sheet validation (cat ....)
- (maybe, bec how about testing???) make public fields from rulebook (e.g. Hindrance.Name etc) private again and provide getters
- create validation struct with internal counter
    - we have to mutate it over the running process
    - redux???
    - afterwards change modifier to function which can change the internal counter
        - fixing ModifierKindEdgeSlot ModifierKindAddHindrance issues
- races + racial effects
- after we added ModifierKindEdgeSlot we also need a new kind of EdgeValidator which ensures e.g. at least one edge (depending on all collected validators) is from level novize (human racial ability) ?????
- hindrance effects
- edges + edges effects + edged requirements
- DerivedStatistics
- gear + gear effects
- ??? gold validation + inventory + gear IF init validation ???
- build a frontend app with go and the wasm compiler or go and gopherjs/vecty
- use opaque types (example creating hindrances)

## out of scope v1.0
- racial abilities
  - android
    - hindrance: OUTSIDER (Major): Androids subtract 2 from Persuasion rolls when interacting with anyone besides other androids, and have no legal rights in most areas (they’re generally considered property).
- subskills (roadmap v1.1)
  - in racial ability pace modifiers
  - edge adjustments
  - hindrance adjustment
- pace
  - character enrichment with pace info about diffrent skills (roadmap v2)
