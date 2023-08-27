# executable for linux based systems
linux-cli:
	go build -o family-tree main.go

# executable for Windows
win-cli:
	go build -o family-tree.exe main.go	