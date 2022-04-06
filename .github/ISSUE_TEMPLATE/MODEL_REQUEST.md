---
name: modelタスク
about: model追加はこちら
labels: develop
---

## 🎉 概要
ユーザーに紐づく企業情報を追加できるようにCompanyモデルを追加してください。


## 💪 なぜ必要なのか？
ユーザーに紐づく企業情報を取り扱うため

## 📖 参考資料 (optional)

(参考リンクなどあれば、なければ消す)


## 📎 TODO

- [ ] `Company`モデルの作成


## 🎁 バリデーションチェック
- [ ] `name`が必須であること


## 🔑 Rspecテスト

- [ ] Rspecを実行するためにfactoryを作成すること
- [ ] モデルテストで`name`がないと保存できないこと、のテスト追加
- [ ] モデルテストで`name`がない場合はバリデーションエラーが出ること、のテスト追加

## 🖼 Lint/テストチェック

このタスク完了後にLintやテストコマンドをターミナルで走らせてエラーが出ないことを確認してください。

- [ ] `bundle exec rspec`が通ること
- [ ] `bundle exec scss-lint app/javascript/stylesheets/`が通ること
- [ ] `bundle exec haml-lint app/views/`が通ること
- [ ] `bundle exec rubocop`が通ること
- [ ] `bundle exec brakeman`が通ること
- [ ] `bundle exec rails_best_practices`が通ること


## 🎁 review app最終確認
コミットをpushしてpull requestを作成すると自動でHerokuにこのissue(PR)に紐づくreview appが生成されます。
このタスクで完了すべきリストを確認し、全て完了していたらコードレビューに出しましょう！
※review appとはissueごとにアプリをHerokuにデプロイ→生成して、テストの環境のアプリのことです。developやmasterにマージしてからだとエラー検知が遅れる可能性があるので、事前にテストの環境を作成して、バグを取り除く意図で作成しています。
※review appはpull requestを作成すると自動的に作成されて、pull requestの画面下部の`View deployment`から確認が可能です。

[画像で確認する](https://github.com/gen8888/cosmeticdb/blob/main/review_app.png)



- [ ] review appで本タスクの要件を全て満たしていること


## DB設計
timestampは自動で入るので記載していません。

カラム名   |  型
----- | -----
id  | pk
name   |  string NOT NULL
user_id | FK NOT NULL
