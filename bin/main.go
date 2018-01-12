package main

import "github.com/xiaoxiaoyijian/cjiebago"
import "log"

func main() {
	jieba := cjiebago.NewJieba("/home/steven/var/code/go/src/github.com/xiaoxiaoyijian/cjiebago/dict/jieba.dict.utf8", "/home/steven/var/code/go/src/github.com/xiaoxiaoyijian/cjiebago/dict/hmm_model.utf8", "")
	defer jieba.Close()
	words := jieba.CutAll("小明硕士毕业于中国科学院计算所，后在日本京都大学深造")
	log.Println(words)
}
