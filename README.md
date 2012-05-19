Metaphone package for Go
======================

A port of `metaphone` module from the Natural javascript package: <https://github.com/NaturalNode/natural/blob/master/lib/natural/phonetics/metaphone.js>

Also see the wikipedia article on metaphone: <http://en.wikipedia.org/wiki/Metaphone>

Installation
-------------

    go install github.com/sanjayc77/metaphone 

This will install the `metaphone` package.

Example
-------

    import "metaphone"

    st := metaphone.Process("phonetics")  // => FNTKS

Tests
-----

    go test metaphone

TODO
----
Precompile the regular expressions for faster processing.