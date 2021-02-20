package config

import (
    "time"
    "log"
)

const TICK_RATE = time.Minute * 2
const DEBUG_TICK_RATE = time.Second * 1

var DEBUG         bool = false
var DEBUG_VERBOSE bool = false
var DEBUG_STOCK   bool = false
var DEBUG_PERF    bool = false

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

func PerfLog(msg string) {
    if DEBUG_PERF {
        genLog("PERF", msg)
    }
}

func genLog(flag string, msg string) {
    log.Printf("[%s]%s\n", flag, msg)
}

func EnableDebug() {
    DEBUG = true
}

func EnableAllDebug() {
    DEBUG         = true
    DEBUG_VERBOSE = true
    DEBUG_STOCK   = true
    DEBUG_PERF    = true
}

func Ticker() *time.Ticker {
    if DEBUG {
            return time.NewTicker(DEBUG_TICK_RATE)
    } else {
            return time.NewTicker(TICK_RATE)
    }
}
