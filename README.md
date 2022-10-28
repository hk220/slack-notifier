# slack-notifier

標準入力の内容をslackに書き込むコマンドです。

## 使い方

例えば、現在の時間をslackの`#random`チャネルに`slack-notifier`として書き込む場合は以下のようにします。

```bash
# 設定ファイルを使わない場合
date | slack-notifier --token "xorb-XXXXXXX" --channel "#random" --username "slack-notifier"

# 設定ファイルを使う場合
date | slack-notifier
```

## 設定ファイルの例

```toml
token=""
username="slack-notifier"
channel="#random"
```
`$HOME/.slack-notifier.toml`に作成してください。
