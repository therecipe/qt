#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamAttributes_NewQXmlStreamAttributes();
void QXmlStreamAttributes_Append(void* ptr, char* namespaceUri, char* name, char* value);
void QXmlStreamAttributes_Append2(void* ptr, char* qualifiedName, char* value);
int QXmlStreamAttributes_HasAttribute2(void* ptr, void* qualifiedName);
int QXmlStreamAttributes_HasAttribute3(void* ptr, char* namespaceUri, char* name);
int QXmlStreamAttributes_HasAttribute(void* ptr, char* qualifiedName);
void* QXmlStreamAttributes_Value3(void* ptr, void* namespaceUri, void* name);
void* QXmlStreamAttributes_Value5(void* ptr, void* qualifiedName);
void* QXmlStreamAttributes_Value2(void* ptr, char* namespaceUri, void* name);
void* QXmlStreamAttributes_Value(void* ptr, char* namespaceUri, char* name);
void* QXmlStreamAttributes_Value4(void* ptr, char* qualifiedName);

#ifdef __cplusplus
}
#endif