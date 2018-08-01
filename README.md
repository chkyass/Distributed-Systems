# Distributed-Systems
This repo contains solutions to the labs of the MIT distributed systems [course](https://pdos.csail.mit.edu/6.824/). The advantage of these labs is that they contain pre-existing tests to verify the accuracy of the solution code

## MapReduce [lab1](https://pdos.csail.mit.edu/6.824/labs/lab-1.html)
In this lab, I built a MapReduce library.  
The files for this lab are in the mapreduce folder, except one in the main folder.:
- common_map.go contains the code executed for each user supplied Map function. It partitions the output of the different Map tasks.
- common_reduce.go contains the code executed for each user supplied Reduce function. It gathers the output of Map tasks to apply the Reduce task to them.
- main/wc.go contains an exemple of user supplied Map and Reduce function for counting occurences of each word in multiple files in parallel.
- sheduler.go contains the code responsible for distribution of the different Map and Reduce tasks on different workers. It also handle workers failure.

The lab is based on this [article](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)

## Raft [lab2](https://pdos.csail.mit.edu/6.824/labs/lab-raft.html) (Currently working on it)
First in a serie of labs in which i will build a fault-tolerant key/value storage system.  
Implementation of the Raft protocol.  
Based on this article [article](https://pdos.csail.mit.edu/6.824/papers/raft-extended.pdf)
