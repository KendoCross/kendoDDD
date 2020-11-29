# kendoDDD

微服务的强大之处在于清晰地定义了它们的职责并划定了它们之间的边界。它的目的是在边界内建立高内聚，在边界外建立低耦合。也就是说，倾向于一起改变的事物应该放在一起。正如现实生活中的许多问题一样，但这说起来容易做起来难，业务在不断发展，设想也随之改变。因此，重构能力是设计系统时考虑的另一项关键问题。

那应该怎么将一个大的系统合理有效的拆分成微服务呢？你需要了解领域驱动设计（DDD）。领域驱动设计因 Eric Evans 的著作而出名，它是一组思想、原则和模式，可以帮助我们基于业务领域的底层模型设计软件系统。开发人员和领域专家一起使用统一的通用语言创建业务模型。然后将这些模型绑定到有意义的系统上，在这些系统和处理这些服务的团队之间建立协作协议。更重要的是，它们设计了系统之间的概念轮廓或边界也就是上下文。微服务设计从这些概念中汲取了灵感，因为所有这些原理都有助于构建可以独立变更和发展的模块化系统。

但是DDD也仍然是一个很大的很宽泛的方法论，需要了解的东西很多。本文尝试直接从代码分层入手，简单的入门一下。

## 一、分层说明

分层架构有一个重要的原则：每层只能与位于其下方的层发生耦合。详细说来又可以分为：严格分层架构，某层只能与直接位于其下方的层发生耦合；松散分层架构，则允许任意上方层与任意下方层发生耦合。严格分层，自然是最理想化的，但这样肯定会导致大量繁琐的适配代码出现，故在严格与松散之间，追寻和把握恰达好处的均衡。

![DDD经典分层](./asset/01.png)
如上图，这是经典的DDD分层方式。由于领域驱动包含的概念实在太多了，不论是聚合根、领域对象、领域服务、领域事件、仓储、贫血充血模型、界限上下文、通用语言，任何一个深入起来都又太多内容了，所以一开始换一个角度，从MVC/MVP/MVVM的分层架构入手，类比理解DDD经典的四层。然后融合自己已有的编码习惯和认知，按照各层的主要功能定位，可以写的代码范围约束，慢慢再结合理解DDD的概念完善代码编写。

1.presentation 表层。负责向用户显示信息和解释用户指令。这里的用户不一定是使用GUI的人，也可以是另一个系统。 该层系统的出入口主要的功用逻辑也尽量的简单，主要承接不同“表现”形式采集到的指令/出入参等，并进行转发给应用层。该层的核心价值在于多样化，而不在于功能有多强大。表现形式可以是 ①命令行②gRPC③HTTP(Gin)④HTTP(Beego)等,不涉及到具体的业务逻辑，彼此之间有了很好的替代。

2.application 应用层。定义软件要完成的任务，并且指挥表达领域概念的对象来解决问题。应用层要尽量简单，不包含业务规则，只为下一层中的领域对象协调任务，分配工作。 我的理解应用层是很薄的一层，只作为计算机领域到业务领域的过渡层。比如计算机能够识别和传输的肯定是2进制字节流，这一层可以充当翻译，把这些晦涩难懂的机器“语言”，转化为领域业务人员建模出来的语言。或者说是高级计算机编程语言，这里一般会有专门的xxxxVM来承接所需的出参、入参数据。这一层直接消费领域层，并且开始记录一些系统型功能，比如运行日志、事件溯源。
这一层的也应该尽可能的业务无关，以公用代码逻辑为主。

3.domain领域层。负责表达业务概念，业务状态信息以及业务规则。尽管保存业务状态的技术细节由基础设施层提供，但反应业务情况的状态是由本层控制并使用的。 领域驱动设计里最核心的部分了，可以细拆分为聚合根、实体，领域服务等一大堆其他概念。这里不展开详细说明了，简单的理解下 聚合根，负责整个聚合业务的所有功能就行了。比如项目中的fileAggregate，该类直接负责与平台系统管理员相关的所有操作业务，对内依赖调用领域服务、其他实体，或封装一些不对外的方法、函数等，完成所有所需的功能，由聚合根对外统一提供方法。可以把之前MVP里面的Presenter的主要代码需要完成的功能转移过来，再按照领贫血模型来分离出领域服务+实体+值对象等。

4.infrastructure基础设施层。为上面各层提供通用的技术能力：为应用层传递消息，为领域层提供持久化机制，为表现层绘制屏幕组件，等等。 这一层也是讲求和业务逻辑无关，只重点提供通用的功能。未来最主要需要编码的部分是仓储功能，和数据库打交道的这部分逻辑了，其他的基本都是基础功能或帮助类，变更的概率不大。
仓储层标准实践，也尽可能的做到与具体业务无关，单纯的（增删改查）数据库持久化等功能。

## 二、业务实践

以最简单的文件管理业务为例，实现两个接口：1.添加文件，2.获取文件。

1.presentation 表现层。主要和不同的Web框架有一定的耦合，不同的框架代码不全一样。但核心功能是相同的，就是进行HTTP请求和应用层的视图模型的双向转换，并且处理HTTP的状态等。

```golang
    //获取到文件的基本信息，和保存路径，并返回给请求方
    func GetFile(c *gin.Context) {
        rst, has, err := application.GetFileById(c.Param("id"))
        if err != nil {
            c.Error(err)
            return
        }
        if !has {
            c.AbortWithStatus(http.StatusNotFound)
            return
        }
        c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", rst.FileName))
        c.Header("Content-Type", rst.ContentType)
        c.File(rst.FilePath)
    }

    //新增文件，通过web框架将HTTP请求转换为应用层的视图模型
    func AddFile(c *gin.Context) {
        var parm application.AddFileForm
        if err := c.ShouldBind(&parm); err != nil {
            c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
            return
        }
        if parm.UpFile == nil {
            c.Error(errorext.NewCodeError(101, "文件无效", nil))
            return
        }
        fileId, err := application.AddFile(parm)
        if err != nil {
            c.Error(err)
            return
        }
        c.JSON(http.StatusOK, fileId)
    }
```

2.application应用层。通用的日志、打点等。通过直接持有领域层的聚合根，仓储层等直接进行业务表达。并将不常变化的领域模型，转换为可能经常变化的视图模型。

```go
    //从数据库获取文件的基本信息，不涉及到
    func GetFileById(id string) (file FileInfo, has bool, err error) {
        obj, has, err := infrastructure.RepoFac.FilesRepo.GetById(id)
        if err != nil {
            logs.Error("FilesRepo GetById ERR:%v", err)
            return
        }
        if !has {
            return
        }

        file.ContentType = obj.ContentType
        file.FileName = obj.FileName
        file.FilePath = obj.FilePath
        file.Size = obj.Size
        return
    }

    func AddFile(parm AddFileForm) (fileId string, err error) {
        fileInfo := files.FileInfos{}
        f, err := parm.UpFile.Open()
        if err != nil {
            return
        }
        defer f.Close()
        fileInfo.FileBody, err = ioutil.ReadAll(f)
        if err != nil {
            return
        }
        fileInfo.ContentType = parm.UpFile.Header.Get("Content-Type")
        fileInfo.Size = int(parm.UpFile.Size)
        fileInfo.FileName = parm.UpFile.Filename
        fileId, err = files.SingleFilesAgg.AddNewFile(fileInfo)
        if err != nil {
            logs.Error("SingleFilesAgg AddFile ERR:", err.Error())
            return
        }
        return
    }

    //领域模型不应该直接暴露给用户，专门为可视化定制的视图模型，视图模型可能是多变的
    type FileInfo struct {
        FileName    string `json:"file_name" xorm:"comment('文件名称') VARCHAR(255)"`
        FilePath    string `json:"file_path" xorm:"comment('文件目录') VARCHAR(256)"`
        ContentType string `json:"content_type" xorm:"comment('文件类型') VARCHAR(64)"`
        Size        int    `json:"size" xorm:"comment('文件大小') INT"`
    }

    type AddFileForm struct {
        UpFile *multipart.FileHeader `form:"up_file"`
        Remark string                `form:"remark"`
    }
```

3. domain领域层。负责表达业务概念，业务状态信息以及业务规则。尽管保存业务状态的技术细节由基础设施层提供，但反应业务情况的状态是由本层控制并使用的。 领域驱动设计里最核心的部分了，可以细拆分为聚合根、实体，领域服务等一大堆其他概念。这里不展开详细说明了，简单的理解下 聚合根，负责整个聚合业务的所有功能就行了。比如项目中的filesAggregate，该类直接负责文件管理相关的所有操作业务，对内依赖调用领域服务、相关实体等，或封装一些不对外的方法、函数，完成所有所需的功能，由聚合根对外统一提供方法。可以把之前MVP里面的Presenter的主要代码需要完成的功能转移过来，再按照领贫血模型来分离出领域服务+实体+值对象等。

```go
    //聚合根作为该领域上下文唯一对外开放的出入口
    //聚合根持有相关的实体，指挥实体完成任务
    func (a *filesAggregate) AddNewFile(fileInfo FileInfos) (fileId string, err error) {
        fileId = uuid.New().String()
        fileInfo.Id = fileId
        en := newfileEnById(fileInfo.Id)
        err = en.AddFile(fileInfo.FileBody)
        return
    }

    //真正实际处理业务，存储文件，并写入数据库进行记录。由文件实体来完成
    func (en *fileEntity) AddFile(body []byte) (err error) {
        filePath := FileRootPath + time.Now().Format("2006/01/02/") + en.Id + en.FileName
        helper.MakesureFileExist(filePath)
        err = ioutil.WriteFile(filePath, body, 0755)
        if err != nil {
            return
        }
        en.FilePath = filePath
        en.Status = 1
        _, err = fileRepos.AddObj(en.FileInfo)
        return
    }
```

4. infrastructure基础设施层。为上面各层提供通用的技术能力：为应用层传递消息，为领域层提供持久化机制，为表现层绘制屏幕组件，等等。 这一层也是讲求和业务逻辑无关，只重点提供通用的功能。未来最主要需要编码的部分是仓储功能，和数据库打交道的这部分逻辑了，其他的基本都是基础功能或帮助类，变更的概率不大。
仓储层标准实践，也尽可能的做到与具体业务无关，单纯的（增删改查）数据库持久化等功能。
本项目里fileRepo实现对文件管理数据持久化的一些CRUD通用业务代码。

```go
    //新增，
    func (r *fileRepo) AddObj(obj *interfaces.FileInfo) (num int64, err error) {
        now := time.Now().Unix()
        obj.Created = now
        obj.Updated = now
        _, err = writeEngine.
            Insert(obj)
        return
    }

    //单条查询
    func (r *fileRepo) GetById(id string) (obj interfaces.FileInfo, has bool, err error) {
        has, err = readEngine.
            Where("id=?", id).
            Get(&obj)
        return
    }

    func (r *fileRepo) Find(parm interfaces.FindParmFiles) (objs []interfaces.FileInfo, total int64, err error) {
        objs = make([]interfaces.FileInfo, 0)
        err = readEngine.
            Desc("created").
            Limit(parm.PageSize, parm.Page*parm.PageSize).
            Find(&objs)
        return
    }
```

领域服务 再我自己的项目实践中，领域服务主要用来完成本服务与其他服务之间的数据交流，比如调用其他HTTP接口、RPC请求等。

5.crossutting、interfaces这些都不是DDD经典分层里的，主要是用来方便实现DDD所增加和分离出来的一些接口、和基础概念实现。interfaces层 这一层，在标准的分层里其实没有的。这里主要将仓储功能与领域层进一步解耦，利用依赖注入的方式来反转。该层主要定义ORM实体类，也就是数据库表结构的映射，再将持久化的功能抽象成接口的定义。具体的实现，可以有多种，这样持久化的方式可以是彼此无依赖。
  
以上即为采用经典四层来实践领域驱动设计最简单的方式，DDD的六边形分层，理解了经典分层，也很容易实践的。

## 三、进阶之CQRS+ES

![CQRS](./asset/02.png)
命令查询职责分离，是由Betrand Meyer（Eiffel语言之父，OCP提出者）提出来的。命令(Command):不返回任何结果(void)，但会改变对象的状态。查询(Query):返回结果，但是不会改变对象的状态，对系统没有副作用。在我的实践过程中，其实还是让命令返回了一些主键之类的。项目采用了 github.com/looplab/eventhorizon 库作为具体的实现，并进行了二次封装，在应用层和表现层，几乎不用写太多代码了。由CQRS的概念可知，系统很方便对数据库进行读写分离。

ES事件溯源：在CQRS中，每一个确定的命令操作，不论成功还是失败，只要执行之后就产生相应的事件（Event）。这样只需要把所有系统运行的Event，以及其产生时的数据记录起来，这就是系统的历史记录了，并且能够方便的回滚到某一历史状态。Event Sourcing就是用来进行存储和管理事件的。

接下来采用DDD+CQRS来完成上面的文件管理业务。
查询可以从应用层进行分离，直接操作仓储曾获取业务不是特别复杂的查询，这与没有引入CQRS的代码可以保持一致。重点关注——命令面即可。

1.首先抽象Command。诚如项目中domain/files/commands.go 里面AddFileCmd 所示，①实现eh.Command接口即可，这里的AggregateID需要特别留意，此处的ID代表的是执行该命令的聚合根,可以留空；②实现infrastructure/ddd/ddd.Validator接口，对参数有效性进行校验。③eh.RegisterCommand向eventhorizon里注册该命令。

```go
    func init() {
        eh.RegisterCommand(func() eh.Command { return &AddFileCmd{} })
    }

    const (
        AddFileCmdType eh.CommandType = "AddFileCmd"
    )

    //添加文件,所需要的一些参数
    type AddFileCmd struct {
        FileName    string `json:"file_name"`
        ContentType string `json:"content_type"`
        Size        int    `json:"size"`
        FileBody    []byte `json:"file_body" eh:"optional"`
    }

    func (c *AddFileCmd) AggregateID() uuid.UUID {
        return uuid.Nil
    }
    func (c *AddFileCmd) CommandType() eh.CommandType     { return AddFileCmdType }
    func (c *AddFileCmd) AggregateType() eh.AggregateType { return "" }
    func (c *AddFileCmd) Verify() error                   { return nil }
```

2.命令抽象出来之后，可以串一下表现层/应用层对该命令的消费。先看直接消费者，应用层。比较宽泛的封装在infrastructure/ddd/app_common.go，将表现层接受来的请求主体，根据命令Key转换为命令，并且进行参数校验、再进行命令总线进行发布。只需要在const里做好cmdtype和表现层cmdkey的转换即可。 具体业务代码见：application/files_app.go  AddFileCommand函数。 表现层，再消费应用层，并尽量做到对领域层的无感知。采用了CQRS对于Command型只需要返回HTTP请求结果即可。

```go
    func AddFileCommand(parm AddFileForm) (err error) {
        //looplab框架根据命令类型，创建命令实例
        cmd, err := eh.CreateCommand(files.AddFileCmdType)
        if err != nil {
            err = fmt.Errorf("could not create command: %w", err)
            return err
        }

        //命令进行校验
        if vldt, ok := cmd.(ddd.Validator); ok {
            err = vldt.Verify()
            if err != nil {
                return
            }
        }
        //通过命令总线，将命令发布出去，至于谁订阅了该命令则不关系
        if err = bus.HandleCommand(context.Background(), cmd); err != nil {
            return err
        }
        return
    }
```

3.命令的Handler，一般都有相关领域上下文的聚合根来承担。聚合根实现 eh.Aggregate接口即可，HandleCommand方法实现方式也基本类似的。具体业务代码的实现，见代码注释不做赘述。遵循面向对象的几大基本原则即可，单一职责（可以多一些实体类）、开放封闭。这样代码变更，或者与他人协作，也只改动单一的源代码文件。 领域层会直接消费基础设施层，一些常用的帮助工具代码先忽略，主要关注即infrastructure/repoFac.go的RepoFac能提供哪些持久化能力即可。

```go
    //Command异步执行，不需要返回值的
    //识别是哪个CMD的请求，并取出相关数据，转交Entity进行业务表达。
    func (a *filesAggregate) HandleCommand(ctx context.Context, cmd eh.Command) (err error) {
        switch cmd := cmd.(type) {
        case *AddFileCmd:
            ov := interfaces.FileInfo{
                Id:          uuid.New().String(),
                FileName:    cmd.FileName,
                Size:        cmd.Size,
                ContentType: cmd.ContentType,
            }
            en := newfileEnByOV(ov)
            err = en.AddFile(cmd.FileBody)
            if err != nil {
                logs.Error("新增文件出错：%s ", err.Error())
            }
            // 文件新增成功之后，激活领域事件，并发布事件，由事件订阅者继续完成后续操作
            a.AppendEvent(FileAddedEvent, en.FileInfo, time.Now())
            bus.RaiseEvents(context.Background(), a.AggregateBase, 1)
        default:
            err = fmt.Errorf("couldn't handle command")
        }
        return
    }
```

4.事件以及事件的Handler，某个命令处理完毕之后，也即某个事件发生了，有可能需要短信通知、邮件通知等等。事件订阅者（filesEventHandler）继续后续的逻辑。

```go
    // 事件订阅者收到事件发生信号，以及附带的数据
    func (a *filesEventHandler) HandleEvent(ctx context.Context, event eh.Event) (err error) {
        switch evtData := event.Data().(type) {
        case *interfaces.FileInfo:
            println("订阅者发现了，文件新增成功！ 文件ID：", evtData.Id, "订阅者继续业务逻辑...")
        }

    return fmt.Errorf("couldn't handle Event")
}
```

以上即为CQRS模式下处理CMD的一般步骤和实践，具体某个业务发生之后，即可激活领域事件，将所有的事件依次持久化下来，便能够进行事件溯源（ES），暂时不深入讨论，先实践最基本的微服务拆解和实际入手代码编写。
