package main

import ("fmt"
	sara "github.com/Shopify/sarama"
	"time"
)

func main()  {
	var addr  =  []string{"192.168.157.130:9092"}
	config:= sara.NewConfig()
	err := config.Validate()
	config.Consumer.Offsets.CommitInterval = 100 * time.Millisecond
	config.Consumer.Offsets.Retention  = 3 * time.Minute

	if(err!=nil){
		panic(err)
	}
	c , err:= sara.NewClient(addr,config)
	if(err!=nil) {
		panic(err)
	}
	fmt.Println(c)

	consumer,err := sara.NewConsumer(addr,config)
	fmt.Println("consumer",consumer)
	if(err!=nil){
		panic(err)
	}

	//partitionConsumer, err := consumer.ConsumePartition("TutorialTopic", 0, sara.OffsetOldest)
	partitionConsumer, err := consumer.ConsumePartition("TutorialTopic", 0, sara.OffsetNewest)
	if(err!=nil){
		panic(err)
	}
	fmt.Println("partitionConsumer=",partitionConsumer)

 	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Partition:\t%d\n", msg.Partition)
		fmt.Printf("Offset:\t%d\n", msg.Offset)
		fmt.Printf("Key:\t%s\n", string(msg.Key))
		fmt.Printf("Value:\t%s\n", string(msg.Value))
		fmt.Println()
 	}

}