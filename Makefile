# misc/sample.csv を利用してプレビューを実施
FILE := misc/sample
SAMPLEDIR := misc/sample
preview_sample:
	go run . preview $(SAMPLEDIR)

preview_help:
	go run . help preview

lint_sample:
	go run . lint $(SAMPLEDIR)

lint_help:
	go run . help lint

# 直近の日付に合わせてmis/samplecちょっかのファイルを更新する
update_sample:
	echo "1,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-4d '+%Y-%m-%d')" > $(SAMPLEDIR)/1.csv
	echo "2,$(shell date -v-5d '+%Y-%m-%d')," >> $(SAMPLEDIR)/1.csv
	echo "3,$(shell date -v-2d '+%Y-%m-%d'),$(shell date -v-1d '+%Y-%m-%d')" >> $(SAMPLEDIR)/1.csv
	echo "5,$(shell date -v-3d '+%Y-%m-%d'),$(shell date -v-3d '+%Y-%m-%d')" > $(SAMPLEDIR)/2.csv
	echo "1,$(shell date -v-7d '+%Y-%m-%d')," >> $(SAMPLEDIR)/2.csv
	echo "2,$(shell date -v-4d '+%Y-%m-%d'),$(shell date -v-1d '+%Y-%m-%d')" >> $(SAMPLEDIR)/2.csv
	echo "5,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-5d '+%Y-%m-%d')" > $(SAMPLEDIR)/3.csv
	echo "1,$(shell date -v-5d '+%Y-%m-%d')," >> $(SAMPLEDIR)/3.csv
	echo "2,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-4d '+%Y-%m-%d')" >> $(SAMPLEDIR)/3.csv
	echo "3,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-3d '+%Y-%m-%d')" >> $(SAMPLEDIR)/3.csv
	echo "2,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-2d '+%Y-%m-%d')" >> $(SAMPLEDIR)/3.csv
	echo "2,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-4d '+%Y-%m-%d')" > $(SAMPLEDIR)/4.csv
	echo "3,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-3d '+%Y-%m-%d')" >> $(SAMPLEDIR)/4.csv
	echo "2,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-2d '+%Y-%m-%d')" >> $(SAMPLEDIR)/4.csv
	echo "3,$(shell date -v-5d '+%Y-%m-%d'),$(shell date -v-4d '+%Y-%m-%d')" > $(SAMPLEDIR)/5.csv
	echo "1,$(shell date -v-5d '+%Y-%m-%d')," >> $(SAMPLEDIR)/5.csv
	echo "3,$(shell date -v-2d '+%Y-%m-%d'),$(shell date -v-1d '+%Y-%m-%d')" >> $(SAMPLEDIR)/5.csv

test:
	go test ./... -v
