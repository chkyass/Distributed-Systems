# Distributed-Systems
This repo contains solutions to the labs of the MIT distributed systems [course](https://pdos.csail.mit.edu/6.824/). The avantage of this labs is that they contains pre existent tests to verify the correctness of the code
## MapReduce [lab1](https://pdos.csail.mit.edu/6.824/labs/lab-1.html)
In this lab i built a MapReduce library and fault tolerant distributed systems.
The files for this lab are in the mapreduce folder, excepting one in the main folder:
- common_map.go contains the code executed for every user supplied map function. It partition the output of differents generic map tasks.
- common_reduce.go contains the code executed for every user supplied reduce function. It gather the output of the map tasks to apply them the reduce task.
- main/wc.go contains an exemple of user supplied map and reduce function to count occurences of each word in multiple files.
- sheduler.go contains the code responsible of distributing the different map and reduce task on different workers. It also handle worker failure.

The lab is based on this [article](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)

## Raft [lab2](https://pdos.csail.mit.edu/6.824/labs/lab-raft.html) (Currently working on it)
First in a serie of labs in which i'll build a fault-tolerant key/value storage system.  
Implementation of the Raft protocol.  
Based on this article [article](https://pdos.csail.mit.edu/6.824/papers/raft-extended.pdf)
