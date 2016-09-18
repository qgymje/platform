package main

import (
	"tech.cloudzen/tests/arango/db"
	"github.com/leesper/go_rng"
	"fmt"
	"strconv"
	"time"
	"math/rand"
	"sync"
	"github.com/diegogub/aranGO"
)

func main() {
	db.Connect()
	// poplution()
	aranGO.Graph{}.Traverse()
}

func poplution() {
	db.DB.TruncateCollection(db.RelCollName)
	db.DB.TruncateCollection(db.NodeColName)
	randGen := rng.NewPoissonGenerator(time.Now().UnixNano())
	var vnum  = 10000
	//fmt.Print("Input Node Count for the test: ")
	//fmt.Scanf("%d", vnum)
	fmt.Printf("Creating %d vertices \n", vnum)
	for i := 0; i < vnum; i++{
		node := db.Node{Id: i}
		node.Document.Key = strconv.Itoa(i)
		node.Created = time.Now().Unix()
		db.NodeColl.Save(node)
	}
	fmt.Println("Vertices created")
	fmt.Println("Create Edges")
	edges := 0
	parallelCounter := make(chan bool, 32)
	var wg sync.WaitGroup
	for i := 0; i < vnum; i++{
		nodeKey := strconv.Itoa(i)
		edgeCount := int(randGen.Poisson(20))
		if edgeCount >= vnum {
			edgeCount = vnum - 1;
		}
		for j := 0; j < edgeCount; j++{
			oppisiteKey := strconv.Itoa(rand.Intn(vnum))
			wg.Add(1)
			go func () {
				parallelCounter <- true
				db.RelColl.Relate("tNode/" + nodeKey, "tNode/" + oppisiteKey, db.Rel{Created: time.Now().Unix()})
				<- parallelCounter
				wg.Done()
			} ()
			edges ++
		}
	}
	wg.Wait()
	fmt.Println("Edges created: " + strconv.Itoa(edges))
}