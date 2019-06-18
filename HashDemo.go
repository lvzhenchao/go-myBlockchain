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
