![logo](https://chaosblade.oss-cn-hangzhou.aliyuncs.com/doc/image/chaosblade-logo.png)

# Chaosblade-exec-cplus: Chaosblade executor for chaos experiments on c++ applications


## Introduction
The project is a chaosblade executor based on [GDB] for chaos 
experiments on c++ applications. The drill can be implemented through the blade cli, see 
[chaosblade](https://github.com/chaosblade-io/chaosblade) project for details.


## Compiling
In the project root directory, execute the following command to compile
```bash
make
```

The compilation result will be stored in the target directory.


## Deploying
When deploying Chaosblade-exec-cplus, the shell script files under the source’s folders [/src/main/resources/script] should be deployed on the disk of server separately. Create a new folder in server, put all the shell script files under the new folder, and when you startup chaosblade-exec-cplus jar, you can use the command just like: 
```bash
nohup java -jar chaosblade-exec-cplus.jar --server.port=8908 --script.location=/home/admin/cplus/ & 
```
the value of parameter "script.location" can be set the location of the new created folder. So that the  chaosblade-exec-cplus jar know where to get the script files.


## Restriction
1. If the C++ application doesn’t add '-g' when compiling, Chaosblade-exec-cplus can’t work. 
2. If the linux system can’t support yum command to install software,  the [GDB], and [expect] software should be manual installed, before use Chaosblade-exec-cplus.


## Contributing
We welcome every contribution, even if it is just a punctuation. See details of [CONTRIBUTING](CONTRIBUTING.md)


## Bugs and Feedback
For bug report, questions and discussions please submit [GitHub Issues](https://github.com/chaosblade-io/chaosblade-exec-cplus/issues).

Contact us: chaosblade.io.01@gmail.com

Gitter room: [chaosblade community](https://gitter.im/chaosblade-io/community)


## License
Chaosblade-exec-cplus is licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for the full license text.
