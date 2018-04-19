package dotray

import "sync"

var node *Node
var lock = &sync.Mutex{}
var reqID int64

const maxBackupSeedLen = 10 // the max length of the seed backups

const seedMaxRetry = 3 // the max retry times when a seed failed to connect

const syncBackupSeedInterval = 30 // the seed backup interval

const pingInterval = 30 // second
const maxPingAllowed = 8

const maxBackupSeedAlive = 240

const maxResendStayTime = 120
const minResendStayTime = 20
