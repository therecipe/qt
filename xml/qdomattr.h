#ifdef __cplusplus
extern "C" {
#endif

void* QDomAttr_NewQDomAttr();
void* QDomAttr_NewQDomAttr2(void* x);
char* QDomAttr_Name(void* ptr);
int QDomAttr_NodeType(void* ptr);
void QDomAttr_SetValue(void* ptr, char* v);
int QDomAttr_Specified(void* ptr);
char* QDomAttr_Value(void* ptr);

#ifdef __cplusplus
}
#endif