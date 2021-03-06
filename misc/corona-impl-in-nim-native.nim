{.passc:"-g -O0"}

import os
import asyncdispatch
import asyncfutures
#import threadpool
import tables
import deques
import locks

include "nimlog.nim"
include "nimplus.nim"
include "coro.nim"
include "hook.nim"

const dftstksz = 128 * 1024

proc GC_addStack(bottom: pointer) {.cdecl, importc.}
proc GC_removeStack(bottom: pointer) {.cdecl, importc.}
proc GC_setActiveStack(bottom: pointer) {.cdecl, importc.}
# proc GC_setStackBottom(bottom: pointer) {.cdecl, importc.}

type
    Grstate = enum
        Waiting, Runnable, Executing, Finished

    Stack = ref object
        sp*: pointer
        sz*: int

    PGoroutine = ptr Goroutine
    Goroutine = ref object
        id*: int
        fproc*: proc()
        stk*: Stack
        state*: Grstate

    PProcessor = ptr Processor
    Processor = ref object
        id*: int

    PMachine = ptr Machine
    Machine = ref object
        id*: int
        ngrs*: Table[int, proc()] # 新任务，未分配栈
        grs*: Table[int,PGoroutine] # grid => # Deque[PGoroutine]
        gr*: PGoroutine # 当前在执行的
        pklk*: Lock # park lock
        pkcd*: Cond

    PRootenv = ptr Rootenv
    Rootenv = ref object
        gridno*: int
        mths*: Table[int, Thread[PMachine]]
        mcs*: Table[int, Machine]
        grrefs*: Table[int, Goroutine] # 用于持有所有goroutine的引用, grid =>
        goroinited*: bool
        goroinitev*: AsyncEvent
        goroinitlk*: Lock
        goroinitcd*: Cond


const GOMAXPROC = 3 # testing with 3
proc goro_machine_new():Machine =
    var mc = new Machine
    mc.ngrs = initTable[int, proc()]()
    mc.grs = initTable[int, PGoroutine]()
    mc.pklk.initLock()
    mc.pkcd.initCond()
    return mc

proc goro_rootenv_new():Rootenv =
    var re = new Rootenv
    re.mcs = initTable[int, Machine]()
    re.mths = initTable[int, Thread[PMachine]]()
    re.grrefs = initTable[int, Goroutine]()
    re.goroinitev = newAsyncEvent()
    re.goroinitlk.initLock()
    re.goroinitcd.initCond()
    return re

var rtenv : Rootenv = goro_rootenv_new()
var pre  = rtenv.addr

var mcid {.threadvar.} : int

proc goro_nxtid() : int =
    pre.gridno += 1
    return pre.gridno

proc goro_create(id : int) : Goroutine =
    # pre.gridno += 1
    var gr = new Goroutine
    gr.id = id
    gr.stk = new Stack
    gr.stk.sp = allocShared0(dftstksz)
    gr.stk.sz = dftstksz
    return gr

# 添加到0 goro
proc goro_post0(fn:proc) =
    var gr = goro_create(goro_nxtid())
    pre.grrefs.add(gr.id, gr)
    var pgr = pre.grrefs[gr.id].addr
    var mc = pre.mcs[0]
    mc.grs.add(gr.id, pgr)
    mc.pkcd.signal()
    return

proc goro_post(fn:proc) =
    var id = goro_nxtid()
    var mc = pre.mcs[0]
    mc.ngrs.add(id, fn)
    mc.pkcd.signal()
    return

proc goro_remove(fn:proc) =

    return

proc goro_move(fm, to : int) =
    return

proc goro_yield() =
    return

proc goro_machine_init(id:int) =
    mcid = id
    return
proc goro_current_machine() : int = return mcid

proc processor_proc0(pm : PMachine) =
    goro_machine_init(pm.id)
    linfo "started...", pm.id, pre.mcs.len
    linfo "currmcid", goro_current_machine()
    linfo repr(pre.goroinitev)
    pre.goroinited = true
    pre.goroinitev.trigger()
    pre.goroinitcd.signal()
    linfo "proc0 signaled"
    while true:
        if pm.grs.len == 0: pm.pkcd.wait(pm.pklk)

        linfo "grcnt", pm.grs.len
        # move to ready queue
        for id, fn in pm.ngrs:
            var gr = goro_create(id)
            gr.state = Runnable
            pre.grrefs.add(id, gr)
            pm.grs.add(id, pre.grrefs[id].addr)

        # clean temp queue
        var grids : seq[int]
        for id,_ in pm.ngrs: grids.add(id)
        for id in grids: pm.ngrs.del(id)

        # find free m and ready task

    return

proc processor_proc(pm : PMachine) =
    goro_machine_init(pm.id)
    linfo "started...", pm.id, pre.mcs.len
    linfo "currmcid", goro_current_machine()
    while true: sleep(5000)
    return

proc goro_init() =
    var maxproc : int = GOMAXPROC
    if maxproc == 0: maxproc = 3
    for mid in 0..maxproc-1:
        var mc = goro_machine_new()
        mc.id = maxproc-1-mid
        pre.mcs.add(mid, mc)

        pre.mths.add(mid, Thread[PMachine]())
        if mid == 0:
            createThread(pre.mths[mid], processor_proc0, pre.mcs[mid].addr)
        else:
            createThread(pre.mths[mid], processor_proc, pre.mcs[mid].addr)
        # linfo "isrun?", running(pre.mths[mid])
    return

proc umain() =
    goro_post(proc() = echo 123)
    return

proc atrivaltofn(fd:AsyncFD):bool = return false
addTimer(5000, false, atrivaltofn)
proc ongoroinitdone(fd:AsyncFD):bool =
    linfo "goro inited done", fd
    return false
addEvent(pre.goroinitev, atrivaltofn) # why cannot catch the event?
linfo repr(pre.goroinitev)

goro_init()

linfo "wait proc0 ..."
if not pre.goroinited: pre.goroinitcd.wait(pre.goroinitlk)
linfo "goro inited done"
pre.goroinitcd.deinitCond()
pre.goroinitlk.deinitLock()

if isMainModule:
    umain()
    while true: poll(50000)

