# sgtast
## feature
- [x] : 対象のcsvからバーンダウンチャートを表示する
- [ ] : csvのフォーマットについてヘルプを表示する
- [ ] : ディレクトリ指定で複数のバーンダウンチャートを表示できるようにする

## Dependency
- go 1.19
- labstack/echo
- Google Charts(cdn)

## Usage
```
$ go run . preview <csv file path>
```
##### example
```
$ make preview_sample
```
