#ifdef __cplusplus
extern "C" {
#endif

void* QXmlItem_NewQXmlItem();
void* QXmlItem_NewQXmlItem4(void* atomicValue);
void* QXmlItem_NewQXmlItem2(void* other);
void* QXmlItem_NewQXmlItem3(void* node);
int QXmlItem_IsNode(void* ptr);
int QXmlItem_IsNull(void* ptr);
void QXmlItem_DestroyQXmlItem(void* ptr);

#ifdef __cplusplus
}
#endif