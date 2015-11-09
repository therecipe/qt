#ifdef __cplusplus
extern "C" {
#endif

void* QDomNotation_NewQDomNotation();
void* QDomNotation_NewQDomNotation2(void* x);
int QDomNotation_NodeType(void* ptr);
char* QDomNotation_PublicId(void* ptr);
char* QDomNotation_SystemId(void* ptr);

#ifdef __cplusplus
}
#endif