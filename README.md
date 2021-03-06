# intcode

## Usage

```go
intcode <flags> <path>
```

## Intcode Language Specifications

### Opcodes

The Opcodes 01-09 and 99 are defined in [Advent of Code 2019](https://adventofcode.com/2019).

The Opcodes 10-19 and 80 are an extension by me.

| Opcode | Params | Description                                                  |
| ------ | ------ | ------------------------------------------------------------ |
| 01     | 3      | first arg + second arg = third arg                           |
| 02     | 3      | first arg * second arg = third arg                           |
| 03     | 1      | Inputs an integer and stores it in the first arg             |
| 04     | 1      | Outputs the first arg                                        |
| 05     | 2      | If the first arg is non-zero, sets the instruction pointer to the second arg |
| 06     | 2      | If the first arg is zero, sets the instruction pointer to the second arg. |
| 07     | 3      | If the first arg is less than the second argument, sets the third arg to 1. If not less, sets it to 0 |
| 08     | 3      | If the first arg is equal to the second argument, sets the third arg to 1. If not equal, sets it to 0 |
| 09     | 1      | Adds the first arg to the relative base register             |
| **—**  | **—**  | **Additional section below**                                 |
| 10     | 3      | Bitwise AND: first arg \& second arg = third arg             |
| 11     | 3      | Bitwise OR: first arg \| second arg = third arg              |
| 12     | 3      | Bitwise XOR: first arg ^ second arg = third arg              |
| 13     | 3      | Integer Division: first arg / second arg = third arg         |
| 14     | 3      | Modulo: first arg % second arg = third arg                   |
| 15     | 3      | Left shift: first arg << second arg = third arg              |
| 16     | 3      | Right shift: first arg >> second arg = third arg             |
| 17     | 3      | Negate: If first arg is 0, return 1. Otherwise return 0      |
| 18     | 3      | Unix Timestamp: Sets the first arg to the current unix timestamp |
| 19     | 3      | Random: Sets the first arg to a random positive number       |
| *80*   | *3*    | **Future:** *Syscall: Perform a syscall with eax=first arg, ebx=second arg, ecx=third arg* |
| 99     | 0      | Ends the program                                             |

> From [esolangs.org/wiki/Intcode](https://esolangs.org/wiki/Intcode)

### Parameter Modes

| Mode | Name           | Description                                                  |
| ---- | -------------- | ------------------------------------------------------------ |
| 0    | Position Mode  | The parameter is the address of the value.                   |
| 1    | Immediate Mode | The parameter is the value itself (Not used for writing).    |
| 2    | Relative Mode  | The parameter is added to the relative base register, which results in the memory address of the value. |

> From [esolangs.org/wiki/Intcode](https://esolangs.org/wiki/Intcode)

```bash
ABCDE
 1002

DE - two-digit opcode,      02 == opcode 2
 C - mode of 1st parameter,  0 == position mode
 B - mode of 2nd parameter,  1 == immediate mode
 A - mode of 3rd parameter,  0 == position mode,
                                  omitted due to being a leading zero
```

> From [adventofcode.com/2019/day/5](https://adventofcode.com/2019/day/5)

### Examples

```bash
1, 3, 4, 2, 99
```
