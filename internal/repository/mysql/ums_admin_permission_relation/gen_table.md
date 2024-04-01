#### mall.ums_admin_permission_relation 
后台用户和权限关系表(除角色中定义的权限以外的加减权限)

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | admin_id |  | bigint |  | YES |  |  |
| 3 | permission_id |  | bigint |  | YES |  |  |
| 4 | type |  | int |  | YES |  |  |
