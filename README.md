Intro
-----

Toy client-server example that demonstrates Golang's nice networking libraries. 

Context
-------

NIST maintains official timekeeping computers at [time.nist.gov](time.nist.gov). They are for public use, and respond with the current time when upon request.

Let's setup a TCP connection between our computer and the official NIST servers on `PORT 13`.

Usage
-----

```bash
# client demo
make client
./client time.nist.gov:13
```
