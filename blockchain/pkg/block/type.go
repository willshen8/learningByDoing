package block

type Blockchain struct {
	Blockchain []Block
}

type Block struct {
	Index         int // position of the record in the blockchain
	Timestamp     string
	TransactionID int
	Hash          string // SHA256 representation of the record
	PrevHash      string // Prev hash record
}
