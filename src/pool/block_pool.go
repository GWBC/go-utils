package pool

import (
	"sync"
	"sync/atomic"
)

type Block struct {
	Pkg           []byte //整个数据包
	PayloadOffset int    //有效载荷偏移

	pool     *sync.Pool
	refCount atomic.Int32
}

func (b *Block) init(size int, payloadOffset int, pool *sync.Pool) {
	b.refCount.Store(1)
	b.pool = pool

	b.Pkg = make([]byte, size)
	b.PayloadOffset = payloadOffset
}

func (b *Block) reset(pkgSize int) {
	b.refCount.Store(1)
	b.Pkg = b.Pkg[:pkgSize]
}

func (b *Block) AddRef() *Block {
	b.refCount.Add(1)
	return b
}

func (b *Block) Release() {
	if b.refCount.Add(-1) <= 0 {
		b.pool.Put(b)
	}
}

func (b *Block) SetPkgSize(size int) *Block {
	b.Pkg = b.Pkg[:size]
	return b
}

func (b *Block) SetPayload(data []byte) *Block {
	b.Pkg = b.Pkg[:b.PayloadOffset+len(data)]
	b.Pkg = append(b.Pkg[:b.PayloadOffset], data...)
	return b
}

func (b *Block) Payload() []byte {
	return b.Pkg[b.PayloadOffset:]
}

//////////////////////////////////////////////////////////////////

type BlockPool struct {
	pool          *sync.Pool
	BlockSize     int
	PayloadOffset int
}

func (c *BlockPool) init(blockSize int, payloadOffset int) {
	c.BlockSize = blockSize
	c.PayloadOffset = payloadOffset

	c.pool = &sync.Pool{}
	c.pool.New = func() any {
		blk := &Block{}
		blk.init(c.BlockSize, c.PayloadOffset, c.pool)
		return blk
	}
}

func (c *BlockPool) Get() *Block {
	blk := c.pool.Get().(*Block)
	blk.reset(c.BlockSize)

	return blk
}

//////////////////////////////////////////////////////////////////

func CreateBlockPool(blockSize int, payloadOffset int) *BlockPool {
	pool := &BlockPool{}
	pool.init(blockSize, payloadOffset)

	return pool
}
