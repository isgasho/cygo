#ifndef _YIELD_TYPES_H_
#define _YIELD_TYPES_H_

typedef enum yield_type
    {
     YIELD_TYPE_NONE = 0,
     YIELD_TYPE_CHAN_SEND,
     YIELD_TYPE_CHAN_RECV,
     YIELD_TYPE_CHAN_SELECT,
     YIELD_TYPE_CHAN_SELECT_NOCASE,
     YIELD_TYPE_CONNECT,
     YIELD_TYPE_READ,
     YIELD_TYPE_READV,
     YIELD_TYPE_RECV,
     YIELD_TYPE_RECVFROM,
     YIELD_TYPE_WRITE,
     YIELD_TYPE_WRITEV,
     YIELD_TYPE_SEND,
     YIELD_TYPE_SENDTO,
     YIELD_TYPE_SENDMSG,

     YIELD_TYPE_POLL,
     YIELD_TYPE_SELECT,
     YIELD_TYPE_ACCEPT,

     YIELD_TYPE_SLEEP,
     YIELD_TYPE_USLEEP,
     YIELD_TYPE_NANOSLEEP,

     YIELD_TYPE_GETHOSTBYNAMER,
     YIELD_TYPE_GETHOSTBYNAME2R,
     YIELD_TYPE_GETHOSTBYADDR,

     YIELD_TYPE_MAX,
    } yeield_type;


#endif

