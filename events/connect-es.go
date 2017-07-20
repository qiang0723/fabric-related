package main

import (
//  "encoding/json"
  "fmt"
//  "reflect"
  "time"

  "golang.org/x/net/context"

  elastic "gopkg.in/olivere/elastic.v5"
)

// create a struct: Event, an index:eventhub, and an document event0411
// eventhub/event0411 
// Event is a structure used for serializing/deserializing data in Elasticsearch.
type Event struct {
  Channel_header     string                `json:"channel_header"`
  Signature_header   string                `json:"signature_header"`
  Data               string                `json:"data"`
  Created            time.Time             `json:"created,omitempty"`
}

func main() {
  // Starting with elastic.v5, you must pass a context to execute each service
  ctx := context.Background()

  // Create a client and connect to 127.16.3.9, default is 127.0.0.1
  client, err := elastic.NewClient(elastic.SetURL("http://172.16.3.9:9200"))
  if err != nil {
    // Handle error
    panic(err)
  }

  // Ping the Elasticsearch server to get e.g. the version number
  info, code, err := client.Ping("http://172.16.3.9:9200").Do(ctx)
  if err != nil {
    // Handle error
    panic(err)
  }
  fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

  // Getting the ES version number is quite common, so there's a shortcut
  esversion, err := client.ElasticsearchVersion("http://172.16.3.9:9200")
  if err != nil {
    // Handle error
    panic(err)
  }
  fmt.Printf("Elasticsearch version %s\n", esversion)

  // Use the IndexExists service to check if a specified index exists.
  exists, err := client.IndexExists("eventhub").Do(ctx)
  if err != nil {
    // Handle error
    panic(err)
  }
  if !exists {
    // Create a new index.
    createIndex, err := client.CreateIndex("eventhub").Do(ctx)
    if err != nil {
      // Handle error
      panic(err)
    }   
    if !createIndex.Acknowledged {
      // Not acknowledged
    }   
  }

  createtime := time.Now()

  // Add a document to the index
  events := Event{Channel_header: "testchainid", Signature_header: " test ", Data: "this is a test", Created: createtime}
  output, err := client.Index().
    Index("eventhub").
    Type("event0411").
    Id("1").
    BodyJson(events).
    Do(ctx)
  if err != nil {
    // Handle error
    panic(err)
  }
  fmt.Printf("Indexed eventhub %s to index %s, type %s\n", output.Id, output.Index, output.Type)

  // Get event0411 with specified ID
  get1, err := client.Get().
    Index("eventhub").
    Type("event0411").
    Id("1").
    Do(ctx)
  if err != nil {
    // Handle error
    panic(err)
  }
  if get1.Found {
    fmt.Printf("Got source %s\n", get1.Source)
  }

}
