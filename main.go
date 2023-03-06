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
	"sync"
)

var a int
var once sync.Once

func main() {

	readCsv()
	//readTXT2w()

	//initPhoneNumber(1234567, 2)
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	core.RunWindowsServer()           //设置路由,启动端口监听

	//测试git更新

}

func readCsv() {

	//创建一个新文件，写入内容
	filePath := "./ws3000.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()

	// Open the file
	csvfile, err := os.Open("2w-发送到-筛性别年龄(全部数据)-2023_3_6.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("Record has %d columns.\n", len(record))
		//city, _ := iconv.ConvertString(record[2], "gb2312", "utf-8")

		if record[3] == "女" {
			fmt.Printf("%s %s %s %s \n", record[0], record[1], record[2], record[3])

			writer.WriteString(record[0] + "\n")

		}

	}

	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()

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

func readTXT2w() {
	//打开文件
	file, err := os.Open("./俄罗斯20w.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	var name = 0
	filePath := "./output.txt"
	files, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	//写入内容
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(files)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	for {
		if name == 20000 {
			break
		}
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}

		writer.WriteString(str)

		name++
	}
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
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

func writerTxtOneLine(str string) {
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
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	writer.WriteString(str)

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
