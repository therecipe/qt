#ifdef __cplusplus
extern "C" {
#endif

void* QTextFragment_NewQTextFragment();
void* QTextFragment_NewQTextFragment3(void* other);
int QTextFragment_CharFormatIndex(void* ptr);
int QTextFragment_Contains(void* ptr, int position);
int QTextFragment_IsValid(void* ptr);
int QTextFragment_Length(void* ptr);
int QTextFragment_Position(void* ptr);
char* QTextFragment_Text(void* ptr);

#ifdef __cplusplus
}
#endif