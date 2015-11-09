#ifdef __cplusplus
extern "C" {
#endif

void* QProcessEnvironment_NewQProcessEnvironment();
void* QProcessEnvironment_NewQProcessEnvironment2(void* other);
void QProcessEnvironment_Clear(void* ptr);
int QProcessEnvironment_Contains(void* ptr, char* name);
int QProcessEnvironment_IsEmpty(void* ptr);
char* QProcessEnvironment_Keys(void* ptr);
void QProcessEnvironment_Swap(void* ptr, void* other);
char* QProcessEnvironment_ToStringList(void* ptr);
char* QProcessEnvironment_Value(void* ptr, char* name, char* defaultValue);
void QProcessEnvironment_DestroyQProcessEnvironment(void* ptr);

#ifdef __cplusplus
}
#endif