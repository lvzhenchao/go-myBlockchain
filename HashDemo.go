package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func calculateHash(toBeHashed string) string{
	hashInBytes := sha256.Sum256([]byte(toBeHashed))//哈希运算，转成字节切片，返回的是字节数组
	hashInStr   := hex.EncodeToString(hashInBytes[:])//转成字符串

	log.Printf("%s, %s", toBeHashed, hashInStr)//记录一下日志
	return hashInStr

}

func main() {
	calculateHash("test1")
	calculateHash("test1")
	calculateHash("test2")
}

/*
输出：
2019/06/18 18:36:13 test1, 1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014
2019/06/18 18:36:13 test1, 1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014
2019/06/18 18:36:13 test2, 60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752

  

*/
