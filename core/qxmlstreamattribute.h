#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamAttribute_NewQXmlStreamAttribute();
void* QXmlStreamAttribute_NewQXmlStreamAttribute3(char* namespaceUri, char* name, char* value);
void* QXmlStreamAttribute_NewQXmlStreamAttribute2(char* qualifiedName, char* value);
void* QXmlStreamAttribute_NewQXmlStreamAttribute4(void* other);
int QXmlStreamAttribute_IsDefault(void* ptr);
void* QXmlStreamAttribute_Name(void* ptr);
void* QXmlStreamAttribute_NamespaceUri(void* ptr);
void* QXmlStreamAttribute_Prefix(void* ptr);
void* QXmlStreamAttribute_QualifiedName(void* ptr);
void* QXmlStreamAttribute_Value(void* ptr);
void QXmlStreamAttribute_DestroyQXmlStreamAttribute(void* ptr);

#ifdef __cplusplus
}
#endif