CREATE TABLE IF NOT EXISTS rider(
    number varchar(20) PRIMARY KEY,
    name varchar(50) not null,
    rarity int not null,
    category varchar(50)[] not null,
    attribute varchar(2) not null,
    attack bigint not null,
    defence bigint not null,
    hitpoint bigint not null,
    finishingattack bigint not null,
    skillname varchar(50) not null,
    attackteki bigint not null,
    defenceteki bigint not null,
    hitpointteki bigint not null,
    finishingattackteki bigint not null
);

INSERT INTO rider(
    number,
    name,
    rarity,
    category,
    attribute,
    attack,
    defence,
    hitpoint,
    finishngattack,
    skillname,
    attackteki,
    defenceteki,
    hitpointteki,
    finishngattackteki
) VALUES
    ('1-001','仮面ライダーキバ キバフォーム',5,'{"キバ"}','技',450,400,600,2500,'秘められた魔皇力',1,0,2,3),
    ('1-002','仮面ライダーキバ キバフォーム',1,'{"キバ"}','技',400,300,500,2000,'ヘルズゲート解放',0,1,2,0),
    ('1-003','仮面ライダーキバ ガルルフォーム',2,'{"キバ","フォームチェンジ","ぶき"}','速',400,350,550,2300,'ワイルドブルー',0,3,1,0),
    ('1-004','仮面ライダーキバ ガルルフォーム',1,'{"キバ","フォームチェンジ","ぶき"}','速',350,400,550,2200,'猛々しき咆吼',2,0,0,1),
    ('1-005','仮面ライダーキバ バッシャーフォーム',2,'{"キバ","フォームチェンジ","ぶき"}','技',350,400,550,2300,'碧の魔皇力',1,0,1,2),
    ('1-006','仮面ライダーキバ バッシャーフォーム',1,'{"キバ","フォームチェンジ","ぶき"}','技',350,450,500,1900,'水棲の超感覚',1,1,1,0),
    ('1-007','仮面ライダーキバ ドッガフォーム',1,'{"キバ","フォームチェンジ","ぶき"}','力',500,400,450,2300,'パープルサンダー',2,0,1,0),
    ('1-008','仮面ライダーキバ ドッガフォーム',1,'{"キバ","フォームチェンジ","ぶき"}','力',450,400,400,2100,'紫の魔皇力',0,2,1,0),
    ON CONFLICT DO NOTHING;


CREATE TABLE IF NOT EXISTS riderSkill(
    skillname varchar(50) PRIMARY KEY,
    skilltype smallint not null,
    skillconditions varchar(50) not null,
    myattack bigint not null,
    mydefence bigint not null,
    myhitpoint bigint not null,
    myfinishingattack bigint not null,
    myattackpoint bigint not null,
    enemyattack bigint not null,
    enemydefence bigint not null,
    enemyhitpoint bigint not null,
    enemyfinishingattack bigint not null,
    enemyattackpoint bigint not null
);


INSERT INTO riderSkill(
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
) VALUES
    ('秘められた魔皇力',3,'',0,0,150,0,0,0,0,0,0,0),
    ('ヘルズゲート解放',2,'',0,0,150,0,0,0,0,0,0,0),
    ('ワイルドブルー',3,'',0,0,0,0,20,0,0,0,0,0),
    ('猛々しき咆吼',1,'',0,0,0,0,0,0,0,0,0,-20),
    ('碧の魔皇力',2,'',150,0,0,0,0,0,0,0,0,0),
    ('水棲の超感覚',3,'',100,0,0,0,0,0,0,0,0,0),
    ('パープルサンダー',9,'',100,0,0,0,0,0,0,0,0,0),
    ('パープルサンダー',3,'',150,0,0,0,0,0,0,0,0,0),
    ON CONFLICT DO NOTHING;
