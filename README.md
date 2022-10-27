# slack-sender

標準入力の内容をslackに書き込むコマンドです。

## 使い方

例えば、現在の時間をslackの`#random`チャネルに`slack-sender`として書き込む場合は以下のようにします。

```bash
# 設定ファイルを使わない場合
date | slack-sender --token "xorb-XXXXXXX" --channel "#random" --username "slack-sender"

# 設定ファイルを使う場合
date | slack-sender
```

## 設定ファイルの例

```toml
token=""
username="slack-sender"
channel="#random"
```
`$HOME/.slack-sender.toml`に作成してください。
