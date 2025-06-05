package yml

import (
	"context"
	"fmt"
	"io"
	"time"
)

func (o *offersImport) startImport(ctx context.Context, authOpt entity.AuthOptions, feedUrl string, maxSizeLimitFeed, httpRequestTimeout int) (result entity.FeedHistory, err error) {
	var fileId string
	fileId, err = o.downloadFile(authOpt, feedUrl, maxSizeLimitFeed, httpRequestTimeout)
	if err != nil {
		logger.ErrorWithErr(ctx, err, "error parsing offers")
		return result, err
	}

	defer func() {
		//Подчищаем файлы
		//новый сtx на удаление
		newCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		deferErr := o.mongoRepo.DeleteFeedFileFromGridFs(newCtx, fileId)
		if deferErr != nil {
			logger.ErrorWithErr(newCtx, deferErr, "cant delete file from mongo")
		}
		cancel()
	}()

	var fl io.ReadCloser
	fl, err = o.mongoRepo.GetFeedFileFromGridFs(fileId)
	if err != nil {
		return
	}
	defer fl.Close()
	iter, err := yml.NewFromReader(ctx, fl)
	if err != nil {
		logger.ErrorWithErr(ctx, err, "error parsing offers")
		return
	}

	defer func() {
		err = iter.Close()
		if err != nil {
			logger.ErrorWithErr(ctx, err, fmt.Sprintf("Error close file %s", fileId))
		}
	}()
	result, err = o.baseImportYML(ctx, iter)
	if err != nil {
		return result, err
	}
	if *o.isServiceStop || o.isImportStop {
		return
	}

	return result, nil
}
