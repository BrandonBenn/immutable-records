<h1 align="center">Immutable Records using One-Way Hash</h1>

<br/>

<h2 align="center">DESCRIPTION</h2>

Given a directory of files, a snapshot of its contents is built using a Merkle
hash tree. The hash of root the is stored in a file called CHECKSUM. If any of
the files in the directory, or the CHECKSUM was tampered with, the verification
will give a warning.

<h2 align="center">SYNOPSIS</h2>

```sh
irow generate [directory]
irow verify   [directory]
```
