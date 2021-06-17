go run main.go<br>
frontend側のパッケージ管理にはyarnを使ってみる<br>
yarn init = npm init<br>
yarn add = npm install<br>
yarn add -D @babel/cli @babel/core @babel/preset-env babel-loader webpack webpack-cli@3.11.2 webpack-dev-server typescript ts-loader react react-dom @types/react @types/react-dom bootstrap css-loader style-loader @popperjs/core<br>
@types/reactをインストールするとreactをインポートせずにReact.FCの型宣言が可能になる<br>
webpack-cliはバージョン4だとdev-serverが起動できないのでバージョン3をインストールしている<br>
tsconfing.jsonのallowSyntheticDefaultImportsはreact関連のライブラリをインポートするために必要<br>
x := 1とvar x = 1は同じ(変数宣言)<br>