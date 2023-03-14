package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"main.go/core"
	"main.go/global"
	"main.go/initialize"
	"os"
	"strconv"
)

func main() {

	//initPhoneNumber(1234567, 2)

	//WriteCsv()
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	core.RunWindowsServer()           //设置路由,启动端口监听

	//测试git更新

}

func readTXT() {
	//打开文件
	file, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}

func writerTxt() {
	//创建一个新文件，写入内容
	filePath := "./output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	//写入内容
	str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt文件要看到换行效果要用 \r\n
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()

}

func initPhoneNumber(start int, diff int) {
	//41
	s := []int{903, 905, 906, 909, 960, 961, 962, 963, 964, 910, 911, 912, 913, 914, 915, 916, 917, 918, 919, 920, 921, 922, 923, 924, 925, 926, 927, 928, 929, 930, 931, 937, 980, 981, 982, 983, 984, 985, 986, 987, 988}

	//创建一个新文件，写入内容
	filePath := "./output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	writer := bufio.NewWriter(file)
	var i = start
	for i = start; i < (start + diff); i++ {
		for _, v := range s {
			number := strconv.Itoa(v) + strconv.Itoa(i) + "\n"
			//写入时，使用带缓存的 *Writer
			writer.WriteString(number)
		}

	}

	writer.Flush()
}

func WriteCsv() {
	//创建文件
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "刘备", "23"},
		{"2", "张飞", "23"},
		{"3", "关羽", "23"},
		{"4", "赵云", "23"},
		{"5", "黄忠", "23"},
		{"6", "马超", "23"},
	}
	//写入数据
	w.WriteAll(data)
	w.Flush()
}

func ReadCsv() {
	//准备读取文件
	fileName := "test.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
	}
	fmt.Println("\n---------------------------\n")
	// 针对小文件，也可以一次性读取所有的文件。注意，r要重新赋值，因为readall是读取剩下的
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	for _, row := range content {
		fmt.Println(row)
	}
}
