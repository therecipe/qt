#ifdef __cplusplus
extern "C" {
#endif

void* QFileSelector_NewQFileSelector(void* parent);
char* QFileSelector_AllSelectors(void* ptr);
char* QFileSelector_ExtraSelectors(void* ptr);
char* QFileSelector_Select(void* ptr, char* filePath);
void QFileSelector_SetExtraSelectors(void* ptr, char* list);
void QFileSelector_DestroyQFileSelector(void* ptr);

#ifdef __cplusplus
}
#endif