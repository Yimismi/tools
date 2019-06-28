go build bootstrap.go
go build main.go
mkdir target
mv bootstrap.exe target
cd target
mkdir config
mkdir web
mkdir log
cp -r /e "../config" "config"
cp -r /e "../web" "web"
