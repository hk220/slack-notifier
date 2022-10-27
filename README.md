# slack-sender

標準入力の内容をslackに書き込むコマンドです。

## 設定ファイル

下記のフォーマットの設定ファイルを`$HOME/.slack-sender.toml`に作成してください。

```toml
token=""
username="slack-sender"
channel="#random"
```

## 使い方

例えば、現在の時間をslackの`#random`チャネルに`slack-sender`として書き込む場合は以下のようにします。

```bash
date | slack-sender --token "xorb-XXXXXXX" --channel "#random" --username "slack-sender"
```
