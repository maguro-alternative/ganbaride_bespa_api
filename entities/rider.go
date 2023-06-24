package entities


/*
Name
	ライダーの名前
Category
	カテゴリー
Attribute
	属性(力、技、速)
Attack
	コウゲキ
Defence
	ボウギョ
HitPoint
	タイリョク
FinishingAttack
	ヒッサツ
SkillName
	ライダースキル名
SkillType
	ライダースキルの条件
SkillConditions
	仲間が○○の場合、敵が○○の場合、などのカテゴリー名(ディケイド、電王など)
SkillBonus
	スキルボーナス
AttackTeki
	コウゲキ適正
DefenceTeki
	ボウギョ適正
HitPointTeki
	タイリョク適正
FinishingAttackTeki
	ヒッサツ適正
*/
type Rider struct{
	Name				string					`json:"name"`
	Category			[]string				`json:"category"`
	Attribute			string					`json:"attribute"`
	Attack				int						`json:"attack"`
	Defence				int						`json:"defence"`
	HitPoint			int						`json:"hitpoint"`
	FinishingAttack		int						`json:"finishingAttack"`
	SkillName			string					`json:"skillname"`
	SkillType			int						`json:"skilltype"`
	SkillConditions		string					`json:"skillconditions"`
	SkillBonus			SkillBonus				`json:"skillbonus"`
	AttackTeki			int						`json:"suitableAttack"`
	DefenceTeki			int						`json:"suitableDefence"`
	HitPointTeki		int						`json:"suitableHitpoint"`
	FinishingAttackTeki	int						`json:"suitableFinishingAttack"`
}