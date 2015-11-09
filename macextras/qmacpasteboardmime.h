#ifdef __cplusplus
extern "C" {
#endif

char* QMacPasteboardMime_ConvertorName(void* ptr);
int QMacPasteboardMime_Count(void* ptr, void* mimeData);
char* QMacPasteboardMime_FlavorFor(void* ptr, char* mime);
char* QMacPasteboardMime_MimeFor(void* ptr, char* flav);
void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr);

#ifdef __cplusplus
}
#endif