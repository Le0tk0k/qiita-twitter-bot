# Qiita Twitter Bot
QiitaのGo記事お届けBot ([@BotQiita](https://twitter.com/BotQiita))

## Explanation Entry
[Go×Qiita API×LambdaでTwitter Bot作った](https://qiita.com/Le0tk0k/items/7ce7f13514de93bac050)

## Architecture

![Architecture](./docs/architecture.png)

- Qiita APIを利用してGoのタグがついた記事を取得します
- Twitter APIを利用して認証＆記事をツイートします
- lambda関数はGoで書かれており、cloudwatchにより１時間毎に定期実行されます。
