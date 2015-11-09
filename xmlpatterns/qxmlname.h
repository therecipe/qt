#ifdef __cplusplus
extern "C" {
#endif

void* QXmlName_NewQXmlName();
void* QXmlName_NewQXmlName2(void* namePool, char* localName, char* namespaceURI, char* prefix);
int QXmlName_QXmlName_IsNCName(char* candidate);
int QXmlName_IsNull(void* ptr);
char* QXmlName_LocalName(void* ptr, void* namePool);
char* QXmlName_NamespaceUri(void* ptr, void* namePool);
char* QXmlName_Prefix(void* ptr, void* namePool);
char* QXmlName_ToClarkName(void* ptr, void* namePool);

#ifdef __cplusplus
}
#endif