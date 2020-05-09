# Levenshtein Distance

The Levenshtein distance is a string metric for measuring the difference between two sequences. Informally, the Levenshtein distance between two words is the minimum number of 
single-character edits (insertions, deletions or substitutions) required to change one word into the other.

![Levenshtein Distance](https://www.educative.io/api/edpresso/shot/4635258232242176/image/5657809267982336)

## Example

For example, the Levenshtein distance between `kitten` and `sitting` is `3`, since the following three edits change one into the other, and there is no way to do it with fewer than three edits:

1. **k**itten → **s**itten (substitution of "s" for "k")
2. sitt**e**n → sitt**i**n (substitution of "i" for "e")
3. sittin → sittin**g** (insertion of "g" at the end).

## Implementation 

- [Levenshtein Distance](https://github.com/TannerGabriel/learning-go/blob/master/algorithms/string/HammingDistance/levenshtein.go)

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Levenshtein_distance)
- [Educative](https://www.educative.io/edpresso/the-levenshtein-distance-algorithm)
