# Sieve of Eratosthenes

The Sieve of Eratosthenes is an algorithm for finding all prime numbers up to some limit `n`.

Steps: 

1. Create a boolean array of `n + 1` positions to represent the numbers `0` through `n`.
2. Set positions `0` and `1` to `false`, and the rest to `true`.
3. Start at the first prime number `p = 2`.
4. Mark as `false` all the multiples of `p`.
5. Find the first position greater than `p` that is `true` in the array. If there is no such position, stop. Otherwise, let `p` equal this new number (which is the next prime), and repeat from step 4.

When the algorithm terminates, the numbers remaining `true` in the array are all the primes below `n`.

![Sieve](https://upload.wikimedia.org/wikipedia/commons/b/b9/Sieve_of_Eratosthenes_animation.gif)

## Complexity

The algorithm has a complexity of `O(n log(log n))`.

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)