### Ward Monitor Photo/Profile Importer

Port of a java command line tool to Go. For importing staff photos and profiles directly into a Ward Monitor application. 
Currently being used by another script that scans a file system, resizes large photos and bulk imports them using this tool.
Looking at adding the features of this script into this tool, using the many Golang image libraries. 

--

###### Why Port to Go

- Should be faster without jvm overhead
- Compiles to executable so target system won't require Java to be installed
- Decent excuse to use Go

--

###### Todo

- Write some tests
- Add photo resizing features
