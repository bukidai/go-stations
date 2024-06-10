#!/bin/bash

# BASIC認証のテスト

username=admin
password=pass

echo "正常系項目"

# 正常系
echo "正しいユーザIDとパスワードを入力"
echo "runcommand: curl -i -u $username:$password http://localhost:8080/basic"
echo "期待値: 200 OK"
curl -i -u $username:$password http://localhost:8080/basic

echo "BASIC認証対象外にアクセス"
echo "runcommand: curl -i http://localhost:8080/show-os"
echo "期待値: 200 OK"
curl -i http://localhost:8080/show-os

echo "異常系項目"

# 異常系
echo "ユーザIDが間違っている"
echo "runcommand: curl -i -u wrong:$password http://localhost:8080/basic"
echo "期待値: 401 Unauthorized"
curl -i -u wrong:$password http://localhost:8080/basic

echo "パスワードが間違っている"
echo "runcommand: curl -i -u $username:wrong http://localhost:8080/basic"
echo "期待値: 401 Unauthorized"
curl -i -u $username:wrong http://localhost:8080/basic

echo "ユーザIDとパスワードが間違っている"
echo "runcommand: curl -i -u wrong:wrong http://localhost:8080/basic"
echo "期待値: 401 Unauthorized"
curl -i -u wrong:wrong http://localhost:8080/basic

echo "ユーザIDが空"
echo "runcommand: curl -i -u :$password http://localhost:8080/basic"
echo "期待値: 401 Unauthorized"
curl -i -u :$password http://localhost:8080/basic

echo "パスワードが空"
echo "runcommand: curl -i -u $username: http://localhost:8080/basic"
echo "期待値: 401 Unauthorized"
curl -i -u $username: http://localhost:8080/basic

echo "ユーザIDとパスワードが空"
echo "runcommand: curl -i -u : http://localhost:8080/basic"
echo "期待値: 401 Unauthorized"
curl -i -u : http://localhost:8080/basic

echo "認証情報がない"
echo "runcommand: curl -i http://localhost:8080/basic"
echo "期待値: 401 Unauthorized"
curl -i http://localhost:8080/basic