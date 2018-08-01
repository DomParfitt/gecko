#!/bin/bash 
gecko_root=$(pwd)
gecko_path="$HOME/gecko"

mkdir -p $gecko_path/frontend
echo "Created target directory $gecko_path"

echo "Building core components"
for dir in $gecko_root/cmd/*; do
    if [ -d $dir ]; then
        echo $dir/
        echo "Building ${dir##*/}"
        go build $dir/*.go
    fi
done

mv $gecko_root/*.exe $gecko_path

cd $gecko_root/frontend

echo "Installing any missing dependencies"
npm install

echo "Building front end"
npm run build

cp -r $gecko_root/frontend/build/ $gecko_path/frontend