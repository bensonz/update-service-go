# update-service-go

An update service written in go. Uses usbmuxd (connect to iPad) and systemd
(Linux services manager)

## Avaliable commands

### for every single command, conclude it as [[::Real message here::]]

| commands   | full example                   | usage                                                                              |
| ---------- | ------------------------------ | ---------------------------------------------------------------------------------- |
| deletefile | deletefile:filename            | used to delete a file with filename                                                |
| writefile  | writefile:filename:filecontent | used to write a certain file with filename and filecontent (this appends the file) |
| exec       | exec:filename                  | execute a file with filename, only if executable                                   |
