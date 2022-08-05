# learn-go-lang

To use different package in go lang, first create module with

```
go mod init <modulename>
// Best practice is to put as github.com/username
// For ex. github.com/cw-sarvesh
```

Then go to main.go file and add path same as your directory. 

To generate folder tree structure in windows run in root dir

```
tree /f > tree.txt
```

Folder structure

```bash
│   go.mod
│   main.go
│   README.md
│   tree.txt
│   
└───service
        service.go
```

in main.go, import like this

```
service "github.com/cw-sarvesh/service"
```

** Always remember to export any function or variable it must start with PascleCase
