# strsim-examples

Real-world examples of [strsim](https://github.com/jcoruiz/strsim), a string similarity library for Go. To try its algorithms on your own text without writing code, use the [playground](https://strsim.than.cl).

## Examples

### [fuzzy-search](fuzzy-search/): product catalog search

Find the best matches for a typo-ridden query using different metrics. Shows how Jaro-Winkler, Levenshtein, and Dice produce different ranking results.

```
Query: "wireles headphone"

--- Jaro-Winkler ---
  #1  0.7372  Wireless Bluetooth Headphones
  #2  0.6933  Wireless Gaming Mouse

--- Levenshtein ---
  #1  0.5517  Wireless Bluetooth Headphones
  #2  0.4286  Wireless Gaming Mouse
```

**Features shown:** `FindTopN`, `Metric` interface, metric comparison.

### [dedup](dedup/): duplicate detection

Detect duplicate contacts using three strategies: edit distance (catches typos), Soundex (catches phonetic equivalents), and Double Metaphone (handles linguistic variations).

```
Strategy 1: Jaro-Winkler (typo detection)
  Group: "John Smith"
    ↳ "Jon Smith" (0.97)
    ↳ "John Smyth" (0.96)

Strategy 2: Soundex (phonetic matching)
  Soundex match: "John Smith" ↔ "John Smyth"  [J525]
```

**Features shown:** `FindAboveThreshold`, `SoundexMatch`, `DoubleMetaphoneMatch`, `ASCIIOnly`.

### [autocomplete](autocomplete/): autocomplete suggestions

Suggest completions for partial input against a dictionary of programming tools. Uses `ASCIIOnly` for maximum throughput on ASCII identifiers.

```
  "go"       → go (1.00), gorm (0.87), golang (0.82)
  "typscrip" → typescript (0.95), javascript (0.71)
  "kuber"    → kubernetes (0.90), vue (0.69)
```

**Features shown:** `FindTopN`, `ASCIIOnly` fast path.

### [spell-checker](spell-checker/): spell correction

Compare how Levenshtein, OSA, and Damerau-Levenshtein handle different types of misspellings. OSA and Damerau-Levenshtein correctly detect transpositions as single edits.

```
  "recieve":    Levenshtein → "believe" (dist: 2)
                OSA         → "receive" (dist: 1)  ← correct!

  "langauge":   Levenshtein → "language" (dist: 2)
                OSA         → "language" (dist: 1)  ← transposition
```

**Features shown:** `DistanceMetric` interface, `FindBestMatch` with an edit distance, `Distance()`, algorithm comparison.

### [name-matcher](name-matcher/): name matching across systems

Match person names across systems using a pipeline that combines Jaro-Winkler similarity with phonetic encoding for boosted scoring. Handles nicknames, spelling variations, and accented characters.

```
  Bob Johnson       → Robert Johnson       (score: 0.87)
  Katherine O'Brian → Catherine O'Brien    (score: 1.00)
  Micheal Smyth     → Michael Smith        (score: 1.00)
  Sarah Williams    → NO MATCH             (score: 0.54)
```

**Features shown:** `FindBestMatch`, `SoundexMatch`, `MetaphoneMatch`, `DoubleMetaphoneMatch`, combining edit distance + phonetic in a scoring pipeline.

## Running

```bash
go run ./fuzzy-search/
go run ./dedup/
go run ./autocomplete/
go run ./spell-checker/
go run ./name-matcher/
```

Requires Go 1.22.2 or newer, the version in `go.mod`.
