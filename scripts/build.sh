#!/bin/sh

prepareBuild(){   
    echo -n "Preparing Build Enviroment...\n"
    rm -rf ./build
    mkdir -p ../build ./build/windows ./build/linux
    echo -n "Done\n"
}

# Build Linux
buildLinux(){
    echo -n "Building Linux Binary...\n"
    go build -o ../build/linux/TES3MP-CRC32 ../src/main.go 
    echo -n "Done\n"
}
buildWindows(){
    echo -n "Building Windows Binary...\n"
    GOOS=windows GOARCH=386 go build -o ../build/windows/TES3MP-CRC32.exe ../src/main.go 
    echo -n "Done\n"
}

prepareBuild
buildLinux
buildWindows