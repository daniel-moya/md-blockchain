package main

import (
  "crypto/sha256"
  "fmt"
)

type Block struct {
  Data string 
  Hash []byte 
  PreviousHash []byte
}

type MDBlockchain struct {
  chain []Block
}

func (b *MDBlockchain) addBlock (newBlock *Block) {
  // has the addBlocki
  h := sha256.New()
  mergedSlice := newBlock.Data + string(newBlock.PreviousHash)
  h.Write([]byte(mergedSlice))
  newBlock.Hash = h.Sum(nil) 
  b.chain = append(b.chain, *newBlock)
}

func check(err error) {
  if err != nil {
    fmt.Println(err)
    return
  }
}

func main() {
  var mdBlockchain MDBlockchain
  var genesisBlock Block
  genesisBlock.Data = "Genesis Block."
  genesisBlock.PreviousHash = []byte{'0'}
  mdBlockchain.addBlock(&genesisBlock)
  
  // Block1 
  var newBlock Block
  newBlock.Data = "Block 1"
  newBlock.PreviousHash = genesisBlock.Hash
  mdBlockchain.addBlock(&newBlock)

  var newBlock2 Block
  newBlock2.Data = "Block 2"
  newBlock2.PreviousHash = newBlock.Hash
  mdBlockchain.addBlock(&newBlock2)

  for _, block := range mdBlockchain.chain {
    fmt.Printf("%v : %x \n", block.Data, block.Hash)
  }
}
