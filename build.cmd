go build bootstrap.go
go build main.go
md target
move bootstrap.exe target
cd target
md config
md web
md log
xcopy /e "../config" "config"
xcopy /e "../web" "web"
