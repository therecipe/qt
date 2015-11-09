#ifdef __cplusplus
extern "C" {
#endif

void* QMimeDatabase_NewQMimeDatabase();
void QMimeDatabase_DestroyQMimeDatabase(void* ptr);
char* QMimeDatabase_SuffixForFileName(void* ptr, char* fileName);

#ifdef __cplusplus
}
#endif