gophersay: main.go
	go build main.go -o gophersay

install: gophersay
	echo "Checking for sudo perms.."
	sudo echo "Done"

	sudo cp gophersay /usr/local/bin/gophersay