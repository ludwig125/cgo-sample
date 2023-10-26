all: libmessage.so test

libmessage.so: message.c
	gcc -shared -fPIC -o $@ $^

test: test.o
	gcc -L./ -o $@ $^ -lmessage

%.o : %.c
	gcc -o $@ -c $<

clean:
	-rm -f *.so *.o
