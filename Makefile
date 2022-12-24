# misc/sample.csv を利用してプレビューを実施
FILE := misc/sample
SAMPLEDIR := misc/sample
preview_sample:
	go run . preview $(SAMPLEDIR)

# 直近の日付に合わせてsample.csvファイルを更新する
update_sample:
	echo "1,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-4d '+%Y-%m-%d')" > $(FILE)
	echo "2,$(shell date -v-5d '+%Y-%m-%d')," >> $(FILE)
	echo "3,$(shell date -v-2d '+%Y-%m-%d'),$(shell date -v-1d '+%Y-%m-%d')" >> $(FILE)

test:
	go test ./... -v
