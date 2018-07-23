package mapreduce

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	kv1 := KeyValue{"a", "b"}
	kv2 := KeyValue{"c", "d"}

	//Create a file if don't exist in append mode
	fd, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}

	//Jsonfy the structure
	encodedKV1, _ := json.Marshal(&kv1)
	encodedKV2, _ := json.Marshal(&kv2)
	if err != nil {
		panic(err)
	}

	fd.Write(append(encodedKV1, byte('\n')))
	fd.Write(append(encodedKV2, byte('\n')))
	fd.Close()

	file, _ := os.Open("test.txt")
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	KV := make([]KeyValue, 2)
	i := 0
	for fscanner.Scan() {
		//fmt.Println(fscanner.Text())
		json.Unmarshal(fscanner.Bytes(), &KV[i])
		i++
	}
	//json.Unmarshal(content, &KVs)
	//fmt.Printf("%v", KV)

	kv := make(map[string][]string)
	kv["a"] = append(kv["a"], "b")
	kv["a"] = append(kv["a"], "c")
	fmt.Println(kv["a"])

}
