# Pascal's Triangle

In mathematics, **Pascal's triangle** is a triangular array of the [binomial coefficients](https://en.wikipedia.org/wiki/Binomial_coefficient).

The rows of Pascal's triangle are conventionally enumerated starting with row `n = 0` at the top (the `0th` row). The entries in each row are numbered from the left beginning 
with `k = 0` and are usually staggered relative to the numbers in the adjacent rows. The triangle may be constructed in the following manner: In row `0` (the topmost row), there 
is a unique nonzero entry `1`. Each entry of each subsequent row is constructed by adding the number above and to the left with the number above and to the right, treating blank 
entries as `0`. For example, the initial number in the first (or any other) row is `1` (the sum of `0` and `1`), whereas the numbers `1` and `3` in the third row are added to produce the number `4` in the fourth row.

![Pascal's Triangle](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

## Implementation

- [Pascal's Triangle](https://github.com/TannerGabriel/learning-go/blob/master/algorithms/math/PascalTriangle/pascal.go)

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Pascal%27s_triangle)
- [GeeksForGeeks](https://www.geeksforgeeks.org/pascal-triangle/)