#ifdef __cplusplus
extern "C" {
#endif

int QFontComboBox_FontFilters(void* ptr);
void QFontComboBox_SetCurrentFont(void* ptr, void* font);
void QFontComboBox_SetFontFilters(void* ptr, int filters);
void QFontComboBox_SetWritingSystem(void* ptr, int script);
int QFontComboBox_WritingSystem(void* ptr);
void* QFontComboBox_NewQFontComboBox(void* parent);
void QFontComboBox_DestroyQFontComboBox(void* ptr);

#ifdef __cplusplus
}
#endif