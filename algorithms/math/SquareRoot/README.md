# Square Root (Newton's Method)

**Newton's square root method** (also known as the Newton–Raphson method), named after _Isaac Newton_ and _Joseph Raphson_, is one example of a root-finding algorithm. It is a method for finding successively better approximations to the roots of a real-valued function.

## Newton's Method General Idea

The method starts with a function `f` defined over the real numbers `x`, the function's derivative `f'`, and an initial guess `x0` for a root of the function `f`. If the function satisfies the assumptions made in the derivation of the formula and the initial guess is close, then a better approximation `x1` is:

![](https://wikimedia.org/api/rest_v1/media/math/render/svg/52c50eca0b7c4d64ef2fdca678665b73e944cb84)

Geometrically, `(x1, 0)` is the intersection of the `x`-axis and the tangent of the graph of `f` at `(x0, f (x0))`.

The process is repeated as:

![](https://wikimedia.org/api/rest_v1/media/math/render/svg/710c11b9ec4568d1cfff49b7c7d41e0a7829a736)

Until the result reaches a sufficiently accurate value. The method can also be extended to complex functions and to systems of equations.

![](https://upload.wikimedia.org/wikipedia/commons/e/e0/NewtonIteration_Ani.gif)

## References

- [Wikipedia square roots](https://en.wikipedia.org/wiki/Methods_of_computing_square_roots)
- [Newton's square root method on Wikipedia](https://en.wikipedia.org/wiki/Newton%27s_method)