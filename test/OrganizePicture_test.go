package algorithm_test

import (
	"mojiayi-golang-algorithm/tree"
	"testing"
)

func TestOrganizePicture(t *testing.T) {
	destDir := "D:\\照片备份\\"
	// D:\BaiduNetdiskDownload\来自：MHA-AL00\DCIM\Camera  IMG_20180101_162404.jpg
	// D:\BaiduNetdiskDownload\来自：iPhone   2014-04-28 085009.jpg
	// D:\BaiduNetdiskDownload\来自：iPad   2014-02-11 214138.jpg
	// D:\BaiduNetdiskDownload\来自：H30-T00\DCIM\Camera   IMG_20171005_132841.jpg
	// D:\BaiduNetdiskDownload\meta9备份  IMG_20200130_213650.jpg
	// D:\BaiduNetdiskDownload\来自：微信备份\微信图片备份   mmexport1559126324147.jpg
	// D:\BaiduNetdiskDownload\美图   QQ图片20140408134347.jpg
	// D:\BaiduNetdiskDownload\元宝成长图片   2019-7-22 113335 2.jpg

	srcDirList := []string{"D:\\BaiduNetdiskDownload\\元宝成长图片", "D:\\BaiduNetdiskDownload\\美图", "D:\\BaiduNetdiskDownload\\来自：微信备份\\微信图片备份", "D:\\BaiduNetdiskDownload\\meta9备份",
		"D:\\BaiduNetdiskDownload\\来自：H30-T00\\DCIM\\Camera", "D:\\BaiduNetdiskDownload\\来自：iPad", "D:\\BaiduNetdiskDownload\\来自：iPhone",
		"D:\\BaiduNetdiskDownload\\来自：MHA-AL00\\DCIM\\Camera"}
	tree.OrganizePictureByDate(srcDirList, destDir)
}
