package loadbalancer

import (
	"crypto/sha512"
	"encoding/binary"
	"mojiayi-golang-algorithm/domain"
	"sort"
	"strconv"
)

type ConsistentHashScheduler struct {
	ServerList  *[]domain.ServerInfo
	hashHostMap map[uint64]domain.ServerInfo
	hashRing    []uint64
}

func (s *ConsistentHashScheduler) InitHashRing() {
	s.hashHostMap = make(map[uint64]domain.ServerInfo, len(*s.ServerList))
	s.hashRing = make([]uint64, len(*s.ServerList))
	for index, host := range *s.ServerList {
		checksum := sha512.Sum512([]byte(strconv.Itoa(host.ServiceId)))
		var hashCode = binary.BigEndian.Uint64(checksum[:])
		s.hashHostMap[hashCode] = host
		s.hashRing[index] = hashCode
	}
}

func (s *ConsistentHashScheduler) Choose(key string) domain.ServerInfo {
	checksum := sha512.Sum512([]byte(key))
	var hashCode = binary.BigEndian.Uint64(checksum[:])
	index := sort.Search(len(s.hashRing), func(i int) bool {
		return s.hashRing[i] > hashCode
	})
	if index >= len(s.hashRing) {
		index = 0
	}
	return s.hashHostMap[s.hashRing[index]]
}
