package unix

// +build ignore

/*
Input to cgo -godefs.  See README.md
*/

/*
#include <sys/stat.h>
#include <sys/mount.h>
#include <sys/time.h>
*/
import "C"

type Stat_t C.struct_stat

type Statfs_t C.struct_statfs

// Time
type Timespec C.struct_timespec

type Timeval C.struct_timeval


const MNT_RDONLY=C.MS_RDONLY
const MNT_FSS=C.MS_FSS
const MNT_DATA=C.MS_DATA
const MNT_NOSUID=C.MS_NOSUID
const MNT_REMOUNT=C.MS_REMOUNT
const MNT_NOTRUNC=C.MS_NOTRUNC
const MNT_OVERLAY=C.MS_OVERLAY
const MNT_OPTIONSTR=C.MS_OPTIONSTR
const MNT_GLOBAL=C.MS_GLOBAL
const MNT_FORCE=C.MS_FORCE
const MNT_NOMNTTAB=C.MS_NOMNTTAB

// for compatibility with linux options
const MNT_DETACH=C.MS_FORCE
