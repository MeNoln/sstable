package main

import (
	"os"

	"github.com/MeNoln/sstable/sstable"
)

func main() {
	testBlockWrite()
	rf, _ := os.Open("temp_test_blocks.sst")
	defer rf.Close()

	t, _ := sstable.NewTable(rf)

	val, _ := t.Search("middle")
	println(val)
}

func testBlockWrite() {
	blocks := sstable.NewBlocks()
	blocks = blocks.AppendBlock([]byte("zero"), []byte("zero1"))
	blocks = blocks.AppendBlock([]byte("middle"), []byte("middle1"))
	blocks = blocks.AppendBlock([]byte("at"), []byte("at1"))

	file, err := os.Create("temp_test_blocks.sst")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sstable.WriteTable(blocks, file)
}

func testBlockRead() {
	rf, _ := os.Open("temp_test_blocks.sst")
	defer rf.Close()

	_, _ = sstable.NewTable(rf)
}
