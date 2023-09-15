#include <complex.h>
#include <errno.h>
#include <fenv.h>
#include <float.h>
#include <inttypes.h>
#include <limits.h>
#include <locale.h>
#include <math.h>
#include <setjmp.h>
#include <signal.h>
#include <stdatomic.h>
#include <stddef.h>
#include <stdbool.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include <time.h>
#include <wchar.h>

#ifdef __APPLE__
typedef uint16_t char16_t;
typedef uint32_t char32_t;
#else
#include <uchar.h>
#endif

// we need to represent the C struct in Go
// so we sort all fields by offset and then 
// print the type and name of each field
// for structures with implementation-defined
// order.
typedef struct {
    const char *source;
    size_t offset;
} field_t;

int compare_fields(const void *a, const void *b) {
    const field_t *fa = a;
    const field_t *fb = b;
    return fa->offset - fb->offset;
}

void structure(const char *name, field_t fields[]) {
    int length = 0;
    for (int i = 0; 1; i++) {   
        length = i;
        if (*fields[i].source == 0) {
            break;
        }
     }
    qsort(fields, length, sizeof(field_t), compare_fields);
    printf("type %s struct {\n", name);
    for (int i = 0; i < length; i++) {
        printf("\t%s\n", fields[i].source);
    }
    printf("}\n\n");
}

char* goType(size_t bytes, bool sign) {
    if (sign) {
        switch (bytes*CHAR_BIT) {
            case 8: return "int8";
            case 16: return "int16";
            case 32: return "int32";
            case 64: return "int64";
        }
        char* buf = calloc(100, sizeof(char));
        sprintf(buf, "[%d]byte", (int)bytes);
        return buf;
    }
    switch (bytes*CHAR_BIT) {
        case 8: return "uint8";
        case 16: return "uint16";
        case 32: return "uint32";
        case 64: return "uint64";
    }
    char* buf = calloc(100, sizeof(char));
    sprintf(buf, "[%d]byte", (int)bytes);
    return buf;
}

int main(int argc, char *argv[]) {
    fesetenv(FE_DFL_ENV);

    printf("// Code generated by gen/gen.c; DO NOT EDIT.\n\n");
    printf("package ffi\n\n");

    // Standard Library Constants
    printf("const (\n");
    printf("\tc_CHAR_BIT               = %d\n", CHAR_BIT);
    printf("\tc_MB_LEN_MAX             = %d\n", MB_LEN_MAX);
    printf("\tc_CHAR_MIN               = %d\n", CHAR_MIN);
    printf("\tc_CHAR_MAX               = %d\n", CHAR_MAX);
    printf("\tc_SCHAR_MIN              = %d\n", SCHAR_MIN);
    printf("\tc_SHRT_MIN               = %d\n", SHRT_MIN);
    printf("\tc_INT_MIN                = %d\n", INT_MIN);
    printf("\tc_LONG_MIN               = %ld\n", LONG_MIN);
    printf("\tc_LLONG_MIN              = %lld\n", LLONG_MIN);
    printf("\tc_SCHAR_MAX              = %d\n", SCHAR_MAX);
    printf("\tc_SHRT_MAX               = %d\n", SHRT_MAX);
    printf("\tc_INT_MAX                = %d\n", INT_MAX);
    printf("\tc_LONG_MAX               = %ld\n", LONG_MAX);
    printf("\tc_LLONG_MAX              = %lld\n", LLONG_MAX);
    printf("\tc_UCHAR_MAX              = %u\n", UCHAR_MAX);
    printf("\tc_USHRT_MAX              = %u\n", USHRT_MAX);
    printf("\tc_UINT_MAX               = %u\n", UINT_MAX);
    printf("\tc_ULONG_MAX              = %lu\n", ULONG_MAX);
    printf("\tc_ULLONG_MAX             = %llu\n", ULLONG_MAX);
    printf("\tc_PTRDIFF_MIN            = %ld\n", PTRDIFF_MIN);
    printf("\tc_PTRDIFF_MAX            = %ld\n", PTRDIFF_MAX);
    printf("\tc_SIZE_MAX               = %lu\n", SIZE_MAX);
    printf("\tc_WINT_MIN               = %d\n", WINT_MIN);
    printf("\tc_WINT_MAX               = %d\n", WINT_MAX);
    printf("\tc_WCHAR_MIN              = %d\n", WCHAR_MIN);
    printf("\tc_WCHAR_MAX              = %d\n", WCHAR_MAX);
    printf("\tc_INT8_MIN               = %d\n", INT8_MIN);
    printf("\tc_INT16_MIN              = %d\n", INT16_MIN);
    printf("\tc_INT32_MIN              = %d\n", INT32_MIN);
    printf("\tc_INT64_MIN              = %lld\n", (long long)INT64_MIN);
    printf("\tc_INT8_MAX               = %d\n", INT8_MAX);
    printf("\tc_INT16_MAX              = %d\n", INT16_MAX);
    printf("\tc_INT32_MAX              = %d\n", INT32_MAX);
    printf("\tc_INT64_MAX              = %lld\n", (long long)INT64_MAX);
    printf("\tc_UINT8_MAX              = %u\n", UINT8_MAX);
    printf("\tc_UINT16_MAX             = %u\n", UINT16_MAX);
    printf("\tc_UINT32_MAX             = %u\n", UINT32_MAX);
    printf("\tc_UINT64_MAX             = %llu\n", (unsigned long long)UINT64_MAX);
    printf("\tc_INT_FAST8_MIN          = %lld\n", (long long)INT_FAST8_MIN);
    printf("\tc_INT_FAST16_MIN         = %lld\n", (long long)INT_FAST16_MIN);
    printf("\tc_INT_FAST32_MIN         = %lld\n", (long long)INT_FAST32_MIN);
    printf("\tc_INT_FAST64_MIN         = %lld\n", (long long)INT_FAST64_MIN);
    printf("\tc_INT_FAST8_MAX          = %lld\n", (long long)INT_FAST8_MAX);
    printf("\tc_INT_FAST16_MAX         = %lld\n", (long long)INT_FAST16_MAX);
    printf("\tc_INT_FAST32_MAX         = %lld\n", (long long)INT_FAST32_MAX);
    printf("\tc_INT_FAST64_MAX         = %lld\n", (long long)INT_FAST64_MAX);
    printf("\tc_UINT_FAST8_MAX         = %llu\n", (unsigned long long)UINT_FAST8_MAX);
    printf("\tc_UINT_FAST16_MAX        = %llu\n", (unsigned long long)UINT_FAST16_MAX);
    printf("\tc_UINT_FAST32_MAX        = %llu\n", (unsigned long long)UINT_FAST32_MAX);
    printf("\tc_UINT_FAST64_MAX        = %llu\n", (unsigned long long)UINT_FAST64_MAX);
    printf("\tc_INT_LEAST8_MIN         = %d\n", INT_LEAST8_MIN);
    printf("\tc_INT_LEAST16_MIN        = %d\n", INT_LEAST16_MIN);
    printf("\tc_INT_LEAST32_MIN        = %d\n", INT_LEAST32_MIN);
    printf("\tc_INT_LEAST64_MIN        = %lld\n", (long long)INT_LEAST64_MIN);
    printf("\tc_INT_LEAST8_MAX         = %d\n", INT_LEAST8_MAX);
    printf("\tc_INT_LEAST16_MAX        = %d\n", INT_LEAST16_MAX);
    printf("\tc_INT_LEAST32_MAX        = %d\n", INT_LEAST32_MAX);
    printf("\tc_INT_LEAST64_MAX        = %lld\n", (long long)INT_LEAST64_MAX);
    printf("\tc_UINT_LEAST8_MAX        = %u\n", UINT_LEAST8_MAX);
    printf("\tc_UINT_LEAST16_MAX       = %u\n", UINT_LEAST16_MAX);
    printf("\tc_UINT_LEAST32_MAX       = %u\n", UINT_LEAST32_MAX);
    printf("\tc_UINT_LEAST64_MAX       = %llu\n", (unsigned long long)UINT_LEAST64_MAX);
    printf("\tc_INTMAX_MIN             = %lld\n", (long long)INTMAX_MIN);
    printf("\tc_INTMAX_MAX             = %lld\n", (long long)INTMAX_MAX);
    printf("\tc_UINTMAX_MAX            = %llu\n", (unsigned long long)UINTMAX_MAX);
    printf("\tc_INTPTR_MIN             = %ld\n", INTPTR_MIN);
    printf("\tc_INTPTR_MAX             = %ld\n", INTPTR_MAX);
    printf("\tc_UINTPTR_MAX            = %lu\n", UINTPTR_MAX);

    printf("\tc_SIG_ATOMIC_MIN         = %lld\n", (long long)SIG_ATOMIC_MIN);
    printf("\tc_SIG_ATOMIC_MAX         = %lld\n", (long long)SIG_ATOMIC_MAX);

    printf("\tc_FLT_RADIX              = %d\n", FLT_RADIX);
    printf("\tc_DECIMAL_DIG            = %d\n", DECIMAL_DIG);
    printf("\tc_FLT_DECIMAL_DIG        = %d\n", FLT_DECIMAL_DIG);
    printf("\tc_DBL_DECIMAL_DIG        = %d\n", DBL_DECIMAL_DIG);
    printf("\tc_LDBL_DECIMAL_DIG       = %d\n", LDBL_DECIMAL_DIG);
    printf("\tc_FLT_MIN                = %e\n", FLT_MIN);
    printf("\tc_DBL_MIN                = %e\n", DBL_MIN);
    printf("\tc_LDBL_MIN               = %Le\n", LDBL_MIN);
    printf("\tc_FLT_TRUE_MIN           = %e\n", FLT_TRUE_MIN);
    printf("\tc_DBL_TRUE_MIN           = %e\n", DBL_TRUE_MIN);
    printf("\tc_LDBL_TRUE_MIN          = %Le\n", LDBL_TRUE_MIN);
    printf("\tc_FLT_MAX                = %e\n", FLT_MAX);
    printf("\tc_DBL_MAX                = %e\n", DBL_MAX);
    printf("\tc_LDBL_MAX               = %Le\n", LDBL_MAX);
    printf("\tc_FLT_EPSILON            = %e\n", FLT_EPSILON);
    printf("\tc_DBL_EPSILON            = %e\n", DBL_EPSILON);
    printf("\tc_LDBL_EPSILON           = %Le\n", LDBL_EPSILON);
    printf("\tc_FLT_DIG                = %d\n", FLT_DIG);
    printf("\tc_DBL_DIG                = %d\n", DBL_DIG);
    printf("\tc_LDBL_DIG               = %d\n", LDBL_DIG);
    printf("\tc_FLT_MANT_DIG           = %d\n", FLT_MANT_DIG);
    printf("\tc_DBL_MANT_DIG           = %d\n", DBL_MANT_DIG);
    printf("\tc_LDBL_MANT_DIG          = %d\n", LDBL_MANT_DIG);
    printf("\tc_FLT_MIN_EXP            = %d\n", FLT_MIN_EXP);
    printf("\tc_DBL_MIN_EXP            = %d\n", DBL_MIN_EXP);
    printf("\tc_LDBL_MIN_EXP           = %d\n", LDBL_MIN_EXP);
    printf("\tc_FLT_MIN_10_EXP         = %d\n", FLT_MIN_10_EXP);
    printf("\tc_DBL_MIN_10_EXP         = %d\n", DBL_MIN_10_EXP);
    printf("\tc_LDBL_MIN_10_EXP        = %d\n", LDBL_MIN_10_EXP);
    printf("\tc_FLT_MAX_EXP            = %d\n", FLT_MAX_EXP);
    printf("\tc_DBL_MAX_EXP            = %d\n", DBL_MAX_EXP);
    printf("\tc_LDBL_MAX_EXP           = %d\n", LDBL_MAX_EXP);
    printf("\tc_FLT_MAX_10_EXP         = %d\n", FLT_MAX_10_EXP);
    printf("\tc_DBL_MAX_10_EXP         = %d\n", DBL_MAX_10_EXP);
    printf("\tc_LDBL_MAX_10_EXP        = %d\n", LDBL_MAX_10_EXP);
    printf("\tc_FLT_ROUNDS             = %d\n", FLT_ROUNDS);
    //printf("\tc_FLT_EVAL_METHOD        = %d\n", FLT_EVAL_METHOD);
    printf("\tc_FLT_HAS_SUBNORM        = %d\n", FLT_HAS_SUBNORM);
    printf("\tc_DBL_HAS_SUBNORM        = %d\n", DBL_HAS_SUBNORM);
    printf("\tc_LDBL_HAS_SUBNORM       = %d\n", LDBL_HAS_SUBNORM);

    printf("\tc_EDOM                   = %d\n", EDOM);
    printf("\tc_ERANGE                 = %d\n", ERANGE);
    printf("\tc_EILSEQ                 = %d\n", EILSEQ);

    printf("\tc_FE_DFL_ENV             = %d\n", fetestexcept(FE_ALL_EXCEPT));
    printf("\tc_FE_DIVBYZERO           = %d\n", FE_DIVBYZERO);
    printf("\tc_FE_INEXACT             = %d\n", FE_INEXACT);
    printf("\tc_FE_INVALID             = %d\n", FE_INVALID);
    printf("\tc_FE_OVERFLOW            = %d\n", FE_OVERFLOW);
    printf("\tc_FE_UNDERFLOW           = %d\n", FE_UNDERFLOW);
    printf("\tc_FE_ALL_EXCEPT          = %d\n", FE_ALL_EXCEPT);

    printf("\tc_fegetround             = %d\n", fegetround());   
    printf("\tc_FE_DOWNWARD            = %d\n", FE_DOWNWARD);
    printf("\tc_FE_TONEAREST           = %d\n", FE_TONEAREST);
    printf("\tc_FE_TOWARDZERO          = %d\n", FE_TOWARDZERO);
    printf("\tc_FE_UPWARD              = %d\n", FE_UPWARD);

    printf("\tc_FP_NORMAL              = %d\n", FP_NORMAL);
    printf("\tc_FP_SUBNORMAL           = %d\n", FP_SUBNORMAL);
    printf("\tc_FP_ZERO                = %d\n", FP_ZERO);
    printf("\tc_FP_INFINITE            = %d\n", FP_INFINITE);
    printf("\tc_FP_NAN                 = %d\n", FP_NAN);

    printf("\tc_SIGTERM                = %d\n", SIGTERM);
    printf("\tc_SIGSEGV                = %d\n", SIGSEGV);
    printf("\tc_SIGINT                 = %d\n", SIGINT);
    printf("\tc_SIGILL                 = %d\n", SIGILL);
    printf("\tc_SIGABRT                = %d\n", SIGABRT);
    printf("\tc_SIGFPE                 = %d\n", SIGFPE);

    printf("\tc_LC_ALL                 = %d\n", LC_ALL);
    printf("\tc_LC_COLLATE             = %d\n", LC_COLLATE);
    printf("\tc_LC_CTYPE               = %d\n", LC_CTYPE);
    printf("\tc_LC_MONETARY            = %d\n", LC_MONETARY);
    printf("\tc_LC_NUMERIC             = %d\n", LC_NUMERIC);
    printf("\tc_LC_TIME                = %d\n", LC_TIME);

    printf("\tc_MATH_ERRNO             = %d\n", MATH_ERRNO);
    printf("\tc_MATH_ERREXCEPT         = %d\n", MATH_ERREXCEPT);
    printf("\tc_math_errhandling       = %d\n", math_errhandling);

    printf("\tc_EXIT_SUCCESS           = %d\n", EXIT_SUCCESS);
    printf("\tc_EXIT_FAILURE           = %d\n", EXIT_FAILURE);

    printf("\tc_true                   = %d\n", true);
    printf("\tc_false                  = %d\n", false);

    printf("\tc_ATOMIC_BOOL_LOCK_FREE  = %d\n", ATOMIC_BOOL_LOCK_FREE);
    printf("\tc_ATOMIC_CHAR_LOCK_FREE  = %d\n", ATOMIC_CHAR_LOCK_FREE);
    printf("\tc_ATOMIC_CHAR16_T_LOCK_FREE  = %d\n", ATOMIC_CHAR16_T_LOCK_FREE);
    printf("\tc_ATOMIC_CHAR32_T_LOCK_FREE  = %d\n", ATOMIC_CHAR32_T_LOCK_FREE);
    printf("\tc_ATOMIC_WCHAR_T_LOCK_FREE  = %d\n", ATOMIC_WCHAR_T_LOCK_FREE);
    printf("\tc_ATOMIC_SHORT_LOCK_FREE  = %d\n", ATOMIC_SHORT_LOCK_FREE);
    printf("\tc_ATOMIC_INT_LOCK_FREE  = %d\n", ATOMIC_INT_LOCK_FREE);
    printf("\tc_ATOMIC_LONG_LOCK_FREE  = %d\n", ATOMIC_LONG_LOCK_FREE);
    printf("\tc_ATOMIC_LLONG_LOCK_FREE  = %d\n", ATOMIC_LLONG_LOCK_FREE);
    printf("\tc_ATOMIC_POINTER_LOCK_FREE  = %d\n", ATOMIC_POINTER_LOCK_FREE);

    printf("\tc_EOF                   = %d\n", EOF);
    printf("\tc_FOPEN_MAX             = %d\n", FOPEN_MAX);
    printf("\tc_FILENAME_MAX          = %d\n", FILENAME_MAX);
    printf("\tc_L_tmpnam              = %d\n", L_tmpnam);
    printf("\tc_TMP_MAX               = %d\n", TMP_MAX);
    printf("\tc__IOFBF                 = %d\n", _IOFBF);
    printf("\tc__IOLBF                 = %d\n", _IOLBF);
    printf("\tc__IONBF                 = %d\n", _IONBF);
    printf("\tc_BUFSIZ                = %d\n", BUFSIZ);
    printf("\tc_SEEK_SET              = %d\n", SEEK_SET);
    printf("\tc_SEEK_CUR              = %d\n", SEEK_CUR);
    printf("\tc_SEEK_END              = %d\n", SEEK_END);

    printf("\tc_CLOCKS_PER_SEC        = %lld\n", (long long)CLOCKS_PER_SEC);

    printf(")\n\n");

    printf("type (\n");
    printf("\tc_char                %s\n", goType(sizeof(char), CHAR_MIN < 0));
    printf("\tc_signed_char         %s\n", goType(sizeof(signed char), true));
    printf("\tc_unsigned_char       %s\n", goType(sizeof(unsigned char), false));
    printf("\tc_short               %s\n", goType(sizeof(short), SHRT_MIN < 0));
    printf("\tc_unsigned_short      %s\n", goType(sizeof(unsigned short), false));
    printf("\tc_int                 %s\n", goType(sizeof(int), INT_MIN < 0));
    printf("\tc_unsigned_int        %s\n", goType(sizeof(unsigned int), false));
    printf("\tc_long                %s\n", goType(sizeof(long), LONG_MIN < 0));
    printf("\tc_unsigned_long       %s\n", goType(sizeof(unsigned long), false));
    printf("\tc_longlong            %s\n", goType(sizeof(long long), LLONG_MIN < 0));
    printf("\tc_unsigned_longlong   %s\n", goType(sizeof(unsigned long long), false));
    printf("\tc_float               float%d\n", (unsigned)sizeof(float)*CHAR_BIT);
    printf("\tc_double              float%d\n", (unsigned)sizeof(double)*CHAR_BIT);
    if (sizeof(double)*CHAR_BIT > 64) {
        printf("\tc_long_double         [%d]byte\n", (unsigned)sizeof(double));
    } else {
        printf("\tc_long_double         float%d\n", (unsigned)sizeof(double)*CHAR_BIT);
    }
    if (sizeof(float_t)*CHAR_BIT > 64) {
        printf("\tc_float_t             [%d]byte\n", (unsigned)sizeof(float_t));
    } else {
        printf("\tc_float_t             float%d\n", (unsigned)sizeof(float_t)*CHAR_BIT);
    }
    if (sizeof(double_t)*CHAR_BIT > 64) {
        printf("\tc_double_t            [%d]byte\n", (unsigned)sizeof(double_t));
    } else {
        printf("\tc_double_t            float%d\n", (unsigned)sizeof(double_t)*CHAR_BIT);
    }

    printf("\tc_int8_t              %s\n", goType(sizeof(int8_t), true));
    printf("\tc_int16_t             %s\n", goType(sizeof(int16_t), true));
    printf("\tc_int32_t             %s\n", goType(sizeof(int32_t), true));
    printf("\tc_int64_t             %s\n", goType(sizeof(int64_t), true));
    printf("\tc_uint8_t             %s\n", goType(sizeof(uint8_t), false));
    printf("\tc_uint16_t            %s\n", goType(sizeof(uint16_t), false));
    printf("\tc_uint32_t            %s\n", goType(sizeof(uint32_t), false));
    printf("\tc_uint64_t            %s\n", goType(sizeof(uint64_t), false));

    printf("\tc_char16_t            %s\n", goType(sizeof(char16_t), WCHAR_MIN < 0));
    printf("\tc_char32_t            %s\n", goType(sizeof(char32_t), WCHAR_MIN < 0));
    printf("\tc_wchar_t             %s\n", goType(sizeof(wchar_t), WCHAR_MIN < 0));
    printf("\tc_wint_t              %s\n", goType(sizeof(wint_t), WCHAR_MIN < 0));   


    printf("\tc_size_t               %s\n", goType(sizeof(size_t), false));
    printf("\tc_time_t               %s\n", goType(sizeof(time_t), true));
    printf("\tc_clock_t              %s\n", goType(sizeof(clock_t), true));

    printf("\tc_bool                 %s\n", goType(sizeof(_Bool), false));

    printf("\tc_uintptr_t           %s\n", goType(sizeof(uintptr_t), false));
    printf("\tc_ptrdiff_t           %s\n", goType(sizeof(ptrdiff_t), true));
    printf("\tc_intptr_t               %s\n", goType(sizeof(intptr_t), true));

    printf("\tc_max_align_t         %s\n", goType(sizeof(max_align_t), false));
    printf("\tc_sig_atomic_t        %s\n", goType(sizeof(sig_atomic_t), true));

    printf("\tc_intmax_t                %s\n", goType(sizeof(intmax_t), true));
    printf("\tc_uintmax_t               %s\n", goType(sizeof(uintmax_t), false));

    printf("\tc_int_fast8_t             %s\n", goType(sizeof(int_fast8_t), true));
    printf("\tc_int_fast16_t            %s\n", goType(sizeof(int_fast16_t), true));
    printf("\tc_int_fast32_t            %s\n", goType(sizeof(int_fast32_t), true));
    printf("\tc_int_fast64_t            %s\n", goType(sizeof(int_fast64_t), true));
    printf("\tc_uint_fast8_t            %s\n", goType(sizeof(uint_fast8_t), false));
    printf("\tc_uint_fast16_t           %s\n", goType(sizeof(uint_fast16_t), false));
    printf("\tc_uint_fast32_t           %s\n", goType(sizeof(uint_fast32_t), false));
    printf("\tc_uint_fast64_t           %s\n", goType(sizeof(uint_fast64_t), false));

    printf("\tc_int_least8_t             %s\n", goType(sizeof(int_least8_t), true));
    printf("\tc_int_least16_t            %s\n", goType(sizeof(int_least16_t), true));
    printf("\tc_int_least32_t            %s\n", goType(sizeof(int_least32_t), true));
    printf("\tc_int_least64_t            %s\n", goType(sizeof(int_least64_t), true));
    printf("\tc_uint_least8_t            %s\n", goType(sizeof(uint_least8_t), false));
    printf("\tc_uint_least16_t           %s\n", goType(sizeof(uint_least16_t), false));
    printf("\tc_uint_least32_t           %s\n", goType(sizeof(uint_least32_t), false));
    printf("\tc_uint_least64_t           %s\n", goType(sizeof(uint_least64_t), false));
    printf(")\n\n");

    printf("type c_fenv_t [%u]byte\n\n", (unsigned)sizeof(fenv_t));
    printf("type c_jmp_buf [%u]byte\n\n", (unsigned)sizeof(jmp_buf));
    printf("type c_FILE [%u]byte\n\n", (unsigned)sizeof(FILE));
    printf("type c_fpos_t [%u]byte\n\n", (unsigned)sizeof(fpos_t));
    printf("type c_va_list [%u]byte\n\n", (unsigned)sizeof(va_list));
    printf("type c_mbstate_t [%u]byte\n\n", (unsigned)sizeof(mbstate_t));

    /*
    structure("c_lconv", (field_t[]){
        {"DecimalPoint                   String", offsetof(struct lconv, decimal_point)},
        {"ThousandsSeperator             String", offsetof(struct lconv, thousands_sep)},
        {"Grouping                       String", offsetof(struct lconv, grouping)},
        {"MonetaryDecimalPoint           String", offsetof(struct lconv, mon_decimal_point)},
        {"MonetaryThousandsSeperator     String", offsetof(struct lconv, mon_thousands_sep)},
        {"MonetaryGrouping               String", offsetof(struct lconv, mon_grouping)},
        {"PositiveSign                   String", offsetof(struct lconv, positive_sign)},
        {"NegativeSign                   String", offsetof(struct lconv, negative_sign)},
        {"CurrencySymbol                 String", offsetof(struct lconv, currency_symbol)},
        {"FractionDigits                 Char", offsetof(struct lconv, frac_digits)},
        {"LocalCurrencyPrefixesPositive  Char", offsetof(struct lconv, p_cs_precedes)},
        {"LocalCurrencyPrefixesNegative  Char", offsetof(struct lconv, n_cs_precedes)},
        {"LocalCurrencyPositiveSpacing   Char", offsetof(struct lconv, p_sep_by_space)},
        {"LocalCurrencyNegativeSpacing   Char", offsetof(struct lconv, n_sep_by_space)},
        {"LocalCurrencyPositiveSignPos   Char", offsetof(struct lconv, p_sign_posn)},
        {"LocalCurrencyNegativeSignPos   Char", offsetof(struct lconv, n_sign_posn)},

        {"CurrencyName                   String", offsetof(struct lconv, int_curr_symbol)},
        {"MonetaryFractionalDigits       Char", offsetof(struct lconv, int_frac_digits)},
        {"CurrencyPrefixesPositive       Char", offsetof(struct lconv, int_p_cs_precedes)},
        {"CurrencyPrefixesNegative       Char", offsetof(struct lconv, int_n_cs_precedes)},
        {"CurrencyPositiveSpacing        Char", offsetof(struct lconv, int_p_sep_by_space)},
        {"CurrencyNegativeSpacing        Char", offsetof(struct lconv, int_n_sep_by_space)},
        {"CurrencyPositiveSignPos        Char", offsetof(struct lconv, int_p_sign_posn)},
        {"CurrencyNegativeSignPos        Char", offsetof(struct lconv, int_n_sign_posn)},  
        {"", 0},
    });

    structure("c_div_t", (field_t[]){
        {"Quotient Int", offsetof(div_t, quot)},
        {"Remainder Int", offsetof(div_t, rem)},
        {"", 0},
    });
    structure("c_ldiv_t", (field_t[]){
        {"Quotient Long", offsetof(ldiv_t, quot)},
        {"Remainder Long", offsetof(ldiv_t, rem)},
        {"", 0},
    });
    structure("c_lldiv_t", (field_t[]){
        {"Quotient LongLong", offsetof(ldiv_t, quot)},
        {"Remainder LongLong", offsetof(ldiv_t, rem)},
        {"", 0},
    });
    structure("c_imaxdiv_t", (field_t[]){
        {"Quotient Longest", offsetof(imaxdiv_t, quot)},
        {"Remainder Longest", offsetof(imaxdiv_t, rem)},
        {"", 0},
    });
    structure("c_timespec", (field_t[]){
        {"Seconds     Time", offsetof(struct timespec, tv_sec)},
        {"Nanoseconds Time", offsetof(struct timespec, tv_nsec)},
        {"", 0},
    });
    structure("c_tm", (field_t[]){
        {"Seconds         Int", offsetof(struct tm, tm_sec)},
        {"Minutes         Int", offsetof(struct tm, tm_min)},
        {"Hours           Int", offsetof(struct tm, tm_hour)},
        {"Days            Int", offsetof(struct tm, tm_mday)},
        {"Months          Int", offsetof(struct tm, tm_mon)},
        {"Years           Int", offsetof(struct tm, tm_year)},
        {"Weekdays        Int", offsetof(struct tm, tm_wday)},
        {"DaysThisYear    Int", offsetof(struct tm, tm_yday)},
        {"DaylightSavings Int", offsetof(struct tm, tm_isdst)},
        {"", 0},
    });
    */
}
