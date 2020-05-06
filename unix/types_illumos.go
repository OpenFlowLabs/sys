// Copyright 2020 OpenFlowLabs GmbH. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

/*
Input to cgo -godefs.  See README.md
*/

package unix

/*
#include <sys/mount.h>
*/
import "C"

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
