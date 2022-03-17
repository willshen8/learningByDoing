package block

func isBlockValid(newBlock, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// replaceChain replaces the old blockchain with new one
func (b *Blockchain) replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(b.Blockchain) {
		b.Blockchain = newBlocks
	}
}
