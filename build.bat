@echo off
set build_path=%cd%
set install_path=%USERPROFILE%\gecko

if not exist %install_path%\frontend mkdir %install_path%\frontend\build
echo Created target directory %install_path%

echo Building core components
for /R %build_path%\cmd\ %%f in (*.go) do (
    echo Building %%f
    go.exe build %%f
)

echo Installing binaries in target directory
move %build_path%\*.exe %install_path%

cd %build_path%\frontend

echo Installing any missing dependencies
call npm install

echo Building front end
call npm run build

echo Installing resources in target directory
xcopy /s/e/y %build_path%\frontend\build %install_path%\frontend\build