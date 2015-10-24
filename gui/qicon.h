#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QIcon_NewQIcon();
QtObjectPtr QIcon_NewQIcon4(QtObjectPtr other);
QtObjectPtr QIcon_NewQIcon6(QtObjectPtr engine);
QtObjectPtr QIcon_NewQIcon3(QtObjectPtr other);
QtObjectPtr QIcon_NewQIcon2(QtObjectPtr pixmap);
QtObjectPtr QIcon_NewQIcon5(char* fileName);
void QIcon_AddFile(QtObjectPtr ptr, char* fileName, QtObjectPtr size, int mode, int state);
void QIcon_AddPixmap(QtObjectPtr ptr, QtObjectPtr pixmap, int mode, int state);
int QIcon_QIcon_HasThemeIcon(char* name);
int QIcon_IsNull(QtObjectPtr ptr);
char* QIcon_Name(QtObjectPtr ptr);
void QIcon_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int alignment, int mode, int state);
void QIcon_Paint2(QtObjectPtr ptr, QtObjectPtr painter, int x, int y, int w, int h, int alignment, int mode, int state);
void QIcon_QIcon_SetThemeName(char* name);
void QIcon_QIcon_SetThemeSearchPaths(char* paths);
void QIcon_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QIcon_QIcon_ThemeName();
char* QIcon_QIcon_ThemeSearchPaths();
void QIcon_DestroyQIcon(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif