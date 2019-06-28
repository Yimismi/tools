rm -rf ./target
go build bootstrap.go
go build main.go
mkdir target
mkdir target/tools
mv bootstrap target
cd target/tools
cp  -r ../../config ./
cp  -r ../../web ./
cd -