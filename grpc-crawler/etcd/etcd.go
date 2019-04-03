package etcd

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client"
	"golang.org/x/net/context"
	"log"
	"os"
	"sync"
	"time"
)

type EtcdService struct {
	KeyApi client.KeysAPI
	Peers  map[string]*Peers
	Lock   sync.Mutex
}

type Peers struct {
	Pid  int
	Ip   string
	Name string
}

type Service interface {
	Add(peers *Peers)
	Update(peers *Peers)
	Remove(peers *Peers)
	Watch()
}

func RegisterEtcdService(endpoints []string) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Panicf("client Init Error:%v\n", err)
	}
	service := &EtcdService{
		Peers:  make(map[string]*Peers),
		KeyApi: client.NewKeysAPI(c),
	}
	go service.Watch()
}

func RespNodeToPeers(node *client.Node) *Peers {
	var peer *Peers
	err := json.Unmarshal([]byte(node.Value), &peer)
	if err != nil {
		fmt.Print(err)
		log.Print(err)
	}
	return peer
}

func (e *EtcdService) Watch() {
	api := e.KeyApi
	w := api.Watcher("workers/", &client.WatcherOptions{
		Recursive: true,
	})
	peer := &Peers{}
	for {
		resp, err := w.Next(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		println(resp.Action)
		println(resp)
		switch resp.Action {
		case "expire":
			peer = RespNodeToPeers(resp.PrevNode)
			e.Remove(peer)
		case "set":
			if _, ok := e.Peers[peer.Name]; ok {
				e.Update(peer)
				log.Println("Update peer ", peer.Name)
			} else {
				peer = RespNodeToPeers(resp.Node)
				e.Add(peer)
				log.Println("Add peer ", peer.Name)
			}
		case "delete":
			peer = RespNodeToPeers(resp.Node)
			e.Remove(peer)
			log.Println("Remove peer ", peer.Name)
		default:
			log.Print("返回参数不正确:%v", resp.Action)
		}
	}

}

func (e *EtcdService) Add(peers *Peers) {
	p := &Peers{
		Pid:  os.Getpid(),
		Ip:   peers.Ip,
		Name: peers.Name,
	}
	e.Lock.Lock()
	defer e.Lock.Unlock()
	e.Peers[peers.Name] = p
}

func (e *EtcdService) Update(peers *Peers) {
	e.Peers[peers.Name] = peers
	e.Lock.Lock()
	defer e.Lock.Unlock()
}

func (e *EtcdService) Remove(peers *Peers) {
	delete(e.Peers, peers.Name)
}
