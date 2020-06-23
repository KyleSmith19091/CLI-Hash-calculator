# CLI-Hash-Calculator
This is a command line interface program which can be used to quickly calculate a
hash for either a file or any literal value such as 12 or "Bed" etc. I designed this
due to the fact that I have had the desire to calculate a hash for a file or value so to create
this program was a 'no-brainer'

## Usage
The structure of the command is ->
```bash
$ cli hash --[value or filename] value or filename --hash [md5, sha1, sha 256, sha512]
```
Example
```
$ cli hash --value 3.1439593920 --hash md5
```
Output
```
$ Md5: 52e75cb1f7cf22b1a0aecec8a9f62818
```



