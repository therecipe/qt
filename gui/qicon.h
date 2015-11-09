#ifdef __cplusplus
extern "C" {
#endif

void* QIcon_NewQIcon();
void* QIcon_NewQIcon4(void* other);
void* QIcon_NewQIcon6(void* engine);
void* QIcon_NewQIcon3(void* other);
void* QIcon_NewQIcon2(void* pixmap);
void* QIcon_NewQIcon5(char* fileName);
void QIcon_AddFile(void* ptr, char* fileName, void* size, int mode, int state);
void QIcon_AddPixmap(void* ptr, void* pixmap, int mode, int state);
int QIcon_QIcon_HasThemeIcon(char* name);
int QIcon_IsNull(void* ptr);
char* QIcon_Name(void* ptr);
void QIcon_Paint(void* ptr, void* painter, void* rect, int alignment, int mode, int state);
void QIcon_Paint2(void* ptr, void* painter, int x, int y, int w, int h, int alignment, int mode, int state);
void QIcon_QIcon_SetThemeName(char* name);
void QIcon_QIcon_SetThemeSearchPaths(char* paths);
void QIcon_Swap(void* ptr, void* other);
char* QIcon_QIcon_ThemeName();
char* QIcon_QIcon_ThemeSearchPaths();
void QIcon_DestroyQIcon(void* ptr);

#ifdef __cplusplus
}
#endif