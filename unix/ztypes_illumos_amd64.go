// cgo -godefs types_illumos.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build amd64,illumos

package unix

const MNT_RDONLY = 0x1
const MNT_FSS = 0x2
const MNT_DATA = 0x4
const MNT_NOSUID = 0x10
const MNT_REMOUNT = 0x20
const MNT_NOTRUNC = 0x40
const MNT_OVERLAY = 0x80
const MNT_OPTIONSTR = 0x100
const MNT_GLOBAL = 0x200
const MNT_FORCE = 0x400
const MNT_NOMNTTAB = 0x800

const MNT_DETACH = 0x400
