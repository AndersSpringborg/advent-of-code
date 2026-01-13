package internal

import (
	"log"
	"strings"
)

type Server struct {
	id      string
	devices []*Server
}

func NewServer(id string) *Server {
	return &Server{
		id:      id,
		devices: make([]*Server, 0),
	}
}

func (s *Server) AddDevice(device *Server) {
	s.devices = append(s.devices, device)
}

type ServerGraph struct {
	server map[string]*Server
}

func NewServerGraph() *ServerGraph {
	return &ServerGraph{
		server: make(map[string]*Server),
	}
}

func (sg *ServerGraph) ToString() string {
	var sb strings.Builder
	for id, server := range sg.server {
		sb.WriteString(id)
		sb.WriteString(": ")
		for _, device := range server.devices {
			sb.WriteString(device.id)
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (sg *ServerGraph) ParseLine(line string) {

	parts := strings.SplitN(line, ": ", 2)
	server := parts[0]
	devicesString := parts[1]

	devices := strings.Split(devicesString, " ")

	// Make suer all devices exist in the graph
	for _, deviceID := range devices {
		if deviceID == "" {
			continue
		}
		if _, exists := sg.server[deviceID]; !exists {
			sg.server[deviceID] = NewServer(deviceID)
		}
	}

	newServer := NewServer(server)
	if _, exists := sg.server[server]; exists {
		newServer = sg.server[server]
	}
	for _, deviceID := range devices {
		if deviceID == "" {
			continue
		}
		newServer.AddDevice(sg.server[deviceID])
	}

	sg.server[server] = newServer
}

const start = "you"
const srv = "svr"
const end = "out"
const fft = "fft"
const dac = "dac"

func (s *Server) dfsToEnd(server *Server) int {
	if server.id == end {
		return 1
	}

	totalPaths := 0
	for _, device := range server.devices {
		totalPaths += s.dfsToEnd(device)
	}
	return totalPaths
}

type Cache struct {
	paths []map[struct {
		id     string
		hasFCC bool
		hasDAC bool
	}]int
}

func NewCache(depth int) Cache {
	paths := make([]map[struct {
		id     string
		hasFCC bool
		hasDAC bool
	}]int, depth)
	for i := 0; i < depth; i++ {
		paths[i] = make(map[struct {
			id     string
			hasFCC bool
			hasDAC bool
		}]int)
	}
	return Cache{paths: paths}
}

func (s *Server) dfsToEndWithFCCandDAC(server *Server, hasFCC bool, hasDAC bool, cache Cache) int {
	if server.id == end {
		if hasFCC && hasDAC {
			return 1
		}
		return 0
	}
	// Check cache
	if _, exists := cache.paths[len(cache.paths)-1][struct {
		id     string
		hasFCC bool
		hasDAC bool
	}{server.id, hasFCC, hasDAC}]; exists {
		return cache.paths[len(cache.paths)-1][struct {
			id     string
			hasFCC bool
			hasDAC bool
		}{server.id, hasFCC, hasDAC}]
	}

	totalPaths := 0
	for _, device := range server.devices {
		newHasFCC := hasFCC || device.id == fft
		newHasDAC := hasDAC || device.id == dac
		totalPaths += s.dfsToEndWithFCCandDAC(device, newHasFCC, newHasDAC, cache)
	}

	// Store in cache
	cache.paths[len(cache.paths)-1][struct {
		id     string
		hasFCC bool
		hasDAC bool
	}{server.id, hasFCC, hasDAC}] = totalPaths

	return totalPaths
}

func (sg *ServerGraph) TotalPathsFromYouToOutWithFCCandDAC() int {
	startServer, exists := sg.server[srv]
	if !exists {
		log.Fatal("Start server 'you' not found in the graph")
	}
	return startServer.dfsToEndWithFCCandDAC(startServer, false, false, NewCache(1000))
}

func (sg *ServerGraph) TotalPathsFromYouToOut() int {
	startServer, exists := sg.server[start]
	if !exists {
		log.Fatal("Start server 'you' not found in the graph")
	}
	return startServer.dfsToEnd(startServer)
}
