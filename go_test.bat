@echo off
setlocal
set PATH=%PATH%;%~dp0%\build\Release;
call go test . ./apis/transaction ./apis/block ./tests -v
