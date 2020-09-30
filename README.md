# Qiita Twitter Bot
[@BotQiita](https://twitter.com/BotQiita)  

## Architecture

![Architecture](./docs/architecture.png)

- Qiita APIを利用してGoのタグがついた記事を取得します
- Twitter APIを利用して記事をツイートします
- lambda関数はGoで書かれており、cloudwatchにより１時間毎に定期実行されます。
