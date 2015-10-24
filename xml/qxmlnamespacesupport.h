#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlNamespaceSupport_NewQXmlNamespaceSupport();
void QXmlNamespaceSupport_PopContext(QtObjectPtr ptr);
char* QXmlNamespaceSupport_Prefix(QtObjectPtr ptr, char* uri);
char* QXmlNamespaceSupport_Prefixes(QtObjectPtr ptr);
char* QXmlNamespaceSupport_Prefixes2(QtObjectPtr ptr, char* uri);
void QXmlNamespaceSupport_PushContext(QtObjectPtr ptr);
void QXmlNamespaceSupport_Reset(QtObjectPtr ptr);
void QXmlNamespaceSupport_SetPrefix(QtObjectPtr ptr, char* pre, char* uri);
char* QXmlNamespaceSupport_Uri(QtObjectPtr ptr, char* prefix);
void QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif