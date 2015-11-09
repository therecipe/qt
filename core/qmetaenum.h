#ifdef __cplusplus
extern "C" {
#endif

int QMetaEnum_IsFlag(void* ptr);
int QMetaEnum_IsValid(void* ptr);
int QMetaEnum_KeyCount(void* ptr);
int QMetaEnum_KeyToValue(void* ptr, char* key, int ok);
int QMetaEnum_KeysToValue(void* ptr, char* keys, int ok);
int QMetaEnum_Value(void* ptr, int index);
void* QMetaEnum_ValueToKeys(void* ptr, int value);

#ifdef __cplusplus
}
#endif