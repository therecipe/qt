#ifdef __cplusplus
extern "C" {
#endif

void* QHelpContentItem_Child(void* ptr, int row);
int QHelpContentItem_ChildCount(void* ptr);
int QHelpContentItem_ChildPosition(void* ptr, void* child);
void* QHelpContentItem_Parent(void* ptr);
int QHelpContentItem_Row(void* ptr);
char* QHelpContentItem_Title(void* ptr);
void QHelpContentItem_DestroyQHelpContentItem(void* ptr);

#ifdef __cplusplus
}
#endif