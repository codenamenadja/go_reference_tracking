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
