# Protoq master test  1asd
Grpc exercise 
  
在hello-world文件夹下执行
protoc -I proto/ proto/helloworld.proto --go_out=plugins=grpc:proto

protocol编译器就会在一系列目录中查找需要被导入的文件，这些目录通过protocol编译器的命令行参数-I/–import_path指定。如果不提供参数，编译器就在其调用目录下查找。