#ifdef __cplusplus
extern "C" {
#endif

void* QSpacerItem_NewQSpacerItem(int w, int h, int hPolicy, int vPolicy);
void QSpacerItem_ChangeSize(void* ptr, int w, int h, int hPolicy, int vPolicy);
int QSpacerItem_ExpandingDirections(void* ptr);
int QSpacerItem_IsEmpty(void* ptr);
void QSpacerItem_SetGeometry(void* ptr, void* r);
void* QSpacerItem_SpacerItem(void* ptr);
void QSpacerItem_DestroyQSpacerItem(void* ptr);

#ifdef __cplusplus
}
#endif