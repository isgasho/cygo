A library that implemention golang's core features,
goroutines, schedulers, channels, goselect and garbage collect.

### Requirements
* [ ] libgc-dev 8.0.4+

### Features

* [x] stackful coroutine
* [x] multiple threads
* [x] channels
* [x] golang select semantic
* [x] garbage collect
* [x] syscall hook
* [ ] explict syscall functions wrapper

### Usage

All exported functions is in `crnpub.h`

### Todos

* [x] channel select semantic
* [x] wait reason
* [ ] goroutines stats, count, memory
* [ ] goroutines stack info
* [ ] scheduler switch to goroutine
* [ ] native main function switch to goroutine
* [ ] improve send/recv bool flag
* [ ] mutex lock/unlock yield?
* [ ] dynamic increase/decrease processor(P)
* [x] sockfd timeout support
* [ ] dynamic increase/decrease stack size
* [ ] copy stack, copy stack 解决了栈大小溢出的问题了没？
* [x] detect stack size in use in yeild
* [ ] stat info for fibers/schedulers
* [ ] lock to os thread
* [ ] drop thread local storage usage


### Difference with Go
* Ours gosched is a sleep, Go's Gosched is long parking and wait resched
* Go have Sudog, we haven't

### Thirdpartys

Thanks all the contributors.

* libcoro 
* libchan https://github.com/tylertreat/chan
* libcollectc
* rixlog
* libgc >= 8.0.0
* libevent >= 2.1
* picoev
* libelf/plthook

### BUGS
* [x] SIGABRT: coro\_init () at /noro/coro.c:104
    when resume a finished fiber, this crash occurs
* [ ] SIGSEGV: bdwgc: mark.c:1581 GC\_push\_all\_eager
    https://www.mail-archive.com/ecls-list@lists.sourceforge.net/msg00161.html
* [ ] ASSERT: bdwgc: misc.c: 1986 GC\_disable
* [ ] gopackany:goroutine\_post, seems many times just after this: Assertion failure: extra/../misc.c:1986
* [ ] ==17317==  Access not within mapped region at address 0x0
    ==17317==    at 0x168C14: kind\_j8p3h8iojolY0RbF1nkaZgxmltree (stdlib\_xmltree.nim.c:310)
* [x] current curl upload data 100%, but wait long time, then throw error: Error in the HTTP2 framing layer
* [ ] libcurl slow response(26s). so wtf, 上网高峰出现，否则表现挺好的？？？并不，隔时间长了执行则消息，间隔时间短执行则出现。
* [ ] hang forever on ppoll ()
* [ ] sometimes GC will stop work
* [x] hang forever on __lll_lock_wait_private () from /usr/lib/libc.so.6
    occurs when call linfo/log write in push_other_roots callback
    says that linfo/log has some where not safe point
    related with signal handler
* [ ] GC_clear_fl_marks infinite loop
* [ ] set finalizer to non-empty would cause crash at GC_is_marked()???
   if manually GC_FREE(obj) it???
* [ ] GC_mark_from crash
    maybe stack corrupt reason
* [ ] ==24352==WARNING: ASan doesn't fully support makecontext/swapcontext functions and may produce false positives in some cases!
* [ ] https://blog.csdn.net/fyyyr/article/details/79389272
   安全点的选定基本上是以“是否具有让程序长时间执行的特征”为标准进行选定的  https://www.jianshu.com/p/952fc6c4501c
   1. 循环的末尾
   2. 方法临返回前 / 调用方法的call指令后
   3. 可能抛异常的位置   

### Note
* A program entering an infinite recursion or running out of space in the stack memory is known as a stack overflow

GC_NPROCS=1 ./prog to set gc thread count

### 同类
* https://github.com/canonical/libco pure C, single thread, no dynamic stack size
* 
