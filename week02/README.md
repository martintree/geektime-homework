## 问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
### 应该Wrap这个error，抛给上层，在上层对此类型的error做相应的处理 