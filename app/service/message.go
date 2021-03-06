package service

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cast"

	"gorm.io/gorm"

	"github.com/go-pay/gopay/pkg/util"

	"github.com/silenceper/wechat/v2/officialaccount/basic"

	"github.com/silenceper/wechat/v2/officialaccount"

	"github.com/zngue/go_user_login/app/model"

	"github.com/silenceper/wechat/v2/officialaccount/message"

	"github.com/gin-gonic/gin"
	"github.com/zngue/go_user_login/app/request"
)

type ScanResp struct {
	Type    int
	Action  string
	UserStr string
	Code    string
}
type DataLogin struct {
	ImgUrl     string `json:"imgUrl"`
	RandStr    string `json:"randStr"`
	ExpireTime int    `json:"expireTime"`
}
type Message interface {
	SetMessage(ctx *gin.Context, id string)
	QrcodeCreate(id string, action string) (data *DataLogin, err error)
}
type wechatMessage struct {
}

func (m *wechatMessage) UserLogin(office *officialaccount.OfficialAccount, openid string) (*model.User, error) {

	var request request.UserRequest
	request.OpenID = openid
	detail, err2 := NewUser().Detail(&request)
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return nil, err2
	}
	if detail != nil && detail.ID > 0 {
		return detail, nil
	}
	info, err3 := office.GetUser().GetUserInfo(openid)
	if err3 != nil {
		return nil, err3
	}
	userinfo := &model.User{
		Subscribe:     info.Subscribe,
		OpenID:        info.OpenID,
		Nickname:      info.Nickname,
		Sex:           info.Sex,
		City:          info.City,
		Country:       info.Country,
		Province:      info.Province,
		Headimgurl:    info.Headimgurl,
		SubscribeTime: info.SubscribeTime,
		UnionID:       info.UnionID,
	}
	return userinfo, nil

}

func (m *wechatMessage) MixMessageAnalysis(mixMessage *message.MixMessage, office *officialaccount.OfficialAccount) (*message.Reply, error) {

	info, err := office.GetUser().GetUserInfo(string(mixMessage.FromUserName))
	fmt.Println(info, err)
	if mixMessage.Event == message.EventScan || mixMessage.Event == message.EventSubscribe {

		messageStr := ""
		if mixMessage.Event == message.EventScan {
			messageStr = mixMessage.EventKey
		} else {
			messageStr = info.QrSceneStr
		}
		if len(messageStr) > 0 {
			var re ScanResp
			err2 := json.Unmarshal([]byte(messageStr), &re)
			if err2 != nil {
				return nil, err2
			}
			userInfo, err3 := m.UserLogin(office, string(mixMessage.FromUserName))
			if err3 != nil {
				return nil, err3
			}
			if userInfo == nil {
				return nil, nil
			}
			if re.Action == "login" {
				if err := NewUser().EditUserAndAction(userInfo, &re); err != nil {
					return nil, err
				}
				text := message.NewText(fmt.Sprintf("??????%s????????????", userInfo.Nickname))
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}, nil
			}
			if re.Action == "code" {
				re.Code = util.GetRandomNumber(6)
				if err := NewUser().EditUserAndAction(userInfo, &re); err != nil {
					return nil, err
				}
				text := message.NewText(fmt.Sprintf("???zngue???????????????????????????%s?????????????????????????????????5????????????????????????????????????????????????????????????", re.Code))
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}, nil
			}
		}
	}
	return nil, nil
}

func (m *wechatMessage) QrcodeCreate(id string, action string) (dataLogin *DataLogin, err error) {
	var accountReq request.AccountRequest
	accountReq.ID = cast.ToInt(id)
	office, errs := m.GetOffice(&accountReq)
	if errs != nil {
		return nil, errs
	}
	basics := office.GetBasic()
	expireTime := 300
	var data = basic.Request{
		ExpireSeconds: int64(expireTime),
		ActionName:    "QR_STR_SCENE",
	}
	randomString := util.GetRandomString(10)
	SceneStr := fmt.Sprintf(`{"type":1,"action":"%s","userstr":"%s"}`, action, randomString)
	data.ActionInfo.Scene.SceneID = 1
	data.ActionInfo.Scene.SceneStr = SceneStr
	ticket, errQr := basics.GetQRTicket(&data)
	if errQr != nil {
		return nil, errQr
	}
	imageUrl := basic.ShowQRCode(ticket)
	return &DataLogin{
		ImgUrl:     imageUrl,
		RandStr:    randomString,
		ExpireTime: expireTime,
	}, nil

}
func (m *wechatMessage) GetOffice(accentReq *request.AccountRequest) (*officialaccount.OfficialAccount, error) {
	accountOne, err := m.GetAccount(accentReq)
	if err != nil {
		return nil, err
	}
	office := NewWechat().AccountOffice(accountOne)
	return office, nil
}

func (m *wechatMessage) GetAccount(accentReq *request.AccountRequest) (*model.Account, error) {
	return NewAccount().Detail(accentReq)
}

func (m *wechatMessage) SetMessage(ctx *gin.Context, id string) {
	var accountReq request.AccountRequest
	accountReq.ID = cast.ToInt(id)
	office, err := m.GetOffice(&accountReq)
	if err != nil {
		return
	}
	server := office.GetServer(ctx.Request, ctx.Writer)
	server.SetMessageHandler(func(mixMessage *message.MixMessage) *message.Reply {

		reply := &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: message.NewText("????????????????????????????????????"),
		}
		analysis, err2 := m.MixMessageAnalysis(mixMessage, office)
		if err2 != nil {
			return reply
		}
		reply = analysis
		return reply
	})
	if err = server.Serve(); err != nil {
		fmt.Println(err)
		return
	}
	if err = server.Send(); err != nil {
		fmt.Println(err)
		return
	}
}

func NewMessage() Message {
	return new(wechatMessage)
}
