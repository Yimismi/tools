rm -rf ./target
go build bootstrap.go
go build main.go
mkdir target
mv bootstrap target
cd target
cp  ../config ./
cp  ../web ./
