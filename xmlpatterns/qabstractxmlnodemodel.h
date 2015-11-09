#ifdef __cplusplus
extern "C" {
#endif

int QAbstractXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2);
int QAbstractXmlNodeModel_Kind(void* ptr, void* ni);
char* QAbstractXmlNodeModel_StringValue(void* ptr, void* n);
void* QAbstractXmlNodeModel_TypedValue(void* ptr, void* node);
void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(void* ptr);

#ifdef __cplusplus
}
#endif