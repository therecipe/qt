#ifdef __cplusplus
extern "C" {
#endif

int QTextList_ItemNumber(void* ptr, void* block);
char* QTextList_ItemText(void* ptr, void* block);
void QTextList_Add(void* ptr, void* block);
int QTextList_Count(void* ptr);
void QTextList_RemoveItem(void* ptr, int i);
void QTextList_SetFormat(void* ptr, void* format);

#ifdef __cplusplus
}
#endif