#include "runtime.h"

void .Add(int a, int b, int ret)
{
	ret = a + b;
	FLUSH(&ret)
}
