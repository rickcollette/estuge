# ESTUGE - S2G (Static Site Generator - ess-too-gee)

Pre-reqs:

golang >= 1.22.4

building:

```sh
go mod tidy
make build
```

running:

```sh
estuge SITENAME [template]
```

RPM Build:

```sh
make rpm
```

Debian Build:

```sh
make deb
```

Installation:

```sh
make install
```

Manual:

```sh
man estuge
```

More Help: https://govamp.org 

# Documentation

- [Create a basic website](Documentation/CreatingABasicWebsite.md)
- [File folder hierarchy](Documentation/FileFolderHierarchy.md)

# License
***MIT***
```
Copyright 2024 <megalith@root.sh>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```