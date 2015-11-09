#ifdef __cplusplus
extern "C" {
#endif

void* QTextLength_NewQTextLength();
void* QTextLength_NewQTextLength2(int ty, double value);
double QTextLength_RawValue(void* ptr);
int QTextLength_Type(void* ptr);
double QTextLength_Value(void* ptr, double maximumLength);

#ifdef __cplusplus
}
#endif