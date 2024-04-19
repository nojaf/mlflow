package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mlflow/mlflow/mlflow/go/pkg/protos"
	"github.com/mlflow/mlflow/mlflow/go/pkg/protos/artifacts"
)

type MlflowService interface {
	GetExperimentByName(input *protos.GetExperimentByName) (protos.GetExperimentByName_Response, *MlflowError)
	CreateExperiment(input *protos.CreateExperiment) (protos.CreateExperiment_Response, *MlflowError)
	SearchExperiments(input *protos.SearchExperiments) (protos.SearchExperiments_Response, *MlflowError)
	GetExperiment(input *protos.GetExperiment) (protos.GetExperiment_Response, *MlflowError)
	DeleteExperiment(input *protos.DeleteExperiment) (protos.DeleteExperiment_Response, *MlflowError)
	RestoreExperiment(input *protos.RestoreExperiment) (protos.RestoreExperiment_Response, *MlflowError)
	UpdateExperiment(input *protos.UpdateExperiment) (protos.UpdateExperiment_Response, *MlflowError)
	CreateRun(input *protos.CreateRun) (protos.CreateRun_Response, *MlflowError)
	UpdateRun(input *protos.UpdateRun) (protos.UpdateRun_Response, *MlflowError)
	DeleteRun(input *protos.DeleteRun) (protos.DeleteRun_Response, *MlflowError)
	RestoreRun(input *protos.RestoreRun) (protos.RestoreRun_Response, *MlflowError)
	LogMetric(input *protos.LogMetric) (protos.LogMetric_Response, *MlflowError)
	LogParam(input *protos.LogParam) (protos.LogParam_Response, *MlflowError)
	SetExperimentTag(input *protos.SetExperimentTag) (protos.SetExperimentTag_Response, *MlflowError)
	SetTag(input *protos.SetTag) (protos.SetTag_Response, *MlflowError)
	DeleteTag(input *protos.DeleteTag) (protos.DeleteTag_Response, *MlflowError)
	GetRun(input *protos.GetRun) (protos.GetRun_Response, *MlflowError)
	SearchRuns(input *protos.SearchRuns) (protos.SearchRuns_Response, *MlflowError)
	ListArtifacts(input *protos.ListArtifacts) (protos.ListArtifacts_Response, *MlflowError)
	GetMetricHistory(input *protos.GetMetricHistory) (protos.GetMetricHistory_Response, *MlflowError)
	GetMetricHistoryBulkInterval(input *protos.GetMetricHistoryBulkInterval) (protos.GetMetricHistoryBulkInterval_Response, *MlflowError)
	LogBatch(input *protos.LogBatch) (protos.LogBatch_Response, *MlflowError)
	LogModel(input *protos.LogModel) (protos.LogModel_Response, *MlflowError)
	LogInputs(input *protos.LogInputs) (protos.LogInputs_Response, *MlflowError)
}
type ModelRegistryService interface {
	CreateRegisteredModel(input *protos.CreateRegisteredModel) (protos.CreateRegisteredModel_Response, *MlflowError)
	RenameRegisteredModel(input *protos.RenameRegisteredModel) (protos.RenameRegisteredModel_Response, *MlflowError)
	UpdateRegisteredModel(input *protos.UpdateRegisteredModel) (protos.UpdateRegisteredModel_Response, *MlflowError)
	DeleteRegisteredModel(input *protos.DeleteRegisteredModel) (protos.DeleteRegisteredModel_Response, *MlflowError)
	GetRegisteredModel(input *protos.GetRegisteredModel) (protos.GetRegisteredModel_Response, *MlflowError)
	SearchRegisteredModels(input *protos.SearchRegisteredModels) (protos.SearchRegisteredModels_Response, *MlflowError)
	GetLatestVersions(input *protos.GetLatestVersions) (protos.GetLatestVersions_Response, *MlflowError)
	CreateModelVersion(input *protos.CreateModelVersion) (protos.CreateModelVersion_Response, *MlflowError)
	UpdateModelVersion(input *protos.UpdateModelVersion) (protos.UpdateModelVersion_Response, *MlflowError)
	TransitionModelVersionStage(input *protos.TransitionModelVersionStage) (protos.TransitionModelVersionStage_Response, *MlflowError)
	DeleteModelVersion(input *protos.DeleteModelVersion) (protos.DeleteModelVersion_Response, *MlflowError)
	GetModelVersion(input *protos.GetModelVersion) (protos.GetModelVersion_Response, *MlflowError)
	SearchModelVersions(input *protos.SearchModelVersions) (protos.SearchModelVersions_Response, *MlflowError)
	GetModelVersionDownloadUri(input *protos.GetModelVersionDownloadUri) (protos.GetModelVersionDownloadUri_Response, *MlflowError)
	SetRegisteredModelTag(input *protos.SetRegisteredModelTag) (protos.SetRegisteredModelTag_Response, *MlflowError)
	SetModelVersionTag(input *protos.SetModelVersionTag) (protos.SetModelVersionTag_Response, *MlflowError)
	DeleteRegisteredModelTag(input *protos.DeleteRegisteredModelTag) (protos.DeleteRegisteredModelTag_Response, *MlflowError)
	DeleteModelVersionTag(input *protos.DeleteModelVersionTag) (protos.DeleteModelVersionTag_Response, *MlflowError)
	SetRegisteredModelAlias(input *protos.SetRegisteredModelAlias) (protos.SetRegisteredModelAlias_Response, *MlflowError)
	DeleteRegisteredModelAlias(input *protos.DeleteRegisteredModelAlias) (protos.DeleteRegisteredModelAlias_Response, *MlflowError)
	GetModelVersionByAlias(input *protos.GetModelVersionByAlias) (protos.GetModelVersionByAlias_Response, *MlflowError)
}
type MlflowArtifactsService interface {
	DownloadArtifact(input *artifacts.DownloadArtifact) (artifacts.DownloadArtifact_Response, *MlflowError)
	UploadArtifact(input *artifacts.UploadArtifact) (artifacts.UploadArtifact_Response, *MlflowError)
	ListArtifacts(input *artifacts.ListArtifacts) (artifacts.ListArtifacts_Response, *MlflowError)
	DeleteArtifact(input *artifacts.DeleteArtifact) (artifacts.DeleteArtifact_Response, *MlflowError)
	CreateMultipartUpload(input *artifacts.CreateMultipartUpload) (artifacts.CreateMultipartUpload_Response, *MlflowError)
	CompleteMultipartUpload(input *artifacts.CompleteMultipartUpload) (artifacts.CompleteMultipartUpload_Response, *MlflowError)
	AbortMultipartUpload(input *artifacts.AbortMultipartUpload) (artifacts.AbortMultipartUpload_Response, *MlflowError)
}

func registerMlflowServiceRoutes(service MlflowService, app *fiber.App) {
	app.Get("/api/2.0/mlflow/experiments/get-by-name", func(ctx *fiber.Ctx) error {
		var input *protos.GetExperimentByName
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetExperimentByName(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/experiments/create", func(ctx *fiber.Ctx) error {
		var input *protos.CreateExperiment
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.CreateExperiment(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/experiments/search", func(ctx *fiber.Ctx) error {
		var input *protos.SearchExperiments
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.SearchExperiments(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/experiments/search", func(ctx *fiber.Ctx) error {
		var input *protos.SearchExperiments
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.SearchExperiments(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/experiments/get", func(ctx *fiber.Ctx) error {
		var input *protos.GetExperiment
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetExperiment(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/experiments/delete", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteExperiment
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteExperiment(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/experiments/restore", func(ctx *fiber.Ctx) error {
		var input *protos.RestoreExperiment
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.RestoreExperiment(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/experiments/update", func(ctx *fiber.Ctx) error {
		var input *protos.UpdateExperiment
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.UpdateExperiment(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/create", func(ctx *fiber.Ctx) error {
		var input *protos.CreateRun
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.CreateRun(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/update", func(ctx *fiber.Ctx) error {
		var input *protos.UpdateRun
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.UpdateRun(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/delete", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteRun
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteRun(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/restore", func(ctx *fiber.Ctx) error {
		var input *protos.RestoreRun
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.RestoreRun(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/log-metric", func(ctx *fiber.Ctx) error {
		var input *protos.LogMetric
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.LogMetric(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/log-parameter", func(ctx *fiber.Ctx) error {
		var input *protos.LogParam
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.LogParam(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/experiments/set-experiment-tag", func(ctx *fiber.Ctx) error {
		var input *protos.SetExperimentTag
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.SetExperimentTag(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/set-tag", func(ctx *fiber.Ctx) error {
		var input *protos.SetTag
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.SetTag(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/delete-tag", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteTag
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteTag(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/runs/get", func(ctx *fiber.Ctx) error {
		var input *protos.GetRun
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetRun(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/search", func(ctx *fiber.Ctx) error {
		var input *protos.SearchRuns
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.SearchRuns(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/artifacts/list", func(ctx *fiber.Ctx) error {
		var input *protos.ListArtifacts
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.ListArtifacts(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/metrics/get-history", func(ctx *fiber.Ctx) error {
		var input *protos.GetMetricHistory
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetMetricHistory(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/metrics/get-history-bulk-interval", func(ctx *fiber.Ctx) error {
		var input *protos.GetMetricHistoryBulkInterval
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetMetricHistoryBulkInterval(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/log-batch", func(ctx *fiber.Ctx) error {
		var input *protos.LogBatch
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.LogBatch(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/log-model", func(ctx *fiber.Ctx) error {
		var input *protos.LogModel
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.LogModel(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/runs/log-inputs", func(ctx *fiber.Ctx) error {
		var input *protos.LogInputs
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.LogInputs(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
}
func registerModelRegistryServiceRoutes(service ModelRegistryService, app *fiber.App) {
	app.Post("/api/2.0/mlflow/registered-models/create", func(ctx *fiber.Ctx) error {
		var input *protos.CreateRegisteredModel
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.CreateRegisteredModel(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/registered-models/rename", func(ctx *fiber.Ctx) error {
		var input *protos.RenameRegisteredModel
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.RenameRegisteredModel(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Patch("/api/2.0/mlflow/registered-models/update", func(ctx *fiber.Ctx) error {
		var input *protos.UpdateRegisteredModel
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.UpdateRegisteredModel(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Delete("/api/2.0/mlflow/registered-models/delete", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteRegisteredModel
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteRegisteredModel(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/registered-models/get", func(ctx *fiber.Ctx) error {
		var input *protos.GetRegisteredModel
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetRegisteredModel(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/registered-models/search", func(ctx *fiber.Ctx) error {
		var input *protos.SearchRegisteredModels
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.SearchRegisteredModels(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/registered-models/get-latest-versions", func(ctx *fiber.Ctx) error {
		var input *protos.GetLatestVersions
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.GetLatestVersions(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/registered-models/get-latest-versions", func(ctx *fiber.Ctx) error {
		var input *protos.GetLatestVersions
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetLatestVersions(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/model-versions/create", func(ctx *fiber.Ctx) error {
		var input *protos.CreateModelVersion
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.CreateModelVersion(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Patch("/api/2.0/mlflow/model-versions/update", func(ctx *fiber.Ctx) error {
		var input *protos.UpdateModelVersion
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.UpdateModelVersion(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/model-versions/transition-stage", func(ctx *fiber.Ctx) error {
		var input *protos.TransitionModelVersionStage
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.TransitionModelVersionStage(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Delete("/api/2.0/mlflow/model-versions/delete", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteModelVersion
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteModelVersion(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/model-versions/get", func(ctx *fiber.Ctx) error {
		var input *protos.GetModelVersion
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetModelVersion(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/model-versions/search", func(ctx *fiber.Ctx) error {
		var input *protos.SearchModelVersions
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.SearchModelVersions(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/model-versions/get-download-uri", func(ctx *fiber.Ctx) error {
		var input *protos.GetModelVersionDownloadUri
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetModelVersionDownloadUri(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/registered-models/set-tag", func(ctx *fiber.Ctx) error {
		var input *protos.SetRegisteredModelTag
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.SetRegisteredModelTag(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/model-versions/set-tag", func(ctx *fiber.Ctx) error {
		var input *protos.SetModelVersionTag
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.SetModelVersionTag(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Delete("/api/2.0/mlflow/registered-models/delete-tag", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteRegisteredModelTag
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteRegisteredModelTag(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Delete("/api/2.0/mlflow/model-versions/delete-tag", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteModelVersionTag
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteModelVersionTag(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow/registered-models/alias", func(ctx *fiber.Ctx) error {
		var input *protos.SetRegisteredModelAlias
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.SetRegisteredModelAlias(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Delete("/api/2.0/mlflow/registered-models/alias", func(ctx *fiber.Ctx) error {
		var input *protos.DeleteRegisteredModelAlias
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteRegisteredModelAlias(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow/registered-models/alias", func(ctx *fiber.Ctx) error {
		var input *protos.GetModelVersionByAlias
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.GetModelVersionByAlias(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
}
func registerMlflowArtifactsServiceRoutes(service MlflowArtifactsService, app *fiber.App) {
	app.Get("/api/2.0/mlflow-artifacts/artifacts/:path", func(ctx *fiber.Ctx) error {
		var input *artifacts.DownloadArtifact
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.DownloadArtifact(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Put("/api/2.0/mlflow-artifacts/artifacts/:path", func(ctx *fiber.Ctx) error {
		var input *artifacts.UploadArtifact
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.UploadArtifact(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Get("/api/2.0/mlflow-artifacts/artifacts", func(ctx *fiber.Ctx) error {
		var input *artifacts.ListArtifacts
		if err := ctx.QueryParser(&input); err != nil {
			return err
		}
		output, err := service.ListArtifacts(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Delete("/api/2.0/mlflow-artifacts/artifacts/:path", func(ctx *fiber.Ctx) error {
		var input *artifacts.DeleteArtifact
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.DeleteArtifact(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow-artifacts/mpu/create/:path", func(ctx *fiber.Ctx) error {
		var input *artifacts.CreateMultipartUpload
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.CreateMultipartUpload(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow-artifacts/mpu/complete/:path", func(ctx *fiber.Ctx) error {
		var input *artifacts.CompleteMultipartUpload
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.CompleteMultipartUpload(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
	app.Post("/api/2.0/mlflow-artifacts/mpu/abort/:path", func(ctx *fiber.Ctx) error {
		var input *artifacts.AbortMultipartUpload
		if err := ctx.BodyParser(&input); err != nil {
			return err
		}
		output, err := service.AbortMultipartUpload(input)
		if err != nil && err.ErrorCode == protos.ErrorCode_NOT_IMPLEMENTED {
			return ctx.Next()
		}
		if err != nil {
			return err
		}
		return ctx.JSON(&output)
	})
}
