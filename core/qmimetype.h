#ifdef __cplusplus
extern "C" {
#endif

void* QMimeType_NewQMimeType();
void* QMimeType_NewQMimeType2(void* other);
char* QMimeType_FilterString(void* ptr);
char* QMimeType_GenericIconName(void* ptr);
char* QMimeType_GlobPatterns(void* ptr);
char* QMimeType_IconName(void* ptr);
int QMimeType_Inherits(void* ptr, char* mimeTypeName);
int QMimeType_IsDefault(void* ptr);
int QMimeType_IsValid(void* ptr);
char* QMimeType_Name(void* ptr);
void QMimeType_DestroyQMimeType(void* ptr);
char* QMimeType_Aliases(void* ptr);
char* QMimeType_AllAncestors(void* ptr);
char* QMimeType_Comment(void* ptr);
char* QMimeType_ParentMimeTypes(void* ptr);
char* QMimeType_PreferredSuffix(void* ptr);
char* QMimeType_Suffixes(void* ptr);
void QMimeType_Swap(void* ptr, void* other);

#ifdef __cplusplus
}
#endif