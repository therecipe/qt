#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFontComboBox_FontFilters(QtObjectPtr ptr);
void QFontComboBox_SetCurrentFont(QtObjectPtr ptr, QtObjectPtr font);
void QFontComboBox_SetFontFilters(QtObjectPtr ptr, int filters);
void QFontComboBox_SetWritingSystem(QtObjectPtr ptr, int script);
int QFontComboBox_WritingSystem(QtObjectPtr ptr);
QtObjectPtr QFontComboBox_NewQFontComboBox(QtObjectPtr parent);
void QFontComboBox_DestroyQFontComboBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif