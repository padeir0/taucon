# Taucon

Mostly made to generate a list of all possible TAUtologies and CONtradictions up to a limit.

The program only considers expressions where every operand is a variable.
So expressions like `T`, `T -> F`, `T <-> T` are valid, but not counted.

Operators considered are `∨`, `∧`, `->`, `<->` and `~`.
There are only two precedence levels, one for binary operators and one for
unary operators, all order of evaluation is expressed using parenthesis

The following table presents the results, the function `f`
represents the whole program, it's shown as
`f(i, j)` where `i` = number of operators, `j` = number of variables.
The columns are:
 - `inconst` column represents the number of logically contingent
expressions (ie. the number of expressions that are not a tautology or
contradiction)
 - `const` column represents the number of tautologies
and contradictions (contradictions + tautologies)
 - `con` column represents the number of contradictions
 - `tau` column represents the numbe of tautologies

|           |  inconst | const |  con  |   tau   |
|:---------:|:--------:|:-----:|:-----:|:-------:|
| `f(0, 0)` | 0        | 0     | 0     |  0      |
| `f(0, 1)` | 1        | 0     | 0     |  0      |
| `f(0, 2)` | 0        | 0     | 0     |  0      |
| `f(0, 3)` | 0        | 0     | 0     |  0      |
| `f(0, 4)` | 0        | 0     | 0     |  0      |
| `f(1, 0)` | 0        | 0     | 0     |  0      |
| `f(1, 1)` | 3        | 2     | 0     |  2      |
| `f(1, 2)` | 8        | 0     | 0     |  0      |
| `f(1, 3)` | 0        | 0     | 0     |  0      |
| `f(1, 4)` | 0        | 0     | 0     |  0      |
| `f(2, 0)` | 0        | 0     | 0     |  0      |
| `f(2, 1)` | 23       | 22    | 6     |  16     |
| `f(2, 2)` | 190      | 26    | 0     |  26     |
| `f(2, 3)` | 192      | 0     | 0     |  0      |
| `f(2, 4)` | 0        | 0     | 0     |  0      |
| `f(3, 0)` | 0        | 0     | 0     |  0      |
| `f(3, 1)` | 263      | 242   | 48    |  194    |
| `f(3, 2)` | 4374     | 1114  | 78    |  1036   |
| `f(3, 3)` | 11724    | 756   | 0     |  756    |
| `f(3, 4)` | 7680     | 0     | 0     |  0      |
| `f(4, 0)` | 0        | 0     | 0     |  0      |
| `f(4, 1)` | 3199     | 3146  | 710   |  2436   |
| `f(4, 2)` | 108566   | 33274 | 3740  |  29534  |
| `f(4, 3)` | 557970   | 63150 | 2244  |  60906  |
| `f(4, 4)` | 880776   | 33144 | 0     |  33144  |
