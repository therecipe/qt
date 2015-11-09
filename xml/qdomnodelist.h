#ifdef __cplusplus
extern "C" {
#endif

void* QDomNodeList_NewQDomNodeList();
void* QDomNodeList_NewQDomNodeList2(void* n);
int QDomNodeList_Count(void* ptr);
int QDomNodeList_IsEmpty(void* ptr);
int QDomNodeList_Length(void* ptr);
int QDomNodeList_Size(void* ptr);
void QDomNodeList_DestroyQDomNodeList(void* ptr);

#ifdef __cplusplus
}
#endif