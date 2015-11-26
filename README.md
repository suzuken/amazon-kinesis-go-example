# Amazon Kinesis Stream consumer and producer example using aws-sdk-go

Demonstration for

* create stream
* wait until Kinesis stream available
* describe stream
* put record
* put multiple records (using PutRecords API)
* get records using shard iterator
* delete stream

Authorization is depends on aws-sdk-go. I use `~/.aws/config` for this purpose.

Example output

```
-> % go run main.go -stream my-kinesis-stream -region ap-northeast-1
{

}
{
  StreamDescription: {
    HasMoreShards: false,
    RetentionPeriodHours: 24,
    Shards: [{
        HashKeyRange: {
          EndingHashKey: "340282366920938463463374607431768211455",
          StartingHashKey: "0"
        },
        SequenceNumberRange: {
          StartingSequenceNumber: "49556707348180705259475312504965645807584792562676793346"
        },
        ShardId: "shardId-000000000000"
      }],
    StreamARN: "arn:aws:kinesis:ap-northeast-1:xxxxxxxxxxxx:stream/my-kinesis-stream",
    StreamName: "my-kinesis-stream",
    StreamStatus: "ACTIVE"
  }
}
{
  SequenceNumber: "49556707348180705259475312505169954271099667229664346114",
  ShardId: "shardId-000000000000"
}
[{
  Data: [
    104,
    111,
    103,
    101,
    48
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    49
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    50
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    51
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    52
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    53
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    54
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    55
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    56
  ],
  PartitionKey: "key2"
} {
  Data: [
    104,
    111,
    103,
    101,
    57
  ],
  PartitionKey: "key2"
}]
{
  FailedRecordCount: 0,
  Records: [
    {
      SequenceNumber: "49556707348180705259475312505171163196919281858839052290",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505172372122738896488013758466",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505173581048558511117188464642",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505174789974378125746363170818",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505175998900197740375537876994",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505177207826017355004712583170",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505178416751836969633887289346",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505179625677656584263061995522",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505180834603476198892236701698",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49556707348180705259475312505182043529295813521411407874",
      ShardId: "shardId-000000000000"
    }
  ]
}
{
  ShardIterator: "AAAAAAAAAAE8llzsV1gPXCWRLV7UcXSisncE5jyOHgs2yEDWRWT5q8VpFZ7vFW8qncqwvxHeqavSnkW+jdAUjBRYeNrDC0eBdgRMYHLnaIoIuolr78xS/ooobWxH8RZwG0IiyDYn/dOUDLv32PoB7skzAZ/1AYjFAYBnIVPp5AUq771NKh8WQ5nSABzI5jxJDJceH2ZYTZdwAfn95aOBeMXWwCvQoZMr"
}
{
  MillisBehindLatest: 0,
  NextShardIterator: "AAAAAAAAAAE9dau92rwQBbwfw+VFECaAlJsgOwdYE3ZXfQufLdXqsorcrvsFl0t+MqB+KwkROo8t1PaYwDBFH1uBR7R/4ZGb5m2W04p8MjdYbl9rkkyMkF5x356V2IyXIZqchzHBzYo34FM9C2+ndN+VNY/NRiKqtONp4q0W9hLcFT1WtbTkCJThq5E4RefhQuyX8D6p6A4c0wUpk2KadAZrZDmb0a44",
  Records: [
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101
      ],
      PartitionKey: "key1",
      SequenceNumber: "49556707348180705259475312505169954271099667229664346114"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        48
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505171163196919281858839052290"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        49
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505172372122738896488013758466"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        50
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505173581048558511117188464642"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        51
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505174789974378125746363170818"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        52
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505175998900197740375537876994"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        53
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505177207826017355004712583170"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        54
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505178416751836969633887289346"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        55
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505179625677656584263061995522"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        56
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505180834603476198892236701698"
    },
    {
      ApproximateArrivalTimestamp: 2015-11-26 07:19:31 +0000 UTC,
      Data: [
        104,
        111,
        103,
        101,
        57
      ],
      PartitionKey: "key2",
      SequenceNumber: "49556707348180705259475312505182043529295813521411407874"
    }
  ]
}
{
  MillisBehindLatest: 0,
  NextShardIterator: "AAAAAAAAAAH+BPEFSUYT2vcKcZBldMWp14lGxZbnLxF0zrbeLO3UOmN6Lc2NavT4hPVxivCzxUHY8bctYaj7l3F95y9z/zIs7dytvsVI5kvcYr23OFCJfRjhSYnXutz64GKPLwlxoN2sZg38LZSpcRUMqKROD/5AJHwx1OWA9RAIvI6n25gd6u0rX7SgGXbQHko/8CjoVMMUzzBofLNr7YBRyDl59ErA",
  Records: []
}
{

}
```

For handling errors, see [handling errors · aws/aws-sdk-go Wiki](https://github.com/aws/aws-sdk-go/wiki/handling-errors).

## LICENSE

MIT

## Author

Kenta Suzuki (a.k.a. suzuken)
