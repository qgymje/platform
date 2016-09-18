package main

import ("fmt"
	 sara "github.com/Shopify/sarama"
)

func main() {
 	var addr  =  []string{"192.168.157.130:9092"}
	config:= sara.NewConfig()
	err := config.Validate()
	if(err!=nil){
		panic(err)
	}
	c , err:= sara.NewClient(addr,config)
	if(err!=nil) {
		panic(err)
	}
	fmt.Println(c)
	topic,err :=c.Topics()
	if(err!=nil) {
		panic(err)
	}
	fmt.Println(topic)
	partitions,err :=c.Partitions("TutorialTopic")
	if(err!=nil) {
		panic(err)
	}
	fmt.Println(partitions)

	partition,err :=c.WritablePartitions("TutorialTopic")
	if(err!=nil) {
		panic(err)
	}
	fmt.Println(partition)

 	newAsyncProducer,err := sara.NewAsyncProducer(addr,config)
	if(err!=nil){
		panic(err)
	}
	fmt.Println(newAsyncProducer)
	newAsyncProducer,err = sara.NewAsyncProducerFromClient(c)
	fmt.Println(newAsyncProducer)


	newSyncProducerFromClient,err := sara.NewSyncProducerFromClient(c)
	if(err!=nil){
		panic(err)
	}
	fmt.Println("newSyncProducerFromClient=",newSyncProducerFromClient)

	var msg sara.ProducerMessage
	msg.Topic="TutorialTopic"
	msg.Value=sara.ByteEncoder([]byte("i am test"))

 	partition1 , offset , err  := newSyncProducerFromClient.SendMessage(&msg)
	if(err!=nil){
		panic(err)
	}
	fmt.Println(partition1,offset)


}


