1. 使用 `godoc -http :8000` 打开localhost的8000端口后，在浏览器中访问 `localhost:8000/pkg` 可以查看本地的 **package** 信息，可以看到分 **standard library** 和 **third-party** 两部分，后者就是 ***$GOPATH/src*** 目录下的包（package）。

2. **src** 下每个目录中的 Go 源文件只能属于一个包（package），一般而言文件中是 `package main` 的没有过多的描述信息。

3. 本目录定义的是 `package integers` ，Go Doc（http://localhost:8000/pkg/TDD/integers/) 中的信息如下图：

https://github.com/ZikeWang/TDD-go-codes/blob/master/pics/godoc_package_integers.png

4. 可以利用注释为函数添加文档，这些将出现在 Go Doc 中，就像你查看标准库的文档一样。注释写在函数定义前，必须以函数名起头。