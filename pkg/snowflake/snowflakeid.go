package snowflakeid

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

const (
	NodeBits       = 10                        // 节点 ID 的位数
	TimestampBits  = 41                        // 时间戳的位数
	SequenceBits   = 13                        // 序列号的位数
	NodeMax        = -1 ^ (-1 << NodeBits)     // 最大节点 ID
	SequenceMask   = -1 ^ (-1 << SequenceBits) // 序列号掩码
	NodeShift      = SequenceBits              // 节点 ID 左移位数
	TimestampShift = SequenceBits + NodeBits   // 时间戳左移位数
	EpochOffset    = 1622198400000             // 时间戳的起始偏移量（2021-05-29）
)

// SnowflakeIDGenerator 生成器
type SnowflakeIDGenerator struct {
	mu         sync.Mutex // 互斥锁
	nodeID     int64      // 节点 ID
	lastTime   int64      // 上次生成 ID 的时间戳
	sequence   int64      // 序列号
	randSource *rand.Rand // 随机数生成器
}

// NewSnowflakeIDGenerator 创建新的雪花算法 ID 生成器
func NewSnowflakeIDGenerator(nodeID int64) (*SnowflakeIDGenerator, error) {
	if nodeID < 0 || nodeID >= (1<<NodeBits) {
		return nil, errors.New("invalid node ID")
	}

	// 创建一个随机数生成器
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &SnowflakeIDGenerator{
		nodeID:     nodeID,
		lastTime:   -1,
		sequence:   0,
		randSource: randSource,
	}, nil
}

// Generate 生成一个新的 Snowflake ID
func (g *SnowflakeIDGenerator) Generate() (int64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 获取当前时间戳
	now := time.Now().UnixNano() / 1000000 // 毫秒

	// 如果时间回拨了，则返回错误
	if now < g.lastTime {
		return 0, errors.New("time is moving backwards")
	}

	// 如果是同一毫秒生成的，则增加序列号
	if now == g.lastTime {
		g.sequence++
	} else {
		// 如果是新的一毫秒，则重置序列号
		g.sequence = g.randSource.Int63n(1<<SequenceBits - 1) // 随机生成序列号
		g.lastTime = now
	}

	// 检查序列号是否超出范围
	if g.sequence >= 1<<SequenceBits {
		return 0, errors.New("sequence overflow")
	}

	// 生成 ID
	id := ((now - EpochOffset) << TimestampShift) |
		(g.nodeID << NodeShift) |
		g.sequence

	return id, nil
}
