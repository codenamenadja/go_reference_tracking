HOW TO write go code
====================

Code organization
-----------------

Go programs are organized into packages.

Package
   Collection of source files in the same directory that are compiled together.

Functions, types, variables and constants defined in one source file 
are visible to all other source files within the same package.

A repository contains one or more modules.

Module
   Collection of related Go packages that are released together.

Repository
   Typically contains only one module located at the root of the repository.

A file, named ``go.mod`` there declares the module path::

   the import path prefix for all packages within the module.

The module contains the packages in the directory containing its ``go.mod`` file,
as well as subdirectories of that directory,
up to next subdirectory containing another ``go.mod`` file(if any).

Note that you don't need to publish your code to remote repository before you can build it.
A module can be defined locally without belonging to a repository.
However, it's good habit to organize your code as if you will publish it someday.

Each module's path not only serves as an import path prefix for its packages,
but also indicates where the ``go`` command should look to download it.
For ie, in order to download the module ``golang.org/x/tools`` ,
the ``go`` command would consult the repository indicated by ``https//golang.org/x/tools`` .

Import path
   Is a string used to import a package.

A package's import path is its module path joined with its subdirectory within the module.
For ie, the package's import path is ``github.com/google/go-cmp/cmp`` .
Packages in the standard library do not have a module path prefix.


Your First Program
------------------

To compile and run a simple program, first choose a module path and create ``go.mod`` file that declares it:

.. code-blcok:: bash

   $ mkdir hello
   $ cd hello
   $ go mod init exmple.com/usr/hello
   go: creating new go.mod: module example.com/user/hello
   $ cat go.mod
   module example.com/user/hello

   go 1.16
   $

The first statement in Go source file must be package name.
Executable commands must always use package ``main``

Next, create a file named hello.go inside that directory containing the following code:

.. code-block:: go

   package main

   import (
       "fmt"
   )

   func main()     {
       fmt.Println("Hello, world.");
   }

Now you can build and install that program with ``go`` tool::

   ``$ go install example.com/user/hello``

This command builds ``hello`` command, producing an executable binary.
It then, installs that binary as ``$HOME/go/bin/hello`` .

This install directory is controlled by the ``GOPATH`` and ``GOBIN`` environment variables.
If ``GOBIN`` is set, binaries are installed to that directory.
If ``GOPATH`` is set, binaries are installed to the ``bin`` subdirectory of the first directory in the ``GOPATH`` list.
Otherwise, binaries are installed to the ``bin`` subdirectory of the default ``GOPATH`` ``($HOME/go || %USERPROFILE%/go)`` .

You can use the ``go env`` command to portably set the default value for an environment variable for future ``go`` commands::

   ``$ go env -w GOBIN=/somedir/else/bin``
   ``$ go env -u GOBIN`` // unset the variable.

Command like ``go install`` apply within the context of the moudle containing the currect working directory.
If the working directory is not within the ``example.com/user/hello`` module, ``go install`` may fuil.

For convenience, go commands accepts pths relative to the working directory,
and default to the package in the current working directory if no other path is given.
So in our working directory, the following commands are all equivalent:

   ``$ go install example.com/user/hello``
   ``$ go install .``
   ``$ go install``



