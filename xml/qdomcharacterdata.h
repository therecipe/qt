#ifdef __cplusplus
extern "C" {
#endif

void* QDomCharacterData_NewQDomCharacterData();
void* QDomCharacterData_NewQDomCharacterData2(void* x);
void QDomCharacterData_AppendData(void* ptr, char* arg);
char* QDomCharacterData_Data(void* ptr);
int QDomCharacterData_Length(void* ptr);
int QDomCharacterData_NodeType(void* ptr);
void QDomCharacterData_SetData(void* ptr, char* v);

#ifdef __cplusplus
}
#endif