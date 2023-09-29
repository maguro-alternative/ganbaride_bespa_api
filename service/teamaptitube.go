package service

import (
	"database/sql"

	"github.com/maguro-alternative/ganbaride_bespa_api/model"
)

type TeamAptitubeService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTeamAptitubeService(db *sql.DB) *TeamAptitubeService {
	return &TeamAptitubeService{
		db: db,
	}
}

func (s TeamAptitubeService) ZeneiRiderSkill(
	zenei model.Rider,
	kouei model.Rider,
	eneZen model.Rider,
	eneKou model.Rider,
) model.SkillBonus {
	var bonus model.SkillBonus = *model.NewSkillBonus()
	var zeneiSkillType int = s.skillTypeSearch(zenei.SkillName)
	var riderSkill model.RiderSkill = s.skillSearch(zenei.SkillName)
	var skillBonus model.SkillBonus = model.SkillBonus{
		MyAttack:             riderSkill.MyAttack,
		MyDefence:            riderSkill.MyDefence,
		MyHitPoint:           riderSkill.MyHitPoint,
		MyFinishingAttack:    riderSkill.MyFinishingAttack,
		MyAttackPoint:        riderSkill.MyAttackPoint,
		EnemyAttack:          riderSkill.EnemyAttack,
		EnemyDefence:         riderSkill.EnemyDefence,
		EnemyHitPoint:        riderSkill.EnemyHitPoint,
		EnemyFinishingAttack: riderSkill.EnemyFinishingAttack,
		EnemyAttackPoint:     riderSkill.EnemyAttackPoint,
	}

	// 無条件 ゼンエイの場合
	if (zeneiSkillType == 0 || zeneiSkillType == 1) {
		bonus = skillBonus
	// ゼンエイ時、仲間と属性が一緒の場合
	}else if (zeneiSkillType == 3 && kouei.Attribute == zenei.Attribute){
		bonus = skillBonus
	// 仲間と属性が一緒の場合
	}else if (zeneiSkillType == 5 && kouei.Attribute == zenei.Attribute){
		bonus = skillBonus
	// ゼンエイ時、仲間が○○の場合
	}else if (zeneiSkillType == 6 && contains(kouei.Category,riderSkill.SkillConditions)){
		bonus = skillBonus
	// 相手と属性が同じ場合
	}else if (zeneiSkillType == 8 && zenei.Attribute == eneZen.Attribute){
		bonus = skillBonus
	// ゼンエイ時、相手と属性が同じ場合
	}else if (zeneiSkillType == 9 && zenei.Attribute == eneZen.Attribute){
		bonus = skillBonus
	// 相手が○○の場合
	}else if (zeneiSkillType == 11 && (contains(eneZen.Category,riderSkill.SkillConditions) || contains(eneKou.Category,riderSkill.SkillConditions))){
		bonus = skillBonus
	// ゼンエイジ相手に○○がいる場合
	}else if (zeneiSkillType == 12 && (contains(eneZen.Category,riderSkill.SkillConditions) || contains(eneKou.Category,riderSkill.SkillConditions))){
		bonus = skillBonus
	}
	return bonus
}

func (s TeamAptitubeService) KoueiRiderSkill(
	zenei model.Rider,
	kouei model.Rider,
	eneZen model.Rider,
	eneKou model.Rider,
) model.SkillBonus {
	var bonus model.SkillBonus = *model.NewSkillBonus()
	var koueiSkillType int = s.skillTypeSearch(kouei.SkillName)
	var riderSkill model.RiderSkill = s.skillSearch(kouei.SkillName)
	var skillBonus model.SkillBonus = model.SkillBonus{
		MyAttack:             riderSkill.MyAttack,
		MyDefence:            riderSkill.MyDefence,
		MyHitPoint:           riderSkill.MyHitPoint,
		MyFinishingAttack:    riderSkill.MyFinishingAttack,
		MyAttackPoint:        riderSkill.MyAttackPoint,
		EnemyAttack:          riderSkill.EnemyAttack,
		EnemyDefence:         riderSkill.EnemyDefence,
		EnemyHitPoint:        riderSkill.EnemyHitPoint,
		EnemyFinishingAttack: riderSkill.EnemyFinishingAttack,
		EnemyAttackPoint:     riderSkill.EnemyAttackPoint,
	}


	// 無条件 コウエイの場合
	if (koueiSkillType == 0 || koueiSkillType == 2) {
		bonus = skillBonus
	// コウエイ時、仲間と属性が一緒の場合
	}else if (koueiSkillType == 4 && kouei.Attribute == zenei.Attribute){
		bonus = skillBonus
	// 仲間と属性が一緒の場合
	}else if (koueiSkillType == 5 && kouei.Attribute == zenei.Attribute){
		bonus = skillBonus
	// コウエイ時、仲間が○○の場合
	}else if (koueiSkillType == 7 && contains(zenei.Category,riderSkill.SkillConditions)){
		bonus = skillBonus
	// 相手と属性が同じ場合
	}else if (koueiSkillType == 8 && kouei.Attribute == eneZen.Attribute){
		bonus = skillBonus
	// コウエイ時、相手と属性が同じ場合
	}else if (koueiSkillType == 10 && kouei.Attribute == eneZen.Attribute){
		bonus = skillBonus
	// 相手が○○の場合
	}else if (koueiSkillType == 11 && (contains(eneZen.Category,riderSkill.SkillConditions) || contains(eneKou.Category,riderSkill.SkillConditions))){
		bonus = skillBonus
	// コウエイ時相手に○○がいる場合
	}else if (koueiSkillType == 13 && (contains(eneZen.Category,riderSkill.SkillConditions) || contains(eneKou.Category,riderSkill.SkillConditions))){
		bonus = skillBonus
	}
	return bonus
}

func (s TeamAptitubeService) skillTypeSearch(skillName string) int {
	sqlQuery := `select skilltype from riderSkill where skill_name = ?`
	var skillType int
	err := s.db.QueryRow(sqlQuery, skillName).Scan(&skillType)
	if err != nil {
		panic(err.Error())
	}
	return skillType
}

func (s TeamAptitubeService) skillSearch(skillName string) model.RiderSkill {
	sqlQuery := `
	select
		skillname,
		skilltype,
		skillconditions,
		myattack,
		mydefence,
		myhitpoint,
		myfinishingattack,
		myattackpoint,
		enemyattack,
		enemydefence,
		enemyhitpoint,
		enemyfinishingattack,
		enemyattackpoint
	from riderSkill where skill_name = ?
	`
	var skill model.RiderSkill
	err := s.db.QueryRow(sqlQuery, skillName).Scan(&skill)
	if err != nil {
		panic(err.Error())
	}
	return skill
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