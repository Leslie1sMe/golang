//golang文件操作练习记录
package main

//输入流:数据从数据源（文件）到程序（内存）的路径
//输出流:数据从程序（内存）到数据源（文件）的路径
type countType struct {
	spaceNum int
	wordNum  int
	num int
	otherNum int
}

func main() {
	//1. 打开文件
	//file, error := os.Open("/Users/leslie/go/src/go_code/grammar/test.txt")
	//if error != nil {
	//	log.Fatal(error)
	//}
	//fmt.Printf("file=%v", file)
	//err :=file.Close()
	//file.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//2. 读文件操作
	//file, error := os.Open("/Users/leslie/go/src/go_code/grammar/test.txt")
	//if error != nil {
	//	log.Fatal(error)
	//}
	//
	////结束时关闭file句柄，防止内存泄漏
	//defer file.Close()
	//
	//reader := bufio.NewReader(file)
	//for {
	//	str, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		log.Fatal(err)
	//	}
	//	fmt.Print(str)
	//}
	//3. 一次行读取文件(适合文件不太大的情况)
	//file :="/Users/leslie/go/src/go_code/grammar/test.txt"
	//reader,error:=ioutil.ReadFile(file)
	//if error != nil{
	//	log.Fatal(error)
	//}
	////输出byte切片
	//fmt.Println(reader)
	////输出字符
	//fmt.Println(string(reader))
	//4. 创建一个文件输入(缓冲读取需要使用Flush)
	//file, error := os.OpenFile("/Users/leslie/go/src/go_code/grammar/test2.txt", (os.O_WRONLY|os.O_CREATE), 0660)
	//if error != nil {
	//	log.Fatal(error)
	//}
	//
	//defer file.Close()
	//
	//str := "hello,Leslie"
	//writer := bufio.NewWriter(file)
	//for x := 0; x < 5; x++ {
	//	writer.WriteString(str)
	//}
	////使用bufio.Writer数据会先写入缓存，所以需要使用Flush才可以真正写入到文件。
	//writer.Flush()
	//5. 打开一个已经有的文件，将原来的内容覆盖成新的内容十句"hello，Leslie"
	//file, error := os.OpenFile("/Users/leslie/go/src/go_code/grammar/test.txt", os.O_TRUNC|os.O_WRONLY, 660)
	//if error != nil {
	//	log.Fatal(error)
	//}
	//writer := bufio.NewWriter(file)
	//for x := 0; x < 10; x++ {
	//	writer.WriteString("hello,Leslie")
	//}
	//writer.Flush()
	//file.Close()
	//6. 追加操作
	//file, error := os.OpenFile("/Users/leslie/go/src/go_code/grammar/test.txt", os.O_APPEND|os.O_WRONLY, 660)
	//if error != nil {
	//	log.Fatal(error)
	//}
	//writer := bufio.NewWriter(file)
	//writer.WriteString("ABC,Le")
	//writer.Flush()
	//file.Close()
	//7.打开一个存在的文件，将原来的内容读出在终端，然后追加123....
	//file, error := os.OpenFile("/Users/leslie/go/src/go_code/grammar/test.txt", os.O_RDONLY|os.O_APPEND|os.O_WRONLY, 660)
	//if error == io.EOF {
	//	log.Fatal(error)
	//}
	//reader:=bufio.NewReader(file)
	//fmt.Println(reader)
	//writer := bufio.NewWriter(file)
	//writer.WriteString("123131313")
	//writer.Flush()
	//file.Close()
	//8. 两个文件，将一个文件的内容写到另一个文件当中去
	//file := "/Users/leslie/go/src/go_code/grammar/test.txt"
	//reader, error := ioutil.ReadFile(file)
	//if error != nil {
	//	log.Fatal(error)
	//}
	//file1 := "/Users/leslie/go/src/go_code/grammar/test1.txt"
	//ioutil.WriteFile(file1,reader,660)
	//9. 判断文件是否存在
	//fileInfo,error:=os.Stat("/Users/leslie/go/src/go_code/grammar/te2st.txt")
	//if error != nil {
	//	if os.IsNotExits(error){
	//return false
	//}else{
	// return true
	// }
	//}
	//fmt.Println(fileInfo)
	//10.文件copy
	//src, error := os.Open("/Users/leslie/Downloads/cpts_1638_dlq/images/2.jpg")
	//if error != nil {
	//	log.Fatal(error)
	//}
	//file := bufio.NewReader(src)
	//defer src.Close()
	//dest := "/Users/leslie/go/src/go_code/grammar/golang.jpg"
	//file2, error := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0660)
	//if error != nil {
	//	log.Fatal(error)
	//}
	//descFile := bufio.NewWriter(file2)
	//defer file2.Close()
	//res, error := io.Copy(descFile, file)
	//if error !=nil {
	//	log.Fatal(error)
	//}
	//fmt.Println(res)
	//11.统计一下英文数字空格等的个数
	//filename := "/Users/leslie/go/src/go_code/grammar/test.txt"
	//file, error := os.Open(filename)
	//if error != nil {
	//	log.Fatal(error)
	//}
	//reader := bufio.NewReader(file)
	//var fileType countType
	//for {
	//	str, error := reader.ReadString('\n')
	//	if error == io.EOF {
	//		break
	//	}
	//
	//	for _, v := range str {
	//
	//		switch {
	//		case v >= 'a' && v <= 'z':
	//			fallthrough
	//		case v >='A' && v <= 'Z':
	//			fileType.wordNum++
	//		case v == ' ' || v == '\t':
	//			fileType.spaceNum++
	//		case v >= '1' && v <= '9':
	//			fileType.num++
	//		default:
	//			fileType.otherNum++
	//
	//		}
	//	}
	//}
	//fmt.Printf("字母个数=%v 空格个数=%v 数字个数=%v 其他个数=%v",fileType.wordNum,fileType.spaceNum,fileType.num,fileType.otherNum)

}



