package tree

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func OrganizePictureByDate(srcDirList []string, destDir string) {
	for _, srcDir := range srcDirList {
		organizeOneDirPicture(srcDir, destDir)
	}
}

func organizeOneDirPicture(srcDir string, destDir string) {
	file, err := os.Stat(srcDir)
	if err != nil {
		fmt.Printf("查看文件夹出现异常,srcDir=%v,err=%v\n", srcDir, err)
		return
	}
	if file.IsDir() {
		subFiles, _ := ioutil.ReadDir(srcDir)
		for _, subFile := range subFiles {
			organizeOneDirPicture(srcDir+"\\"+subFile.Name(), destDir)
		}
	} else {
		suffix := ".jpg"
		if !strings.HasSuffix(file.Name(), suffix) {
			fmt.Printf("不是jpg文件,fileName=%v\n", file.Name())
			return
		}
		creationDate, err := getCreationDate(file)
		if err != nil {
			fmt.Printf("获取文件所属日期失败,fileName=%v,err=%v\n", file.Name(), err)
			return
		}
		if len(creationDate) <= 6 {
			fmt.Printf("文件所属日期错误,fileName=%v,creationDate=%v\n", file.Name(), creationDate)
			return
		}
		month := creationDate[:6]
		monthDirPath := destDir + month
		_, err = os.Stat(monthDirPath)
		if os.IsNotExist(err) {
			os.Mkdir(monthDirPath, os.ModePerm)
		}
		srcFile, _ := ioutil.ReadFile(srcDir)
		destFilepath := monthDirPath + "\\" + creationDate + "_" + strconv.Itoa(rand.Int()) + suffix
		err = ioutil.WriteFile(destFilepath, srcFile, 0644)
		if err != nil {
			fmt.Printf("复制文件失败,srcFile=%v,err=%v\n", srcFile, err)
		}
	}

}

func getCreationDate(file os.FileInfo) (string, error) {
	// 默认为文件本身的创建时间
	creationDate := time.Unix(file.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds()/1e9, 0).Format("20060102")

	// D:\BaiduNetdiskDownload\来自：MHA-AL00\DCIM\Camera  IMG_20180101_162404.jpg
	// D:\BaiduNetdiskDownload\来自：iPhone   2014-04-28 085009.jpg
	// D:\BaiduNetdiskDownload\来自：iPad   2014-02-11 214138.jpg
	// D:\BaiduNetdiskDownload\来自：H30-T00\DCIM\Camera   IMG_20171005_132841.jpg
	// D:\BaiduNetdiskDownload\meta9备份  IMG_20200130_213650.jpg
	// D:\BaiduNetdiskDownload\来自：微信备份\微信图片备份   mmexport1559126324147.jpg
	// D:\BaiduNetdiskDownload\美图   QQ图片20140408134347.jpg
	// D:\BaiduNetdiskDownload\元宝成长图片   2019-7-22 113335 2.jpg   2019-11-1 91040 7.jpg
	fileName := file.Name()
	if strings.HasPrefix(fileName, "IMG_") {
		fromIndex := strings.Index(fileName, "_")
		toIndex := strings.LastIndex(fileName, "_")
		creationDate = fileName[(fromIndex + 1):toIndex]
	} else if strings.HasPrefix(fileName, "QQ图片") {
		creationDate = strings.ReplaceAll(fileName, "QQ图片", "")[0:8]
	} else if strings.HasPrefix(fileName, "mmexport") {
		timeInMills, err := strconv.ParseInt(strings.ReplaceAll(fileName, "mmexport", "")[0:13], 10, 64)
		if err == nil {
			creationDate = time.UnixMilli(timeInMills).Format("20060102")
		} else {
			fmt.Printf("文件名转换成日期出现异常,fileName=%v,err=%v\n", fileName, err)
			return creationDate, err
		}
	} else {
		toIndex := strings.Index(fileName, " ")
		if toIndex > 0 {
			creationDate = fileName[0:toIndex]
			firstDash := strings.Index(creationDate, "-")
			lastDash := strings.LastIndex(creationDate, "-")
			if toIndex > 0 && firstDash > 0 && lastDash > 0 {
				year := creationDate[0:firstDash]
				month := creationDate[(firstDash + 1):lastDash]
				day := creationDate[(lastDash + 1):]
				creationDate = year
				if len(month) == 1 {
					month = "0" + month
				}
				creationDate += month
				if len(day) == 1 {
					day = "0" + day
				}
				creationDate += day
			} else {
				return creationDate, errors.New("不是要重新整理的照片")
			}
		}
	}

	return creationDate, nil
}
