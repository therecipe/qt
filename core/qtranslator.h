#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTranslator_NewQTranslator(QtObjectPtr parent);
int QTranslator_IsEmpty(QtObjectPtr ptr);
int QTranslator_Load2(QtObjectPtr ptr, QtObjectPtr locale, char* filename, char* prefix, char* directory, char* suffix);
int QTranslator_Load(QtObjectPtr ptr, char* filename, char* directory, char* search_delimiters, char* suffix);
char* QTranslator_Translate(QtObjectPtr ptr, char* context, char* sourceText, char* disambiguation, int n);
void QTranslator_DestroyQTranslator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif