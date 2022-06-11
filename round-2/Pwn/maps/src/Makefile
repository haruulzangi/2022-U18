all: compile 

compile:
	gcc maps.c -m32 -fno-stack-protector -no-pie -z execstack -Wl,-z,norelro -o maps
