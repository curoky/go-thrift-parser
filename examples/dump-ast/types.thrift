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

struct InnerStructType {
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
    10:InnerStructType var_struct_type;

    // conatiner
    100:list<string> var_string_list;
    101:list<binary> var_binary_list;
    102:set<string> var_string_set;
    103:map<string, binary> var_string_binary_map;
    104:list<InnerStructType> var_struct_list;
    105:set<InnerStructType> var_struct_set;
    106:map<string, InnerStructType> var_string_struct_map;

    // enum
    201:EnumType var_enum;
    202:set<EnumType> var_enum_set;

    // union
    301:UnionType var_union;

    // Field Requiredness
    401:required i32 var_required_i32;
    402:optional i32 var_optional_i32;
}

exception ExceptionType {
  1: string msg;
}

struct MethodReq {
}

struct MethodResponse {
}

service ServiceV1 {
  MethodResponse method1(1: MethodReq req) (tag.v1='xxx', tag.v2='xxx');
}
