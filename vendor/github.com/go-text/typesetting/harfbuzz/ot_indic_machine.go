package harfbuzz

// Code generated with ragel -Z -o ot_indic_machine.go ot_indic_machine.rl ; sed -i '/^\/\/line/ d' ot_indic_machine.go ; goimports -w ot_indic_machine.go  DO NOT EDIT.

// ported from harfbuzz/src/hb-ot-shape-complex-indic-machine.rl Copyright © 2015 Google, Inc. Behdad Esfahbod

// indic_syllable_type_t
const (
	indicConsonantSyllable = iota
	indicVowelSyllable
	indicStandaloneCluster
	indicSymbolCluster
	indicBrokenCluster
	indicNonIndicCluster
)

const indSM_ex_A = 9
const indSM_ex_C = 1
const indSM_ex_CM = 16
const indSM_ex_CS = 18
const indSM_ex_DOTTEDCIRCLE = 11
const indSM_ex_H = 4
const indSM_ex_M = 7
const indSM_ex_MPst = 13
const indSM_ex_N = 3
const indSM_ex_PLACEHOLDER = 10
const indSM_ex_RS = 12
const indSM_ex_Ra = 15
const indSM_ex_Repha = 14
const indSM_ex_SM = 8
const indSM_ex_Symbol = 17
const indSM_ex_V = 2
const indSM_ex_VD = 9
const indSM_ex_X = 0
const indSM_ex_ZWJ = 6
const indSM_ex_ZWNJ = 5

var _indSM_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 6,
	1, 7, 1, 8, 1, 9, 1, 10,
	1, 11, 1, 12, 1, 13, 1, 14,
	1, 15, 1, 16, 1, 17, 1, 18,
	2, 2, 3, 2, 2, 4, 2, 2,
	5,
}

var _indSM_key_offsets []int16 = []int16{
	0, 1, 7, 12, 17, 18, 24, 31,
	37, 38, 43, 48, 49, 55, 62, 69,
	76, 77, 82, 87, 88, 94, 100, 107,
	108, 113, 118, 119, 125, 131, 136, 137,
	155, 165, 174, 182, 188, 192, 195, 196,
	198, 205, 211, 217, 225, 232, 238, 242,
	249, 253, 258, 262, 271, 281, 291, 300,
	308, 314, 324, 333, 341, 347, 350, 351,
	353, 360, 366, 374, 381, 387, 391, 398,
	402, 406, 411, 415, 424, 434, 440, 449,
	458, 466, 472, 482, 488, 491, 492, 494,
	501, 507, 515, 522, 528, 532, 541, 548,
	552, 556, 561, 565, 575, 582, 588, 598,
	607, 615, 621, 631, 637, 640, 641, 643,
	650, 656, 664, 671, 677, 681, 690, 697,
	701, 705, 710, 714, 729, 739, 753, 761,
	765, 769, 770, 772, 782, 787, 791, 794,
	795, 797,
}

var _indSM_trans_keys []byte = []byte{
	8, 4, 7, 8, 13, 5, 6, 7,
	8, 13, 5, 6, 7, 8, 13, 5,
	6, 13, 4, 7, 8, 13, 5, 6,
	4, 7, 8, 12, 13, 5, 6, 4,
	7, 8, 13, 5, 6, 8, 7, 8,
	13, 5, 6, 7, 8, 13, 5, 6,
	13, 4, 7, 8, 13, 5, 6, 4,
	7, 8, 12, 13, 5, 6, 4, 7,
	8, 12, 13, 5, 6, 4, 7, 8,
	12, 13, 5, 6, 8, 7, 8, 13,
	5, 6, 7, 8, 13, 5, 6, 13,
	4, 7, 8, 13, 5, 6, 4, 7,
	8, 13, 5, 6, 4, 7, 8, 12,
	13, 5, 6, 8, 7, 8, 13, 5,
	6, 7, 8, 13, 5, 6, 13, 4,
	7, 8, 13, 5, 6, 4, 7, 8,
	13, 5, 6, 7, 8, 13, 5, 6,
	8, 1, 2, 3, 4, 5, 6, 7,
	8, 9, 12, 13, 14, 15, 16, 17,
	18, 10, 11, 3, 4, 5, 6, 7,
	8, 9, 12, 13, 16, 3, 4, 7,
	8, 9, 13, 16, 5, 6, 4, 7,
	8, 9, 13, 16, 5, 6, 1, 5,
	6, 8, 9, 15, 8, 9, 5, 6,
	5, 8, 9, 9, 5, 9, 1, 3,
	8, 9, 15, 5, 6, 1, 8, 9,
	15, 5, 6, 1, 5, 6, 8, 9,
	15, 3, 4, 7, 8, 9, 13, 5,
	6, 4, 7, 8, 9, 13, 5, 6,
	7, 8, 9, 13, 5, 6, 5, 8,
	9, 13, 4, 7, 8, 9, 13, 5,
	6, 5, 6, 8, 9, 3, 8, 9,
	5, 6, 5, 6, 8, 9, 3, 4,
	7, 8, 9, 13, 16, 5, 6, 3,
	4, 5, 6, 7, 8, 9, 12, 13,
	16, 3, 4, 5, 6, 7, 8, 9,
	12, 13, 16, 3, 4, 5, 6, 7,
	8, 9, 13, 16, 4, 5, 6, 7,
	8, 9, 13, 16, 1, 5, 6, 8,
	9, 15, 3, 4, 5, 6, 7, 8,
	9, 12, 13, 16, 3, 4, 7, 8,
	9, 13, 16, 5, 6, 4, 7, 8,
	9, 13, 16, 5, 6, 1, 5, 6,
	8, 9, 15, 5, 8, 9, 9, 5,
	9, 1, 3, 8, 9, 15, 5, 6,
	1, 8, 9, 15, 5, 6, 3, 4,
	7, 8, 9, 13, 5, 6, 4, 7,
	8, 9, 13, 5, 6, 7, 8, 9,
	13, 5, 6, 5, 8, 9, 13, 4,
	7, 8, 9, 13, 5, 6, 5, 6,
	8, 9, 8, 9, 5, 6, 3, 8,
	9, 5, 6, 5, 6, 8, 9, 3,
	4, 7, 8, 9, 13, 16, 5, 6,
	3, 4, 5, 6, 7, 8, 9, 12,
	13, 16, 4, 7, 8, 13, 5, 6,
	3, 4, 5, 6, 7, 8, 9, 13,
	16, 3, 4, 7, 8, 9, 13, 16,
	5, 6, 4, 7, 8, 9, 13, 16,
	5, 6, 1, 5, 6, 8, 9, 15,
	3, 4, 5, 6, 7, 8, 9, 12,
	13, 16, 1, 5, 6, 8, 9, 15,
	5, 8, 9, 9, 5, 9, 1, 3,
	8, 9, 15, 5, 6, 1, 8, 9,
	15, 5, 6, 3, 4, 7, 8, 9,
	13, 5, 6, 4, 7, 8, 9, 13,
	5, 6, 7, 8, 9, 13, 5, 6,
	5, 8, 9, 13, 3, 4, 7, 8,
	9, 13, 16, 5, 6, 4, 7, 8,
	9, 13, 5, 6, 5, 6, 8, 9,
	8, 9, 5, 6, 3, 8, 9, 5,
	6, 5, 6, 8, 9, 3, 4, 5,
	6, 7, 8, 9, 12, 13, 16, 4,
	7, 8, 12, 13, 5, 6, 4, 7,
	8, 13, 5, 6, 3, 4, 5, 6,
	7, 8, 9, 12, 13, 16, 3, 4,
	7, 8, 9, 13, 16, 5, 6, 4,
	7, 8, 9, 13, 16, 5, 6, 1,
	5, 6, 8, 9, 15, 3, 4, 5,
	6, 7, 8, 9, 12, 13, 16, 1,
	5, 6, 8, 9, 15, 5, 8, 9,
	9, 5, 9, 1, 3, 8, 9, 15,
	5, 6, 1, 8, 9, 15, 5, 6,
	3, 4, 7, 8, 9, 13, 5, 6,
	4, 7, 8, 9, 13, 5, 6, 7,
	8, 9, 13, 5, 6, 5, 8, 9,
	13, 3, 4, 7, 8, 9, 13, 16,
	5, 6, 4, 7, 8, 9, 13, 5,
	6, 5, 6, 8, 9, 8, 9, 5,
	6, 3, 8, 9, 5, 6, 5, 6,
	8, 9, 1, 2, 3, 4, 5, 6,
	7, 8, 9, 12, 13, 15, 16, 10,
	11, 3, 4, 5, 6, 7, 8, 9,
	12, 13, 16, 1, 2, 3, 4, 5,
	6, 7, 8, 9, 11, 12, 13, 15,
	16, 4, 7, 8, 9, 12, 13, 5,
	6, 5, 8, 9, 13, 5, 8, 9,
	13, 9, 5, 9, 1, 3, 4, 7,
	8, 9, 13, 15, 5, 6, 3, 8,
	9, 5, 6, 8, 9, 5, 6, 5,
	8, 9, 9, 5, 9, 1, 10, 15,
}

var _indSM_single_lengths []byte = []byte{
	1, 4, 3, 3, 1, 4, 5, 4,
	1, 3, 3, 1, 4, 5, 5, 5,
	1, 3, 3, 1, 4, 4, 5, 1,
	3, 3, 1, 4, 4, 3, 1, 16,
	10, 7, 6, 6, 2, 3, 1, 2,
	5, 4, 6, 6, 5, 4, 4, 5,
	4, 3, 4, 7, 10, 10, 9, 8,
	6, 10, 7, 6, 6, 3, 1, 2,
	5, 4, 6, 5, 4, 4, 5, 4,
	2, 3, 4, 7, 10, 4, 9, 7,
	6, 6, 10, 6, 3, 1, 2, 5,
	4, 6, 5, 4, 4, 7, 5, 4,
	2, 3, 4, 10, 5, 4, 10, 7,
	6, 6, 10, 6, 3, 1, 2, 5,
	4, 6, 5, 4, 4, 7, 5, 4,
	2, 3, 4, 13, 10, 14, 6, 4,
	4, 1, 2, 8, 3, 2, 3, 1,
	2, 3,
}

var _indSM_range_lengths []byte = []byte{
	0, 1, 1, 1, 0, 1, 1, 1,
	0, 1, 1, 0, 1, 1, 1, 1,
	0, 1, 1, 0, 1, 1, 1, 0,
	1, 1, 0, 1, 1, 1, 0, 1,
	0, 1, 1, 0, 1, 0, 0, 0,
	1, 1, 0, 1, 1, 1, 0, 1,
	0, 1, 0, 1, 0, 0, 0, 0,
	0, 0, 1, 1, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 0, 1, 0,
	1, 1, 0, 1, 0, 1, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 1,
	1, 1, 1, 1, 0, 1, 1, 0,
	1, 1, 0, 0, 1, 1, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 1,
	1, 1, 1, 1, 0, 1, 1, 0,
	1, 1, 0, 1, 0, 0, 1, 0,
	0, 0, 0, 1, 1, 1, 0, 0,
	0, 0,
}

var _indSM_index_offsets []int16 = []int16{
	0, 2, 8, 13, 18, 20, 26, 33,
	39, 41, 46, 51, 53, 59, 66, 73,
	80, 82, 87, 92, 94, 100, 106, 113,
	115, 120, 125, 127, 133, 139, 144, 146,
	164, 175, 184, 192, 199, 203, 207, 209,
	212, 219, 225, 232, 240, 247, 253, 258,
	265, 270, 275, 280, 289, 300, 311, 321,
	330, 337, 348, 357, 365, 372, 376, 378,
	381, 388, 394, 402, 409, 415, 420, 427,
	432, 436, 441, 446, 455, 466, 472, 482,
	491, 499, 506, 517, 524, 528, 530, 533,
	540, 546, 554, 561, 567, 572, 581, 588,
	593, 597, 602, 607, 618, 625, 631, 642,
	651, 659, 666, 677, 684, 688, 690, 693,
	700, 706, 714, 721, 727, 732, 741, 748,
	753, 757, 762, 767, 782, 793, 808, 816,
	821, 826, 828, 831, 841, 846, 850, 854,
	856, 859,
}

var _indSM_indicies []byte = []byte{
	1, 0, 2, 4, 5, 4, 3, 0,
	4, 6, 4, 3, 0, 4, 5, 4,
	3, 0, 4, 0, 7, 4, 5, 4,
	3, 0, 2, 4, 5, 8, 4, 3,
	0, 10, 12, 13, 12, 11, 9, 14,
	9, 12, 15, 12, 11, 9, 12, 13,
	12, 11, 9, 12, 9, 16, 12, 13,
	12, 11, 9, 10, 12, 13, 17, 12,
	11, 9, 10, 12, 13, 18, 12, 11,
	9, 20, 22, 23, 24, 22, 21, 19,
	25, 19, 22, 27, 22, 21, 26, 22,
	23, 22, 21, 19, 22, 26, 20, 22,
	23, 22, 21, 19, 28, 22, 23, 22,
	21, 19, 30, 32, 33, 34, 32, 31,
	29, 35, 29, 32, 36, 32, 31, 29,
	32, 33, 32, 31, 29, 32, 29, 30,
	32, 33, 32, 31, 29, 37, 32, 33,
	32, 31, 29, 22, 38, 22, 21, 0,
	40, 39, 42, 43, 44, 45, 46, 47,
	22, 23, 48, 24, 22, 50, 51, 52,
	53, 54, 49, 41, 56, 57, 58, 59,
	4, 5, 60, 8, 4, 61, 55, 62,
	57, 4, 5, 60, 4, 61, 63, 55,
	57, 4, 5, 60, 4, 61, 63, 55,
	42, 64, 65, 1, 60, 42, 55, 1,
	60, 66, 55, 60, 67, 60, 55, 60,
	55, 60, 60, 55, 42, 68, 1, 60,
	42, 66, 55, 42, 1, 60, 42, 66,
	55, 42, 66, 65, 1, 60, 42, 55,
	69, 70, 4, 5, 60, 4, 71, 55,
	70, 4, 5, 60, 4, 71, 55, 4,
	5, 60, 4, 71, 55, 60, 67, 60,
	4, 55, 72, 4, 5, 60, 4, 73,
	55, 64, 74, 1, 60, 55, 64, 1,
	60, 66, 55, 66, 74, 1, 60, 55,
	56, 57, 4, 5, 60, 4, 61, 63,
	55, 56, 57, 58, 63, 4, 5, 60,
	8, 4, 61, 55, 76, 77, 78, 79,
	12, 13, 80, 18, 12, 81, 75, 82,
	77, 83, 79, 12, 13, 80, 12, 81,
	75, 77, 83, 79, 12, 13, 80, 12,
	81, 75, 84, 85, 86, 14, 80, 84,
	75, 87, 77, 88, 89, 12, 13, 80,
	17, 12, 81, 75, 90, 77, 12, 13,
	80, 12, 81, 83, 75, 77, 12, 13,
	80, 12, 81, 83, 75, 84, 91, 86,
	14, 80, 84, 75, 80, 92, 80, 75,
	80, 75, 80, 80, 75, 84, 93, 14,
	80, 84, 91, 75, 84, 14, 80, 84,
	91, 75, 94, 95, 12, 13, 80, 12,
	96, 75, 95, 12, 13, 80, 12, 96,
	75, 12, 13, 80, 12, 96, 75, 80,
	92, 80, 12, 75, 97, 12, 13, 80,
	12, 98, 75, 85, 99, 14, 80, 75,
	14, 80, 91, 75, 85, 14, 80, 91,
	75, 91, 99, 14, 80, 75, 87, 77,
	12, 13, 80, 12, 81, 83, 75, 87,
	77, 88, 83, 12, 13, 80, 17, 12,
	81, 75, 10, 12, 13, 12, 11, 75,
	76, 77, 83, 79, 12, 13, 80, 12,
	81, 75, 101, 45, 22, 23, 48, 22,
	52, 102, 100, 45, 22, 23, 48, 22,
	52, 102, 100, 103, 104, 105, 25, 48,
	103, 100, 44, 45, 106, 107, 22, 23,
	48, 24, 22, 52, 100, 103, 108, 105,
	25, 48, 103, 100, 48, 109, 48, 100,
	48, 100, 48, 48, 100, 103, 110, 25,
	48, 103, 108, 100, 103, 25, 48, 103,
	108, 100, 111, 112, 22, 23, 48, 22,
	113, 100, 112, 22, 23, 48, 22, 113,
	100, 22, 23, 48, 22, 113, 100, 48,
	109, 48, 22, 100, 44, 45, 22, 23,
	48, 22, 52, 102, 100, 114, 22, 23,
	48, 22, 115, 100, 104, 116, 25, 48,
	100, 25, 48, 108, 100, 104, 25, 48,
	108, 100, 108, 116, 25, 48, 100, 44,
	45, 106, 102, 22, 23, 48, 24, 22,
	52, 100, 20, 22, 23, 24, 22, 21,
	117, 20, 22, 23, 22, 21, 117, 119,
	120, 121, 122, 32, 33, 123, 34, 32,
	124, 118, 125, 120, 32, 33, 123, 32,
	124, 122, 118, 120, 32, 33, 123, 32,
	124, 122, 118, 126, 127, 128, 35, 123,
	126, 118, 119, 120, 121, 49, 32, 33,
	123, 34, 32, 124, 118, 126, 129, 128,
	35, 123, 126, 118, 123, 130, 123, 118,
	123, 118, 123, 123, 118, 126, 131, 35,
	123, 126, 129, 118, 126, 35, 123, 126,
	129, 118, 132, 133, 32, 33, 123, 32,
	134, 118, 133, 32, 33, 123, 32, 134,
	118, 32, 33, 123, 32, 134, 118, 123,
	130, 123, 32, 118, 119, 120, 32, 33,
	123, 32, 124, 122, 118, 135, 32, 33,
	123, 32, 136, 118, 127, 137, 35, 123,
	118, 35, 123, 129, 118, 127, 35, 123,
	129, 118, 129, 137, 35, 123, 118, 42,
	43, 44, 45, 106, 102, 22, 23, 48,
	24, 22, 42, 52, 49, 100, 56, 138,
	58, 59, 4, 5, 60, 8, 4, 61,
	55, 42, 43, 44, 45, 139, 140, 22,
	141, 142, 49, 24, 22, 42, 52, 55,
	20, 22, 141, 60, 24, 22, 143, 55,
	60, 67, 60, 22, 55, 142, 144, 142,
	22, 55, 142, 55, 142, 142, 55, 42,
	68, 20, 22, 141, 60, 22, 42, 143,
	55, 146, 40, 148, 147, 145, 40, 148,
	147, 145, 148, 149, 148, 145, 148, 145,
	148, 148, 145, 42, 49, 42, 117,
}

var _indSM_trans_targs []byte = []byte{
	31, 37, 42, 2, 43, 46, 4, 50,
	51, 31, 60, 9, 66, 69, 61, 11,
	74, 75, 78, 31, 83, 17, 89, 92,
	93, 84, 31, 19, 98, 31, 107, 24,
	113, 116, 117, 108, 26, 122, 127, 31,
	134, 31, 32, 53, 79, 81, 100, 101,
	85, 102, 123, 124, 94, 132, 137, 31,
	33, 35, 6, 52, 38, 47, 34, 1,
	36, 40, 0, 39, 41, 44, 45, 3,
	48, 5, 49, 31, 54, 56, 14, 77,
	62, 70, 55, 7, 57, 72, 64, 58,
	13, 76, 59, 8, 63, 65, 67, 68,
	10, 71, 12, 73, 31, 80, 20, 82,
	96, 87, 15, 99, 16, 86, 88, 90,
	91, 18, 95, 21, 97, 31, 31, 103,
	105, 22, 27, 109, 118, 104, 106, 120,
	111, 23, 110, 112, 114, 115, 25, 119,
	28, 121, 125, 126, 131, 128, 129, 29,
	130, 31, 133, 30, 135, 136,
}

var _indSM_trans_actions []byte = []byte{
	21, 0, 5, 0, 5, 0, 0, 5,
	5, 23, 5, 0, 5, 0, 0, 0,
	5, 5, 5, 29, 5, 0, 36, 0,
	36, 0, 31, 0, 5, 25, 5, 0,
	5, 0, 5, 0, 0, 5, 0, 27,
	0, 7, 5, 5, 36, 0, 39, 39,
	0, 5, 36, 5, 36, 5, 0, 9,
	5, 0, 0, 5, 0, 5, 5, 0,
	5, 5, 0, 0, 5, 5, 5, 0,
	0, 0, 5, 11, 5, 0, 0, 5,
	0, 5, 5, 0, 5, 5, 5, 5,
	0, 5, 5, 0, 0, 5, 5, 5,
	0, 0, 0, 5, 17, 36, 0, 36,
	5, 5, 0, 36, 0, 0, 5, 36,
	36, 0, 0, 0, 5, 19, 13, 5,
	0, 0, 0, 0, 5, 5, 5, 5,
	5, 0, 0, 5, 5, 5, 0, 0,
	0, 5, 0, 33, 33, 0, 0, 0,
	0, 15, 5, 0, 0, 0,
}

var _indSM_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0,
}

var _indSM_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 3,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0,
}

var _indSM_eof_trans []int16 = []int16{
	1, 1, 1, 1, 1, 1, 1, 10,
	10, 10, 10, 10, 10, 10, 10, 20,
	20, 27, 20, 27, 20, 20, 30, 30,
	30, 30, 30, 30, 30, 1, 40, 0,
	56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 76, 76, 76,
	76, 76, 76, 76, 76, 76, 76, 76,
	76, 76, 76, 76, 76, 76, 76, 76,
	76, 76, 76, 76, 76, 76, 76, 101,
	101, 101, 101, 101, 101, 101, 101, 101,
	101, 101, 101, 101, 101, 101, 101, 101,
	101, 101, 101, 101, 118, 118, 119, 119,
	119, 119, 119, 119, 119, 119, 119, 119,
	119, 119, 119, 119, 119, 119, 119, 119,
	119, 119, 119, 101, 56, 56, 56, 56,
	56, 56, 56, 56, 146, 146, 146, 146,
	146, 118,
}

const indSM_start int = 31
const indSM_first_final int = 31
const indSM_error int = -1

const indSM_en_main int = 31

func findSyllablesIndic(buffer *Buffer) {
	var p, ts, te, act, cs int
	info := buffer.Info

	{
		cs = indSM_start
		ts = 0
		te = 0
		act = 0
	}

	pe := len(info)
	eof := pe

	var syllableSerial uint8 = 1

	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if p == pe {
			goto _test_eof
		}
	_resume:
		_acts = int(_indSM_from_state_actions[cs])
		_nacts = uint(_indSM_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _indSM_actions[_acts-1] {
			case 1:
				ts = p

			}
		}

		_keys = int(_indSM_key_offsets[cs])
		_trans = int(_indSM_index_offsets[cs])

		_klen = int(_indSM_single_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case (info[p].complexCategory) < _indSM_trans_keys[_mid]:
					_upper = _mid - 1
				case (info[p].complexCategory) > _indSM_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_indSM_range_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case (info[p].complexCategory) < _indSM_trans_keys[_mid]:
					_upper = _mid - 2
				case (info[p].complexCategory) > _indSM_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_indSM_indicies[_trans])
	_eof_trans:
		cs = int(_indSM_trans_targs[_trans])

		if _indSM_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_indSM_trans_actions[_trans])
		_nacts = uint(_indSM_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _indSM_actions[_acts-1] {
			case 2:
				te = p + 1

			case 3:
				act = 1
			case 4:
				act = 5
			case 5:
				act = 6
			case 6:
				te = p + 1
				{
					foundSyllableIndic(indicNonIndicCluster, ts, te, info, &syllableSerial)
				}
			case 7:
				te = p
				p--
				{
					foundSyllableIndic(indicConsonantSyllable, ts, te, info, &syllableSerial)
				}
			case 8:
				te = p
				p--
				{
					foundSyllableIndic(indicVowelSyllable, ts, te, info, &syllableSerial)
				}
			case 9:
				te = p
				p--
				{
					foundSyllableIndic(indicStandaloneCluster, ts, te, info, &syllableSerial)
				}
			case 10:
				te = p
				p--
				{
					foundSyllableIndic(indicSymbolCluster, ts, te, info, &syllableSerial)
				}
			case 11:
				te = p
				p--
				{
					foundSyllableIndic(indicBrokenCluster, ts, te, info, &syllableSerial)
					buffer.scratchFlags |= bsfHasBrokenSyllable
				}
			case 12:
				te = p
				p--
				{
					foundSyllableIndic(indicNonIndicCluster, ts, te, info, &syllableSerial)
				}
			case 13:
				p = (te) - 1
				{
					foundSyllableIndic(indicConsonantSyllable, ts, te, info, &syllableSerial)
				}
			case 14:
				p = (te) - 1
				{
					foundSyllableIndic(indicVowelSyllable, ts, te, info, &syllableSerial)
				}
			case 15:
				p = (te) - 1
				{
					foundSyllableIndic(indicStandaloneCluster, ts, te, info, &syllableSerial)
				}
			case 16:
				p = (te) - 1
				{
					foundSyllableIndic(indicSymbolCluster, ts, te, info, &syllableSerial)
				}
			case 17:
				p = (te) - 1
				{
					foundSyllableIndic(indicBrokenCluster, ts, te, info, &syllableSerial)
					buffer.scratchFlags |= bsfHasBrokenSyllable
				}
			case 18:
				switch act {
				case 1:
					{
						p = (te) - 1
						foundSyllableIndic(indicConsonantSyllable, ts, te, info, &syllableSerial)
					}
				case 5:
					{
						p = (te) - 1
						foundSyllableIndic(indicBrokenCluster, ts, te, info, &syllableSerial)
						buffer.scratchFlags |= bsfHasBrokenSyllable
					}
				case 6:
					{
						p = (te) - 1
						foundSyllableIndic(indicNonIndicCluster, ts, te, info, &syllableSerial)
					}
				}

			}
		}

	_again:
		_acts = int(_indSM_to_state_actions[cs])
		_nacts = uint(_indSM_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _indSM_actions[_acts-1] {
			case 0:
				ts = 0

			}
		}

		p++
		if p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
		if p == eof {
			if _indSM_eof_trans[cs] > 0 {
				_trans = int(_indSM_eof_trans[cs] - 1)
				goto _eof_trans
			}
		}

	}

	_ = act // needed by Ragel, but unused
}
