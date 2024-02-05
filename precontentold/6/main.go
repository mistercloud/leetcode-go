package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Solution(read, writer)
}

var (
	subNetTotal   int = 256
	subNet16Total int = 65536
)

type Subnet struct {
	name         int
	totalBlocked int
	ips          []string
	takeIps      bool
}

func Solution(reader *bufio.Reader, writer *bufio.Writer) {
	defer writer.Flush()

	var total, maxOutput int
	fmt.Fscanln(reader, &total, &maxOutput)
	i := 0
	mask16 := subNet16Total
	links := make(map[int]*Subnet, subNetTotal)
	subnets := []*Subnet{}
	for total > 0 {
		var ip string
		fmt.Fscanln(reader, &ip)
		mask16--
		ip3, _ := strconv.Atoi(strings.Split(ip, `.`)[2])
		_, ok := links[ip3]
		subnet := &Subnet{ip3, subNetTotal, []string{}, false}
		if ok {
			subnet = links[ip3]
		} else {
			links[ip3] = subnet
			subnets = append(subnets, subnet)
		}

		subnet.ips = append(subnet.ips, ip)
		subnet.totalBlocked--

		total--
		i++
	}

	if len(links) > maxOutput {
		//единственный способ - запретить все ip
		fmt.Fprintln(writer, mask16)
		fmt.Fprintln(writer, "1")
		fmt.Fprintln(writer, "100.200.0.0/16")
	} else {
		//для каждой маски выбираем либо конкретный ip либо всю подсеть
		sort.Slice(subnets, func(i, j int) bool {
			return subnets[i].totalBlocked >= subnets[j].totalBlocked
		})

		blockedTotal, lines := minimumBlocked(subnets, 0, maxOutput)
		fmt.Fprintln(writer, blockedTotal)
		fmt.Fprintln(writer, lines)

		for i := 0; i <= 256; i++ {
			_, ok := links[i]
			if ok {
				if links[i].takeIps {
					for _, ip := range links[i].ips {
						fmt.Fprintln(writer, ip)
					}
				} else {
					fmt.Fprintln(writer, "100.200."+strconv.Itoa(i)+".0/24")
				}
			}
		}
	}

}

func minimumBlocked(subnets []*Subnet, index int, maxLines int) (int, int) {

	if index >= len(subnets) {
		return 0, 0
	}

	blocked, lines := subnets[index].totalBlocked, 1
	if len(subnets[index].ips) <= (maxLines - len(subnets) + index + 1) {
		subnets[index].takeIps = true
		blocked = 0
		lines = len(subnets[index].ips)
	}
	nextBlocked, nextLines := minimumBlocked(subnets, index+1, maxLines-lines)

	return blocked + nextBlocked, lines + nextLines
}
