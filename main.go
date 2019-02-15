package main

import "BlockchainProject/block"

func main() {
	bc := block.NewBlockchain()
	defer bc.Db.Close()

	cli := &block.CLI{bc}
	cli.Run()
}
