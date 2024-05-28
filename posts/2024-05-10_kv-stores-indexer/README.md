---
publication_date: 2024-05-10T13:37:00Z
slug: kv-stores-indexer
tags: [blog, post, tx-indexer, dev]
authors: [ajnavarro]
---

# Key/Value Stores: How We Improved the Performance of Our tx-indexer by 10x

In this article, we'll discuss how we achieved a tenfold increase in the processing speed of the tx-indexer by applying four key concepts related to our use of key/value storage:

・ [Key/Value Stores: How We Improved the Performance of Our tx-indexer by 10x](#keyvalue-stores-how-we-improved-the-performance-of-our-tx-indexer-by-10x)  
  ・ [Understanding Key/Value Store Variability](#understanding-keyvalue-store-variability)  
  ・ [The Importance of Efficient Data Encoding](#the-importance-of-efficient-data-encoding)  
  ・ [Implementing Secondary Indexes on a Key/Value Store](#implementing-secondary-indexes-on-a-keyvalue-store)  
  ・ [The Role of Batch Inserts in Enhancing Performance](#the-role-of-batch-inserts-in-enhancing-performance)  
    ・ [Data consistency](#data-consistency)  
    ・ [Speed](#speed)  
      ・ [Old](#old)  
      ・ [New](#new)  
  ・ [Conclusion](#conclusion)  

The Transaction Indexer ([tx-indexer](https://github.com/gnolang/tx-indexer)) is the primary tool Gno.land uses to index its networks. It is in charge of keeping up with block production, fetching new data, indexing it, and serving it to users while providing filtering and subscription capabilities. The tx-indexer creates versatility and ease of use when using on-chain data, which is one of the key aspects of a fully functioning decentralized application.

## Understanding Key/Value Store Variability

Not all key/value storages are created equal. Each varies significantly, and depending on their internal data structures, some are better suited for certain use cases than others. A deep understanding of the key/value store you plan to use will help you better organize data for efficient writing and reading and assist in choosing the best store for your specific needs.

While [PebbleDB](https://github.com/cockroachdb/pebble) is based on [RocksDB](https://github.com/facebook/rocksdb/wiki/RocksDB-Overview), the two databases differ significantly. Both utilize LSM Trees built upon SSTables; however, PebbleDB supports only a subset of the features available in RocksDB. For instance, PebbleDB lacks built-in transaction capabilities, but these can be alternatively implemented through the use of Batches and/or Snapshots.

## The Importance of Efficient Data Encoding

Our indexing involved elements defined by consecutive integers, with Blocks on one side and Transactions within a Block on the other.

Initially, Blocks were indexed using a combination of `block_prefix` and block_id encoded in little endian. This method wasn't allowing us to use iterators for ordered data retrieval, forcing us to fetch elements individually, resulting in excessive and inefficient database queries.

After refactoring, we adopted a binary encoding scheme that allowed for custom encoding of strings and integers. This flexibility enabled ascending or descending order iterations, which significantly improved our ability to read data sequentially through iterators and, consequently, reduced query times dramatically.

Small example about how we encoded uint32 values in ascending order:

```go!
func encodeUint32Ascending(b []byte, v uint32) []byte {
	return append(b, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}
```

## Implementing Secondary Indexes on a Key/Value Store

While most filters are applied on the fly due to their low cost, we implemented secondary indexes to fetch Transactions by Hash efficiently.

Secondary indexes are specialized key groups that directly reference the primary index key where the data resides. For example, a transaction with ID `3` in block `42` is indexed as `/index/txs/[uint64]42[uint32]3`. These transactions are also uniquely identified by a hash representing the entire transaction content.

To fetch transactions by hash, we created a secondary index that points to the primary index:

`/index/txh/[HASH] -> /data/txs/[uint64]42[uint32]3`  

Although our secondary indexes do not require ordered iteration, this capability remains available, allowing us to apply additional filters as necessary. For instance, we could index transactions by year:

`/index/txYear/[uint16]2024[uint64]42[uint32]3 -> /data/txs/[uint64]42[uint32]3`

This format allows us to iterate through transactions within a specific year, from the start to the end of 2023, for example:

・ from: `/index/txYear/[uint16]2023[uint64]0[uint32]0`  
・ to: `/index/txYear/[uint16]2023[uint64]MAX_UINT64[uint32]MAX_UINT32`  

## The Role of Batch Inserts in Enhancing Performance

The advantages of write batches are often overlooked but crucial. Inserting elements individually can lead to data consistency issues and slower operations.

### Data consistency

Batches ensure atomicity—either all elements are persisted, or none are. Without batches, a failure during insertion could result in a block being saved without some of its transactions.

### Speed

Each insertion involves internal processes that slow down the operation. By grouping several entries in one batch, we significantly enhance insertion speed. These are new benchmarks comparing the old and new way of writting elements without and with batches. Note that these are just synthetic benchmarks and the 10x improvement was measured when using the indexer as it is (we came from speding 30 mins to 3 mins with the new storage changes):

#### Old

```go!
func BenchmarkPebbleWrites(b *testing.B) {
	store, err := NewDB(b.TempDir())
	require.NoError(b, err)
	defer store.Close()

	pairs := generateRandomPairs(b, b.N)

	b.ResetTimer()
	for k, v := range pairs {
		err := store.Set([]byte(k), v)

		b.StopTimer()
		require.NoError(b, err)
		b.StartTimer()
	}
}
```

```
goos: linux
goarch: amd64
pkg: github.com/gnolang/tx-indexer/storage/pebble
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkPebbleWrites
BenchmarkPebbleWrites-12    	    1316	    928941 ns/op	      33 B/op	       0 allocs/op
PASS
ok  	github.com/gnolang/tx-indexer/storage/pebble	1.384s
```

#### New

```go!
func BenchmarkPebbleWrites(b *testing.B) {
	store, err := NewPebble(b.TempDir())
	require.NoError(b, err)
	defer store.Close()

	pairs := generateRandomBlocks(b, b.N)

	batch := store.WriteBatch()

	b.ResetTimer()
	for _, v := range pairs {
		err := batch.SetBlock(v)

		b.StopTimer()
		require.NoError(b, err)
		b.StartTimer()
	}

	err = batch.Commit()

	b.StopTimer()
	require.NoError(b, err)
	b.StartTimer()
}
```

```
goos: linux
goarch: amd64
pkg: github.com/gnolang/tx-indexer/storage
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkPebbleWrites
BenchmarkPebbleWrites-12          249462              4730 ns/op            1704 B/op         43 allocs/op
PASS
ok      github.com/gnolang/tx-indexer/storage   4.669s
```

## Conclusion

If you find out this interesting and want to have a deeper look about [how it is done](https://github.com/gnolang/tx-indexer/tree/main/storage), or just try our indexer, it is as simple as ramping up a docker image:

```
docker run -it -p 8546:8546 ghcr.io/gnolang/tx-indexer:latest start -remote http://test3.gno.land:36657
```

And start playing with it through its GraphQL interface at `http://localhost:8546/graphql`
