#include <stdio.h>
#include <stdlib.h>

#define LINE_MAX 256

int
main(int argc, char** argv)
{
	if (argc != 2) {
		printf("Please specify input text file as only argument.\n");
		return EXIT_FAILURE;
	}

	FILE* input;
	input = fopen(argv[1], "r");
	if (!input) {
		printf("Not a valid file: %s\n", argv[1]);
		return EXIT_FAILURE;
	}

	char line[LINE_MAX];
	int frequency = 0;
	while (fgets(line, LINE_MAX, input)) {
		frequency += atoi(line);
	}

	printf("Frequency: %d\n", frequency);

	return EXIT_SUCCESS;
}
