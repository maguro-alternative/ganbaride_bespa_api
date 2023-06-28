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
// ライダーステータス
type Aptitude struct {
	Name				string
	Category			[]string
	Attribute			string
	Attack 				int
	Defence 			int
	HitPoint 			int
	FinishingAttack 	int
	SkillName			string
	SkillType			int
	SkillConditions		string
	SkillBonus			SkillBonus
	AttackTeki 			int
	DefenceTeki 		int
	HitPointTeki 		int
	FinishingAttackTeki int
}

/*
0	無条件
1	ゼンエイのとき
2	コウエイのとき
3	ゼンエイのとき、仲間の属性が一緒の場合
4	コウエイのとき、仲間の属性が一緒の場合
5	仲間の属性が一緒の場合
6	ゼンエイのとき、仲間が○○のとき
7	コウエイのとき、仲間が○○のとき
8	あいての属性が一緒のとき
9	ゼンエイのとき、あいての属性が一緒のとき
10	コウエイのとき、あいての属性が一緒のとき
11	あいてが○○のとき
12	ゼンエイのとき、あいてが○○のとき
13	コウエイのとき、あいてが○○のとき
*/
// func (変数名 構造体) 関数名(変数名 型) 型
func (zenei Aptitude) ZeneiRiderSkill(
	kouei Aptitude,
	eneZen Aptitude,
	eneKou Aptitude,
) SkillBonus {
	var zeneiSkillType int = zenei.SkillType
	var bonus SkillBonus = *NewSkillBonus()
	
	// 無条件 ゼンエイの場合
	if (zeneiSkillType == 0 || zeneiSkillType == 1) {
		bonus = zenei.SkillBonus
	// ゼンエイ時、仲間と属性が一緒の場合
	}else if (zeneiSkillType == 3 && kouei.Attribute == zenei.Attribute){
		bonus = zenei.SkillBonus
	// 仲間と属性が一緒の場合
	}else if (zeneiSkillType == 5 && kouei.Attribute == zenei.Attribute){
		bonus = zenei.SkillBonus
	// ゼンエイ時、仲間が○○の場合
	}else if (zeneiSkillType == 6 && contains(kouei.Category,zenei.SkillConditions)){
		bonus = zenei.SkillBonus
	// 相手と属性が同じ場合
	}else if (zeneiSkillType == 8 && zenei.Attribute == eneZen.Attribute){
		bonus = zenei.SkillBonus
	// ゼンエイ時、相手と属性が同じ場合
	}else if (zeneiSkillType == 9 && zenei.Attribute == eneZen.Attribute){
		bonus = zenei.SkillBonus
	// 相手が○○の場合
	}else if (zeneiSkillType == 11 && (contains(eneZen.Category,zenei.SkillConditions) || contains(eneKou.Category,zenei.SkillConditions))){
		bonus = zenei.SkillBonus
	// ゼンエイジ相手に○○がいる場合
	}else if (zeneiSkillType == 12 && (contains(eneZen.Category,zenei.SkillConditions) || contains(eneKou.Category,zenei.SkillConditions))){
		bonus = zenei.SkillBonus
	}
	return bonus
}

func (kouei Aptitude) KoueiRiderSkill(
	zenei Aptitude,
	eneZen Aptitude,
	eneKou Aptitude,
) SkillBonus {
	var koueiSkillType int = kouei.SkillType
	var bonus SkillBonus = *NewSkillBonus()
	
	// 無条件 コウエイの場合
	if (koueiSkillType == 0 || koueiSkillType == 2) {
		bonus = kouei.SkillBonus
	// コウエイ時、仲間と属性が一緒の場合
	}else if (koueiSkillType == 4 && 
		kouei.Attribute == zenei.Attribute){

		bonus = kouei.SkillBonus
	// 仲間と属性が一緒の場合
	}else if (koueiSkillType == 5 && 
		kouei.Attribute == zenei.Attribute){

		bonus = kouei.SkillBonus
	// コウエイ時、仲間が○○の場合
	}else if (koueiSkillType == 7 && contains(
		zenei.Category,kouei.SkillConditions)){

		bonus = kouei.SkillBonus
	// 相手と属性が同じ場合
	}else if (koueiSkillType == 8 && 
		kouei.Attribute == eneZen.Attribute){

		bonus = kouei.SkillBonus
	// コウエイ時、相手と属性が同じ場合
	}else if (koueiSkillType == 10 && 
		kouei.Attribute == eneZen.Attribute){

		bonus = kouei.SkillBonus
	// 相手が○○の場合
	}else if (koueiSkillType == 11 && (
		contains(eneZen.Category,kouei.SkillConditions) || 
		contains(eneKou.Category,kouei.SkillConditions))){

		bonus = kouei.SkillBonus
	// コウエイ時相手に○○がいる場合
	}else if (koueiSkillType == 13 && (
		contains(eneZen.Category,kouei.SkillConditions) || 
		contains(eneKou.Category,kouei.SkillConditions))){

		bonus = kouei.SkillBonus
	}
	return bonus
}

// 配列に指定したものが含まれているか確認
func contains(elems []string, v string) bool {
    for _, s := range elems {
        if v == s {
            return true
        }
    }
    return false
}

// 配列に指定したものが含まれていた場合、場所を返す
func ContainsPoint(elems []int, v int) int {
	var i int = 0
    for _, s := range elems {
        if v == s {
            return i
        }
		i += 1
    }
    return -1
}