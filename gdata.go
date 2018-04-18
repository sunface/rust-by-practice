package dotray

import "sync"

var lock = &sync.Mutex{}

const maxBackupSeedLen = 10
