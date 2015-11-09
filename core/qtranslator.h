#ifdef __cplusplus
extern "C" {
#endif

void* QTranslator_NewQTranslator(void* parent);
int QTranslator_IsEmpty(void* ptr);
int QTranslator_Load2(void* ptr, void* locale, char* filename, char* prefix, char* directory, char* suffix);
int QTranslator_Load(void* ptr, char* filename, char* directory, char* search_delimiters, char* suffix);
char* QTranslator_Translate(void* ptr, char* context, char* sourceText, char* disambiguation, int n);
void QTranslator_DestroyQTranslator(void* ptr);

#ifdef __cplusplus
}
#endif