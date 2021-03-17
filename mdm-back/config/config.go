package config

import (
    "time"
    "log"
)

const TICK_RATE = time.Second
const DEBUG_TICK_RATE = time.Millisecond * 100

var DEBUG         bool = false
var DEBUG_VERBOSE bool = false
var DEBUG_STOCK   bool = false
var DEBUG_PERF    bool = false
var DEBUG_FLAG    bool = false

func DebugLog(msg string) {
    if DEBUG {
        genLog("DEBG", msg)
    }
}

func VerboseLog(msg string) {
    if DEBUG_VERBOSE {
        genLog("VERB", msg)
    }
}

func StockLog(msg string) {
    if DEBUG_STOCK {
        genLog("STCK", msg)
    }
}

func PerfLog(msg string) {
    if DEBUG_PERF {
        genLog("PERF", msg)
    }
}

func FlagLog(msg string) {
    if DEBUG_FLAG {
        genLog("FLAG", msg)
    }
}

func MainLog(msg string) {
    genLog("Main", msg)
}

func TestLog(msg string) {
    log.Println("")
    genLog("!!TEST!!", msg)
}

func genLog(flag string, msg string) {
    // Formatting check; If it starts with [ don't use a space becuase there's more tags
    if msg[0] == byte('[') {
        log.Printf("[%s]%s\n", flag, msg)
    } else {
        log.Printf("[%s] %s\n", flag, msg)
    }
}

func EnableDebug() {
    DEBUG = true
}

func EnableAllDebug() {
    DEBUG         = true
    DEBUG_VERBOSE = true
    DEBUG_STOCK   = true
    DEBUG_PERF    = true
    DEBUG_FLAG    = true
}

func Ticker() *time.Ticker {
    if DEBUG_STOCK {
            return time.NewTicker(DEBUG_TICK_RATE)
    } else {
            return time.NewTicker(TICK_RATE)
    }
}
