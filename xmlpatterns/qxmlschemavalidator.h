#ifdef __cplusplus
extern "C" {
#endif

void* QXmlSchemaValidator_NewQXmlSchemaValidator();
void* QXmlSchemaValidator_NewQXmlSchemaValidator2(void* schema);
void* QXmlSchemaValidator_MessageHandler(void* ptr);
void* QXmlSchemaValidator_NetworkAccessManager(void* ptr);
void QXmlSchemaValidator_SetMessageHandler(void* ptr, void* handler);
void QXmlSchemaValidator_SetNetworkAccessManager(void* ptr, void* manager);
void QXmlSchemaValidator_SetSchema(void* ptr, void* schema);
void QXmlSchemaValidator_SetUriResolver(void* ptr, void* resolver);
void* QXmlSchemaValidator_UriResolver(void* ptr);
int QXmlSchemaValidator_Validate2(void* ptr, void* source, void* documentUri);
int QXmlSchemaValidator_Validate3(void* ptr, void* data, void* documentUri);
int QXmlSchemaValidator_Validate(void* ptr, void* source);
void QXmlSchemaValidator_DestroyQXmlSchemaValidator(void* ptr);

#ifdef __cplusplus
}
#endif