package controller

import (
	"github.com/yagizdagabak/answer/internal/base/handler"
	"github.com/yagizdagabak/answer/internal/base/middleware"
	"github.com/yagizdagabak/answer/internal/schema"
	"github.com/yagizdagabak/answer/internal/service/activity"
	"github.com/yagizdagabak/answer/internal/service/activity_common"
	"github.com/yagizdagabak/answer/internal/service/role"
	"github.com/yagizdagabak/answer/pkg/uid"
	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	activityCommonService *activity_common.ActivityCommon
	activityService       *activity.ActivityService
}

// NewActivityController new activity controller.
func NewActivityController(
	activityCommonService *activity_common.ActivityCommon,
	activityService *activity.ActivityService) *ActivityController {
	return &ActivityController{activityCommonService: activityCommonService, activityService: activityService}
}

// GetObjectTimeline get object timeline
// @Summary get object timeline
// @Description get object timeline
// @Tags Comment
// @Produce json
// @Param object_id query string false "object id"
// @Param tag_slug_name query string false "tag slug name"
// @Param object_type query string false "object type" Enums(question, answer, tag)
// @Param show_vote query boolean false "is show vote"
// @Success 200 {object} handler.RespBody{data=schema.GetObjectTimelineResp}
// @Router /answer/api/v1/activity/timeline [get]
func (ac *ActivityController) GetObjectTimeline(ctx *gin.Context) {
	req := &schema.GetObjectTimelineReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.ObjectID = uid.DeShortID(req.ObjectID)

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	if userInfo := middleware.GetUserInfoFromContext(ctx); userInfo != nil {
		req.IsAdmin = userInfo.RoleID == role.RoleAdminID
	}

	resp, err := ac.activityService.GetObjectTimeline(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetObjectTimelineDetail get object timeline detail
// @Summary get object timeline detail
// @Description get object timeline detail
// @Tags Comment
// @Produce json
// @Param revision_id query string true "revision id"
// @Success 200 {object} handler.RespBody{data=schema.GetObjectTimelineResp}
// @Router /answer/api/v1/activity/timeline/detail [get]
func (ac *ActivityController) GetObjectTimelineDetail(ctx *gin.Context) {
	req := &schema.GetObjectTimelineDetailReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	resp, err := ac.activityService.GetObjectTimelineDetail(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
