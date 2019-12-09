package mircodec

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/yenkeia/mirgo/common"
	"github.com/yenkeia/mirgo/proto/server"
)

// MirUserInformationCodec ...
type MirUserInformationCodec struct{}

// Name 返回名字
func (m *MirUserInformationCodec) Name() string {
	return "MirUserInformationCodec"
}

// MimeType 我也不知道是干嘛的
func (m *MirUserInformationCodec) MimeType() string {
	return "application/binary"
}

// Encode 将数据转换为字节数组
func (*MirUserInformationCodec) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {
	return encode(msgObj)
}

// TODO Decode 将字节数组转换为数据
func (*MirUserInformationCodec) Decode(data interface{}, msgObj interface{}) error {
	ui := msgObj.(*server.UserInformation)
	bytes := data.([]byte)
	reader := &BytesWrapper{Bytes: &bytes}
	ui.ObjectID = reader.ReadUInt32()
	ui.RealId = reader.ReadUInt32()
	ui.Name = reader.ReadString()
	ui.GuildName = reader.ReadString()
	ui.GuildRank = reader.ReadString()
	ui.NameColour = reader.ReadUInt32()
	ui.Class = common.MirClass(reader.ReadByte())
	ui.Gender = common.MirGender(reader.ReadByte())
	ui.Level = reader.ReadUInt16()
	x := reader.ReadInt32()
	y := reader.ReadInt32()
	ui.Location = common.Point{X: uint32(x), Y: uint32(y)}
	ui.Direction = common.MirDirection(reader.ReadByte())
	ui.Hair = reader.ReadUInt8()
	ui.HP = reader.ReadUInt16()
	ui.MP = reader.ReadUInt16()
	ui.Experience = reader.ReadInt64()
	ui.MaxExperience = reader.ReadInt64()
	ui.LevelEffect = common.LevelEffects(reader.ReadUInt8())

	// Inventory
	if reader.ReadBoolean() {
		//count := reader.ReadInt32()
		//last := reader.Last()
		//for i := 0; i < int(count); i++ {
		//	if reader.ReadBoolean() {
		//		//common.UserItem{}
		//		last = decodeValue(nil, last)
		//	}
		//}
	}

	// Equipment
	if reader.ReadBoolean() {
	}

	// QuestInventory
	if reader.ReadBoolean() {
	}

	return decode(msgObj, bytes)
}

func init() {
	codec.RegisterCodec(new(MirUserInformationCodec))
}
