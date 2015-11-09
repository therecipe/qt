#ifdef __cplusplus
extern "C" {
#endif

void* QXmlSchema_NewQXmlSchema();
void* QXmlSchema_NewQXmlSchema2(void* other);
int QXmlSchema_IsValid(void* ptr);
int QXmlSchema_Load2(void* ptr, void* source, void* documentUri);
int QXmlSchema_Load3(void* ptr, void* data, void* documentUri);
int QXmlSchema_Load(void* ptr, void* source);
void* QXmlSchema_MessageHandler(void* ptr);
void* QXmlSchema_NetworkAccessManager(void* ptr);
void QXmlSchema_SetMessageHandler(void* ptr, void* handler);
void QXmlSchema_SetNetworkAccessManager(void* ptr, void* manager);
void QXmlSchema_SetUriResolver(void* ptr, void* resolver);
void* QXmlSchema_UriResolver(void* ptr);
void QXmlSchema_DestroyQXmlSchema(void* ptr);

#ifdef __cplusplus
}
#endif