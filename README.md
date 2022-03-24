Testing in Go
===

This repo conatins examples of packages in Go with a unit test suite for functions in each. 

The branch `solutions` includes the same packages but with the test solutions covered in the Testing in Go workshop

Packages
---

**viewcounter**

Basic video view counter which increments a 32 bit integer represting the number of views for a video. Test suite demonstrates basics of table testing.



**pwdvaalidator**

Validator function which returns an error message for invalid inputs. Test suite demonstrates adding full test coveraage for a function under test.

**mapiss**

Application that pull the current coorrdinates of the International Space Station and renders a map image with that location marked. Test suite demonstrates mocking functions called by a package.


Resources
---
- [Golang testing package](https://pkg.go.dev/testing)
- [Prefer Table Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- [Accept Interfaces Return Structs](https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b)