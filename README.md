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
- races + racial effects
- hindrance effects
- edges + edges effects + edged requirements
- DerivedStatistics
- gear + gear effects
- ??? gold validation + inventory + gear IF init validation ???
- subskills optional field in skills
- build a frontend app with go and the wasm compiler or go and gopherjs/vecty
- use opaque types (example creating hindrances)
