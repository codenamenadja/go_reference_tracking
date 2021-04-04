Understanding The Go
====================

The origins of go
-----------------

*C-like language, C for 21th centry*
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Inherited from C
   1. Expression syntax
   #. Control flow statements
   #. Basic data types
   #. call-by-value parameter passing
   #. pointers

And above all, C's emphasis on programs that,
   - *Compile to efficient machine code* 
   - *Cooperate naturally with the abstractions of current operating systems*

The Characteristics Of Go
-------------------------

Simplicity
^^^^^^^^^^
   As *Rob Pike* put it, **"Complexity is multiplicative":**
      fixing problem by *making one part of system more complex slowly but surely adds complexity to other parts.*
   Simplicity requrires more work at the beginning of a project
      To reduce an idea to its essence and more discipline over the lifetime of project
         To distinguish good changes from bad ones.

As a High-level language
^^^^^^^^^^^^^^^^^^^^^^^^
   Benefits of hindsight, Go has,
      - GC
      - Package system
      - First class functions
      - Lexical scope
      - SysCall Interface
      - Immutable strings in generally encoded in UTF-8

   Diff from others, Go has,
      - No implicit numeric conversions
      - No constructors or Destructors
      - No Operator overloading
      - No default parameter values
      - No inheritance
      - No generics
      - No exceptions
      - No macros
      - No function annotations
      - No thread-local storage

Type System
^^^^^^^^^^^
   Go has enough of Type system to avoid most of areless mistakes that plague in dynamic languages
   But has Simpler type system than comparable typed languages.

Aware of Computer System design
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
   Go encourages an awareness of comtemporary computer sys design,
   particularly the **importance of locality.**

   Built-in data types and data structures are crafted to work narually without,
      - explicit initialization
      - implicit constructores.

   So relatively,
   few memory allocations and memory writes are hidden in the code.

   Go's aggregate types, *structs* & *arrays,*
   hold their element directly, requireing
      - less storage
      - fewer allocations
      - pointer indirections
   than languages that uses indirect fields.

*batteries included* Standard library
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
   Go's stdlib provides,
      - clean building blocks
      - clean APIS for I/O
      - Text processing
      - Graphics
      - Cryptography
      - Networking
      - Distributed applications
   with supprot for many *standard file formats and protocols.*
