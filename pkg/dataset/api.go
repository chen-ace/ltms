package dataset

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"llm_training_management_system/pkg/ltms_config"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var httpConfig *ltms_config.HttpConfig

func BindRouter(config *ltms_config.HttpConfig, r *gin.Engine) {
	httpConfig = config
	r.POST("/dataset/upload", handleUpload)
	r.GET("/dataset/download", download)
	r.GET("/dataset/ls", ls)
	r.GET("/dataset/head", func(c *gin.Context) {
		fileName := c.Query("fileName")
		n := c.Query("n")
		parseInt, err := strconv.ParseInt(n, 10, 64)
		if err != nil || parseInt > 100 {
			parseInt = 20
		}
		c.Set("data", map[string]string{"text": head(fileName, int(parseInt))})
	})
}

func download(c *gin.Context) {
	fileName := c.Query("fileName")
	dst := filepath.Join(httpConfig.DataDir, fileName)
	//打开文件
	_, errByOpenFile := os.Open(dst)
	//非空处理
	if errByOpenFile != nil {
		c.Set("err", errors.New("资源不存在"))
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(dst)
	return
}

func handleUpload(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("文件上传错误：", err)
		c.Set("err", err)
		return
	}
	filename := file.Filename
	log.Println(filename)
	// 将八进制转义序列转换为普通字符串
	filename, err = strconv.Unquote(`"` + filename + `"`)
	if err != nil {
		fmt.Println("解析错误:", err)
		c.Set("err", err)
		return
	}
	dst := filepath.Join(httpConfig.DataDir, filename)
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.Set("resp", "File Upload finished")

	c.Set("data", map[string]string{"filename": filename})
}

func ls(c *gin.Context) {
	fmt.Println("start ls")
	files, err := os.ReadDir(httpConfig.DataDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
			fmt.Println(file.Name())
		}

	}
	c.Set("data", filenames)
}

func head(fileName string, n int) string {
	dst := filepath.Join(httpConfig.DataDir, fileName)
	// 打开文件
	file, err := os.Open(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建一个带缓冲的读取器
	scanner := bufio.NewScanner(file)

	var result string
	// 读取文件的前5行
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			fmt.Println(scanner.Text())
			result = result + "\n" + scanner.Text()
		} else {
			// 如果在读取5行之前已经到达文件末尾，则退出循环
			break
		}
	}

	// 检查是否有非EOF错误发生
	if err := scanner.Err(); err != nil {
		log.Println("文件读取错误", err)
	}
	return result
}
