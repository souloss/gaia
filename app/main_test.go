package main_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/witchc/gaia/entity"
	bboltkvstorage "github.com/witchc/gaia/infrastructure/bbolt_kv_storage"

	bolt "go.etcd.io/bbolt"
)

var TestNodes []entity.SSH_Node = []entity.SSH_Node{
	{Node: &entity.Node{HostName: "192.168.0.100"}, User: "root", Pwd: "100pwd", Private_key: []byte("asfdqwe")},
	{Node: &entity.Node{HostName: "192.168.0.101"}, User: "root", Pwd: "101pwd", Private_key: []byte("asfdqwe")},
	{Node: &entity.Node{HostName: "192.168.0.102"}, User: "root", Pwd: "102pwd", Private_key: []byte("asfdqwe")},
	{Node: &entity.Node{HostName: "192.168.0.103"}, User: "root", Pwd: "103pwd", Private_key: []byte("asfdqwe")},
}

func Test_All(t *testing.T) {
	db, err := bolt.Open("/tmp/gaia.db", 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("\nGet NodeStorage...")
	node_storage := bboltkvstorage.NewNodeStorage(db)
	fmt.Println("\nGet NodeStorage Successful!", node_storage)

	fmt.Println("\nPut Nodes")
	for _, node := range TestNodes {
		node_storage.Put(&node)
	}
	fmt.Println("\nPut Nodes Successful!")

	fmt.Println("\nGet Node....")
	fmt.Println(node_storage.Get("192.168.0.100", reflect.TypeOf(entity.SSH_Node{})))

	fmt.Println("\nList Nodes....")
	nodes, _ := node_storage.List(reflect.TypeOf(entity.SSH_Node{}))
	for i, v := range nodes {
		fmt.Println(i, v)
	}
}
