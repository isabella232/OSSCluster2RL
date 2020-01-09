# osscluster2rl

This program collects usage data from open source redis clusters and obtains the data necessary to size for a Redis Enterprise cluster

## Results
The data is returned in a CSV file similar to the following:

```
name,master_count,replication_factor,total_key_count,total_memory,maxCommands
host1,3,1,53452,8274224,120
host2,3,2,2510,19768564264,10000
```
| stat | description | notes |
|---|---|---|
|master_count|number of master nodes in the cluster||
|replication_factor|the number of slaves per master in the cluster||
|total_key_count|Total number of keys on the cluster|sum of all nodes|
|total_memory|Amount of memory used by Redis|sum of  used_memory from all master nodes, multiply by 2 factor if using HA in Redis Enterprise|

## Usage
0. Download the [.tar.gz binaries](https://github.com/Redislabs-Solution-Architects/OSSCluster2RL/releases) and unzip
1. copy the example_config.toml file and edit
2. Add nodes, you only need to specify a single node in the cluster the script will auto identify the rest
3. Run the binary. eg for Linux: ```./osscluster2rl_linux_amd64 -conf config.toml```

## Building
0. Ensure you have go >= 1.13 and make installed on your machine
1. run ```make``
2. Run the binary. ```./osscluster2rl -conf config.toml```