# Is power of two

Given a positve integer number find out if the number is a power of two or not.

[**Naive solution**](https://github.com/TannerGabriel/learning-go/blob/master/algorithms/math/IsPowerOfTwo/isPowerOfTwo.go)

We keep dividing the integer number by two unless it becomes 1 and check if the remainder of the division is 0. Otherwise the number can't be a power of two.

[**Bitwise solution**](https://github.com/TannerGabriel/learning-go/blob/master/algorithms/math/IsPowerOfTwo/isPowerOfTwoBitwise.go)

Powers of two numbers in binary form always have just one bit. The only exception is with a signed integer (e.g. an 8-bit signed integer with a value of -128 looks like: 10000000)

```
1: 0001
2: 0010
4: 0100
8: 1000
16: 10000
32: 100000
```

After we check the number is greater 0 we can use the Bitwise AND operator to check that exactly one bit is set.

```
num & (num - 1)
```

For example for the number `8` that operations will look like:

```
  1000
- 0001
  ----
  0111
  
  1000
& 0111
  ----
  0000    
```


## References

* [Bitwise Solution on Stanford](http://www.graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2)
* [Bitwise Operations](https://en.wikipedia.org/wiki/Bitwise_operation)