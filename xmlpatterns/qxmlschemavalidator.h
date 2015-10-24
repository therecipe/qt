#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlSchemaValidator_NewQXmlSchemaValidator();
QtObjectPtr QXmlSchemaValidator_NewQXmlSchemaValidator2(QtObjectPtr schema);
QtObjectPtr QXmlSchemaValidator_MessageHandler(QtObjectPtr ptr);
QtObjectPtr QXmlSchemaValidator_NetworkAccessManager(QtObjectPtr ptr);
void QXmlSchemaValidator_SetMessageHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSchemaValidator_SetNetworkAccessManager(QtObjectPtr ptr, QtObjectPtr manager);
void QXmlSchemaValidator_SetSchema(QtObjectPtr ptr, QtObjectPtr schema);
void QXmlSchemaValidator_SetUriResolver(QtObjectPtr ptr, QtObjectPtr resolver);
QtObjectPtr QXmlSchemaValidator_UriResolver(QtObjectPtr ptr);
int QXmlSchemaValidator_Validate2(QtObjectPtr ptr, QtObjectPtr source, char* documentUri);
int QXmlSchemaValidator_Validate3(QtObjectPtr ptr, QtObjectPtr data, char* documentUri);
int QXmlSchemaValidator_Validate(QtObjectPtr ptr, char* source);
void QXmlSchemaValidator_DestroyQXmlSchemaValidator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif