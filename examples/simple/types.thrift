namespace cpp foo
namespace java foo
cpp_include "foo.h"

const string GLOBAL_CONST_VAR_STRING = "123";
typedef string StrType

enum EnumType {
    ZERO= 0
    ONE = 1;
    TWO = 2;
    THREE = 3;
}

union UnionType {
    1:i16 var_i16;
    2:i32 var_i32;
}

struct StructType {
    // basic type
    1:bool var_bool;
    // The "byte" type is a compatibility alias for "i8".
    // Use "i8" to emphasize the signedness of this type.
    2:byte var_byte;
    3:i16 var_i16;
    4:i32 var_i32;
    5:i64 var_i64;
    6:double var_double;
    7:string var_string;
    8:binary var_binary; // equal to string
    9:StrType var_string_type;

    // conatiner
    10:list<string> var_string_list;
    11:list<binary> var_binary_list;
    12:set<string> var_string_set;
    13:map<string, binary> var_string_binary_map;

    // enum
    14:EnumType var_enum;
    15:set<EnumType> var_enum_set;

    // union
    16:UnionType var_union;

    // Field Requiredness
    17:required i32 var_required_i32;
    18:optional i32 var_optional_i32;
}

struct OutterStructType {
    1:StructType req;
}

exception ExceptionType {
  1: string msg;
}
