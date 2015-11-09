#ifdef __cplusplus
extern "C" {
#endif

void* QDomEntity_NewQDomEntity();
void* QDomEntity_NewQDomEntity2(void* x);
int QDomEntity_NodeType(void* ptr);
char* QDomEntity_NotationName(void* ptr);
char* QDomEntity_PublicId(void* ptr);
char* QDomEntity_SystemId(void* ptr);

#ifdef __cplusplus
}
#endif