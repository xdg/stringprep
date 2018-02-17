// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableD1 = Table{
	RuneRange{0x05BE, 0x05BE},
	RuneRange{0x05C0, 0x05C0},
	RuneRange{0x05C3, 0x05C3},
	RuneRange{0x05D0, 0x05EA},
	RuneRange{0x05F0, 0x05F4},
	RuneRange{0x061B, 0x061B},
	RuneRange{0x061F, 0x061F},
	RuneRange{0x0621, 0x063A},
	RuneRange{0x0640, 0x064A},
	RuneRange{0x066D, 0x066F},
	RuneRange{0x0671, 0x06D5},
	RuneRange{0x06DD, 0x06DD},
	RuneRange{0x06E5, 0x06E6},
	RuneRange{0x06FA, 0x06FE},
	RuneRange{0x0700, 0x070D},
	RuneRange{0x0710, 0x0710},
	RuneRange{0x0712, 0x072C},
	RuneRange{0x0780, 0x07A5},
	RuneRange{0x07B1, 0x07B1},
	RuneRange{0x200F, 0x200F},
	RuneRange{0xFB1D, 0xFB1D},
	RuneRange{0xFB1F, 0xFB28},
	RuneRange{0xFB2A, 0xFB36},
	RuneRange{0xFB38, 0xFB3C},
	RuneRange{0xFB3E, 0xFB3E},
	RuneRange{0xFB40, 0xFB41},
	RuneRange{0xFB43, 0xFB44},
	RuneRange{0xFB46, 0xFBB1},
	RuneRange{0xFBD3, 0xFD3D},
	RuneRange{0xFD50, 0xFD8F},
	RuneRange{0xFD92, 0xFDC7},
	RuneRange{0xFDF0, 0xFDFC},
	RuneRange{0xFE70, 0xFE74},
	RuneRange{0xFE76, 0xFEFC},
}

// TableD1 represents RFC-3454 Table D.1.
var TableD1 Table = tableD1
