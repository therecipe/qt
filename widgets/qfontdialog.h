#ifdef __cplusplus
extern "C" {
#endif

int QFontDialog_Options(void* ptr);
void QFontDialog_SetOptions(void* ptr, int options);
void* QFontDialog_NewQFontDialog(void* parent);
void* QFontDialog_NewQFontDialog2(void* initial, void* parent);
void QFontDialog_Open(void* ptr, void* receiver, char* member);
void QFontDialog_SetCurrentFont(void* ptr, void* font);
void QFontDialog_SetOption(void* ptr, int option, int on);
void QFontDialog_SetVisible(void* ptr, int visible);
int QFontDialog_TestOption(void* ptr, int option);

#ifdef __cplusplus
}
#endif