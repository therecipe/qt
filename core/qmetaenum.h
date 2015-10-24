#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMetaEnum_IsFlag(QtObjectPtr ptr);
int QMetaEnum_IsValid(QtObjectPtr ptr);
int QMetaEnum_KeyCount(QtObjectPtr ptr);
int QMetaEnum_KeyToValue(QtObjectPtr ptr, char* key, int ok);
int QMetaEnum_KeysToValue(QtObjectPtr ptr, char* keys, int ok);
int QMetaEnum_Value(QtObjectPtr ptr, int index);

#ifdef __cplusplus
}
#endif