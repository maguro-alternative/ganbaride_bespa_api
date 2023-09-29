package model

type (
	RiderSkill struct {
		SkillName		   string
		SkillType		   int
		SkillConditions	   string
		SkillBonus		   SkillBonus
	}
	// ライダースキルの構造体
	SkillBonus struct {
		MyAttack             int
		MyDefence            int
		MyHitPoint           int
		MyFinishingAttack    int
		MyAttackPoint        int
		EnemyAttack          int
		EnemyDefence         int
		EnemyHitPoint        int
		EnemyFinishingAttack int
		EnemyAttackPoint     int
	}
)