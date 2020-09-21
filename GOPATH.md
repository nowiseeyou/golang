## go env环境查看 ##
用go env 可查看当前go环境变量。

$ go env
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/mac/MyCodeBase/Go-project/master"
GORACE=""
GOROOT="/usr/local/go"
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/dg/1_zsnh6n2md7y5kpjj318fk00000gp/T/go-build668952288=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
 

## GOROOT ##
GOROOT就是go的安装路径
在~/.bash_profile中添加下面语句:

GOROOT="/usr/local/go"
export GOROOT
当然, 要执行go命令和go工具, 就要配置go的可执行文件的路径:
操作如下:
在~/.bash_profile中配置如下:
export $PATH:$GOROOT/bin

## GOPATH ##
        o install/go get和 go的工具等会用到GOPATH环境变量.
        GOPATH是作为编译后二进制的存放目的地和import包时的搜索路径 (其实也是你的工作目录, 你可以在src下创建你自己的go源文件, 然后开始工作)。
            GOPATH之下主要包含三个目录: bin、pkg、src
            bin目录主要存放可执行文件; pkg目录存放编译好的库文件, 主要是*.a文件; src目录下主要存放go的源文件
        不要把GOPATH设置成go的安装路径,
        可以自己在用户目录下面创建一个目录, 如gopath
        操作如下:

        cd ~
        mkdir gopath

        在~/.bash_profile中添加如下语句:
        GOPATH=/Users/username/gopath
        GOPATH可以是一个目录列表, go get下载的第三方库, 一般都会下载到列表的第一个目录里面
        需要把GOPATH中的可执行目录也配置到环境变量中, 否则你自行下载的第三方go工具就无法使用了, 操作如下:
        在~/bash_profile中配置,

        export $PATH:$GOPATH/bin

        创建一个go项目, 并且编译运行:

        mkdir goproject
        cd goproject
        touch hello.go

        在hello.go中输入:

        package main
        import "fmt"
        func main() {
           fmt.Println("Hello, GO !")
        }

        在项目根目录下执行go build命令来构建你的项目, 构建后会生成hello文件
        运行生成的文件./hello, terminal中输出: Hello, GO !
        当然你也可以直接运行命令go run hello.go来执行程序.
 

如果要上述设置生效, 可以执行命令: source ~/.bash_profile, 上述所有操作均为mac系统下的操作, 如果是非mac系统, 请自行变通.

当存在多个路径时，好像优先采用第一个路径。这个无关紧要了，只有需要的第三方包（库）都能正确下载和使用就ok了。

## GOBIN ##
go install编译存放路径。不允许设置多个路径。可以为空。为空时则遵循“约定优于配置”原则，可执行文件放在各自GOPATH目录的bin文件夹中（前提是：package main的main函数文件不能直接放到GOPATH的src下面。

## GOPATH目录结构 ##
goWorkSpace  // (goWorkSpace为GOPATH目录)
  -- bin  // golang编译可执行文件存放路径，可自动生成。
  -- pkg  // golang编译的.a中间文件存放路径，可自动生成。
  -- src  // 源码路径。按照golang默认约定，go run，go install等命令的当前工作路径（即在此路径下执行上述命令）。
### go目录结构1: ###
project1 // (project1添加到GOPATH目录了)
  -- bin
  -- pkg
  -- src  
     -- models       // package
     -- controllers  // package
     -- main.go      // package main［注意，本文所有main.go均指包main的入口函数main所在文件］

 project2 // (project2添加到GOPATH目录了)
      -- bin
      -- pkg
      -- src
         -- models       // package
         -- controllers  // package
         -- main.go      // package main
 

 

package main文件直接在GOPATH目录到src下。

使用go build可以在src文件夹下编译生成名为“src”的可执行文件。这是golang默认约定。一般我个人不怎么用这个命令。因为它会生成可执行文件在src目录下面。

我一般用：go get 和 go install。

### go get ### [main.go所在路径]
参数 [main.go所在路径]：可选。相对GOPATH/src路径。 缺省是.(src自己)。可指定src下面的子文件夹路径。
go get会做2件事：1. 从远程下载需要用到的包。2.执行go install。（从这里也可以看出golang处处为了简洁而遵循的“约定优于配置”原则）

### go install ### [main.go所在路径]
参数 [main.go所在路径]：可选。 相对GOPATH/src路径。缺省是.(即当前所在目录或工作目录)。可指定src下面的子文件夹。
go install编译生成名称为[main.go父文件夹名]的可执行文件，放到GOBIN路径下。当GOBIN为空时，默认约定是：生成的可执行文件放到GOPATH/bin文件夹中。产生的中间文件（.a）放在project/pkg中（没有变化时，不重新生成.a）。

我们再看此go目录结构1，执行go install（以及go get的第二阶段go install）会报错：

**注意**：如果不用额外方式改变环境变量(公司目前用的是sh脚本编译)，是编译不过的。报错：can’t load package: package .: no buildable Go source files in *

### 解决方法1 ###：
曾经我也因为这个错误感到迷惑，认为所有都环境变量都没有问题。网上也没怎么看到直接明确都解答。看了一些帖子后，触类旁通，设置了GOBIN环境变量后解决。（好吧，我至今也没有完整读过英文官方文档。这种默认约定，官方文档上应该有。）

此解决方法有个弊端，多个project会导致多个GOPATH目录。多个project下的目录结构和包一致的话，直接编译会导致编译问题。因为go优先使用第一个GOPATH目录，导致编译冲突。（当然，你也可以每次手工或脚步修改GOPATH环境变量，感觉很麻烦。）不建议多个project直接设置到茫茫多的GOPATH中。当然有解决方法2，我认为是标准合理的解决方法，就是下面go目录结构2了。

go目录结构2:
goWorkSpace     // goWorkSpace为GOPATH目录
  -- bin
     -- myApp1  // 编译生成
     -- myApp2  // 编译生成
     -- myApp3  // 编译生成
  -- pkg
  -- src
     -- common 1
     -- common 2
     -- common utils ...
     -- myApp1     // project1
        -- models
        -- controllers
        -- others
        -- main.go 
     -- myApp2     // project2
        -- models
        -- controllers
        -- others
        -- main.go 
     -- myApp3     // project3
        -- models
        -- controllers
        -- others
        -- main.go 
 

一个solution里面的多个project或工具组件都并列放在GOPATH的src下，如myApp1，myApp2，myApp3，common utils。
这时，GOBIN可以为空，编译时，可以如下：
go install myApp1 或 go get myApp1
go install myApp2 或 go get myApp2
go install myApp3 或 go get myApp3

这时才是大家都认为的，把可执行程序myApp1、myApp2、myApp3生成在goWorkSpace/bin下面。多个GOPATH也就有了上面的“把每个GOPATH下的bin都加入到PATH中”。

提示：相同结构的project下同名包怎么办？
有同事在初学时，习惯按照go目录结构1，了解到go目录结构2后（以为仅仅是把main放到了子文件夹，其他controllers等包结构不变），有这样的疑惑。他们原来就有这样的问题，同时把go目录加入到GOPATH后，编译就出现问题，因为包名和路径相同（相对GOPATH下的src），go只会优先查找第一个符合的GOPATH）。只会每次编译时手工修改GOPATH，或写脚本编译。（我看着就觉得累，还徒增脚本维护烦恼。）

解决方案就是：除了通用的，公有的工具、组件外，属于各个project自己的东西，统统随着main.go一起移到project目录下。go目录结构2就是这样的。
导入各个project下的controllers方法：import myApp1/controllers，import myApp2/controllers这样的。go的import查找的是包的路径，并不是包名。你只用告诉go，你的包都放在哪了，go会从这些路径下查找有没有所需要的包。只是大家一般习惯包名和文件夹名相同，容易误解。只需要注意，一个文件夹下只允许有一个包名，允许有子文件夹定义不同的包。
import 采用的是相对路径写法：路径是 相对GOROOT和GOPATH下的src。

也可以设置GOBIN，而且这时，由于可执行文件名称不同，也不大容易产生覆盖（需要避免的时多个GOPATH用相同的“myApp”project名称。）

具体的还是看个人喜好和实际情况。我个人本地的环境大致是：

      
  -- my-goWorkSpace    // 主要是为了区分自己的鼓捣的一些东西和工作上的项目
  -- work-goWorkSpace
        -- bin
        -- pkg
        -- src                  
           -- myApp1
              -- .git
              -- models
              -- controllers
              -- main.go 
           -- myApp2
              -- .git
              -- models
              -- controllers
              -- main.go 
           -- myApp3
              -- .git
              -- models
              -- controllers
              -- main.go
本地是有多个GOPATH路径的
注意：GOPATH目录和GOPATH下的src不应该添加到源代码管理中，而是各个project目录myApp1、myApp2、myApp3各自时独立的进行源代码管理

我一般不设置GOBIN，把每个GOPATH下的bin都加入到PATH中。