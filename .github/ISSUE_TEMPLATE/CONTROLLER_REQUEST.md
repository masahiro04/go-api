---
name: controllerタスク
about: controller追加はこちら
labels: develop
---

## 🎉 概要
運営会社が行ったセミナー動画を表示するページを作成してください。

## 💪 なぜ必要なのか？
ユーザーに過去に撮影したセミナー動画を提供するため。

## 作成イメージ
Viewの完成後イメージの画像はる
※原則アプリオーナーが作成したパワポ画像


## 📖 参考テンプレ情報
UIはこちらを参考にしてください。
※利用するテンプレファイルのパスを指定する
※パスだけでなくそのテンプレページの画像も貼ること
`public/global_assets/Limitless_2_3/Template/layout_1/LTR/default/full/gallery_titles.html`


## 📎 TODO

- [ ] `/dashboard/movies_controller`を作成すること
- [ ] index/showアクションの設定を行うこと

※TODO は
- index/show
- new/create
- edit/update/delete
でタスクを分けること。
また分ける際は事前に完了すべきタスクをこのタスク上部に記載すること


## 🎁 バリデーションチェック
- [ ] モデルのバリデーションエラーが発生したら、そのエラーメッセージが表示されること


## 🔑 Rspecテスト

- [ ] index/showのリクエストテストを追加すること

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
