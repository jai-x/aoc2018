#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define LINE_MAX 256
#define ARR_MAX 1234567890

static char line[LINE_MAX];
static int frequency = 0;

static int seen[ARR_MAX];
static int seen_end = 0;

static void
store(int val)
{
	seen[seen_end++] = val;
	
	if (seen_end > ARR_MAX) {
		printf("Make array bigger!\n");
		exit(EXIT_FAILURE);
	}
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

	if (argc != 2) {
		printf("Please specify input text file as only argument.\n");
		exit(EXIT_FAILURE);
	}

	input = fopen(argv[1], "r");
	if (!input) {
		printf("Not a valid file: %s\n", argv[1]);
		exit(EXIT_FAILURE);
	}

	while (true) {
		while (fgets(line, sizeof(line), input)) {
			frequency += atoi(line);

			if (already_seen(frequency)) {
				printf("Duplicate frequency found: %d\n", frequency);
				exit(EXIT_SUCCESS);
			} else {
				store(frequency);
			}
		}

		fseek(input, 0, SEEK_END);
		fseek(input, 0, SEEK_SET);
	}
}
