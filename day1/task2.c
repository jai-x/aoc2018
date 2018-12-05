#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define ARR_MAX 1000

static int seen[ARR_MAX];
static int seen_end = 0;

static void
store(int val)
{
	seen[seen_end++] = val;
}

static bool
already_seen(int val)
{
	for (int i = 0; i < seen_end; i++) {
		if (seen[i] == val) {
			return true;
		}
	}
	return false;
}

int
main(int argc, char** argv)
{
	FILE* input;
	char line[ARR_MAX];
	int frequency = 0;

	if (argc != 2) {
		printf("Please specify input text file as only argument.\n");
		return EXIT_FAILURE;
	}

	input = fopen(argv[1], "r");
	if (!input) {
		printf("Not a valid file: %s\n", argv[1]);
		return EXIT_FAILURE;
	}

	while (fgets(line, ARR_MAX, input)) {
		int diff = atoi(line);
		frequency += diff;

		// What's wrong?
		if (already_seen(frequency)) {
			printf("Duplicate frequency found: %d\n", frequency);
			return EXIT_SUCCESS;
		} else {
			store(frequency);
		}
	}

	printf("Solution not found!\n");
	return EXIT_FAILURE;
}
