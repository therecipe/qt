#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomDocument_NewQDomDocument();
QtObjectPtr QDomDocument_NewQDomDocument4(QtObjectPtr x);
QtObjectPtr QDomDocument_NewQDomDocument3(QtObjectPtr doctype);
QtObjectPtr QDomDocument_NewQDomDocument2(char* name);
int QDomDocument_NodeType(QtObjectPtr ptr);
int QDomDocument_SetContent7(QtObjectPtr ptr, QtObjectPtr dev, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent3(QtObjectPtr ptr, QtObjectPtr dev, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent8(QtObjectPtr ptr, QtObjectPtr source, QtObjectPtr reader, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent4(QtObjectPtr ptr, QtObjectPtr source, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent6(QtObjectPtr ptr, QtObjectPtr buffer, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent(QtObjectPtr ptr, QtObjectPtr data, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent5(QtObjectPtr ptr, char* text, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent2(QtObjectPtr ptr, char* text, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
char* QDomDocument_ToString(QtObjectPtr ptr, int indent);
void QDomDocument_DestroyQDomDocument(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif