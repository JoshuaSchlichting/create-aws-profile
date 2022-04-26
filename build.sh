APP_NAME=assume-role

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS=linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
    OS=macos
elif [[ "$OSTYPE" == "cygwin" ]]; then
    OS=cygwin
elif [[ "$OSTYPE" == "msys" ]]; then
    # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
    OS=mysys
elif [[ "$OSTYPE" == "win32" ]]; then
    OS=windows
elif [[ "$OSTYPE" == "freebsd"* ]]; then
    OS=freebsd
else
    OS=unknown
fi

echo $OS
BIN_FILENAME=bin/$APP_NAME\_$(echo $OS)_$(uname -m)
echo FILENAME = $BIN_FILENAME
go build -o $BIN_FILENAME .
chmod +x $BIN_FILENAME

if [ $# -gt 0 ];then
    if [ $1 = "--install" ]; then
        cp $BIN_FILENAME /usr/local/bin/$APP_NAME
    fi
fi