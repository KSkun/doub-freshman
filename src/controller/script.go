package controller

import (
	"github.com/KSkun/doub-freshman/controller/param"
	"github.com/KSkun/doub-freshman/model"
	"github.com/KSkun/doub-freshman/util/context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"regexp"
)

func retreiveOptionBranch(optionBranch param.ReqOptionBranch,mp map[string]primitive.ObjectID)model.OptionBranch{
	var events []model.Event
	for _,event:=range optionBranch.Event{
		events=append(events,model.Event(event))
	}
	res:=model.OptionBranch{
		Next: mp[optionBranch.Next],
		Text: optionBranch.Text,
		Event: events,
	}
	return res
}


func retreiveOption(option param.ReqOption,mp map[string]primitive.ObjectID)model.Option{
	var conditions []model.Condition
	for _,condition:=range option.Condition{
		conditions=append(conditions,model.Condition(condition))
	}
	res:=model.Option{
		Text:      option.Text,
		Success:   retreiveOptionBranch(option.Success,mp),
		Failed:    retreiveOptionBranch(option.Failed,mp),
		Condition: conditions,
	}
	return res
}

func retreiveStage(stage param.ReqStage,mp map[string]primitive.ObjectID)model.Stage{
	var enterConds []model.Condition
	for _,enterCond:=range stage.EnterCond{
		enterConds=append(enterConds,model.Condition(enterCond))
	}
	var options []model.Option
	for _,option:=range stage.Option{
		options=append(options,retreiveOption(option,mp))
	}

	res:=model.Stage{
		ID:        mp[stage.ID],
		Title:     stage.Title,
		Text:      stage.Text,
		Dead:      stage.Dead,
		EnterCond: enterConds,
		Option:    options,
		Tag:       stage.Tag,
		Continue:  stage.Continue,
		Delay:     stage.Delay,
	}
	return res
}

func HandlePostStage(ctx echo.Context)error{
	type ReqStages struct {
		Stage []param.ReqStage `json:"stages"`
	}
	var req ReqStages
	if err:=ctx.Bind(&req);err!=nil{
		return context.Error(ctx,http.StatusBadRequest,"",err)
	}
	if err:=ctx.Validate(&req);err!=nil{
		return context.Error(ctx,http.StatusBadRequest,"",err)
	}
	mp:=make(map[string]primitive.ObjectID)
	for _,stage:=range req.Stage{
		err:=ConvertIdentifierToObjectId(stage.ID,mp)
		if err!=nil{
			return context.Error(ctx,http.StatusBadRequest,"",err)
		}
	}

	var stages []model.Stage
	for _,stage:=range req.Stage{
		stages=append(stages,retreiveStage(stage,mp))
	}
	m := model.GetModel()
	defer m.Close()
	err:=m.AddStages(stages)
	if err!=nil{
		return context.Error(ctx,http.StatusInternalServerError,"",err)
	}
	return context.Success(ctx,nil)
}

func ConvertIdentifierToObjectId(identifier string ,mp map[string]primitive.ObjectID)error{
	if identifier==""{
		return nil
	}
	matched,err:=regexp.MatchString("stage_",identifier)
	if err!=nil{
		return err
	}
	if matched==true{
		_,ok:=mp[identifier]
		if ok==false{
			mp[identifier]=primitive.NewObjectID()
		}
	}else {
		id,err:=primitive.ObjectIDFromHex(identifier)
		if err!=nil{
			return err
		}
		mp[identifier]=id
	}
	return nil
}