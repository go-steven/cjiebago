all: libjieba.a
install:
	mv libjieba.a /usr/local/lib/libjieba.a
libjieba.a:
	g++ -v -Wall -o jieba.o -c -DLOGGING_LEVEL=LL_WARNING -I deps/ lib/jieba.cpp
	ar rs libjieba.a jieba.o
	rm -f *.o
clean:
	rm -f *.a *.o