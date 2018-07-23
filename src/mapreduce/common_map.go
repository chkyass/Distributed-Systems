package mapreduce

import (
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"os"
)

func doMap(
	jobName string, // the name of the MapReduce job
	mapTask int, // which map task this is
	inFile string,
	nReduce int, // the number of reduce task that will be run ("R" in the paper)
	mapF func(filename string, contents string) []KeyValue,
) {
	// Read File content
	content, err := ioutil.ReadFile(inFile)
	if err != nil {
		panic(err)
	}
	// Apply the user defined function to the content
	keysValues := mapF(inFile, string(content[:]))

	// Partition the intermediate output into files
	for _, kv := range keysValues {
		filename := reduceName(jobName, mapTask, ihash(kv.Key)%nReduce)

		//Create a file if don't exist in append mode
		fd, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			panic(err)
		}

		//Jsonfy the structure
		encodedKV, err := json.Marshal(&kv)
		if err != nil {
			panic(err)
		}

		fd.Write(append(encodedKV, byte('\n')))
		fd.Close()
	}
}

func ihash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & 0x7fffffff)
}
