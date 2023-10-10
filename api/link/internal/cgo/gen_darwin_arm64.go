// Code generated by gen/gen.go. DO NOT EDIT.
package cgo

import "reflect"

var Constants = constants{
	CHAR_BIT:                0x8,
	MB_LEN_MAX:              0x6,
	CHAR_MIN:                -128,
	CHAR_MAX:                0x7f,
	SCHAR_MIN:               -128,
	SHRT_MIN:                -32768,
	INT_MIN:                 -2147483648,
	LONG_MIN:                -9223372036854775808,
	LLONG_MIN:               -9223372036854775808,
	SCHAR_MAX:               0x7f,
	SHRT_MAX:                0x7fff,
	INT_MAX:                 0x7fffffff,
	LONG_MAX:                0x7fffffffffffffff,
	LLONG_MAX:               0x7fffffffffffffff,
	UCHAR_MAX:               0xff,
	USHRT_MAX:               0xffff,
	UINT_MAX:                0xffffffff,
	ULONG_MAX:               0xffffffffffffffff,
	ULLONG_MAX:              0xffffffffffffffff,
	PTRDIFF_MIN:             -9223372036854775808,
	PTRDIFF_MAX:             0x7fffffffffffffff,
	SIZE_MAX:                0xffffffffffffffff,
	WINT_MIN:                -2147483648,
	WINT_MAX:                0x7fffffff,
	WCHAR_MIN:               -2147483648,
	WCHAR_MAX:               0x7fffffff,
	INT8_MIN:                -128,
	INT16_MIN:               -32768,
	INT32_MIN:               -2147483648,
	INT64_MIN:               -9223372036854775808,
	INT8_MAX:                0x7f,
	INT16_MAX:               0x7fff,
	INT32_MAX:               0x7fffffff,
	INT64_MAX:               0x7fffffffffffffff,
	UINT8_MAX:               0xff,
	UINT16_MAX:              0xffff,
	UINT32_MAX:              0xffffffff,
	UINT64_MAX:              0xffffffffffffffff,
	INT_FAST8_MIN:           -128,
	INT_FAST16_MIN:          -32768,
	INT_FAST32_MIN:          -2147483648,
	INT_FAST64_MIN:          -9223372036854775808,
	INT_FAST8_MAX:           0x7f,
	INT_FAST16_MAX:          0x7fff,
	INT_FAST32_MAX:          0x7fffffff,
	INT_FAST64_MAX:          0x7fffffffffffffff,
	UINT_FAST8_MAX:          0xff,
	UINT_FAST16_MAX:         0xffff,
	UINT_FAST32_MAX:         0xffffffff,
	UINT_FAST64_MAX:         0xffffffffffffffff,
	INT_LEAST8_MIN:          -128,
	INT_LEAST16_MIN:         -32768,
	INT_LEAST32_MIN:         -2147483648,
	INT_LEAST64_MIN:         -9223372036854775808,
	INT_LEAST8_MAX:          0x7f,
	INT_LEAST16_MAX:         0x7fff,
	INT_LEAST32_MAX:         0x7fffffff,
	INT_LEAST64_MAX:         0x7fffffffffffffff,
	UINT_LEAST8_MAX:         0xff,
	UINT_LEAST16_MAX:        0xffff,
	UINT_LEAST32_MAX:        0xffffffff,
	UINT_LEAST64_MAX:        0xffffffffffffffff,
	INTMAX_MIN:              -9223372036854775808,
	INTMAX_MAX:              0x7fffffffffffffff,
	UINTMAX_MAX:             0xffffffffffffffff,
	INTPTR_MIN:              -9223372036854775808,
	INTPTR_MAX:              0x7fffffffffffffff,
	UINTPTR_MAX:             0xffffffffffffffff,
	SIG_ATOMIMIN:            -2147483648,
	SIG_ATOMIMAX:            0x7fffffff,
	FLT_RADIX:               0x2,
	DECIMAL_DIG:             0x11,
	FLT_DECIMAL_DIG:         0x9,
	DBL_DECIMAL_DIG:         0x11,
	LDBL_DECIMAL_DIG:        0x11,
	FLT_MIN:                 0,
	DBL_MIN:                 0,
	LDBL_MIN:                "2.225074e-308",
	FLT_TRUE_MIN:            0,
	DBL_TRUE_MIN:            0,
	LDBL_TRUE_MIN:           "4.940656e-324",
	FLT_MAX:                 3.4028234663852886e+38,
	DBL_MAX:                 1.7976931348623157e+308,
	LDBL_MAX:                "1.797693e+308",
	FLT_EPSILON:             0,
	DBL_EPSILON:             0,
	LDBL_EPSILON:            0,
	FLT_DIG:                 0x6,
	DBL_DIG:                 0xf,
	LDBL_DIG:                0xf,
	FLT_MANT_DIG:            0x18,
	DBL_MANT_DIG:            0x35,
	LDBL_MANT_DIG:           0x35,
	FLT_MIN_EXP:             -125,
	DBL_MIN_EXP:             -1021,
	LDBL_MIN_EXP:            -1021,
	FLT_MIN_10_EXP:          -37,
	DBL_MIN_10_EXP:          -307,
	LDBL_MIN_10_EXP:         -307,
	FLT_MAX_EXP:             0x80,
	DBL_MAX_EXP:             0x400,
	LDBL_MAX_EXP:            0x400,
	FLT_MAX_10_EXP:          0x26,
	DBL_MAX_10_EXP:          0x134,
	LDBL_MAX_10_EXP:         0x134,
	FLT_ROUNDS:              1,
	FLT_EVAL_METHOD:         0,
	FLT_HAS_SUBNORM:         1,
	DBL_HAS_SUBNORM:         1,
	LDBL_HAS_SUBNORM:        1,
	EDOM:                    33,
	ERANGE:                  34,
	EILSEQ:                  92,
	FE_DFL_ENV:              16,
	FE_DIVBYZERO:            2,
	FE_INEXACT:              16,
	FE_INVALID:              1,
	FE_OVERFLOW:             4,
	FE_UNDERFLOW:            8,
	FE_ALL_EXCEPT:           159,
	fegetround:              0,
	FE_DOWNWARD:             8388608,
	FE_TONEAREST:            0,
	FE_TOWARDZERO:           12582912,
	FE_UPWARD:               4194304,
	FP_NORMAL:               4,
	FP_SUBNORMAL:            5,
	FP_ZERO:                 3,
	FP_INFINITE:             2,
	FP_NAN:                  1,
	SIGTERM:                 15,
	SIGSEGV:                 11,
	SIGINT:                  2,
	SIGILL:                  4,
	SIGABRT:                 6,
	SIGFPE:                  8,
	LALL:                    0,
	LCOLLATE:                1,
	LCTYPE:                  2,
	LMONETARY:               3,
	LNUMERIC:                4,
	LTIME:                   5,
	MATH_ERRNO:              1,
	MATH_ERREXCEPT:          2,
	math_errhandling:        2,
	EXIT_SUCCESS:            0,
	EXIT_FAILURE:            1,
	True:                    1,
	False:                   0,
	ATOMIBOOL_LOCK_FREE:     2,
	ATOMICHAR_LOCK_FREE:     2,
	ATOMICHAR16_T_LOCK_FREE: 2,
	ATOMICHAR32_T_LOCK_FREE: 2,
	ATOMIWCHAR_T_LOCK_FREE:  2,
	ATOMISHORT_LOCK_FREE:    2,
	ATOMIINT_LOCK_FREE:      2,
	ATOMILONG_LOCK_FREE:     2,
	ATOMILLONG_LOCK_FREE:    2,
	ATOMIPOINTER_LOCK_FREE:  2,
	EOF:                     -1,
	FOPEN_MAX:               0x14,
	FILENAME_MAX:            0x400,
	L_tmpnam:                0x400,
	TMP_MAX:                 0x1269ae40,
	_IOFBF:                  0,
	_IOLBF:                  1,
	_IONBF:                  2,
	BUFSIZ:                  0x400,
	SEEK_SET:                0,
	SEEK_CUR:                1,
	SEEK_END:                2,
	CLOCKS_PER_SEC:          1000000,
}

var Types = types{
	char: value{
		kind:         reflect.Int8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	signed_char: value{
		kind:         reflect.Int8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	unsigned_char: value{
		kind:         reflect.Uint8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	short: value{
		kind:         reflect.Int16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	unsigned_short: value{
		kind:         reflect.Uint16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	int: value{
		kind:         reflect.Int32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	unsigned_int: value{
		kind:         reflect.Uint32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	long: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	unsigned_long: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	long_long: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	unsigned_long_long: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	float: value{
		kind:         reflect.Float32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	double: value{
		kind:         reflect.Float64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	long_double: value{
		kind:         reflect.Array,
		size:         0x8,
		align_memory: 1,
		align_struct: 1,
	},
	float_t: value{
		kind:         reflect.Float32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	double_t: value{
		kind:         reflect.Float64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	int8_t: value{
		kind:         reflect.Int8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	int16_t: value{
		kind:         reflect.Int16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	int32_t: value{
		kind:         reflect.Int32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	int64_t: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	uint8_t: value{
		kind:         reflect.Uint8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	uint16_t: value{
		kind:         reflect.Uint16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	uint32_t: value{
		kind:         reflect.Uint32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	uint64_t: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	char16_t: value{
		kind:         reflect.Uint16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	char32_t: value{
		kind:         reflect.Uint32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	wchar_t: value{
		kind:         reflect.Int32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	wint_t: value{
		kind:         reflect.Int32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	size_t: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	time_t: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	clock_t: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	bool: value{
		kind:         reflect.Bool,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	uintptr_t: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	ptrdiff_t: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	intptr_t: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	max_align_t: value{
		kind:         reflect.Array,
		size:         0x8,
		align_memory: 1,
		align_struct: 1,
	},
	sig_atomic_t: value{
		kind:         reflect.Int32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	intmax_t: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	uintmax_t: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	int_fast8_t: value{
		kind:         reflect.Int8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	int_fast16_t: value{
		kind:         reflect.Int16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	int_fast32_t: value{
		kind:         reflect.Int32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	int_fast64_t: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	uint_fast8_t: value{
		kind:         reflect.Uint8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	uint_fast16_t: value{
		kind:         reflect.Uint16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	uint_fast32_t: value{
		kind:         reflect.Uint32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	uint_fast64_t: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	int_least8_t: value{
		kind:         reflect.Int8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	int_least16_t: value{
		kind:         reflect.Int16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	int_least32_t: value{
		kind:         reflect.Int32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	int_least64_t: value{
		kind:         reflect.Int64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
	uint_least8_t: value{
		kind:         reflect.Uint8,
		size:         0x1,
		align_memory: 1,
		align_struct: 1,
	},
	uint_least16_t: value{
		kind:         reflect.Uint16,
		size:         0x2,
		align_memory: 2,
		align_struct: 2,
	},
	uint_least32_t: value{
		kind:         reflect.Uint32,
		size:         0x4,
		align_memory: 4,
		align_struct: 4,
	},
	uint_least64_t: value{
		kind:         reflect.Uint64,
		size:         0x8,
		align_memory: 8,
		align_struct: 8,
	},
}
