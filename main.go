package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var (
	stream = flag.String("stream", "your-stream", "your stream name")
	region = flag.String("region", "ap-northeast-1", "your AWS region")
)

func main() {
	flag.Parse()

	s := session.New(&aws.Config{Region: aws.String(*region)})
	kc := kinesis.New(s)

	streamName := aws.String(*stream)

	out, err := kc.CreateStream(&kinesis.CreateStreamInput{
		ShardCount: aws.Int64(1),
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)

	if err := kc.WaitUntilStreamExists(&kinesis.DescribeStreamInput{StreamName: streamName}); err != nil {
		panic(err)
	}

	streams, err := kc.DescribeStream(&kinesis.DescribeStreamInput{StreamName: streamName})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", streams)

	putOutput, err := kc.PutRecord(&kinesis.PutRecordInput{
		Data:         []byte("hoge"),
		StreamName:   streamName,
		PartitionKey: aws.String("key1"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", putOutput)

	// put 10 records using PutRecords API
	entries := make([]*kinesis.PutRecordsRequestEntry, 10)
	for i := 0; i < len(entries); i++ {
		entries[i] = &kinesis.PutRecordsRequestEntry{
			Data:         []byte(fmt.Sprintf("hoge%d", i)),
			PartitionKey: aws.String("key2"),
		}
	}
	fmt.Printf("%v\n", entries)
	putsOutput, err := kc.PutRecords(&kinesis.PutRecordsInput{
		Records:    entries,
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	// putsOutput has Records, and its shard id and sequece enumber.
	fmt.Printf("%v\n", putsOutput)

	// retrieve iterator
	iteratorOutput, err := kc.GetShardIterator(&kinesis.GetShardIteratorInput{
		// Shard Id is provided when making put record(s) request.
		ShardId:           putOutput.ShardId,
		ShardIteratorType: aws.String("TRIM_HORIZON"),
		// ShardIteratorType: aws.String("AT_SEQUENCE_NUMBER"),
		// ShardIteratorType: aws.String("LATEST"),
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", iteratorOutput)

	// get records use shard iterator for making request
	records, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: iteratorOutput.ShardIterator,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", records)

	// and, you can iteratively make GetRecords request using records.NextShardIterator
	recordsSecond, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: records.NextShardIterator,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", recordsSecond)

	// OK, finally delete your stream
	deleteOutput, err := kc.DeleteStream(&kinesis.DeleteStreamInput{
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", deleteOutput)
}
