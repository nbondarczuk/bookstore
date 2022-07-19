TARGET=bookstore

build:
	(cd cmd; go build -o ../$(TARGET))

clean:
	rm -f *~  $(TARGET)



