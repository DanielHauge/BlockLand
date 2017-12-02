package main

import (
	"log"
	"github.com/boltdb/bolt"
)

// go get github.com/boltdb/bolt



type Blockchain struct {
	top []byte
	db *bolt.DB
}

type BlockChainReader struct {
	cur_hash []byte
	db 		*bolt.DB
}



/*
This is a method to be called with a instanced Blockchain struct.
It will append/add a block the to chain, basicly.
It will call methods like : NewBlock and that will call SetHash = Calculating the Hash.
 */
func (chain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := chain.db.View(func(tx *bolt.Tx) error{
		b:= tx.Bucket([]byte("blocksBucket"))
		lastHash = b.Get([]byte("1"))

		return nil
	})

	newBlock := NewBlock(data, lastHash)

	err = chain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("1"), newBlock.Hash)
		chain.top = newBlock.Hash
		return err
	})

	if err != nil {
		log.Println("Something went wrong with adding a block")
		log.Println(err.Error())
	}

}



/*
This is initializing a Blockchain with the first block which is a empty block with no previusblock hash, but with the data "Genesis Block" and a timestamp.
 */
func NewBlockChain() *Blockchain {
	var top []byte
	db, err := bolt.Open("dbFile", 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))

		if b == nil{
			firstblock := BirthFirstBlock()
			b, err := tx.CreateBucket([]byte("blocksBucket"))
			err = b.Put(firstblock.Hash, firstblock.Serialize())
			err = b.Put([]byte("1"), firstblock.Hash)
			top = firstblock.Hash
			return err
		} else {
			top = b.Get([]byte("1"))
			return err
		}
	})
	if err != nil{
		log.Println("Something went wrong with the DB Connection")
	}

	bc := Blockchain{top, db}


	return &bc
}



func (chain *Blockchain) GetReader() *BlockChainReader {
	chainreader := &BlockChainReader{chain.top, chain.db}
	return chainreader
}

func (reader *BlockChainReader) Next() *Block {
	var block *Block
	err := reader.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))
		BlockInBytes := b.Get(reader.cur_hash)
		block = Deserialize(BlockInBytes)

		return nil
	})
	if err != nil {
		log.Println("Something went wrong with reading next block in block chain")
		log.Println(err.Error())
	}

	reader.cur_hash = block.PrevBlockHash

	return block

}




