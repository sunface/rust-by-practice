package dotray

import "sync"

var node *Node
var lock = &sync.Mutex{}
var reqID int64 = 0

const maxBackupSeedLen = 10

const seedMaxRetry = 3 // seed节点重试次数超过20的，在收到新的backupSeeds后，优先被替换

const syncBackupSeedInterval = 30 // X秒同步一次备份种子信息
