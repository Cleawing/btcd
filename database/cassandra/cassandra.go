package cassandra
import (
	"github.com/btcsuite/btcd/database"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/hailocab/gocassa"
)

type Cassandra struct {
	keySpace gocassa.KeySpace
}

func (db *Cassandra) Close() error {
	return nil
}

func (db *Cassandra) DeleteAddrIndex() error {
	return database.ErrNotImplemented
}

func (db *Cassandra) DropAfterBlockBySha(sha *wire.ShaHash) error {
	return database.ErrNotImplemented
}

func (db *Cassandra) ExistsSha(sha *wire.ShaHash) (bool, error) {
	return false, database.ErrNotImplemented
}

func (db *Cassandra) ExistsTxSha(sha *wire.ShaHash) (bool, error) {
	return false, database.ErrNotImplemented
}

func (db *Cassandra) FetchAddrIndexTip() (*wire.ShaHash, int32, error) {
	return nil, 0, database.ErrNotImplemented
}

func (db *Cassandra) FetchBlockBySha(sha *wire.ShaHash) (*btcutil.Block, error) {
	return nil, database.ErrNotImplemented
}

func (db *Cassandra) FetchBlockHeaderBySha(sha *wire.ShaHash) (*wire.BlockHeader, error) {
	return nil, database.ErrNotImplemented
}

func (db *Cassandra) FetchBlockHeightBySha(sha *wire.ShaHash) (int32, error) {
	return 0, database.ErrNotImplemented
}

func (db *Cassandra) FetchBlockShaByHeight(height int32) (*wire.ShaHash, error) {
	return nil, database.ErrNotImplemented
}

func (db *Cassandra) FetchHeightRange(startHeight, endHeight int32) ([]wire.ShaHash, error) {
	return nil, database.ErrNotImplemented
}

func (db *Cassandra) FetchTxBySha(txHash *wire.ShaHash) ([]*database.TxListReply, error) {
	return nil, database.ErrNotImplemented
}

func (db *Cassandra) FetchTxByShaList(txShaList []*wire.ShaHash) []*database.TxListReply {
	return db.fetchTxByShaList(txShaList, true)
}

func (db *Cassandra) FetchTxsForAddr(btcutil.Address, int, int) ([]*database.TxListReply, error) {
	return nil, database.ErrNotImplemented
}

func (db *Cassandra) FetchUnSpentTxByShaList(txShaList []*wire.ShaHash) []*database.TxListReply {
	return db.fetchTxByShaList(txShaList, false)
}

func (db *Cassandra) InsertBlock(block *btcutil.Block) (int32, error) {
	return 0, database.ErrNotImplemented
}
func (db *Cassandra) NewestSha() (*wire.ShaHash, int32, error) {
	return nil, 0, database.ErrNotImplemented
}

func (db *Cassandra) RollbackClose() error {
	return db.Close()
}

func (db *Cassandra) Sync() error {
	return nil
}

func (db *Cassandra) UpdateAddrIndexForBlock(*wire.ShaHash, int32, database.BlockAddrIndex) error {
	return database.ErrNotImplemented
}


func (db *Cassandra) fetchTxByShaList(txShaList []*wire.ShaHash, includeSpent bool) []*database.TxListReply {
	replyList := make([]*database.TxListReply, 0, len(txShaList))
//	for i, hash := range txShaList {
	for i, _ := range txShaList {
		reply := database.TxListReply{
			Sha: txShaList[i],
			Err: database.ErrTxShaMissing,
		}
		replyList = append(replyList, &reply)
		reply.Err = database.ErrNotImplemented
		continue
	}
	return replyList
}
