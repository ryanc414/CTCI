#include <stdio.h>
#include <time.h>
#include <stdlib.h>

void swap(int *x, int *y);

int main()
{
	srand(time(NULL));

	int x, y = rand(), rand();
	printf("before swap: x = %d, y = %d\n", x, y);

	swap(&x, &y);

	printf("after swap: x = %d, y = %d\n", x, y);

	return 0;
}

void swap(int *x, int *y)
{
	*x = *x ^ *y;
	*y = *x ^ *y;
	*x = *x ^ *y;
}