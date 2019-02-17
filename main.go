package main

//例子程序如下：
/*
type Data struct{
}
func xxx()　*Data{
	var c Data
	return &c
}
func main(){
	xxx();
}
*/
//go run -gcflags "-m -l" main.go //-m表示进行内存分析　-l表示避免程序内联（避免程序优化）
//这样我们可以看到类似如下的输出：
//........
//./main.go: &c escapes to heap
//./main.go: moved to heap :c
//.........
func main() {

	//=========包================
	//mypkg.Mypkg()

	//=========变量================
	//myvar.Myvar()

	//=========基本语法================
	//mygolang.Mygolang()

	//=========集合================
	//mycollection.Mycollection()
	//mycollection.MySort()

	//=========函数================
	//myfunc.Myfunc()

	//=========延迟================
	//myfunc.Mydefer()

	//=========错误处理================
	//myexception.Myexception()

	//=========类================
	//mystruct.Mystruct()

	//=========方法================
	//mymethod.Mymethod()
	//mymethod.Finallizefunc()

	//=========接口================
	//myintf.Myintf()
	//myintf.Myintf2()
	//myintf.Myintf3()
	//myintf.TestAAABBBCCC()
	//myref.Myref()

	//=========基本api================
	//myapi.MyReg()
	//myapi.Mystr()
	//myapi.Mytime()
	//myapi.MyBase64()
	//myapi.MyJson()
	//myapi.MyLineFilter()
	//myapi.MyRandom()
	//myapi.MySha1()
	//myapi.MyString2()
	//myapi.MyUrlParser()

	//=========并发================
	//mygoroutine.Mygoroutine()
	//mygoroutine.Myselect()
	//myruntime.Myruntime()
	//mysync.MyCountDownLunch()
	//mysync.MyMutx()
	//mysync.MyRWLock()
	//mysync.MyAtomic()
	//mysync.MyOnce()
	//mysync.MyOnce2()
	//mysync.MyPooltest()
	//mytime.MyTicker()
	//mytime.MyTimeout()
	//mytime.MyTimer()

	//=========pprof================
	//mypprof.Testpprof()
	//mypprof.Testmypprof2()
	//=========文件io================
	//myfile.Myfile()
	//myfile.MyFile2()

	//=========命令行参数================
	//mycmdline.Mycmdline()

	//=========网络io================
	//mytcp.MyTcpServer()
	//mytcp.MyTcpClient()
	//onetoone.MyUdpServer()
	//onetoone.MyUdpCli()
	//multi.MyMultiServer()
	//multi.MyMultiCli()
	//broadcast.MyBCServer()
	//broadcast.MyBCCli()

	//=========网络io-http相关================
	//mynet.MyHttp1()
	//mynet.MyHttp2()
	//mynet.MyHttp21()
	//mynet.MyHttp22()
	//mynet.MyHttp3()
	//mynet.MyHttp31()
	//mynet.Myhttpcli()

	//myupload.MyUpServer()
	//myupload.MyUpCli()

	//myssl.MySsl()
	//myssl.MySslCli()

	//myhttp2.MyHttp2Server()
	//myhttp2.MyHttp2Cli()

	//myws.MyWsServer()

	//=========系统================
	//mysys.MyExec()
	//mysys.MyExec2()
	//mysys.MySignal()
	//mysys.MySyscall()
	//mysys.MySyscall2()

	//=========go调用c================
	//goandc.Mygoandc()

	//=========虚拟化================
	//mycontainer.MyUts()
	//mycontainer.MyUts2()
	//mycontainer.MyPid()
	//mycontainer.MyCgrouptest()(?P<命名>子表达式)
	//var digitsRegexp = regexp.MustCompile(`^[ac-inform].*`)
	//st := "ac-inform.hb-ac-test00.ablecloud.log.ERROR.20180910-171824.10857"
	//snatch := digitsRegexp.FindStringSubmatch(st)
	//fmt.Printf("%v", snatch)

}
