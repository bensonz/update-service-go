# update-service-go

An update service written in go. Uses usbmuxd (connect to iPad) and systemd
(Linux services manager)

Basically this program acts a role as a file manager from iPad, iPad can
write,change,delete,create new files in our system. Furthermore, this program
allows iPad to send execute instructions and execute files in our system. For
detailed commands, see [Avaliable commands](#avaliable-commands)

# environments

`go version` gives go version go1.9.2 darwin/amd64

# building

```
go build
```

and go will create an executable with the current dir name (update-service-go)
Put this executable in Linux system and run the service as put in systemd

# Avaliable commands

### for every single command, conclude it as [[::Whole message here::]]

filename should also include filepath

| commands   | full example                   | usage                                                                              |
| ---------- | ------------------------------ | ---------------------------------------------------------------------------------- |
| deletefile | deletefile:filename            | used to delete a file with filename                                                |
| writefile  | writefile:filename:filecontent | used to write a certain file with filename and filecontent (this appends the file) |
| exec       | exec:filename                  | execute a file with filename, only if executable                                   |
| readfile   | readfile:filename              | read a file                                                                        |
| appendfile | appendfile:filename:content    | appends to a file                                                                  |
