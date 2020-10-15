Go Get Started
==============

Write some code
---------------

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

04_Exported names
-----------------

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

05_functions && 06_functions
----------------------------

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

.. info::

   no forward declareations needed and no header files.
   go attempts to reduce the amount of typing in both sensed of the word.
   everything is declared exactly once.

When two or more consecutive named function parameters share a type,
you can omit the type from all but the last.

``x int, y int`` -> ``x, y int``

.. code-block:: go

   func add(x, y int) int {
       return (x + y);
   }

07_Multiple_returns
-------------------

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


