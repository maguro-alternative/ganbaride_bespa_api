package entities

/*
ライダースキルの構造体
*/
type SkillBonus struct {
	MyAttack				int
	MyDefence				int
	MyHitPoint				int
	MyFinishingAttack		int
	MyAttackPoint			int
	EnemyAttack				int
	EnemyDefence			int
	EnemyHitPoint			int
	EnemyFinishingAttack	int
	EnemyAttackPoint		int
}

// コンストラクタ関数を定義する
func NewSkillBonus() *SkillBonus {
	// デフォルト値を設定する
	return &SkillBonus{
	  MyAttack: 0,
	  MyDefence: 0,
	  MyHitPoint: 0,
	  MyFinishingAttack: 0,
	  MyAttackPoint: 0,
	  EnemyAttack: 0,
	  EnemyDefence: 0,
	  EnemyHitPoint: 0,
	  EnemyFinishingAttack: 0,
	  EnemyAttackPoint: 0,
	}
}