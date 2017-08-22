package nanovgo

const (
	nvgInitFontImageSize = 512
	nvgMaxFontImageSize  = 2048
	nvgMaxFontImages     = 4

	nvgInitCommandsSize = 256
	nvgInitPointsSize   = 128
	nvgInitPathsSize    = 16
	nvgInitVertsSize    = 256
	nvgMaxStates        = 32
)

type nvgBlendFactor int

const (
	nvgZERO = 1<<0
	nvgONE = 1<<1
	nvgSRC_COLOR = 1<<2
	nvgONE_MINUS_SRC_COLOR = 1<<3
	nvgDST_COLOR = 1<<4
	nvgONE_MINUS_DST_COLOR = 1<<5
	nvgSRC_ALPHA = 1<<6
	nvgONE_MINUS_SRC_ALPHA = 1<<7
	nvgDST_ALPHA = 1<<8
	nvgONE_MINUS_DST_ALPHA = 1<<9
	nvgSRC_ALPHA_SATURATE = 1<<10
)

type nvgCompositeOperation int

const (
	nvgSOURCE_OVER nvgCompositeOperation = iota
	nvgSOURCE_IN
	nvgSOURCE_OUT
	nvgATOP
	nvgDESTINATION_OVER
	nvgDESTINATION_IN
	nvgDESTINATION_OUT
	nvgDESTINATION_ATOP
	nvgLIGHTER
	nvgCOPY
	nvgXOR
)

type nvgCommands int

const (
	nvgMOVETO nvgCommands = iota
	nvgLINETO
	nvgBEZIERTO
	nvgCLOSE
	nvgWINDING
)

type nvgPointFlags int

const (
	nvgPtCORNER     nvgPointFlags = 0x01
	nvgPtLEFT       nvgPointFlags = 0x02
	nvgPtBEVEL      nvgPointFlags = 0x04
	nvgPrINNERBEVEL nvgPointFlags = 0x08
)

type nvgTextureType int

const (
	nvgTextureALPHA nvgTextureType = 1
	nvgTextureRGBA  nvgTextureType = 2
)

type nvgCodePointSize int

const (
	nvgNEWLINE nvgCodePointSize = iota
	nvgSPACE
	nvgCHAR
)
