#ifdef __cplusplus
extern "C" {
#endif

void* QDomNamedNodeMap_NewQDomNamedNodeMap();
void* QDomNamedNodeMap_NewQDomNamedNodeMap2(void* n);
int QDomNamedNodeMap_Contains(void* ptr, char* name);
int QDomNamedNodeMap_Count(void* ptr);
int QDomNamedNodeMap_IsEmpty(void* ptr);
int QDomNamedNodeMap_Length(void* ptr);
int QDomNamedNodeMap_Size(void* ptr);
void QDomNamedNodeMap_DestroyQDomNamedNodeMap(void* ptr);

#ifdef __cplusplus
}
#endif