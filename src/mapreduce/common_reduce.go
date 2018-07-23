package mapreduce

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func sortKVs(filename string, KVs map[string][]string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	kv := KeyValue{}
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		json.Unmarshal(fscanner.Bytes(), &kv)
		KVs[kv.Key] = append(KVs[kv.Key], kv.Value)
	}

}

func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTask int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	sortedKVs := make(map[string][]string)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), strconv.Itoa(reduceTask)) {
			sortKVs(f.Name(), sortedKVs)
		}
	}

	fd, err := os.OpenFile(mergeName(jobName, reduceTask), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	for k, v := range sortedKVs {
		kv := KeyValue{k, reduceF(k, v)}
		encodedKV, err := json.Marshal(&kv)
		if err != nil {
			panic(err)
		}
		fd.Write(append(encodedKV, byte('\n')))
	}
}
