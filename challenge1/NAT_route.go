package main

import (
	"fmt"
	"strings"
)

func convertToMap(arr []string) map[string][]string {
	dict := make(map[string][]string)
	for _, item := range arr {
		parts := strings.Split(item, " - ")
		nat := parts[0]
		region := parts[1]
		dict[region] = append(dict[region], nat)
	}
	return dict
}

func allocateSubnets(NATsMappedByAZ, subnetsMappedByAZ map[string][]string) map[string][]string {
	allocation := make(map[string][]string)
	unallocatedSubnets := [][]string{}

	// First allocation pass
	for zone, subnets := range subnetsMappedByAZ {
		if nats, exists := NATsMappedByAZ[zone]; exists {
			natQueue := append([][]string{}, toNatQueue(nats, zone)...)
			for _, subnet := range subnets {
				if len(natQueue) > 0 {
					natInstance := natQueue[0]
					natQueue = natQueue[1:]
					key := fmt.Sprintf("%s - %s", natInstance[0], natInstance[1])
					allocation[key] = append(allocation[key], fmt.Sprintf("%s - %s", subnet, zone))
					natQueue = append(natQueue, natInstance)
				} else {
					unallocatedSubnets = append(unallocatedSubnets, []string{subnet, zone})
				}
			}
		} else {
			for _, subnet := range subnets {
				unallocatedSubnets = append(unallocatedSubnets, []string{subnet, zone})
			}
		}
	}

	// Second allocation pass
	natQueue := [][]string{}
	for zone, nats := range NATsMappedByAZ {
		natQueue = append(natQueue, toNatQueue(nats, zone)...)
	}
	for _, subnet := range unallocatedSubnets {
		if len(natQueue) > 0 {
			natInstance := natQueue[0]
			natQueue = natQueue[1:]
			key := fmt.Sprintf("%s - %s", natInstance[0], natInstance[1])
			allocation[key] = append(allocation[key], fmt.Sprintf("%s - %s", subnet[0], subnet[1]))
			natQueue = append(natQueue, natInstance)
		}
	}

	return allocation
}

func toNatQueue(nats []string, zone string) [][]string {
	queue := [][]string{}
	for _, nat := range nats {
		queue = append(queue, []string{nat, zone})
	}
	return queue
}

func main() {
	natList := []string{
		"1 - us-west1-a",
		"2 - us-west1-b",
		"3 - us-west1-b",
	}

	subnetList := []string{
		"1 - us-west1-a",
		"2 - us-west1-b",
		"3 - us-west1-b",
		"4 - us-west1-c",
	}

	NATsMappedByAZ := convertToMap(natList)
	subnetsMappedByAZ := convertToMap(subnetList)

	allocation := allocateSubnets(NATsMappedByAZ, subnetsMappedByAZ)

	for nat, subs := range allocation {
		fmt.Printf("Instance (%s):\n", nat)
		for _, sub := range subs {
			fmt.Printf(" subnet (%s)\n", sub)
		}
	}
}