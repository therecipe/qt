#ifdef __cplusplus
extern "C" {
#endif

void* QXmlNamespaceSupport_NewQXmlNamespaceSupport();
void QXmlNamespaceSupport_PopContext(void* ptr);
char* QXmlNamespaceSupport_Prefix(void* ptr, char* uri);
char* QXmlNamespaceSupport_Prefixes(void* ptr);
char* QXmlNamespaceSupport_Prefixes2(void* ptr, char* uri);
void QXmlNamespaceSupport_PushContext(void* ptr);
void QXmlNamespaceSupport_Reset(void* ptr);
void QXmlNamespaceSupport_SetPrefix(void* ptr, char* pre, char* uri);
char* QXmlNamespaceSupport_Uri(void* ptr, char* prefix);
void QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(void* ptr);

#ifdef __cplusplus
}
#endif