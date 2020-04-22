# Fast Powering Algorithm

**The power of a number** says how many times to use the number in a multiplication.

## Naive Algorithm Complexity

To find `a` raised to the power `b` we multiply `a` to itself, `b` times. That is, `a^b = a * a * a * ... * a` (`b` occurrences of `a`).

This operation will take `O(n)` time since we need to do multiplication operation exactly `n` times.

## Fast Power Algorithm

We can get better execution time by using a divide and conquer approach to compute power which can before up to `O(log(n))`.

The idea behind the algorithm is based on the fact that:

For **even** `Y`:

```text
X^Y = X^(Y/2) * X^(Y/2) 
```

For **odd** `Y`:

```text
X^Y = X^(Y/2) * X^(Y/2) * X
```

**For example**

```text
2^4 = (2 * 2) * (2 * 2) = (2^2) * (2^2)
```

```text
2^5 = (2 * 2) * (2 * 2) * 2 = (2^2) * (2^2) * (2)
```

**Time Complexity**

```text
O(log(n))
``` 

## Implementation

* [Fast powering](https://github.com/TannerGabriel/learning-go/blob/master/algorithms/math/FastPowering/fastpowering.go)

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Exponentiation_by_squaring)
- [GeeksForGeeks](https://www.geeksforgeeks.org/exponential-squaring-fast-modulo-multiplication/)
- [StackExchange](https://math.stackexchange.com/questions/2944032/why-is-the-algorithm-for-modular-exponentiation-by-squaring-considered-as-poly-t/2944039)