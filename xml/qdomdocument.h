#ifdef __cplusplus
extern "C" {
#endif

void* QDomDocument_NewQDomDocument();
void* QDomDocument_NewQDomDocument4(void* x);
void* QDomDocument_NewQDomDocument3(void* doctype);
void* QDomDocument_NewQDomDocument2(char* name);
int QDomDocument_NodeType(void* ptr);
int QDomDocument_SetContent7(void* ptr, void* dev, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent3(void* ptr, void* dev, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent8(void* ptr, void* source, void* reader, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent4(void* ptr, void* source, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent6(void* ptr, void* buffer, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent(void* ptr, void* data, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent5(void* ptr, char* text, char* errorMsg, int errorLine, int errorColumn);
int QDomDocument_SetContent2(void* ptr, char* text, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn);
void* QDomDocument_ToByteArray(void* ptr, int indent);
char* QDomDocument_ToString(void* ptr, int indent);
void QDomDocument_DestroyQDomDocument(void* ptr);

#ifdef __cplusplus
}
#endif