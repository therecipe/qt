#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFontDialog_Options(QtObjectPtr ptr);
void QFontDialog_SetOptions(QtObjectPtr ptr, int options);
QtObjectPtr QFontDialog_NewQFontDialog(QtObjectPtr parent);
QtObjectPtr QFontDialog_NewQFontDialog2(QtObjectPtr initial, QtObjectPtr parent);
void QFontDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member);
void QFontDialog_SetCurrentFont(QtObjectPtr ptr, QtObjectPtr font);
void QFontDialog_SetOption(QtObjectPtr ptr, int option, int on);
void QFontDialog_SetVisible(QtObjectPtr ptr, int visible);
int QFontDialog_TestOption(QtObjectPtr ptr, int option);

#ifdef __cplusplus
}
#endif