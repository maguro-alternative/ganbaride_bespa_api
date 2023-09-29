package model

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
type (
	Rider struct {
		Number              string   `json:"number"`
		Name                string   `json:"name"`
		Rarity              int      `json:"rarity"`
		Category            []string `json:"category"`
		Attribute           string   `json:"attribute"`
		Attack              int      `json:"attack"`
		Defence             int      `json:"defence"`
		HitPoint            int      `json:"hitpoint"`
		FinishingAttack     int      `json:"finishingAttack"`
		SkillName           string   `json:"skillName"`
		AttackTeki          int      `json:"attackTeki"`
		DefenceTeki         int      `json:"defenceTeki"`
		HitPointTeki        int      `json:"hitpointTeki"`
		FinishingAttackTeki int      `json:"finishingAttackTeki"`
	}
)
