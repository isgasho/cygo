#ifndef _NORO_GC_H_
#define _NOGO_GC_H_

#define USE_BDWGC

#ifdef USE_BDWGC

#include <gc/gc.h>
extern void GC_push_all_eager(void*, void*);
extern void GC_set_push_other_roots(void*);

const char* crn_gc_event_name(GC_EventType evty);

void* crn_gc_malloc(size_t size);
void* crn_gc_realloc(void* ptr, size_t size);
void crn_gc_free(void* ptr);

#endif

void* crn_raw_malloc(size_t size);
void* crn_raw_realloc(void* ptr, size_t size);
void crn_raw_free(void* ptr);

#define crn_malloc_st(/*typedesc*/st) (st*)crn_raw_malloc(sizeof(st))
#define convto(/*typedesc*/ st, /*var*/ var) (st*)(var)

#endif
