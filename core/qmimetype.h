#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMimeType_NewQMimeType();
QtObjectPtr QMimeType_NewQMimeType2(QtObjectPtr other);
char* QMimeType_FilterString(QtObjectPtr ptr);
char* QMimeType_GenericIconName(QtObjectPtr ptr);
char* QMimeType_GlobPatterns(QtObjectPtr ptr);
char* QMimeType_IconName(QtObjectPtr ptr);
int QMimeType_Inherits(QtObjectPtr ptr, char* mimeTypeName);
int QMimeType_IsDefault(QtObjectPtr ptr);
int QMimeType_IsValid(QtObjectPtr ptr);
char* QMimeType_Name(QtObjectPtr ptr);
void QMimeType_DestroyQMimeType(QtObjectPtr ptr);
char* QMimeType_Aliases(QtObjectPtr ptr);
char* QMimeType_AllAncestors(QtObjectPtr ptr);
char* QMimeType_Comment(QtObjectPtr ptr);
char* QMimeType_ParentMimeTypes(QtObjectPtr ptr);
char* QMimeType_PreferredSuffix(QtObjectPtr ptr);
char* QMimeType_Suffixes(QtObjectPtr ptr);
void QMimeType_Swap(QtObjectPtr ptr, QtObjectPtr other);

#ifdef __cplusplus
}
#endif