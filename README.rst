Go Get Started
==============

01. Write some code
-------------------

.. code-blcok:: go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!");
   }

Desc::

   - Declare a ``main`` package.
   - Import the popular ``fmt`` package, which contains functions for formatting text, including printing to the console.
     This package is one of the standard library packages you got when you installed GO.
   - Implement a ``main`` function to print a message to the console. ``main`` is a function which default executes.

04. Exported names
------------------

a name is exported if it begins with a capital letter.
when importing package, you can refer only its exported names.

.. code-block:: go

   package main

   import (
       "fmt"
       "math"
   )

   func main() {
       fmt.Println(math.Pi);
   }

05. functions && 06_functions
-----------------------------

A function can take zero or more arguments.
in this example, ``add`` takes two parameter of type ``int`` .
Notice that the type comes after the variable name.

.. code-block:: go

   package main

   import (
       "fmt"
   )

   func main() {
       fmt.Println(add(42, 13));
   }

   func add(x int, y int) int {
       return (x + y);
   } // no forward declareations in go.

.. note::

   no forward declarations needed and no header files.
   go attempts to reduce the amount of typing in both sensed of the word.
   everything is declared exactly once.

When two or more consecutive named function parameters share a type,
you can omit the type from all but the last.

``x int, y int`` -> ``x, y int``

.. code-block:: go

   func add(x, y int) int {
       return (x + y);
   }

07. Multiple_returns
--------------------

A function can return any number of results.
The ``swap`` functions returns two strings.

.. code-block:: go

   package main

   import (
       "fmt"
   )

   func main() {
       a, b := swap("hello", "world");
       fmt.Println(a, b);
   }

   func swap(x, y string) (string, string) {
       return y, x;
   }

08. Named return values
-----------------------

Go's return values may be named. if so,
they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A ``return`` statement without arguments returns the named return values.
This is known as a "naked" return.

Naked return statements should be used only in short functions,
as with the example shown here.
They can harm readability in longer functions.

.. code-block:: go

   func main() {
       fmt.Println(split(17));
   }

   func split(sum int) (x, y int) {
       x = sum * 4 / 9;
       y = sum - x;
       return;
   }

.. note::

   guess not recommanded...

09. Variables
-------------

The ``var`` statement declares a list of variables::

   as in function argument lists, the type is last.

A ``var`` statement can be at package or function level.
we see both is this example.

.. code-block:: go

   var c, python, java bool

   func main() {
       var i int;
       fmt.Println(i, c, python, java);
   }

10. Variables with initializers
-------------------------------

A ``var`` declaration can include initializer, one per variable.
If an initializer is present, the type can be omitted.
the variable will take the type of the initializer.

.. code-block:: go

   var i, j int = 1, 2;

   func main() {
       var c, python, java = true, false, "nol";

       fmt.Println(i, j , c, python, java);
   }

11. Basic types
---------------

Go's basic types are::

   - bool
   - string
   - int
   - int8
   - int32
   - int64
   - uint
   - uint8
   - uint32
   - uint64
   - uintptr
   - byte // alias for unint8
   - rune // alias for int32 // represents a Unicode point
   - float32
   - float64
   - complex64
   - complex128

The example shows variables of several types,
and also that variable declarations may be "factored" into blocks, as with import statements.

The ``int`` ``uint`` and ``uintptr`` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
When you need an integer value you should use ``int`` ,
unless you have a specific reason to use a sized or unsigned integer type.

.. code-block:: go

   import (
       "fmt"
       "math/cmplx"
   )

   var (
       ToBe        bool        =   false
       MaxInt      uint64      =   1<<64 -1
       z           complex128  =   cmplx.Sqrt(-5 + 121)
   )

   func main() {
       fmt.Printf("Type: %T\tValue: %v\n", ToBe, ToBe);
       fmt.Printf("Type: %T\tValue: %v\n", MaxInt, MaxInt);
       fmt.Printf("Type: %T\tValue: %v\n", z, z);
   }

.. code-block:: bash

   Type: bool	Value: false
   Type: uint64	Value: 18446744073709551615
   Type: complex128	Value: (10.770329614269007+0i)

12. Zero values
---------------

Variables declared without an explicit initial value are given their ``zero`` value.

The zero value is::

   - ``0`` for numeric types
   - ``false`` for the booleans type
   - ``""`` (empty string) for strings

.. code-block:: go

   func main() {
       var i   int
       var f   float64
       var b   bool
       var s   string
       fmt.Printf(
           "%v %v %v %q\n",
           i, f, b, s,
       );
   }

.. code-block:: bash

   0 0 false ""

13. Type Conversions
--------------------

The expression ``T(v)`` converts the value ``v`` to the type ``T`` .

Some numeric conversions::

   var i int        = 42;
   var f float64    = float64(i);
   var u uint       = uint(f);

Or, put more simply::

   i :=     42;
   f :=     float(64);
   u :=     uint(f);

Unlike in C, in Go, assignment between items of different types requires an explicit conversion.
Try removing the ``float64`` or ``uint`` conversions in the example and see what happens.

.. code-block:: go

   func main() {
       var x, y int        = 3, 4;
       var f float64       = math.Sqrt(float64(x * x + y * y));
       var z uint          = uint(f);
       i                   := 42;
       f2                  := float64(i);
       u                   := uint(f);

       fmt.Println(x, y, z);
       fmt.Println(i, f2, u);
   }
