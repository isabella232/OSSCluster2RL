package osscluster2rl

import (
	"crypto/tls"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

// GetMemory : collect memory usage information
func GetMemory(servers []string, password string, sslConf *tls.Config, dbg bool) int {
	bytes := 0
	for _, server := range servers {
		client := redis.NewClient(&redis.Options{
			Addr:      server,
			Password:  password,
			TLSConfig: sslConf,
		})
		info := client.Info("memory")
		if dbg {
			fmt.Println("DEBUG: Fetching memory usage information from", server)
			if info.Err() != nil {
				fmt.Println("Error fetching memory usage data from ", server, "Error: ", info.Err())
			}
		}
		for _, line := range strings.Split(info.Val(), "\n") {
			r := regexp.MustCompile(`used_memory:(\d+)`)
			res := r.FindStringSubmatch(line)
			if len(res) > 0 {
				j, _ := strconv.Atoi(res[1])
				bytes += j
			}
		}
	}
	return (bytes)
}

// GetKeyspace : collect memory keyspace information
func GetKeyspace(servers []string, password string, sslConf *tls.Config, dbg bool) int {
	keys := 0
	for _, server := range servers {
		client := redis.NewClient(&redis.Options{
			Addr:      server,
			Password:  password,
			TLSConfig: sslConf,
		})
		info := client.Info("keyspace")
		if dbg {
			fmt.Println("DEBUG: Fetching memory keyspace from", server)
			if info.Err() != nil {
				fmt.Println("Error fetching keyspace data from ", server, "Error: ", info.Err())
			}
		}
		for _, line := range strings.Split(info.Val(), "\n") {
			r := regexp.MustCompile(`db\d+:keys=(\d+),`)
			res := r.FindStringSubmatch(line)
			if len(res) > 0 {
				j, _ := strconv.Atoi(res[1])
				keys += j
			}
		}
	}
	return keys
}

// GetCmdStats : collect command stat information
func GetCmdStats(servers []string, password string, sslConf *tls.Config, dbg bool) (map[string]int, map[string]int) {
	cmdstats := make(map[string]int)
	cmdusec := make(map[string]int)
	for _, server := range servers {
		client := redis.NewClient(&redis.Options{
			Addr:      server,
			Password:  password,
			TLSConfig: sslConf,
		})
		info := client.Info("commandstats")
		for _, line := range strings.Split(info.Val(), "\n") {
			r := regexp.MustCompile(`cmdstat_(\w+):calls=(\d+),usec=(\d+),`)
			res := r.FindStringSubmatch(line)
			if len(res) == 4 {
				j, _ := strconv.Atoi(res[2])
				k, _ := strconv.Atoi(res[3])
				cmdstats[res[1]] += j
				cmdusec[res[1]] += k
			}
		}
	}
	return cmdstats, cmdusec
}
