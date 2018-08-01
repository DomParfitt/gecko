#!/bin/bash 
build_path=$(pwd)
install_path="$HOME/gecko"

mkdir -p $install_path/frontend
echo "Created target directory $install_path"

echo "Building core components"
for dir in $build_path/cmd/*; do
    if [ -d $dir ]; then
        echo "Building ${dir##*/}"
        go build $dir/*.go
    fi
done

echo "Installing binaries in target directory"
mv $build_path/*.exe $install_path

cd $build_path/frontend

echo "Installing any missing dependencies"
npm install

echo "Building front end"
npm run build

echo "Installing resources in target directory"
cp -r $build_path/frontend/build/ $install_path/frontend