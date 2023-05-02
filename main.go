package main

import (
	"bufio"
	"main.go/core"
	"main.go/global"
	"main.go/initialize"
	"math/rand"
	"time"

	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

var a int
var once sync.Once

func main() {

	//readCsv()
	//End4Number()
	//readCvsV2()     //筛选 女性
	//readTXT2w()
	//initPhoneNumber(1234567, 2)
	//readCsvDays()
	//startmoxikenumber()

	//sendWhatsappMessage() //发送 ws 消息
	//creatnumber()
	//WriteXLSX()
	//ReadXlsx()
	//FenGeShuJu()

	//网站初始化

	//
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	core.RunWindowsServer()           //设置路由,启动端口监听

	//测试git更新

}

func readCsv() {

	//创建一个新文件，写入内容
	filePath := "./10000.txt"
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

//func WriteXLSX() {
//
//	f := excelize.NewFile()
//	// 创建一个工作表
//	index, _ := f.NewSheet("Sheet1")
//	// 设置单元格的值
//	f.SetCellValue("Sheet1", "A2", 100)
//	f.SetCellValue("Sheet1", "B2", 100)
//	// 设置工作簿的默认工作表
//	f.SetActiveSheet(index)
//	// 根据指定路径保存文件
//	if err := f.SaveAs("俄罗斯(女)2023-04-04+6K.xlsx"); err != nil {
//		println(err.Error())
//	}
//
//}

func ReadXlsx() {
	f, err := excelize.OpenFile("1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	//创建文本

	WF := excelize.NewFile()
	// 创建一个工作表
	index, _ := WF.NewSheet("Sheet1")
	// 设置单元格的值
	//WF.SetCellValue("Sheet1", "A2", 100)
	//WF.SetCellValue("Sheet1", "B2", 100)
	//// 设置工作簿的默认工作表
	WF.SetActiveSheet(index)
	//// 根据指定路径保存文件
	if err := WF.SaveAs("2000女.xlsx"); err != nil {
		println(err.Error())
	}

	// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//println(cell)

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for num, row := range rows { //行数

		if num < 4000 {
			continue
		}

		print(num, "sss", "\t")
		for index, colCell := range row { //某一行 所有列

			if index == 1 {

				//当前时间戳
				timestamp := time.Now().Unix()

				rantimes := rand.Intn(15*60*60) + (10 * 60) //(88-15 )+15
				timestamp = timestamp - int64(rantimes)
				// 再格式化时间戳转化为日期
				datetime := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")

				//日志输出
				//print(datetime, "\t")

				// 设置单元格的值
				WF.SetCellValue("Sheet1", "B"+strconv.Itoa(num), datetime)
				//WF.SetCellValue("Sheet1", "B2", 100)

			}

			if index == 0 {
				WF.SetCellValue("Sheet1", "A"+strconv.Itoa(num), colCell)
			}

			if index == 2 {
				WF.SetCellValue("Sheet1", "C"+strconv.Itoa(num), "1")
			}
			//print(colCell, "\t")

			// 设置工作簿的默认工作表
			WF.SetActiveSheet(index)
			// 根据指定路径保存文件
			if err := WF.SaveAs("2000女.xlsx"); err != nil {
				println(err.Error())
			}

		}
		println()
	}

}

func readCsvDays() {

	//创建一个新文件，写入内容
	filePath := "./3000.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()

	// Open the file
	csvfile, err := os.Open("22.csv")
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

		if record[2] == "1" {
			//fmt.Printf("%s %s %s %s \n", record[0], record[1], record[2], record[3])

			writer.WriteString(record[0] + "\n")

		}

	}

	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()

}

func readCsvDaysData() {

	//创建一个新文件，写入内容
	filePath := "./1.csv"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()

	// Open the file
	csvfile, err := os.Open("./1.csv")
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

		if record[2] == "1" {
			//fmt.Printf("%s %s %s %s \n", record[0], record[1], record[2], record[3])

			writer.WriteString(record[0] + "\n")

		}

	}

	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()

}

func End4Number() {
	//数据表
	numbers_1 := [10]int{1, 3, 5, 4, 9, 0, 7, 8, 6, 2}
	numbers_2 := [10]int{0, 7, 8, 4, 9, 6, 2, 1, 3, 5}
	numbers_3 := [10]int{6, 2, 1, 8, 4, 9, 3, 5, 0, 7}
	numbers_4 := [10]int{9, 1, 8, 5, 7, 6, 2, 3, 4, 0}
	//创建一个新文件，写入内容
	filePath := "./随机4位尾数.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	totalCount := 0
	/*以下为三重循环*/

	for _, s := range numbers_1 {
		for _, j := range numbers_2 {
			for _, k := range numbers_3 {
				for _, m := range numbers_4 {

					/*确保 i 、j 、k 三位互不相同*/
					if s != k && s != j && j != k {
						totalCount++
						stri := strconv.Itoa(s)
						strj := strconv.Itoa(j)
						strk := strconv.Itoa(k)
						strm := strconv.Itoa(m)
						strall := stri + strj + strk + strm
						writer.WriteString(strall + "\n")
						//fmt.Printf("%s \n", strall)

					}
				}

			}
		}
	}

	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
}

func readCvsV2() {

	//创建一个新文件，写入内容
	filePath := "./5000.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()

	// Open the file
	csvfile, err := os.Open("List.csv")
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

		if len(record) != 0 {

			valus := strings.Split(record[0], ";")
			if valus[1] == "Female" {
				fmt.Printf("%s %s \n", valus[0], valus[1])
				writer.WriteString(valus[0] + "\n")
			}

		}

	}

	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()

}

func readTXT() {
	//打开文件
	file, err := os.Open("./2w.txt")
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
	file, err := os.Open("./57w.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	var name = 0
	filePath := "./1万ws女性活跃.txt"
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
		if name <= 30000 {
			name++
			continue
		}

		if name > 100000 {
			break
		}
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}

		writer.WriteString(str + "\r\n")

		name++
	}
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
	fmt.Println("文件读取结束...")
}

func FenGeShuJu() {
	//打开文件
	file, err := os.Open("./57w.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	var name = 0

	//写入内容
	//写入时，使用带缓存的 *Writer
	filePath := "./5w_女_30活跃.txt"
	files, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	writer := bufio.NewWriter(files)

	//第二个分割
	//filePath2 := "./b渠道.txt"
	//files2, err := os.OpenFile(filePath2, os.O_WRONLY|os.O_CREATE, 0666)
	//writer2 := bufio.NewWriter(files2)

	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	for {

		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}

		//  79910676551    79910676540     79910676529   79910676522

		//  79910676502   79910676398    79910676394     79910397806

		if name > 80004 && name < 120004 {

			//if name == 3500 {
			//	writer.WriteString("79910676502")
			//	writer.WriteString("\r\n")
			//}
			//
			//if name == 4000 {
			//	writer.WriteString("79910676398")
			//	writer.WriteString("\r\n")
			//}
			//
			//if name == 5000 {
			//	writer.WriteString("79910676394")
			//	writer.WriteString("\r\n")
			//}
			//
			//if name == 5500 {
			//	writer.WriteString("79910397806")
			//	writer.WriteString("\r\n")
			//}
			writer.WriteString(str)
			//writer.WriteString("\r\n")

		}

		if name > 80004 {
			break
		}

		//if name < 5000 {
		//
		//	if name == 800 {
		//		writer2.WriteString("79910676551")
		//		writer2.WriteString("\r\n")
		//	}
		//
		//	if name == 1400 {
		//		writer2.WriteString("79910676540")
		//		writer2.WriteString("\r\n")
		//	}
		//
		//	if name == 2400 {
		//		writer2.WriteString("79910676529")
		//		writer2.WriteString("\r\n")
		//	}
		//
		//	if name == 3500 {
		//		writer2.WriteString("79910676522")
		//		writer2.WriteString("\r\n")
		//	}
		//	writer2.WriteString(str)

		//}

		name++
	}
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
	//writer2.Flush()
	fmt.Println("文件读取结束...")
}

func writerTxt() {
	//创建一个新文件，写入内容
	filePath := "./2w.txt"
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
	filePath := "./2w.txt"
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
	filePath := "./2w.txt"
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

// 从运营商号段中挑选合适的 号段
func startmoxikenumber() {

	s := []string{"7936", "7983", "7986", "7901", "7916", "7985", "7925", "7926", "7917", "7985", "7936"}
	//打开文件
	file, err := os.Open("./俄罗斯运营商号段.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()

	//创建一个新文件，写入内容
	filePath := "./莫斯科.txt"
	files, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer files.Close()
	writer := bufio.NewWriter(files)

	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		for _, v := range s {

			if strings.HasPrefix(str, v) {
				//符合条件 写入统计表
				fmt.Print(str)
				writer.WriteString(str)
				break
			}
		}

	}

	writer.Flush()

}

func sendWhatsappMessage() {

	var msg string = "hjaha"
	var tel string = "+639289876"

	apiurl := "https://api.ultramsg.com/instance41376/messages/chat"
	data := url.Values{}
	data.Set("token", "warning44FF")
	data.Set("to", tel)
	data.Set("body", msg)

	payload := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", apiurl, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}

// 根据号段生成数据
func creatnumber() {

	//读取号段
	//打开文件
	file_start, err := os.Open("./莫斯科.txt")

	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file_start.Close()

	//读取 随机号码尾号

	file_end, err := os.Open("./随机4位尾数.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file_end.Close()

	//最终写入的文件
	filePath := "./最终数据.txt"
	file_input, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file_input.Close()

	writer := bufio.NewWriter(file_input)

	//创建一个 *Reader ， 是带缓冲的
	reader_start := bufio.NewReader(file_start)
	reader_end := bufio.NewReader(file_end)

	//读入 到 切片

	var start []string
	var end []string

	for {
		st1, err := reader_start.ReadString('\n') //读到一个换行就结束

		if err == io.EOF {
			//io.EOF 表示文件的末尾
			break
		}

		start = append(start, st1)

	}

	for {

		str2, errs := reader_end.ReadString('\n') //读到一个换行就结束
		if errs == io.EOF {                       //io.EOF 表示文件的末尾
			break
		}

		end = append(end, str2)

	}

	var count = 0

	for _, s := range start {

		for _, e := range end {

			if count >= 2000000 {
				goto endfor
			}

			number := s + e
			number = strings.Replace(number, " ", "", -1)
			// 去除换行符
			number = strings.Replace(number, "\n", "", -1)
			number = number + "\n"
			writer.WriteString(number)

			count++

		}
		writer.Flush()
	}

endfor:
	fmt.Printf("结束")

	writer.Flush()

}
